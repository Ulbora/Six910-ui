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

//ShipMethPage ShipMethPage
type ShipMethPage struct {
	Error          string
	ShippingMethod *sdbi.ShippingMethod
}

//StoreAdminAddShippingMethodPage StoreAdminAddShippingMethodPage
func (h *Six910Handler) StoreAdminAddShippingMethodPage(w http.ResponseWriter, r *http.Request) {
	asms, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping method add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(asms) {
			asmErr := r.URL.Query().Get("error")
			var asmpg ShipMethPage
			asmpg.Error = asmErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddShippingMethodPage, &asmpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddShippingMethod StoreAdminAddShippingMethod
func (h *Six910Handler) StoreAdminAddShippingMethod(w http.ResponseWriter, r *http.Request) {
	aasms, suc := h.getSession(r)
	h.Log.Debug("session suc in dist add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(aasms) {
			aasm := h.processShippingMethod(r)
			h.Log.Debug("shipping method add", *aasm)
			hd := h.getHeader(aasms)
			aasmres := h.API.AddShippingMethod(aasm, hd)
			h.Log.Debug("shipping method add resp", *aasmres)
			if aasmres.Success {
				http.Redirect(w, r, adminAddShippingMethodView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddShippingMethodViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditShippingMethodPage StoreAdminEditShippingMethodPage
func (h *Six910Handler) StoreAdminEditShippingMethodPage(w http.ResponseWriter, r *http.Request) {
	esms, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping method edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(esms) {
			hd := h.getHeader(esms)
			eipErr := r.URL.Query().Get("error")
			esmvars := mux.Vars(r)
			idesmstr := esmvars["id"]
			iID, _ := strconv.ParseInt(idesmstr, 10, 64)
			h.Log.Debug("shipping method id in edit", iID)
			var esmgp ShipMethPage
			esmgp.Error = eipErr
			esmgp.ShippingMethod = h.API.GetShippingMethod(iID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditShippingMethodPage, &esmgp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditShippingMethod StoreAdminEditShippingMethod
func (h *Six910Handler) StoreAdminEditShippingMethod(w http.ResponseWriter, r *http.Request) {
	esmms, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping method edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(esmms) {
			esmm := h.processShippingMethod(r)
			h.Log.Debug("shipping method update", *esmm)
			hd := h.getHeader(esmms)
			res := h.API.UpdateShippingMethod(esmm, hd)
			h.Log.Debug("shipping method update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminShippingMethodListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminShippingMethodListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewShippingMethodList StoreAdminViewShippingMethodList
func (h *Six910Handler) StoreAdminViewShippingMethodList(w http.ResponseWriter, r *http.Request) {
	gsmls, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping method view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gsmls) {
			hd := h.getHeader(gsmls)
			smsl := h.API.GetShippingMethodList(hd)
			h.Log.Debug("shipping method  in list", smsl)
			h.AdminTemplates.ExecuteTemplate(w, adminShippingMethodListPage, &smsl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteShippingMethod StoreAdminDeleteShippingMethod
func (h *Six910Handler) StoreAdminDeleteShippingMethod(w http.ResponseWriter, r *http.Request) {
	dsms, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping method list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dsms) {
			hd := h.getHeader(dsms)
			dsmvars := mux.Vars(r)
			idsmstrd := dsmvars["id"]
			idddsm, _ := strconv.ParseInt(idsmstrd, 10, 64)
			res := h.API.DeleteShippingMethod(idddsm, hd)
			h.Log.Debug("shipping method delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminShippingMethodListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminShippingMethodListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processShippingMethod(r *http.Request) *sdbi.ShippingMethod {
	var s sdbi.ShippingMethod
	id := r.FormValue("id")
	s.ID, _ = strconv.ParseInt(id, 10, 64)
	s.Name = r.FormValue("name")
	cost := r.FormValue("cost")
	s.Cost, _ = strconv.ParseFloat(cost, 64)
	maxWeight := r.FormValue("maxWeight")
	s.MaxWeight, _ = strconv.ParseInt(maxWeight, 10, 64)
	handling := r.FormValue("handling")
	s.Handling, _ = strconv.ParseFloat(handling, 64)
	min := r.FormValue("minOrderAmount")
	s.MinOrderAmount, _ = strconv.ParseFloat(min, 64)
	max := r.FormValue("maxOrderAmount")
	s.MaxOrderAmount, _ = strconv.ParseFloat(max, 64)
	storeID := r.FormValue("storeId")
	s.StoreID, _ = strconv.ParseInt(storeID, 10, 64)
	regionID := r.FormValue("regionId")
	s.RegionID, _ = strconv.ParseInt(regionID, 10, 64)
	shippingCarrierID := r.FormValue("shippingCarrierId")
	s.ShippingCarrierID, _ = strconv.ParseInt(shippingCarrierID, 10, 64)
	insuranceID := r.FormValue("insuranceId")
	s.InsuranceID, _ = strconv.ParseInt(insuranceID, 10, 64)

	return &s
}
