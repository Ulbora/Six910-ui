package handlers

import (
	"net/http"
	"strconv"
	"time"

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

//SpiPage SpiPage
type SpiPage struct {
	Error        string
	StorePlugins *sdbi.StorePlugins
}

//StoreAdminAddStorePluginPage StoreAdminAddStorePluginPage
func (h *Six910Handler) StoreAdminAddStorePluginPage(w http.ResponseWriter, r *http.Request) {
	aspis, suc := h.getSession(r)
	h.Log.Debug("session suc in store plugins add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(aspis) {
			aspiErr := r.URL.Query().Get("error")
			var aspipg SpiPage
			aspipg.Error = aspiErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddStorePluginPage, &aspipg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddStorePlugin StoreAdminAddStorePlugin
func (h *Six910Handler) StoreAdminAddStorePlugin(w http.ResponseWriter, r *http.Request) {
	addspis, suc := h.getSession(r)
	h.Log.Debug("session suc in Store Plugin add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(addspis) {
			aspi := h.processStorePlugin(r)
			h.Log.Debug("Store Plugin add", *aspi)
			hd := h.getHeader(addspis)
			spirres := h.API.AddStorePlugin(aspi, hd)
			h.Log.Debug("Store Plugin add resp", *spirres)
			if spirres.Success {
				http.Redirect(w, r, adminStorePluginListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminStorePluginListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditStorePluginPage StoreAdminEditStorePluginPage
func (h *Six910Handler) StoreAdminEditStorePluginPage(w http.ResponseWriter, r *http.Request) {
	espis, suc := h.getSession(r)
	h.Log.Debug("session suc in Store Plugin edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(espis) {
			hd := h.getHeader(espis)
			espipErr := r.URL.Query().Get("error")
			espivars := mux.Vars(r)
			espiidstr := espivars["id"]
			iID, _ := strconv.ParseInt(espiidstr, 10, 64)
			h.Log.Debug("Store Plugin id in edit", iID)
			var espigp SpiPage
			espigp.Error = espipErr
			espigp.StorePlugins = h.API.GetStorePlugin(iID, hd)
			h.AdminTemplates.ExecuteTemplate(w, adminEditStorePluginPage, &espigp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditStorePlugin StoreAdminEditStorePlugin
func (h *Six910Handler) StoreAdminEditStorePlugin(w http.ResponseWriter, r *http.Request) {
	espiis, suc := h.getSession(r)
	h.Log.Debug("session suc in store plugin edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(espiis) {
			espii := h.processStorePlugin(r)
			h.Log.Debug("store plugin update", *espii)
			hd := h.getHeader(espiis)
			res := h.API.UpdateStorePlugin(espii, hd)
			h.Log.Debug("store plugin update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminStorePluginListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminStorePluginListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewStorePluginList StoreAdminViewStorePluginList
func (h *Six910Handler) StoreAdminViewStorePluginList(w http.ResponseWriter, r *http.Request) {
	gspils, suc := h.getSession(r)
	h.Log.Debug("session suc in store plugin view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gspils) {
			hd := h.getHeader(gspils)
			spisl := h.API.GetStorePluginList(hd)
			h.Log.Debug("store plugin  in list", spisl)
			h.AdminTemplates.ExecuteTemplate(w, adminStorePluginListPage, &spisl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteStorePlugin StoreAdminDeleteStorePlugin
func (h *Six910Handler) StoreAdminDeleteStorePlugin(w http.ResponseWriter, r *http.Request) {
	dspis, suc := h.getSession(r)
	h.Log.Debug("session suc in store plugin list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dspis) {
			hd := h.getHeader(dspis)
			dspivars := mux.Vars(r)
			idstrd := dspivars["id"]
			idddspi, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeleteStorePlugin(idddspi, hd)
			h.Log.Debug("store plugin delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminStorePluginListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminStorePluginListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processStorePlugin(r *http.Request) *sdbi.StorePlugins {
	var i sdbi.StorePlugins
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	i.PluginName = r.FormValue("pluginName")
	i.Category = r.FormValue("category")
	active := r.FormValue("active")
	i.Active, _ = strconv.ParseBool(active)
	oauthClientID := r.FormValue("oauthClientId")
	i.OauthClientID, _ = strconv.ParseInt(oauthClientID, 10, 64)
	i.OauthSecret = r.FormValue("oauthSecret")
	i.OauthRedirectURL = r.FormValue("oauthRedirectUrl")
	i.ActivateURL = r.FormValue("activateUrl")
	i.APIKey = r.FormValue("apiKey")
	rekeyTryCount := r.FormValue("rekeyTryCount")
	i.RekeyTryCount, _ = strconv.ParseInt(rekeyTryCount, 10, 64)
	rekeyDate := r.FormValue("rekeyDate")
	i.RekeyDate, _ = time.Parse(timeFormat, rekeyDate)
	i.IframeURL = r.FormValue("iframeUrl")
	i.MenuTitle = r.FormValue("menuTitle")
	i.MenuIconURL = r.FormValue("menuIconUrl")
	isPgw := r.FormValue("isPgw")
	i.IsPGW, _ = strconv.ParseBool(isPgw)
	pluginID := r.FormValue("pluginId")
	i.PluginID, _ = strconv.ParseInt(pluginID, 10, 64)
	storeID := r.FormValue("storeId")
	i.StoreID, _ = strconv.ParseInt(storeID, 10, 64)

	return &i
}
