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

//SubRegionPage SubRegionPage
type SubRegionPage struct {
	Error     string
	Region    *sdbi.Region
	SubRegion *sdbi.SubRegion
}

//StoreAdminAddSubRegionPage StoreAdminAddSubRegionPage
func (h *Six910Handler) StoreAdminAddSubRegionPage(w http.ResponseWriter, r *http.Request) {
	assrs, suc := h.getSession(r)
	h.Log.Debug("session suc in sub region add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(assrs) {
			aessrvars := mux.Vars(r)
			aridssrstr := aessrvars["regionId"]
			riID, _ := strconv.ParseInt(aridssrstr, 10, 64)
			h.Log.Debug("Region id in add", riID)
			asrErr := r.URL.Query().Get("error")
			hd := h.getHeader(assrs)
			var assrpg SubRegionPage
			assrpg.Region = h.API.GetRegion(riID, hd)
			assrpg.Error = asrErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddSubRegionPage, &assrpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddSubRegion StoreAdminAddSubRegion
func (h *Six910Handler) StoreAdminAddSubRegion(w http.ResponseWriter, r *http.Request) {
	adssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Sub Region add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adssrs) {
			assr := h.processSubRegion(r)
			h.Log.Debug("Sub Region add", *assr)
			hd := h.getHeader(adssrs)
			srres := h.API.AddSubRegion(assr, hd)
			h.Log.Debug("Sub Region add resp", *srres)
			if srres.Success {
				http.Redirect(w, r, adminSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditSubRegionPage StoreAdminEditSubRegionPage
func (h *Six910Handler) StoreAdminEditSubRegionPage(w http.ResponseWriter, r *http.Request) {
	essrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Sub Region edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(essrs) {
			hd := h.getHeader(essrs)
			essrErr := r.URL.Query().Get("error")
			essrvars := mux.Vars(r)

			idssrstr := essrvars["id"]
			iID, _ := strconv.ParseInt(idssrstr, 10, 64)
			h.Log.Debug("Sub Region id in edit", iID)

			reidssrstr := essrvars["regionId"]
			eriID, _ := strconv.ParseInt(reidssrstr, 10, 64)
			h.Log.Debug("Region id in add", eriID)
			var srp SubRegionPage
			srp.Error = essrErr

			var wg sync.WaitGroup

			wg.Add(1)
			go func(regionID int64, header *six910api.Headers) {
				defer wg.Done()
				srp.Region = h.API.GetRegion(regionID, header)
			}(eriID, hd)

			wg.Add(1)
			go func(subRegionID int64, header *six910api.Headers) {
				defer wg.Done()
				srp.SubRegion = h.API.GetSubRegion(subRegionID, header)
			}(iID, hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminEditSubRegionPage, &srp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditSubRegion StoreAdminEditSubRegion
func (h *Six910Handler) StoreAdminEditSubRegion(w http.ResponseWriter, r *http.Request) {
	esssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Sub Region edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(esssrs) {
			esssr := h.processSubRegion(r)
			h.Log.Debug("Sub Region update", *esssr)
			hd := h.getHeader(esssrs)
			res := h.API.UpdateSubRegion(esssr, hd)
			h.Log.Debug("Sub Region update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditSubRegionViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewSubRegionList StoreAdminViewSubRegionList
func (h *Six910Handler) StoreAdminViewSubRegionList(w http.ResponseWriter, r *http.Request) {
	gssrls, suc := h.getSession(r)
	h.Log.Debug("session suc in Sub Region view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gssrls) {
			esssrvars := mux.Vars(r)
			ridssrstr := esssrvars["regionId"]
			riID, _ := strconv.ParseInt(ridssrstr, 10, 64)
			h.Log.Debug("Region id in edit", riID)
			hd := h.getHeader(gssrls)
			srsl := h.API.GetSubRegionList(riID, hd)
			h.Log.Debug("Sub Region  in list", srsl)
			h.AdminTemplates.ExecuteTemplate(w, adminSubRegionListPage, &srsl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteSubRegion StoreAdminDeleteSubRegion
func (h *Six910Handler) StoreAdminDeleteSubRegion(w http.ResponseWriter, r *http.Request) {
	dssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Sub Region list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dssrs) {
			hd := h.getHeader(dssrs)
			dssrvars := mux.Vars(r)
			idssrstrd := dssrvars["id"]
			idddssr, _ := strconv.ParseInt(idssrstrd, 10, 64)
			res := h.API.DeleteSubRegion(idddssr, hd)
			h.Log.Debug("Sub Region delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminSubRegionListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processSubRegion(r *http.Request) *sdbi.SubRegion {
	var i sdbi.SubRegion
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	i.SubRegionCode = r.FormValue("subRegionCode")
	i.Name = r.FormValue("name")
	regionID := r.FormValue("regionId")
	i.RegionID, _ = strconv.ParseInt(regionID, 10, 64)

	return &i
}
