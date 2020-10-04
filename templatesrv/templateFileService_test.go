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
)

func TestTemplateFileService_ExtractFile(t *testing.T) {
	var cs Six910TemplateService
	cs.TemplateFilePath = "./testDownloads"

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	tf, err := os.Open("./testUploads/testTxt.tar.gz")
	if err != nil {
		fmt.Println("tar file not found!")
		os.Exit(1)
	}
	defer tf.Close()
	var ts TemplateFile
	ts.OriginalFileName = tf.Name()
	ts.Name = "newTemplate"
	data, err := ioutil.ReadAll(tf)
	if err != nil {
		fmt.Println(err)
	} else {
		ts.FileData = data
	}
	fmt.Print("file name: ")
	fmt.Println(ts.OriginalFileName)
	res := s.ExtractFile(&ts)
	if res != true {
		t.Fail()
	}
}

func TestTemplateFileService_DeleteTemplate(t *testing.T) {
	var cs Six910TemplateService
	cs.TemplateFilePath = "./testDownloads"

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	res := s.DeleteTemplateFile("newTemplate")
	if res != true {
		t.Fail()
	}
}
