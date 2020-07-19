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
	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

//ImageService ImageService
type ImageService interface {
	AddImage(name string, fileData []byte) bool
	GetImagePath(imageName string) string
	GetImageList() *[]Image
	DeleteImage(name string) bool
}

//Six910ImageService Six910ImageService
type Six910ImageService struct {
	Store ds.JSONDatastore

	Log           *lg.Logger
	ImagePath     string
	ImageFullPath string
}

//GetNew GetNew
func (c *Six910ImageService) GetNew() ImageService {
	return c

}
