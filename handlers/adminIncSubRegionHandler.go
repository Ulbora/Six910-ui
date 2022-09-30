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

//IncSubRegionPage IncSubRegionPage
type IncSubRegionPage struct {
	Error              string
	InccludedSubRegion *sdbi.IncludedSubRegion
	Region             *sdbi.Region
	SubRegion          *sdbi.SubRegion
}

//StoreAdminAddIncludedSubRegionPage StoreAdminAddIncludedSubRegionPage
func (h *Six910Handler) StoreAdminAddIncludedSubRegionPage(w http.ResponseWriter, r *http.Request) {
	ainssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Inc sub region add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ainssrs) {
			inssrvars := mux.Vars(r)
			ridinrstr := inssrvars["regionId"]
			sridinrstr := inssrvars["subRegionId"]
			inriID, _ := strconv.ParseInt(ridinrstr, 10, 64)
			insriID, _ := strconv.ParseInt(sridinrstr, 10, 64)
			h.Log.Debug("Region id in edit", inriID)
			h.Log.Debug("Sub Region id in edit", insriID)
			aesrErr := r.URL.Query().Get("error")
			var ainssrpg ExSubRegionPage
			ainssrpg.Error = aesrErr

			hd := h.getHeader(ainssrs)
			var wg sync.WaitGroup

			wg.Add(1)
			go func(regionID int64, header *six910api.Headers) {
				defer wg.Done()
				ainssrpg.Region = h.API.GetRegion(regionID, header.DeepCopy())
			}(inriID, hd)

			wg.Add(1)
			go func(subRegionID int64, header *six910api.Headers) {
				defer wg.Done()
				ainssrpg.SubRegion = h.API.GetSubRegion(subRegionID, header.DeepCopy())
			}(insriID, hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminAddIncSubRegionPage, &ainssrpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddIncludedSubRegion StoreAdminAddIncludedSubRegion
func (h *Six910Handler) StoreAdminAddIncludedSubRegion(w http.ResponseWriter, r *http.Request) {
	adinssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Inc Sub Region add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adinssrs) {
			ainssr := h.processIncSubRegion(r)
			h.Log.Debug("Inc Sub Region add", *ainssr)
			hd := h.getHeader(adinssrs)
			insrres := h.API.AddIncludedSubRegion(ainssr, hd)
			h.Log.Debug("Inc Sub Region add resp", *insrres)
			if insrres.Success {
				http.Redirect(w, r, adminIncSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminIncSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewIncludedSubRegionList StoreAdminViewIncludedSubRegionList
func (h *Six910Handler) StoreAdminViewIncludedSubRegionList(w http.ResponseWriter, r *http.Request) {
	ginssrls, suc := h.getSession(r)
	h.Log.Debug("session suc in In Sub Region view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ginssrls) {
			rinsssrvars := mux.Vars(r)
			rinidssrstr := rinsssrvars["regionId"]
			insriID, _ := strconv.ParseInt(rinidssrstr, 10, 64)
			h.Log.Debug(" In Sub Region id in edit", insriID)
			hd := h.getHeader(ginssrls)
			inssrsl := h.API.GetIncludedSubRegionList(insriID, hd)
			h.Log.Debug("In Sub Region  in list", *inssrsl)
			h.AdminTemplates.ExecuteTemplate(w, adminIncSubRegionListPage, &inssrsl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteIncludedSubRegion StoreAdminDeleteIncludedSubRegion
func (h *Six910Handler) StoreAdminDeleteIncludedSubRegion(w http.ResponseWriter, r *http.Request) {
	dinssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Inc Sub Region list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dinssrs) {
			hd := h.getHeader(dinssrs)
			dssrvars := mux.Vars(r)
			idinssrstrd := dssrvars["id"]
			idddessr, _ := strconv.ParseInt(idinssrstrd, 10, 64)
			ridinssrstrd := dssrvars["regionId"]
			ridddessr, _ := strconv.ParseInt(ridinssrstrd, 10, 64)
			res := h.API.DeleteIncludedSubRegion(idddessr, ridddessr, hd)
			h.Log.Debug("Inc Sub Region delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminIncSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminIncSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processIncSubRegion(r *http.Request) *sdbi.IncludedSubRegion {
	var in sdbi.IncludedSubRegion
	id := r.FormValue("id")
	in.ID, _ = strconv.ParseInt(id, 10, 64)
	regionID := r.FormValue("regionId")
	in.RegionID, _ = strconv.ParseInt(regionID, 10, 64)
	shippingMethodID := r.FormValue("shippingMethodId")
	in.ShippingMethodID, _ = strconv.ParseInt(shippingMethodID, 10, 64)
	subRegionID := r.FormValue("subRegionId")
	in.SubRegionID, _ = strconv.ParseInt(subRegionID, 10, 64)

	return &in
}
