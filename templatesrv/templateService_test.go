package templatesrv

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
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

var cit Six910TemplateService
var csit TemplateService

func TestCmsService_AddTemplate(t *testing.T) {

	var ds ds.DataStore
	ds.Path = "./testTemplateFiles"
	ds.Delete("temp1")
	cit.TemplateStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cit.Log = &l

	csit = cit.GetNew()

	var tmp Template
	tmp.Name = "temp1"
	suc := csit.AddTemplate(&tmp)
	if !suc {
		t.Fail()
	}
}

func TestCmsService_GetTemplateList(t *testing.T) {
	tlist := csit.GetTemplateList()
	if len(*tlist) < 1 {
		t.Fail()
	}
}

func TestCmsService_ActivateTemplate(t *testing.T) {
	suc := csit.ActivateTemplate("temp1")
	if !suc {
		t.Fail()
	}
}

func TestCmsService_DeleteTemplate(t *testing.T) {
	suc := csit.DeleteTemplate("temp1")
	if suc {
		t.Fail()
	}
}

func TestCmsService_ActivateTemplat2e(t *testing.T) {
	suc := csit.ActivateTemplate("temp2")
	if !suc {
		t.Fail()
	}
}

func TestCmsService_DeleteTemplate2(t *testing.T) {
	suc := csit.DeleteTemplate("temp1")
	if !suc {
		t.Fail()
	}
}

func TestCmsService_AddTemplateFile(t *testing.T) {
	var ci Six910TemplateService
	var csi TemplateService
	ci.TemplateFilePath = "./testUploadTemplates"
	//ci.ImageFullPath = "./testUploadTemplates"

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ci.Log = &l

	csi = ci.GetNew()
	tmpfile, err := os.Open("./testUploads/testTxt.tar.gz")
	fmt.Println("tmpfile: ", tmpfile.Name())
	if err != nil {
		fmt.Println("tmp file not found!")
		os.Exit(1)
	}
	defer tmpfile.Close()
	var originalFileName = tmpfile.Name()
	//i := strings.Index(originalFileName, ".")
	//var fileName = string(originalFileName[:i])
	fmt.Println("originalFileName in add template file: ", originalFileName)
	//fmt.Println("fileName in add template file: ", fileName)
	data, err := ioutil.ReadAll(tmpfile)
	if err != nil {
		fmt.Println(err)
	}
	suc := csi.AddTemplateFile("testTxt", originalFileName, data)
	if !suc {
		t.Fail()
	}
}

func TestCmsService_GetActiveTemplateName(t *testing.T) {
	var cs Six910TemplateService

	var ds ds.DataStore
	ds.Path = "./testTemplateFiles"
	cs.TemplateStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	name := s.GetActiveTemplateName()
	fmt.Println("active template name: ", name)
	if name == "" {
		t.Fail()
	}
}
