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

//PluginPage PluginPage
type PluginPage struct {
	Error  string
	Plugin *sdbi.Plugins
}

//StoreAdminAddPluginPage StoreAdminAddPluginPage
func (h *Six910Handler) StoreAdminAddPluginPage(w http.ResponseWriter, r *http.Request) {
	if !h.OAuth2Enabled {
		apls, suc := h.getSession(r)
		h.Log.Debug("session suc in plugin add view", suc)
		if suc {
			if h.isStoreAdminLoggedIn(apls) {
				aplErr := r.URL.Query().Get("error")
				var aplpg PluginPage
				aplpg.Error = aplErr
				h.AdminTemplates.ExecuteTemplate(w, adminAddPluginPage, &aplpg)
			} else {
				http.Redirect(w, r, adminloginPage, http.StatusFound)
			}
		}
	}
}

//StoreAdminAddPlugin StoreAdminAddPlugin
func (h *Six910Handler) StoreAdminAddPlugin(w http.ResponseWriter, r *http.Request) {
	if !h.OAuth2Enabled {
		apiis, suc := h.getSession(r)
		h.Log.Debug("session suc in plugin add", suc)
		if suc {
			if h.isStoreAdminLoggedIn(apiis) {
				apii := h.processPlugin(r)
				h.Log.Debug("Plugin add", *apii)
				hd := h.getHeader(apiis)
				pires := h.API.AddPlugin(apii, hd)
				h.Log.Debug("Plugin add resp", *pires)
				if pires.Success {
					http.Redirect(w, r, adminAddPluginView, http.StatusFound)
				} else {
					http.Redirect(w, r, adminAddPluginViewFail, http.StatusFound)
				}
			} else {
				http.Redirect(w, r, adminloginPage, http.StatusFound)
			}
		}
	}
}

//StoreAdminEditPluginPage StoreAdminEditPluginPage
func (h *Six910Handler) StoreAdminEditPluginPage(w http.ResponseWriter, r *http.Request) {
	if !h.OAuth2Enabled {
		epis, suc := h.getSession(r)
		h.Log.Debug("session suc in plugin edit view", suc)
		if suc {
			if h.isStoreAdminLoggedIn(epis) {
				hd := h.getHeader(epis)
				eipErr := r.URL.Query().Get("error")
				epivars := mux.Vars(r)
				pidstr := epivars["id"]
				iID, _ := strconv.ParseInt(pidstr, 10, 64)
				h.Log.Debug("plugin id in edit", iID)
				var epip PluginPage
				epip.Error = eipErr
				epip.Plugin = h.API.GetPlugin(iID, hd)
				h.AdminTemplates.ExecuteTemplate(w, adminEditPluginPage, &epip)
			} else {
				http.Redirect(w, r, adminloginPage, http.StatusFound)
			}
		}
	}
}

//StoreAdminEditPlugin StoreAdminEditPlugin
func (h *Six910Handler) StoreAdminEditPlugin(w http.ResponseWriter, r *http.Request) {
	if !h.OAuth2Enabled {
		epiis, suc := h.getSession(r)
		h.Log.Debug("session suc in plugin edit", suc)
		if suc {
			if h.isStoreAdminLoggedIn(epiis) {
				epii := h.processPlugin(r)
				h.Log.Debug("Plugin update", *epii)
				hd := h.getHeader(epiis)
				res := h.API.UpdatePlugin(epii, hd)
				h.Log.Debug("Plugin update resp", *res)
				if res.Success {
					http.Redirect(w, r, adminPluginListView, http.StatusFound)
				} else {
					http.Redirect(w, r, adminPluginListViewFail, http.StatusFound)
				}
			} else {
				http.Redirect(w, r, adminloginPage, http.StatusFound)
			}
		}
	}
}

//StoreAdminViewPluginList StoreAdminViewPluginList
func (h *Six910Handler) StoreAdminViewPluginList(w http.ResponseWriter, r *http.Request) {
	gpils, suc := h.getSession(r)
	h.Log.Debug("session suc in plugin view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gpils) {
			hd := h.getHeader(gpils)
			vpivars := mux.Vars(r)
			ststr := vpivars["start"]
			endstr := vpivars["end"]
			vpistart, _ := strconv.ParseInt(ststr, 10, 64)
			vpiend, _ := strconv.ParseInt(endstr, 10, 64)
			pisl := h.API.GetPluginList(vpistart, vpiend, hd)
			h.Log.Debug("Plugin  in list", pisl)
			h.AdminTemplates.ExecuteTemplate(w, adminPluginListPage, &pisl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeletePlugin StoreAdminDeletePlugin
func (h *Six910Handler) StoreAdminDeletePlugin(w http.ResponseWriter, r *http.Request) {
	if !h.OAuth2Enabled {
		dpis, suc := h.getSession(r)
		h.Log.Debug("session suc in plugin list delete", suc)
		if suc {
			if h.isStoreAdminLoggedIn(dpis) {
				hd := h.getHeader(dpis)
				dpivars := mux.Vars(r)
				idstrd := dpivars["id"]
				idddpi, _ := strconv.ParseInt(idstrd, 10, 64)
				res := h.API.DeletePlugin(idddpi, hd)
				h.Log.Debug("plugin delete resp", *res)
				if res.Success {
					http.Redirect(w, r, adminPluginListView, http.StatusFound)
				} else {
					http.Redirect(w, r, adminPluginListViewFail, http.StatusFound)
				}
			} else {
				http.Redirect(w, r, adminloginPage, http.StatusFound)
			}
		}
	}
}

func (h *Six910Handler) processPlugin(r *http.Request) *sdbi.Plugins {
	var i sdbi.Plugins
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	i.PluginName = r.FormValue("pluginName")
	i.Developer = r.FormValue("developer")
	i.ContactPhone = r.FormValue("contactPhone")
	i.DeveloperAddress = r.FormValue("developerAddress")
	fee := r.FormValue("fee")
	i.Fee, _ = strconv.ParseFloat(fee, 64)
	enabled := r.FormValue("enabled")
	i.Enabled, _ = strconv.ParseBool(enabled)
	i.Category = r.FormValue("category")
	i.ActivateURL = r.FormValue("activateUrl")
	i.OauthRedirectURL = r.FormValue("oauthRedirectUrl")
	isPgw := r.FormValue("isPgw")
	i.IsPGW, _ = strconv.ParseBool(isPgw)

	return &i
}
