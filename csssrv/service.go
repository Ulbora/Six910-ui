package csssrv

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

//page css service

import (
	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

//CSSService CSSService
type CSSService interface {
	GetPageCSS(name string) (bool, *PageCSS)
	GetPage(name string) (bool, *Page)
	UpdatePage(page *Page) bool
}

//Six910CSSService Six910CSSService
type Six910CSSService struct {
	CSSStore     ds.JSONDatastore
	CSSStorePath string
	Log          *lg.Logger
}

//GetNew GetNew
func (c *Six910CSSService) GetNew() CSSService {
	return c
}
