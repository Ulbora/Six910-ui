package handlers

import (
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	"github.com/gorilla/mux"
	"net/http"
	"sort"
	"strings"
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

//StoreAdminAddMenuPage StoreAdminAddMenuPage
func (h *Six910Handler) StoreAdminAddMenuPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			h.AdminTemplates.ExecuteTemplate(w, adminAddMenuPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddMenu StoreAdminAddMenu
func (h *Six910Handler) StoreAdminAddMenu(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc add menu", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			act := h.processMenu(r)
			res := h.MenuService.AddMenu(act)
			h.Log.Debug("add menu res", res)
			if res {
				http.Redirect(w, r, adminMenuList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddMenuFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminGetMenu StoreAdminGetMenu
func (h *Six910Handler) StoreAdminGetMenu(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			mvars := mux.Vars(r)
			mname := mvars["name"]

			_, mres := h.MenuService.GetMenu(mname)
			h.Log.Debug("menu in get: ", *mres)
			h.Log.Debug("content and image list in get content: ", mres)

			h.AdminTemplates.ExecuteTemplate(w, adminUpdateMenu, &mres)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUpdateMenu StoreAdminUpdateMenu
func (h *Six910Handler) StoreAdminUpdateMenu(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			uct := h.processMenu(r)
			res := h.MenuService.UpdateMenu(uct)
			h.Log.Debug("update menu res", res)
			if res {
				http.Redirect(w, r, adminMenuList, http.StatusFound)
			} else {
				//go back
				http.Redirect(w, r, adminGetMenu+"/"+uct.Name, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminMenuList StoreAdminMenuList
func (h *Six910Handler) StoreAdminMenuList(w http.ResponseWriter, r *http.Request) {
	gmls, suc := h.getSession(r)
	h.Log.Debug("session suc in menu view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gmls) {
			msl := h.MenuService.GetMenuList()
			sort.Slice(*msl, func(p, q int) bool {
				return (*msl)[p].Name < (*msl)[q].Name
			})
			h.Log.Debug("Menu  in list", msl)
			h.AdminTemplates.ExecuteTemplate(w, adminMenuListPage, &msl)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processMenu(r *http.Request) *musrv.Menu {
	name := r.FormValue("name")
	name = strings.Replace(name, " ", "", -1)
	h.Log.Debug("name in new menu: ", name)

	location := r.FormValue("location")
	h.Log.Debug("location in new menu: ", location)

	active := r.FormValue("active")
	h.Log.Debug("active in new menu: ", active)

	brand := r.FormValue("brand")
	h.Log.Debug("brand in new menu: ", brand)

	brandLink := r.FormValue("brandLink")
	h.Log.Debug("brandLink in new menu: ", brandLink)

	shade := r.FormValue("shade")
	h.Log.Debug("shade in new menu: ", shade)

	background := r.FormValue("background")
	h.Log.Debug("background in new menu: ", background)

	style := r.FormValue("style")
	h.Log.Debug("style in new menu: ", style)

	shade0 := r.FormValue("shade0")
	h.Log.Debug("shade0 in new menu: ", shade0)

	shade1 := r.FormValue("shade1")
	h.Log.Debug("shade1 in new menu: ", shade1)

	shade2 := r.FormValue("shade2")
	h.Log.Debug("shade2 in new menu: ", shade2)

	shade3 := r.FormValue("shade3")
	h.Log.Debug("shade3 in new menu: ", shade3)

	shade4 := r.FormValue("shade4")
	h.Log.Debug("shade4 in new menu: ", shade4)

	shade5 := r.FormValue("shade5")
	h.Log.Debug("shade5 in new menu: ", shade5)

	bg0 := r.FormValue("bg0")
	h.Log.Debug("bg0 in new menu: ", bg0)

	bg1 := r.FormValue("bg1")
	h.Log.Debug("bg1 in new menu: ", bg1)

	bg2 := r.FormValue("bg2")
	h.Log.Debug("bg2 in new menu: ", bg2)

	bg3 := r.FormValue("bg3")
	h.Log.Debug("bg3 in new menu: ", bg3)

	bg4 := r.FormValue("bg4")
	h.Log.Debug("bg4 in new menu: ", bg4)

	bg5 := r.FormValue("bg5")
	h.Log.Debug("bg5 in new menu: ", bg5)

	menuName0 := r.FormValue("menuName0")
	h.Log.Debug("menuName0 in new menu: ", menuName0)

	menuName1 := r.FormValue("menuName1")
	h.Log.Debug("menuName1 in new menu: ", menuName1)

	menuName2 := r.FormValue("menuName2")
	h.Log.Debug("menuName2 in new menu: ", menuName2)

	menuName3 := r.FormValue("menuName3")
	h.Log.Debug("menuName3 in new menu: ", menuName3)

	menuName4 := r.FormValue("menuName4")
	h.Log.Debug("menuName4 in new menu: ", menuName4)

	menuName5 := r.FormValue("menuName5")
	h.Log.Debug("menuName5 in new menu: ", menuName5)

	menuName6 := r.FormValue("menuName6")
	h.Log.Debug("menuName6 in new menu: ", menuName6)

	menuName7 := r.FormValue("menuName7")
	h.Log.Debug("menuName7 in new menu: ", menuName7)

	menuLink0 := r.FormValue("menuLink0")
	h.Log.Debug("menuLink0 in new menu: ", menuLink0)

	menuLink1 := r.FormValue("menuLink1")
	h.Log.Debug("menuLink1 in new menu: ", menuLink1)

	menuLink2 := r.FormValue("menuLink2")
	h.Log.Debug("menuLink2 in new menu: ", menuLink2)

	menuLink3 := r.FormValue("menuLink3")
	h.Log.Debug("menuLink3 in new menu: ", menuLink3)

	menuLink4 := r.FormValue("menuLink4")
	h.Log.Debug("menuLink4 in new menu: ", menuLink4)

	menuLink5 := r.FormValue("menuLink5")
	h.Log.Debug("menuLink5 in new menu: ", menuLink5)

	menuLink6 := r.FormValue("menuLink6")
	h.Log.Debug("menuLink6 in new menu: ", menuLink6)

	menuLink7 := r.FormValue("menuLink7")
	h.Log.Debug("menuLink7 in new menu: ", menuLink7)

	// archived := r.FormValue("archived")
	// h.Log.Debug("archived in new content: ", archived)

	var ct musrv.Menu
	ct.Name = name
	ct.Location = location
	ct.Brand = brand
	ct.BrandLink = brandLink
	ct.Shade = shade
	ct.Background = background
	ct.Style = style

	if active == "on" {
		ct.Active = true
	} else {
		ct.Active = false
	}

	shdlst := []string{shade0, shade1, shade2, shade3, shade4, shade5}
	ct.ShadeList = &shdlst

	bglst := []string{bg0, bg1, bg2, bg3, bg4, bg5}
	ct.BackgroundList = &bglst

	var m1 musrv.MenuItemItem
	m1.Name = menuName0
	m1.Link = menuLink0

	var m2 musrv.MenuItemItem
	m2.Name = menuName1
	m2.Link = menuLink1

	var m3 musrv.MenuItemItem
	m3.Name = menuName2
	m3.Link = menuLink2

	var m4 musrv.MenuItemItem
	m4.Name = menuName3
	m4.Link = menuLink3

	var m5 musrv.MenuItemItem
	m5.Name = menuName4
	m5.Link = menuLink4

	var m6 musrv.MenuItemItem
	m6.Name = menuName5
	m6.Link = menuLink5

	var m7 musrv.MenuItemItem
	m7.Name = menuName6
	m7.Link = menuLink6

	var m8 musrv.MenuItemItem
	m8.Name = menuName7
	m8.Link = menuLink7

	mlst := []musrv.MenuItemItem{m1, m2, m3, m4, m5, m6, m7, m8}
	ct.MenuItemList = &mlst

	return &ct
}
