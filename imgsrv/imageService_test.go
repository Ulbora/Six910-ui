package imgsrv

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
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

var ci Six910ImageService
var csi ImageService

func TestCmsService_AddImage(t *testing.T) {
	ci.ImagePath = "./testUploadImages"
	ci.ImageFullPath = "./testUploadImages"

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ci.Log = &l

	csi = ci.GetNew()
	imgfile, err := os.Open("./testImages/test.jpg")
	fmt.Println("imgfile: ", imgfile.Name())
	if err != nil {
		fmt.Println("jpg file not found!")
		os.Exit(1)
	}
	defer imgfile.Close()
	var originalFileName = imgfile.Name()
	i := strings.LastIndex(originalFileName, ".")
	var ext = string(originalFileName[i:])
	fmt.Println("ext: ", ext)
	data, err := ioutil.ReadAll(imgfile)
	if err != nil {
		fmt.Println(err)
	}
	suc := csi.AddImage("testImage"+ext, data)
	if !suc {
		t.Fail()
	}
}

func TestCmsService_GetImagePath(t *testing.T) {
	fn := csi.GetImagePath("testImage.jpg")
	if fn != "./testUploadImages/testImage.jpg" {
		t.Fail()
	}
}

func TestCmsService_GetImageList(t *testing.T) {
	res := csi.GetImageList()
	fmt.Println("imageList: ", *res)
	if res == nil {
		t.Fail()
	}
}

func TestCmsService_DeleteImage(t *testing.T) {
	suc := csi.DeleteImage("testImage.jpg")
	if !suc {
		t.Fail()
	}
}
