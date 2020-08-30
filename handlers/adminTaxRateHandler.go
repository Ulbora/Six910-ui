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

//TaxRatePage TaxRatePage
type TaxRatePage struct {
	Error       string
	TaxRate     *sdbi.TaxRate
	TaxRateList *[]sdbi.TaxRate
}

//StoreAdminAddTaxRatePage StoreAdminAddTaxRatePage
func (h *Six910Handler) StoreAdminAddTaxRatePage(w http.ResponseWriter, r *http.Request) {
	adtrs, suc := h.getSession(r)
	h.Log.Debug("session suc in tax rate add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adtrs) {
			atrErr := r.URL.Query().Get("error")
			var atrpg TaxRatePage
			atrpg.Error = atrErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddTaxRatePage, &atrpg)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddTaxRate StoreAdminAddTaxRate
func (h *Six910Handler) StoreAdminAddTaxRate(w http.ResponseWriter, r *http.Request) {
	addtrs, suc := h.getSession(r)
	h.Log.Debug("session suc in taxRate add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(addtrs) {
			atr := h.processTaxRate(r)
			h.Log.Debug("TaxRate add", *atr)
			hd := h.getHeader(addtrs)
			atrres := h.API.AddTaxRate(atr, hd)
			h.Log.Debug("TaxRate add resp", *atrres)
			if atrres.Success {
				http.Redirect(w, r, adminTaxRateListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddTaxRateViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditTaxRatePage  StoreAdminEditTaxRatePage
func (h *Six910Handler) StoreAdminEditTaxRatePage(w http.ResponseWriter, r *http.Request) {
	etrs, suc := h.getSession(r)
	h.Log.Debug("session suc in tax rate edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(etrs) {
			hd := h.getHeader(etrs)
			etrpErr := r.URL.Query().Get("error")
			etrvars := mux.Vars(r)
			country := etrvars["country"]
			state := etrvars["state"]
			idtrstr := etrvars["id"]
			iIDtr, _ := strconv.ParseInt(idtrstr, 10, 64)

			h.Log.Debug("tax rate country in edit", country)
			h.Log.Debug("tax rate state in edit", state)
			var trp TaxRatePage
			trp.Error = etrpErr
			trlst := h.API.GetTaxRate(country, state, hd)
			for i := range *trlst {
				if (*trlst)[i].ID == iIDtr {
					trp.TaxRate = &(*trlst)[i]
					break
				}
			}
			h.Log.Debug("tax rate in edit", trp.TaxRate)
			h.AdminTemplates.ExecuteTemplate(w, adminEditTaxRatePage, &trp)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditTaxRate StoreAdminEditTaxRate
func (h *Six910Handler) StoreAdminEditTaxRate(w http.ResponseWriter, r *http.Request) {
	etrrs, suc := h.getSession(r)
	h.Log.Debug("session suc in tax rate edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(etrrs) {
			etrr := h.processTaxRate(r)
			h.Log.Debug("tax rate update", *etrr)
			hd := h.getHeader(etrrs)
			res := h.API.UpdateTaxRate(etrr, hd)
			h.Log.Debug("tax rate update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminTaxRateListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditTaxRateViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewTaxRateList StoreAdminViewTaxRateList
func (h *Six910Handler) StoreAdminViewTaxRateList(w http.ResponseWriter, r *http.Request) {
	gtrls, suc := h.getSession(r)
	h.Log.Debug("session suc in tax rate view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gtrls) {
			hd := h.getHeader(gtrls)
			trl := h.API.GetTaxRateList(hd)
			h.Log.Debug("tax rate  in list", trl)
			h.AdminTemplates.ExecuteTemplate(w, adminTaxRateListPage, &trl)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteTaxRate StoreAdminDeleteTaxRate
func (h *Six910Handler) StoreAdminDeleteTaxRate(w http.ResponseWriter, r *http.Request) {
	dtrs, suc := h.getSession(r)
	h.Log.Debug("session suc in tax rate delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dtrs) {
			hd := h.getHeader(dtrs)
			dtrvars := mux.Vars(r)
			idstrdtr := dtrvars["id"]
			iddtr, _ := strconv.ParseInt(idstrdtr, 10, 64)
			res := h.API.DeleteTaxRate(iddtr, hd)
			h.Log.Debug("tax rate delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminTaxRateListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminTaxRateListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processTaxRate(r *http.Request) *sdbi.TaxRate {
	var i sdbi.TaxRate
	id := r.FormValue("id")
	i.ID, _ = strconv.ParseInt(id, 10, 64)
	i.Country = r.FormValue("country")
	i.State = r.FormValue("state")
	i.ZipStart = r.FormValue("zipStart")
	i.ZipEnd = r.FormValue("zipEnd")
	percentRate := r.FormValue("percentRate")
	i.PercentRate, _ = strconv.ParseInt(percentRate, 10, 64)
	productCategoryID := r.FormValue("productCategoryId")
	i.ProductCategoryID, _ = strconv.ParseInt(productCategoryID, 10, 64)
	includeHandling := r.FormValue("includeHandling")
	i.IncludeHandling, _ = strconv.ParseBool(includeHandling)
	includeShipping := r.FormValue("includeShipping")
	i.IncludeShipping, _ = strconv.ParseBool(includeShipping)
	i.TaxType = r.FormValue("taxType")
	storeID := r.FormValue("storeId")
	i.StoreID, _ = strconv.ParseInt(storeID, 10, 64)

	return &i
}
