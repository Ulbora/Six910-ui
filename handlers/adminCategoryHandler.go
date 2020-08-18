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

//CatPage CatPage
type CatPage struct {
	Error        string
	CategoryList *[]sdbi.Category
	Category     *sdbi.Category
}

//StoreAdminAddCategoryPage StoreAdminAddCategoryPage
func (h *Six910Handler) StoreAdminAddCategoryPage(w http.ResponseWriter, r *http.Request) {
	acs, suc := h.getSession(r)
	h.Log.Debug("session suc in cat add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(acs) {
			hd := h.getHeader(acs)
			acpErr := r.URL.Query().Get("error")
			var cgp CatPage
			cgp.Error = acpErr
			cgp.CategoryList = h.API.GetCategoryList(hd)
			h.AdminTemplates.ExecuteTemplate(w, adminAddCategoryPage, &cgp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddCategory  StoreAdminAddCategory
func (h *Six910Handler) StoreAdminAddCategory(w http.ResponseWriter, r *http.Request) {
	accs, suc := h.getSession(r)
	h.Log.Debug("session suc in category add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(accs) {
			c := h.processCategory(r)
			h.Log.Debug("Cat add", *c)
			hd := h.getHeader(accs)
			prres := h.API.AddCategory(c, hd)
			h.Log.Debug("Category add resp", *prres)
			if prres.Success {
				http.Redirect(w, r, adminCategoryListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddCategoryViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditCategoryPage StoreAdminEditCategoryPage
func (h *Six910Handler) StoreAdminEditCategoryPage(w http.ResponseWriter, r *http.Request) {
	acs, suc := h.getSession(r)
	h.Log.Debug("session suc in cat edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(acs) {
			hd := h.getHeader(acs)
			acpErr := r.URL.Query().Get("error")
			ecvars := mux.Vars(r)
			idstr := ecvars["id"]
			cID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("prod id in edit", cID)
			var cgp CatPage
			cgp.Error = acpErr

			var wg sync.WaitGroup
			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				cgp.CategoryList = h.API.GetCategoryList(header)
			}(hd)

			wg.Add(1)
			go func(catID int64, header *six910api.Headers) {
				defer wg.Done()
				cgp.Category = h.API.GetCategory(catID, header)
			}(cID, hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminEditCategoryPage, &cgp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditCategory StoreAdminEditCategory
func (h *Six910Handler) StoreAdminEditCategory(w http.ResponseWriter, r *http.Request) {
	eccs, suc := h.getSession(r)
	h.Log.Debug("session suc in cat edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(eccs) {
			ecc := h.processCategory(r)
			h.Log.Debug("Cat update", *ecc)
			hd := h.getHeader(eccs)
			res := h.API.UpdateCategory(ecc, hd)
			h.Log.Debug("Cat update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminCategoryListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditCategoryViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewCategoryList StoreAdminViewCategoryList
func (h *Six910Handler) StoreAdminViewCategoryList(w http.ResponseWriter, r *http.Request) {
	gcls, suc := h.getSession(r)
	h.Log.Debug("session suc in cats view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gcls) {
			hd := h.getHeader(gcls)
			cats := h.API.GetCategoryList(hd)
			h.Log.Debug("prods  in edit", cats)
			h.AdminTemplates.ExecuteTemplate(w, adminCategoryListPage, &cats)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteCategory StoreAdminDeleteCategory
func (h *Six910Handler) StoreAdminDeleteCategory(w http.ResponseWriter, r *http.Request) {
	dcs, suc := h.getSession(r)
	h.Log.Debug("session suc in cat list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dcs) {
			hd := h.getHeader(dcs)
			dcvars := mux.Vars(r)
			idstrd := dcvars["id"]
			idddc, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeleteCategory(idddc, hd)
			h.Log.Debug("cat delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminCategoryListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminCategoryListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processCategory(r *http.Request) *sdbi.Category {
	var c sdbi.Category
	id := r.FormValue("id")
	c.ID, _ = strconv.ParseInt(id, 10, 64)
	c.Name = r.FormValue("name")
	c.Description = r.FormValue("desc")
	c.Image = r.FormValue("image")
	c.Thumbnail = r.FormValue("thumbnail")
	storeID := r.FormValue("storeId")
	c.StoreID, _ = strconv.ParseInt(storeID, 10, 64)
	pID := r.FormValue("parentId")
	c.ParentCategoryID, _ = strconv.ParseInt(pID, 10, 64)

	return &c
}
