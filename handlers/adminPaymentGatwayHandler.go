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

//PgwPage PgwPage
type PgwPage struct {
	Error         string
	PaymentGatway *sdbi.PaymentGateway
}

//StoreAdminAddPaymentGatewayPage StoreAdminAddPaymentGatewayPage
func (h *Six910Handler) StoreAdminAddPaymentGatewayPage(w http.ResponseWriter, r *http.Request) {
	apgs, suc := h.getSession(r)
	h.Log.Debug("session suc in PWG add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(apgs) {
			aiErr := r.URL.Query().Get("error")
			var apgpg PgwPage
			apgpg.Error = aiErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddPaymentGatwayPage, &apgpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddPaymentGateway StoreAdminAddPaymentGateway
func (h *Six910Handler) StoreAdminAddPaymentGateway(w http.ResponseWriter, r *http.Request) {
	apgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(apgs) {
			apg := h.processPgw(r)
			h.Log.Debug("pgw add", *apg)
			hd := h.getHeader(apgs)
			prres := h.API.AddPaymentGateway(apg, hd)
			h.Log.Debug("pgw add resp", *prres)
			if prres.Success {
				http.Redirect(w, r, adminPaymentGatewayListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddPaymentGatewayViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditPaymentGatewayPage StoreAdminEditPaymentGatewayPage
func (h *Six910Handler) StoreAdminEditPaymentGatewayPage(w http.ResponseWriter, r *http.Request) {
	epgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(epgs) {
			hd := h.getHeader(epgs)
			epgpErr := r.URL.Query().Get("error")
			epgvars := mux.Vars(r)
			idstr := epgvars["id"]
			pgID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("pgw id in edit", pgID)
			var dgp PgwPage
			dgp.Error = epgpErr
			dgp.PaymentGatway = h.API.GetPaymentGateway(pgID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditPaymentGatwayPage, &dgp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditPaymentGateway StoreAdminEditPaymentGateway
func (h *Six910Handler) StoreAdminEditPaymentGateway(w http.ResponseWriter, r *http.Request) {
	epgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(epgs) {
			epg := h.processPgw(r)
			h.Log.Debug("pgw update", *epg)
			hd := h.getHeader(epgs)
			res := h.API.UpdatePaymentGateway(epg, hd)
			h.Log.Debug("Pgw update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminPaymentGatewayListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditPaymentGatewayViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewPaymentGatewayList StoreAdminViewPaymentGatewayList
func (h *Six910Handler) StoreAdminViewPaymentGatewayList(w http.ResponseWriter, r *http.Request) {
	gpgls, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gpgls) {
			hd := h.getHeader(gpgls)
			pgl := h.API.GetPaymentGateways(hd)
			h.Log.Debug("pgw  in list", pgl)
			h.AdminTemplates.ExecuteTemplate(w, adminPaymentGatwayListPage, &pgl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeletePaymentGateway StoreAdminDeletePaymentGateway
func (h *Six910Handler) StoreAdminDeletePaymentGateway(w http.ResponseWriter, r *http.Request) {
	dpgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dpgs) {
			hd := h.getHeader(dpgs)
			dpgvars := mux.Vars(r)
			idstrd := dpgvars["id"]
			idddpg, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeletePaymentGateway(idddpg, hd)
			h.Log.Debug("pgw delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminPaymentGatewayListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminPaymentGatewayListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processPgw(r *http.Request) *sdbi.PaymentGateway {
	var p sdbi.PaymentGateway
	id := r.FormValue("id")
	p.ID, _ = strconv.ParseInt(id, 10, 64)
	p.CheckoutURL = r.FormValue("checkoutUrl")
	p.ClientID = r.FormValue("clientId")
	p.ClientKey = r.FormValue("clientKey")
	p.LogoURL = r.FormValue("logoUrl")
	p.PostOrderURL = r.FormValue("postOrderUrl")
	storePID := r.FormValue("storePluginId")
	p.StorePluginsID, _ = strconv.ParseInt(storePID, 10, 64)

	return &p
}
