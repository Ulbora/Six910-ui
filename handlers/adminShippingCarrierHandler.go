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

//ShipCarPage ShipCarPage
type ShipCarPage struct {
	Error           string
	ShippingCarrier *sdbi.ShippingCarrier
}

//StoreAdminAddCarrierPage StoreAdminAddCarrierPage
func (h *Six910Handler) StoreAdminAddCarrierPage(w http.ResponseWriter, r *http.Request) {
	ascs, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping carrier add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ascs) {
			ascErr := r.URL.Query().Get("error")
			var ascpg InsPage
			ascpg.Error = ascErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddShippingCarrierPage, &ascpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddCarrier StoreAdminAddCarrier
func (h *Six910Handler) StoreAdminAddCarrier(w http.ResponseWriter, r *http.Request) {
	adscs, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping carrier add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adscs) {
			asc := h.processShippingCarrier(r)
			h.Log.Debug("shipping carrier add", *asc)
			hd := h.getHeader(adscs)
			scres := h.API.AddShippingCarrier(asc, hd)
			h.Log.Debug("shipping carrier add resp", *scres)
			if scres.Success {
				http.Redirect(w, r, adminAddShippingCarrierView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddShippingCarrierViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditCarrierPage StoreAdminEditCarrierPage
func (h *Six910Handler) StoreAdminEditCarrierPage(w http.ResponseWriter, r *http.Request) {
	escs, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping carrier edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(escs) {
			hd := h.getHeader(escs)
			eipErr := r.URL.Query().Get("error")
			escvars := mux.Vars(r)
			idstr := escvars["id"]
			iID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("shipping carrier id in edit", iID)
			var scgp ShipCarPage
			scgp.Error = eipErr
			scgp.ShippingCarrier = h.API.GetShippingCarrier(iID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditShippingCarrierPage, &scgp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditCarrier StoreAdminEditCarrier
func (h *Six910Handler) StoreAdminEditCarrier(w http.ResponseWriter, r *http.Request) {
	escs, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping carrier edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(escs) {
			esc := h.processShippingCarrier(r)
			h.Log.Debug("shipping carrier update", *esc)
			hd := h.getHeader(escs)
			res := h.API.UpdateShippingCarrier(esc, hd)
			h.Log.Debug("shipping carrier update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminShippingCarrierListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminShippingCarrierListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewCarrierList StoreAdminViewCarrierList
func (h *Six910Handler) StoreAdminViewCarrierList(w http.ResponseWriter, r *http.Request) {
	gscls, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping carrier view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gscls) {
			hd := h.getHeader(gscls)
			scsl := h.API.GetShippingCarrierList(hd)
			h.Log.Debug("shipping carrier  in list", *scsl)
			h.AdminTemplates.ExecuteTemplate(w, adminShippingCarrierListView, &scsl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteCarrier StoreAdminDeleteCarrier
func (h *Six910Handler) StoreAdminDeleteCarrier(w http.ResponseWriter, r *http.Request) {
	dscs, suc := h.getSession(r)
	h.Log.Debug("session suc in shipping carrier list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dscs) {
			hd := h.getHeader(dscs)
			divars := mux.Vars(r)
			idscrd := divars["id"]
			idddsc, _ := strconv.ParseInt(idscrd, 10, 64)
			res := h.API.DeleteShippingCarrier(idddsc, hd)
			h.Log.Debug("shipping carrier delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminShippingCarrierListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminShippingCarrierListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processShippingCarrier(r *http.Request) *sdbi.ShippingCarrier {
	var s sdbi.ShippingCarrier
	id := r.FormValue("id")
	s.ID, _ = strconv.ParseInt(id, 10, 64)
	s.Carrier = r.FormValue("carrier")
	s.Type = r.FormValue("type")
	storeID := r.FormValue("storeId")
	s.StoreID, _ = strconv.ParseInt(storeID, 10, 64)

	return &s
}
