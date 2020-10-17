package handlers

import (
	"github.com/Ulbora/Six910-ui/csssrv"
	"net/http"

	"github.com/gorilla/mux"
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

//StoreAdminGetPageCSS StoreAdminGetPageCSS
func (h *Six910Handler) StoreAdminGetPageCSS(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			cvars := mux.Vars(r)
			cname := cvars["name"]
			_, cres := h.CSSService.GetPage(cname)
			h.Log.Debug("css in page css get: ", *cres)
			var ci ContPage
			ci.PageCSS = cres
			h.Log.Debug("css in page: ", ci)

			h.AdminTemplates.ExecuteTemplate(w, adminEditPageCSS, &ci)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUpdatePageCSS StoreAdminUpdatePageCSS
func (h *Six910Handler) StoreAdminUpdatePageCSS(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {

			var csspg csssrv.Page
			var lnk csssrv.Link
			csspg.Link = &lnk

			name := r.FormValue("name")
			h.Log.Debug("css name in update: ", name)

			background := r.FormValue("background")
			h.Log.Debug("css background in update: ", background)

			color := r.FormValue("color")
			h.Log.Debug("css color in update: ", color)

			pageTitle := r.FormValue("pageTitle")
			h.Log.Debug("css pageTitle in update: ", pageTitle)

			linkColor := r.FormValue("linkColor")
			h.Log.Debug("css linkColor in update: ", linkColor)

			linkVisited := r.FormValue("linkVisited")
			h.Log.Debug("css linkVisited in update: ", linkVisited)

			linkHover := r.FormValue("linkHover")
			h.Log.Debug("css linkHover in update: ", linkHover)

			linkActive := r.FormValue("linkActive")
			h.Log.Debug("css linkActive in update: ", linkActive)

			csspg.Name = name
			csspg.Background = background
			csspg.Color = color
			csspg.PageTitle = pageTitle

			csspg.Link.Active = linkActive
			csspg.Link.Color = linkColor
			csspg.Link.Hover = linkHover
			csspg.Link.Visited = linkVisited

			res := h.CSSService.UpdatePage(&csspg)
			h.Log.Debug("update page css res", res)
			if res {
				http.Redirect(w, r, adminIndex, http.StatusFound)
			} else {
				//go back
				http.Redirect(w, r, adminEditPageCSS+"/"+name, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
