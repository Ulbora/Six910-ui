package csssrv

import (
	"encoding/json"
	"html/template"
)

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

//Page Page
type Page struct {
	Name       string
	Background string
	Color      string
	PageTitle  string
}

//PageCSS PageCSS
type PageCSS struct {
	Name       string
	Background template.CSS
	Color      template.CSS
	PageTitle  template.CSS
}

//GetPageCSS GetPageCSS
func (c *Six910CSSService) GetPageCSS(name string) (bool, *PageCSS) {
	var rtn PageCSS
	var suc bool
	ep := c.CSSStore.Read(name)
	if *ep != nil {
		var p Page
		err := json.Unmarshal(*ep, &p)
		if err == nil {
			var bgnd = "background: " + p.Background + " !important; "
			rtn.Background = template.CSS(bgnd)
			c.Log.Debug("bgnd :  ", bgnd)
			var col = "color: " + p.Color + "; "
			rtn.Color = template.CSS(col)
			var tcol = "color: " + p.PageTitle + "; "
			rtn.PageTitle = template.CSS(tcol)
			suc = true
		}
	}
	return suc, &rtn
}

//GetPage GetPage
func (c *Six910CSSService) GetPage(name string) (bool, *Page) {
	var rtn Page
	var suc bool
	ep := c.CSSStore.Read(name)
	if *ep != nil {
		err := json.Unmarshal(*ep, &rtn)
		if err == nil {
			suc = true
		}
	}
	return suc, &rtn
}

//UpdatePage UpdatePage
func (c *Six910CSSService) UpdatePage(page *Page) bool {
	var rtn bool
	ep := c.CSSStore.Read(page.Name)
	if *ep != nil {
		var p Page
		err := json.Unmarshal(*ep, &p)
		c.Log.Debug("found page in update: ", p)
		if err == nil {
			p.Background = page.Background
			p.Color = page.Color
			p.PageTitle = page.PageTitle
			suc := c.CSSStore.Save(page.Name, p)
			rtn = suc
		}
	}
	return rtn
}
