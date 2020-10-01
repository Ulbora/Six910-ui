package handlers

import (
	"net/http"
	"strconv"

	sdbi "github.com/Ulbora/six910-database-interface"
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

//StoreAdminEditStorePage StoreAdminEditStorePage
func (h *Six910Handler) StoreAdminEditStorePage(w http.ResponseWriter, r *http.Request) {
	stacs, suc := h.getSession(r)
	h.Log.Debug("session suc in store view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(stacs) {
			hd := h.getHeader(stacs)

			h.Log.Debug("h.StoreName:", h.StoreName)
			h.Log.Debug("h.LocalDomain:", h.LocalDomain)
			str := h.API.GetStore(h.StoreName, h.LocalDomain, hd)

			h.AdminTemplates.ExecuteTemplate(w, adminEditStorePage, &str)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditStore StoreAdminEditStore
func (h *Six910Handler) StoreAdminEditStore(w http.ResponseWriter, r *http.Request) {
	steeiis, suc := h.getSession(r)
	h.Log.Debug("session suc in store edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(steeiis) {
			estr := h.processStore(r)
			h.Log.Debug("store update", *estr)
			hd := h.getHeader(steeiis)
			res := h.API.UpdateStore(estr, hd)
			h.Log.Debug("Stroe update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processStore(r *http.Request) *sdbi.Store {
	var s sdbi.Store
	id := r.FormValue("id")
	s.ID, _ = strconv.ParseInt(id, 10, 64)
	s.Company = r.FormValue("company")
	s.FirstName = r.FormValue("firstName")
	s.LastName = r.FormValue("lastName")
	s.LocalDomain = r.FormValue("localDomain")
	s.RemoteDomain = r.FormValue("remoteDomain")
	oauthClientID := r.FormValue("oauthClientID")
	s.OauthClientID, _ = strconv.ParseInt(oauthClientID, 10, 64)
	s.OauthSecret = r.FormValue("oauthSecret")
	s.Email = r.FormValue("email")
	s.City = r.FormValue("city")
	s.State = r.FormValue("state")
	s.Zip = r.FormValue("zip")
	s.StoreName = r.FormValue("storeName")
	s.StoreSlogan = r.FormValue("storeSlogan")
	s.Logo = r.FormValue("logo")
	s.Currency = r.FormValue("currency")
	enabled := r.FormValue("enabled")
	s.Enabled, _ = strconv.ParseBool(enabled)

	return &s
}
