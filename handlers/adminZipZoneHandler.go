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

//ZipZonePage ZipZonePage
type ZipZonePage struct {
	Error           string
	ExcludedZipZone *[]sdbi.ZoneZip
	IncludedZipZone *[]sdbi.ZoneZip
	Region          *sdbi.Region
	SubRegion       *sdbi.SubRegion
}

//StoreAdminAddZipZonePage StoreAdminAddZipZonePage
func (h *Six910Handler) StoreAdminAddZipZonePage(w http.ResponseWriter, r *http.Request) {
	azzssrs, suc := h.getSession(r)
	h.Log.Debug("session suc in zip zone add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(azzssrs) {
			ezzrvars := mux.Vars(r)
			ridzzrstr := ezzrvars["regionId"]
			sridzzrstr := ezzrvars["subRegionId"]
			zzriID, _ := strconv.ParseInt(ridzzrstr, 10, 64)
			zzsriID, _ := strconv.ParseInt(sridzzrstr, 10, 64)
			h.Log.Debug("Region id in add", zzriID)
			h.Log.Debug("Sub Region id in add", zzsriID)
			azzrErr := r.URL.Query().Get("error")
			var azzrpg ZipZonePage
			azzrpg.Error = azzrErr

			hd := h.getHeader(azzssrs)
			var wg sync.WaitGroup

			wg.Add(1)
			go func(regionID int64, header *six910api.Headers) {
				defer wg.Done()
				azzrpg.Region = h.API.GetRegion(regionID, header.DeepCopy())
				h.Log.Debug("Done with Region id in add", regionID)
			}(zzriID, hd)

			wg.Add(1)
			go func(subRegionID int64, header *six910api.Headers) {
				defer wg.Done()
				azzrpg.SubRegion = h.API.GetSubRegion(subRegionID, header.DeepCopy())
				h.Log.Debug("Done with Sub Region id in add", subRegionID)
			}(zzsriID, hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminAddZipZonePage, &azzrpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddZipZone StoreAdminAddZipZone
func (h *Six910Handler) StoreAdminAddZipZone(w http.ResponseWriter, r *http.Request) {
	adezzrs, suc := h.getSession(r)
	h.Log.Debug("session suc in ZipZone add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adezzrs) {
			aezzr, regionID := h.processZoneZip(r)
			h.Log.Debug("regionID add", regionID)
			h.Log.Debug("ZipZone add", *aezzr)
			hd := h.getHeader(adezzrs)
			ezzrres := h.API.AddZoneZip(aezzr, hd)
			h.Log.Debug("ZipZone add resp", *ezzrres)
			if ezzrres.Success {
				http.Redirect(w, r, adminSubRegionListView+"/"+regionID, http.StatusFound)
			} else {
				http.Redirect(w, r, adminSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewZipZoneList StoreAdminViewZipZoneList
func (h *Six910Handler) StoreAdminViewZipZoneList(w http.ResponseWriter, r *http.Request) {
	gezzrls, suc := h.getSession(r)
	h.Log.Debug("session suc in Zone Zip view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gezzrls) {
			rezzrvars := mux.Vars(r)

			incIdstr := rezzrvars["incId"]
			incID, _ := strconv.ParseInt(incIdstr, 10, 64)
			h.Log.Debug(" Zone Zip incID", incID)

			excIdstr := rezzrvars["excId"]
			excID, _ := strconv.ParseInt(excIdstr, 10, 64)
			h.Log.Debug(" Zone Zip  exID", excID)

			var zzp ZipZonePage
			hd := h.getHeader(gezzrls)
			if incID != 0 {
				zzp.IncludedZipZone = h.API.GetZoneZipListByInclusion(incID, hd)
				h.Log.Debug(" Zone Zip  zzp.IncludedZipZone", *zzp.IncludedZipZone)
			} else if excID != 0 {
				zzp.ExcludedZipZone = h.API.GetZoneZipListByExclusion(excID, hd)
				h.Log.Debug(" Zone Zip  zzp.ExcludedZipZone", *zzp.ExcludedZipZone)
			}
			h.Log.Debug("Zone Zip", zzp)
			h.AdminTemplates.ExecuteTemplate(w, adminZipZoneListPage, &zzp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteZipZone StoreAdminDeleteZipZone
func (h *Six910Handler) StoreAdminDeleteZipZone(w http.ResponseWriter, r *http.Request) {
	dzzrs, suc := h.getSession(r)
	h.Log.Debug("session suc in Zone Zip delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dzzrs) {
			hd := h.getHeader(dzzrs)
			dzzvars := mux.Vars(r)

			idzzstd := dzzvars["id"]
			idzz, _ := strconv.ParseInt(idzzstd, 10, 64)
			h.Log.Debug(" Zone Zip delete id ", idzz)

			incidzzstd := dzzvars["incId"]
			inczzid, _ := strconv.ParseInt(incidzzstd, 10, 64)
			h.Log.Debug(" Zone Zip delete incId ", inczzid)

			excidzzstd := dzzvars["excId"]
			excidzz, _ := strconv.ParseInt(excidzzstd, 10, 64)
			h.Log.Debug(" Zone Zip delete excId ", excidzz)

			ridstr := dzzvars["regionId"]
			h.Log.Debug(" Zone Zip delete regionId ", ridstr)

			res := h.API.DeleteZoneZip(idzz, inczzid, excidzz, hd)
			h.Log.Debug(" Zone Zip delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminSubRegionListView+"/"+ridstr, http.StatusFound)
			} else {
				http.Redirect(w, r, adminExSubRegionListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processZoneZip(r *http.Request) (*sdbi.ZoneZip, string) {
	var z sdbi.ZoneZip
	regionIDStr := r.FormValue("regionId")
	//regionID, _ := strconv.ParseInt(regionIDStr, 10, 64)
	zipCode := r.FormValue("zipCode")
	includedSubRegionIDStr := r.FormValue("includedSubRegionId")
	includedSubRegionID, _ := strconv.ParseInt(includedSubRegionIDStr, 10, 64)

	excludedSubRegionIDStr := r.FormValue("excludedSubRegionId")
	excludedSubRegionID, _ := strconv.ParseInt(excludedSubRegionIDStr, 10, 64)

	z.ZipCode = zipCode
	z.IncludedSubRegionID = includedSubRegionID
	z.ExcludedSubRegionID = excludedSubRegionID

	return &z, regionIDStr
}
