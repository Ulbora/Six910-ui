package handlers

import (
	"net/http"
	"strconv"
	"sync"

	six910api "github.com/Ulbora/Six910API-Go"
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

//ExSubRegionPage ExSubRegionPage
type ExSubRegionPage struct {
	Error             string
	ExcludedSubRegion *sdbi.ExcludedSubRegion
	Region            *sdbi.Region
	SubRegion         *sdbi.SubRegion
}

//StoreAdminAddExcludedSubRegionPage StoreAdminAddExcludedSubRegionPage
func (h *Six910Handler) StoreAdminAddExcludedSubRegionPage(w http.ResponseWriter, r *http.Request) {
	aessrs, suc := h.getSession(r)
	h.Log.Debug("session suc in ex sub region add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(aessrs) {
			essrvars := mux.Vars(r)
			ridssrstr := essrvars["regionId"]
			sridssrstr := essrvars["subRegionId"]
			riID, _ := strconv.ParseInt(ridssrstr, 10, 64)
			sriID, _ := strconv.ParseInt(sridssrstr, 10, 64)
			h.Log.Debug("Region id in edit", riID)
			h.Log.Debug("Sub Region id in edit", sriID)
			aesrErr := r.URL.Query().Get("error")
			var assrpg ExSubRegionPage
			assrpg.Error = aesrErr

			hd := h.getHeader(aessrs)
			var wg sync.WaitGroup

			wg.Add(1)
			go func(regionID int64, header *six910api.Headers) {
				defer wg.Done()
				assrpg.Region = h.API.GetRegion(regionID, header.DeepCopy())
			}(riID, hd)

			wg.Add(1)
			go func(subRegionID int64, header *six910api.Headers) {
				defer wg.Done()
				assrpg.SubRegion = h.API.GetSubRegion(subRegionID, header.DeepCopy())
			}(sriID, hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminAddExSubRegionPage, &assrpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddExcludedSubRegion StoreAdminAddExcludedSubRegion
func (h *Six910Handler) StoreAdminAddExcludedSubRegion(w http.ResponseWriter, r *http.Request) {
	adessrs, suc := h.getSession(r)
	h.Log.Debug("session suc in ex Sub Region add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adessrs) {
			aessr := h.processExSubRegion(r)
			h.Log.Debug("Ex Sub Region add", *aessr)
			hd := h.getHeader(adessrs)
			esrres := h.API.AddExcludedSubRegion(aessr, hd)
			h.Log.Debug("Ex Sub Region add resp", *esrres)
			if esrres.Success {
				http.Redirect(w, r, adminExSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminExSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewExcludedSubRegionList StoreAdminViewExcludedSubRegionList
func (h *Six910Handler) StoreAdminViewExcludedSubRegionList(w http.ResponseWriter, r *http.Request) {
	gessrls, suc := h.getSession(r)
	h.Log.Debug("session suc in Ex Sub Region view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gessrls) {
			resssrvars := mux.Vars(r)
			reidssrstr := resssrvars["regionId"]
			esriID, _ := strconv.ParseInt(reidssrstr, 10, 64)
			h.Log.Debug(" Ex Sub Region id in edit", esriID)
			hd := h.getHeader(gessrls)
			essrsl := h.API.GetExcludedSubRegionList(esriID, hd)
			h.Log.Debug("Ex Sub Region  in list", *essrsl)
			h.AdminTemplates.ExecuteTemplate(w, adminExSubRegionListPage, &essrsl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteExcludedSubRegion StoreAdminDeleteExcludedSubRegion
func (h *Six910Handler) StoreAdminDeleteExcludedSubRegion(w http.ResponseWriter, r *http.Request) {
	dessrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Ex Sub Region list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dessrs) {
			hd := h.getHeader(dessrs)
			dssrvars := mux.Vars(r)
			idessrstrd := dssrvars["id"]
			idddessr, _ := strconv.ParseInt(idessrstrd, 10, 64)
			ridessrstrd := dssrvars["regionId"]
			ridddessr, _ := strconv.ParseInt(ridessrstrd, 10, 64)
			res := h.API.DeleteExcludedSubRegion(idddessr, ridddessr, hd)
			h.Log.Debug("Ex Sub Region delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminExSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminExSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processExSubRegion(r *http.Request) *sdbi.ExcludedSubRegion {
	var i sdbi.ExcludedSubRegion
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	regionID := r.FormValue("regionId")
	i.RegionID, _ = strconv.ParseInt(regionID, 10, 64)
	shippingMethodID := r.FormValue("shippingMethodId")
	i.ShippingMethodID, _ = strconv.ParseInt(shippingMethodID, 10, 64)
	subRegionID := r.FormValue("subRegionId")
	i.SubRegionID, _ = strconv.ParseInt(subRegionID, 10, 64)

	return &i
}
