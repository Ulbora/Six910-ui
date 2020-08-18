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

//DistPage DistPage
type DistPage struct {
	Error       string
	Distributor *sdbi.Distributor
}

//StoreAdminAddDistributorPage StoreAdminAddDistributorPage
func (h *Six910Handler) StoreAdminAddDistributorPage(w http.ResponseWriter, r *http.Request) {
	ads, suc := h.getSession(r)
	h.Log.Debug("session suc in dist add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ads) {
			adErr := r.URL.Query().Get("error")
			var adpg DistPage
			adpg.Error = adErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddDistributorPage, &adpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddDistributor StoreAdminAddDistributor
func (h *Six910Handler) StoreAdminAddDistributor(w http.ResponseWriter, r *http.Request) {
	adds, suc := h.getSession(r)
	h.Log.Debug("session suc in dist add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adds) {
			d := h.processDistributor(r)
			h.Log.Debug("Dist add", *d)
			hd := h.getHeader(adds)
			prres := h.API.AddDistributor(d, hd)
			h.Log.Debug("Dist add resp", *prres)
			if prres.Success {
				http.Redirect(w, r, adminAddDistributorView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddDistributorViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditDistributorPage StoreAdminEditDistributorPage
func (h *Six910Handler) StoreAdminEditDistributorPage(w http.ResponseWriter, r *http.Request) {
	eds, suc := h.getSession(r)
	h.Log.Debug("session suc in dist edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(eds) {
			hd := h.getHeader(eds)
			edpErr := r.URL.Query().Get("error")
			edvars := mux.Vars(r)
			idstr := edvars["id"]
			dID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("dist id in edit", dID)
			var dgp DistPage
			dgp.Error = edpErr
			dgp.Distributor = h.API.GetDistributor(dID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditDistributorPage, &dgp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditDistributor StoreAdminEditDistributor
func (h *Six910Handler) StoreAdminEditDistributor(w http.ResponseWriter, r *http.Request) {
	edds, suc := h.getSession(r)
	h.Log.Debug("session suc in dist edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(edds) {
			edd := h.processDistributor(r)
			h.Log.Debug("Dist update", *edd)
			hd := h.getHeader(edds)
			res := h.API.UpdateDistributor(edd, hd)
			h.Log.Debug("Dist update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminDistributorListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditDistributorViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewDistributorList StoreAdminViewDistributorList
func (h *Six910Handler) StoreAdminViewDistributorList(w http.ResponseWriter, r *http.Request) {
	gdls, suc := h.getSession(r)
	h.Log.Debug("session suc in dist view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gdls) {
			hd := h.getHeader(gdls)
			dsl := h.API.GetDistributorList(hd)
			h.Log.Debug("Dist  in list", dsl)
			h.AdminTemplates.ExecuteTemplate(w, adminDistributorListPage, &dsl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteDistributor StoreAdminDeleteDistributor
func (h *Six910Handler) StoreAdminDeleteDistributor(w http.ResponseWriter, r *http.Request) {
	dds, suc := h.getSession(r)
	h.Log.Debug("session suc in cat list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dds) {
			hd := h.getHeader(dds)
			ddvars := mux.Vars(r)
			idstrd := ddvars["id"]
			idddd, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeleteDistributor(idddd, hd)
			h.Log.Debug("dist delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminDistributorListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminDistributorListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processDistributor(r *http.Request) *sdbi.Distributor {
	var d sdbi.Distributor
	id := r.FormValue("id")
	d.ID, _ = strconv.ParseInt(id, 10, 64)
	d.Company = r.FormValue("company")
	d.ContactName = r.FormValue("contactName")
	d.Phone = r.FormValue("phone")
	storeID := r.FormValue("storeId")
	d.StoreID, _ = strconv.ParseInt(storeID, 10, 64)

	return &d
}
