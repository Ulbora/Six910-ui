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
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
)

var mu sync.Mutex

//Image Image
type Image struct {
	Name     string
	ImageURL string
}

//AddImage AddImage
func (c *Six910ImageService) AddImage(name string, fileData []byte) bool {
	mu.Lock()
	defer mu.Unlock()
	var rtn bool
	c.Log.Debug("image file name in add: ", name)
	var imageName = c.ImagePath + string(filepath.Separator) + name
	c.Log.Debug("image complete file name in add: ", imageName)
	err := ioutil.WriteFile(imageName, fileData, 0644)
	if err == nil {
		rtn = true
	}
	return rtn
}

//GetImageList GetImageList
func (c *Six910ImageService) GetImageList() *[]Image {
	var rtn []Image
	ifiles, err := ioutil.ReadDir(c.ImagePath)
	if err == nil {
		for _, ifile := range ifiles {
			if !ifile.IsDir() {
				//fmt.Println("sfile: ", sfile)
				var imgfile Image
				imgfile.Name = ifile.Name()
				imgfile.ImageURL = ".." + string(filepath.Separator) + ".." + string(filepath.Separator) + "images" + string(filepath.Separator) + ifile.Name()
				c.Log.Debug("image ImageURL in list: ", imgfile.ImageURL)
				rtn = append(rtn, imgfile)
			}
		}
	}
	return &rtn
}

//GetImagePath GetImagePath
func (c *Six910ImageService) GetImagePath(imageName string) string {
	return c.ImageFullPath + string(filepath.Separator) + imageName
}

//DeleteImage DeleteImage
func (c *Six910ImageService) DeleteImage(name string) bool {
	mu.Lock()
	defer mu.Unlock()
	var rtn bool
	var imageName = c.ImagePath + string(filepath.Separator) + name
	c.Log.Debug("image complete file name in delete: ", imageName)
	jerr := os.Remove(imageName)
	if jerr == nil {
		rtn = true
	}
	return rtn
}
