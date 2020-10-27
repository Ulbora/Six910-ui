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

//Template Service

import (
	b64 "encoding/base64"
	"encoding/json"
	"html/template"
	//lg "github.com/Ulbora/Level_Logger"
	//ds "github.com/Ulbora/json-datastore"
)

//Menu Menu
type Menu struct {
	Name           string
	Location       string
	Active         bool
	Brand          string
	BrandHTML      template.HTML
	BrandLink      string
	Shade          string
	Background     string
	Style          string
	StyleCSS       template.CSS
	ShadeList      *[]string
	BackgroundList *[]string
	MenuItemList   *[]MenuItemItem
	CartCount      int64
	LoggedIn       bool
}

//MenuItemItem MenuItemItem
type MenuItemItem struct {
	Name string
	Link string
}

//Response Response
type Response struct {
	Success  bool   `json:"success"`
	Name     string `json:"name"`
	FailCode int    `json:"failCode"`
}

//AddMenu AddMenu
func (c *Six910MenuService) AddMenu(menu *Menu) bool {
	var rtn bool
	c.Log.Debug("menu in add: ", *menu)
	menu.Brand = b64.StdEncoding.EncodeToString([]byte(menu.Brand))
	em := c.MenuStore.Read(menu.Name)
	if *em == nil {
		suc := c.MenuStore.Save(menu.Name, menu)
		rtn = suc
	}
	return rtn
}

//UpdateMenu UpdateMenu
func (c *Six910MenuService) UpdateMenu(menu *Menu) bool {
	var rtn bool
	menu.Brand = b64.StdEncoding.EncodeToString([]byte(menu.Brand))
	c.Log.Debug("menu in update: ", *menu)
	em := c.MenuStore.Read(menu.Name)
	if *em != nil {
		var m Menu
		err := json.Unmarshal(*em, &m)
		c.Log.Debug("found menu in update: ", m)
		if err == nil {
			m.Active = menu.Active
			m.Brand = menu.Brand
			m.BrandLink = menu.BrandLink
			m.Background = menu.Background
			m.BackgroundList = menu.BackgroundList
			m.Shade = menu.Shade
			m.ShadeList = menu.ShadeList
			m.Style = menu.Style
			m.MenuItemList = menu.MenuItemList
			suc := c.MenuStore.Save(menu.Name, m)
			rtn = suc
		}
	}
	return rtn
}

//GetMenu GetMenu
func (c *Six910MenuService) GetMenu(name string) (bool, *Menu) {
	var rtn Menu
	var suc bool
	em := c.MenuStore.Read(name)
	if *em != nil {
		var m Menu
		err := json.Unmarshal(*em, &m)
		if err == nil {
			bnd, err2 := b64.StdEncoding.DecodeString(m.Brand)
			if err2 == nil {
				m.Brand = string(bnd)
				m.BrandHTML = template.HTML(m.Brand)
				m.StyleCSS = template.CSS(m.Style)
				c.Log.Debug("menu item:  ", m)
				rtn = m
				suc = true
			}
		}
	}
	return suc, &rtn
}

//GetMenuList GetMenuList
func (c *Six910MenuService) GetMenuList() *[]Menu {
	var rtn []Menu
	res := c.MenuStore.ReadAll()
	for r := range *res {
		var m Menu
		err := json.Unmarshal((*res)[r], &m)
		c.Log.Debug("found menu item in list: ", m)
		if err == nil {
			bnd, err2 := b64.StdEncoding.DecodeString(m.Brand)
			if err2 == nil {
				m.Brand = string(bnd)
				m.BrandHTML = template.HTML(m.Brand)
				m.StyleCSS = template.CSS(m.Style)
				rtn = append(rtn, m)
			}
		}
	}
	return &rtn
}

//DeleteMenu DeleteMenu
func (c *Six910MenuService) DeleteMenu(name string) bool {
	var rtn bool
	suc := c.MenuStore.Delete(name)
	if suc {
		rtn = true
	}
	return rtn
}
