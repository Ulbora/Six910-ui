package bkupsrv

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
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

func TestCmsService_DownloadBackups(t *testing.T) {
	var cs Six910BackupService
	cs.CarouselStorePath = "./testBackup/carouselStore"
	cs.ContentStorePath = "./testBackup/contentStore"
	cs.CountryStorePath = "./testBackup/countryStore"
	cs.CSSStorePath = "./testBackup/cssStore"
	cs.MenuStorePath = "./testBackup/menuStore"
	cs.StateStorePath = "./testBackup/stateStore"
	cs.TemplateStorePath = "./testBackup/templateStore"
	cs.ImagePath = "./testBackup/images"
	cs.TemplateFilePath = "./testBackup/templates"

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()
	suc, f := s.DownloadBackups()
	if !suc || f == nil {
		t.Fail()
	}

	fileToWrite, err := os.OpenFile("./testBackupZips/compress.dat", os.O_CREATE|os.O_RDWR, os.FileMode(600))
	fmt.Println("fileToWrite err : ", err)
	var buf = bytes.NewBuffer(*f)

	//fmt.Println("compress file data: ", buf.Bytes())
	_, err2 := io.Copy(fileToWrite, buf)
	fmt.Println("io.copy err : ", err2)
	os.Chmod("./testBackupZips/compress.dat", os.FileMode(0666))

}

func TestCmsService_UploadBackups(t *testing.T) {

	fileData, rerr := ioutil.ReadFile("./testBackupZips/compress.dat")
	fmt.Println(rerr)

	var b bytes.Buffer
	b.Write(fileData)
	r, err := zlib.NewReader(&b)
	if err == nil {
		var out bytes.Buffer
		io.Copy(&out, r)
		r.Close()
		rtn := out.Bytes()
		ioutil.WriteFile("./testBackupZips/uncompress.json", rtn, 0644)
	}
	var cs Six910BackupService
	cs.CarouselStorePath = "./testBackupRestore/carouselStore"
	cs.ContentStorePath = "./testBackupRestore/contentStore"
	cs.CountryStorePath = "./testBackupRestore/countryStore"
	cs.CSSStorePath = "./testBackupRestore/cssStore"
	cs.MenuStorePath = "./testBackupRestore/menuStore"
	cs.StateStorePath = "./testBackupRestore/stateStore"
	cs.TemplateStorePath = "./testBackupRestore/templateStore"
	cs.ImagePath = "./testBackupRestore/images"
	cs.TemplateFilePath = "./testBackupRestore/templates"

	var cds ds.DataStore
	cds.Path = "./testBackupRestore/contentStore"
	cs.Store = cds.GetNew()

	var tds ds.DataStore
	tds.Path = "./testBackupRestore/templateStore"
	cs.TemplateStore = tds.GetNew()

	var crds ds.DataStore
	crds.Path = "./testBackupRestore/carouselStore"
	cs.CarouselStore = crds.GetNew()

	var cyds ds.DataStore
	cyds.Path = "./testBackupRestore/countryStore"
	cs.CountryStore = cyds.GetNew()

	var cssds ds.DataStore
	cssds.Path = "./testBackupRestore/cssStore"
	cs.CSSStore = cssds.GetNew()

	var mds ds.DataStore
	mds.Path = "./testBackupRestore/menuStore"
	cs.MenuStore = mds.GetNew()

	var sds ds.DataStore
	sds.Path = "./testBackupRestore/stateStore"
	cs.StateStore = sds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()
	suc := s.UploadBackups(&fileData)
	if !suc {
		t.Fail()
	}
}
