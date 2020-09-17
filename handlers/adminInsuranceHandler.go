package handlers

import (
	"net/http"
	"strconv"

	sdbi "github.com/Ulbora/six910-database-interface"
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

//InsPage InsPage
type InsPage struct {
	Error     string
	Insurance *sdbi.Insurance
}

//StoreAdminAddInsurancePage StoreAdminAddInsurancePage
func (h *Six910Handler) StoreAdminAddInsurancePage(w http.ResponseWriter, r *http.Request) {
	ads, suc := h.getSession(r)
	h.Log.Debug("session suc in insc add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ads) {
			aiErr := r.URL.Query().Get("error")
			var aipg InsPage
			aipg.Error = aiErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddInsurancePage, &aipg)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddInsurance StoreAdminAddInsurance
func (h *Six910Handler) StoreAdminAddInsurance(w http.ResponseWriter, r *http.Request) {
	adds, suc := h.getSession(r)
	h.Log.Debug("session suc in dist add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adds) {
			ai := h.processInsurance(r)
			h.Log.Debug("Ins add", *ai)
			hd := h.getHeader(adds)
			prres := h.API.AddInsurance(ai, hd)
			h.Log.Debug("Ins add resp", *prres)
			if prres.Success {
				http.Redirect(w, r, adminInsuranceList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminInsuranceListFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditInsurancePage StoreAdminEditInsurancePage
func (h *Six910Handler) StoreAdminEditInsurancePage(w http.ResponseWriter, r *http.Request) {
	eis, suc := h.getSession(r)
	h.Log.Debug("session suc in ins edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(eis) {
			hd := h.getHeader(eis)
			eipErr := r.URL.Query().Get("error")
			eivars := mux.Vars(r)
			idstr := eivars["id"]
			iID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("ins id in edit", iID)
			var dgp InsPage
			dgp.Error = eipErr
			dgp.Insurance = h.API.GetInsurance(iID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditInsurancePage, &dgp)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditInsurance StoreAdminEditInsurance
func (h *Six910Handler) StoreAdminEditInsurance(w http.ResponseWriter, r *http.Request) {
	eiis, suc := h.getSession(r)
	h.Log.Debug("session suc in ins edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(eiis) {
			eii := h.processInsurance(r)
			h.Log.Debug("ins update", *eii)
			hd := h.getHeader(eiis)
			res := h.API.UpdateInsurance(eii, hd)
			h.Log.Debug("Ins update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminInsuranceList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminInsuranceListFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminViewInsuranceList StoreAdminViewInsuranceList
func (h *Six910Handler) StoreAdminViewInsuranceList(w http.ResponseWriter, r *http.Request) {
	gils, suc := h.getSession(r)
	h.Log.Debug("session suc in ins view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gils) {
			hd := h.getHeader(gils)
			isl := h.API.GetInsuranceList(hd)
			h.Log.Debug("Ins  in list", isl)
			h.AdminTemplates.ExecuteTemplate(w, adminInsuranceListPage, &isl)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminDeleteInsurance StoreAdminDeleteInsurance
func (h *Six910Handler) StoreAdminDeleteInsurance(w http.ResponseWriter, r *http.Request) {
	dis, suc := h.getSession(r)
	h.Log.Debug("session suc in ins list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dis) {
			hd := h.getHeader(dis)
			divars := mux.Vars(r)
			idstrd := divars["id"]
			idddi, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeleteInsurance(idddi, hd)
			h.Log.Debug("ins delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminInsuranceList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminInsuranceListFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processInsurance(r *http.Request) *sdbi.Insurance {
	var i sdbi.Insurance
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	cost := r.FormValue("cost")
	i.Cost, _ = strconv.ParseFloat(cost, 64)
	max := r.FormValue("maxOrderAmount")
	i.MaxOrderAmount, _ = strconv.ParseFloat(max, 64)
	min := r.FormValue("minOrderAmount")
	i.MinOrderAmount, _ = strconv.ParseFloat(min, 64)
	storeID := r.FormValue("storeId")
	i.StoreID, _ = strconv.ParseInt(storeID, 10, 64)

	return &i
}
