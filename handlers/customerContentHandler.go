package handlers

import (
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	"github.com/gorilla/mux"
	"net/http"
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

//ViewContent ViewContent
func (h *Six910Handler) ViewContent(w http.ResponseWriter, r *http.Request) {
	ccis, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		vars := mux.Vars(r)
		name := vars["name"]
		hd := h.getHeader(ccis)
		//cipage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(ccis, ml, hd)

		var ccipage CustomerPage
		ccipage.MenuList = ml
		h.Log.Debug("MenuList", *ccipage.MenuList)
		ccisuc, ccont := h.ContentService.GetContent(name)
		if ccisuc {
			ccipage.Content = ccont
		} else {
			var ct conts.Content
			ccipage.Content = &ct
		}

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		h.Log.Debug("PageBody: ", *csspg)
		ccipage.PageBody = csspg
		h.Templates.ExecuteTemplate(w, customerContentPage, &ccipage)

	}
}
