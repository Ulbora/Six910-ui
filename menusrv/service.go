package menusrv

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

//Menu Service

import (
	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

//MenuService MenuService
type MenuService interface {
	AddMenu(menu *Menu) bool
	UpdateMenu(menu *Menu) bool
	GetMenu(name string) (bool, *Menu)
	GetMenuList() *[]Menu
	DeleteMenu(name string) bool
}

//Six910MenuService Six910MenuService
type Six910MenuService struct {
	MenuStore     ds.JSONDatastore
	MenuStorePath string
	Log           *lg.Logger
}

//GetNew GetNew
func (c *Six910MenuService) GetNew() MenuService {
	return c
}
