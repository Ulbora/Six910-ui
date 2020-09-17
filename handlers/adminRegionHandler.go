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

//RegionPage RegionPage
type RegionPage struct {
	Error  string
	Region *sdbi.Region
}

//StoreAdminAddRegionPage StoreAdminAddRegionPage
func (h *Six910Handler) StoreAdminAddRegionPage(w http.ResponseWriter, r *http.Request) {
	asrs, suc := h.getSession(r)
	h.Log.Debug("session suc in region add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(asrs) {
			asrErr := r.URL.Query().Get("error")
			var asrpg RegionPage
			asrpg.Error = asrErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddRegionPage, &asrpg)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddRegion StoreAdminAddRegion
func (h *Six910Handler) StoreAdminAddRegion(w http.ResponseWriter, r *http.Request) {
	adsrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Region add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adsrs) {
			asr := h.processRegion(r)
			h.Log.Debug("Region add", *asr)
			hd := h.getHeader(adsrs)
			srres := h.API.AddRegion(asr, hd)
			h.Log.Debug("Region add resp", *srres)
			if srres.Success {
				http.Redirect(w, r, adminRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditRegionPage StoreAdminEditRegionPage
func (h *Six910Handler) StoreAdminEditRegionPage(w http.ResponseWriter, r *http.Request) {
	esrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Region edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(esrs) {
			hd := h.getHeader(esrs)
			eipErr := r.URL.Query().Get("error")
			esrvars := mux.Vars(r)
			idsrstr := esrvars["id"]
			iID, _ := strconv.ParseInt(idsrstr, 10, 64)
			h.Log.Debug("Region id in edit", iID)
			var srp RegionPage
			srp.Error = eipErr
			srp.Region = h.API.GetRegion(iID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditRegionPage, &srp)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditRegion StoreAdminEditRegion
func (h *Six910Handler) StoreAdminEditRegion(w http.ResponseWriter, r *http.Request) {
	esrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Region edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(esrs) {
			esr := h.processRegion(r)
			h.Log.Debug("Region update", *esr)
			hd := h.getHeader(esrs)
			res := h.API.UpdateRegion(esr, hd)
			h.Log.Debug("Region update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminViewRegionList StoreAdminViewRegionList
func (h *Six910Handler) StoreAdminViewRegionList(w http.ResponseWriter, r *http.Request) {
	gsrls, suc := h.getSession(r)
	h.Log.Debug("session suc in Region view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gsrls) {
			hd := h.getHeader(gsrls)
			srsl := h.API.GetRegionList(hd)
			h.Log.Debug("Region  in list", srsl)
			h.AdminTemplates.ExecuteTemplate(w, adminRegionListPage, &srsl)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminDeleteRegion StoreAdminDeleteRegion
func (h *Six910Handler) StoreAdminDeleteRegion(w http.ResponseWriter, r *http.Request) {
	dsrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Region list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dsrs) {
			hd := h.getHeader(dsrs)
			dsrvars := mux.Vars(r)
			idsrstrd := dsrvars["id"]
			idddsr, _ := strconv.ParseInt(idsrstrd, 10, 64)
			res := h.API.DeleteRegion(idddsr, hd)
			h.Log.Debug("Region delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processRegion(r *http.Request) *sdbi.Region {
	var i sdbi.Region
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	i.RegionCode = r.FormValue("regionCode")
	i.Name = r.FormValue("name")
	storeID := r.FormValue("storeId")
	i.StoreID, _ = strconv.ParseInt(storeID, 10, 64)

	return &i
}
