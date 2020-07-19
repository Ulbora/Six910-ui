package tmptsrv

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2020 Ken Williamson
 All rights reserved.
 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

import (
	"archive/tar"
	"io"
	"os"
	"path/filepath"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

//TemplateService TemplateService
type TemplateService interface {
	AddTemplateFile(name string, originalFileName string, fileData []byte) bool
	AddTemplate(tpl *Template) bool
	ActivateTemplate(name string) bool
	GetActiveTemplateName() string
	GetTemplateList() *[]Template
	DeleteTemplate(name string) bool

	ExtractFile(tFile *TemplateFile) bool
	DeleteTemplateFile(name string) bool
}

//Six910TemplateService Six910TemplateService
type Six910TemplateService struct {
	Store             ds.JSONDatastore
	TemplateStore     ds.JSONDatastore
	ContentStorePath  string
	TemplateStorePath string
	TemplateFilePath  string
	TemplateFullPath  string
	Log               *lg.Logger
}

//GetNew GetNew
func (c *Six910TemplateService) GetNew() TemplateService {
	return c
}

func (c *Six910TemplateService) extractTarGzFile(tr *tar.Reader, h *tar.Header, dest string) error {
	var rtn error
	fname := h.Name
	c.Log.Debug("fname in extractTarGzFile in template service: ", fname)
	switch h.Typeflag {
	case tar.TypeDir:
		err := os.MkdirAll(dest+string(filepath.Separator)+fname, 0775)
		c.Log.Debug("MkdirAll in tar.TypeDir error in extractTarGzFile in template service:: ", err)
		c.Log.Debug("MkdirAll in tar.TypeDir name in extractTarGzFilein template service:: ", dest+string(filepath.Separator)+fname)
		rtn = err
	case tar.TypeReg:
		derr := os.MkdirAll(filepath.Dir(dest+string(filepath.Separator)+fname), 0775)
		rtn = derr
		c.Log.Debug("MkdirAll in tar.TypeReg error in extractTarGzFilein template service:: ", derr)
		c.Log.Debug("MkdirAll in tar.TypeReg dir name in extractTarGzFilein template service:: ", filepath.Dir(dest+string(filepath.Separator)+fname))
		if derr == nil {
			c.Log.Debug("MkdirAll in tar.TypeReg file name in extractTarGzFilein template service:: ", dest+string(filepath.Separator)+fname)
			writer, cerr := os.Create(dest + string(filepath.Separator) + fname)
			rtn = cerr
			c.Log.Debug("os.Create error in extractTarGzFilein template service:: ", cerr)
			if cerr == nil {
				io.Copy(writer, tr)
				err := os.Chmod(dest+string(filepath.Separator)+fname, 0664)
				c.Log.Debug("os.Chmod error in extractTarGzFilein template service:: ", err)
				rtn = err
				writer.Close()
			}
		}
	}
	return rtn
}
