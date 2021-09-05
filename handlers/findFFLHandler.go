package handlers

import (
	"net/http"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	fflsrv "github.com/Ulbora/Six910-ui/findfflsrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
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

//FFLSearchPage FFLSearchPage
type FFLSearchPage struct {
	Zip           string
	FFLList       *[]fflsrv.FFLList
	FFL           *fflsrv.FFL
	ListFound     bool
	ShowNoResults bool
	PageBody      *csssrv.PageCSS
	MenuList      *[]musrv.Menu
	Content       *conts.Content
	HeaderData    *HeaderData
}

//FindFFLZipPage FindFFLZipPage
func (h *Six910Handler) FindFFLZipPage(w http.ResponseWriter, r *http.Request) {
	cfflss, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", cfflss)
	if suc {
		if h.isStoreCustomerLoggedIn(cfflss) {
			zip := r.URL.Query().Get("zip")
			var pg FFLSearchPage
			pg.Zip = zip

			hd := h.getHeader(cfflss)

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			pg.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(cfflss, ml, hd)
			pg.MenuList = ml

			h.Log.Debug("MenuList", *pg.MenuList)

			cisuc, cicont := h.ContentService.GetContent(shoppingCartContent3)
			if cisuc {
				pg.Content = cicont
			} else {
				var ct conts.Content
				pg.Content = &ct
			}
			h.Templates.ExecuteTemplate(w, fflSearchPage, pg)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//FindFFLZip FindFFLZip
func (h *Six910Handler) FindFFLZip(w http.ResponseWriter, r *http.Request) {
	cfflss2, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cfflss2) {
			zip := r.FormValue("zip")
			h.Log.Debug("zip", zip)

			fflLst, code := h.FFLService.GetFFLList(zip)
			h.Log.Debug("code: ", code)
			h.Log.Debug("fflLst: ", *fflLst)

			var fflpage FFLSearchPage
			fflpage.FFLList = fflLst
			if len(*fflLst) > 0 {
				fflpage.ListFound = true
			} else {
				fflpage.ShowNoResults = true
			}
			hd := h.getHeader(cfflss2)

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			fflpage.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(cfflss2, ml, hd)
			fflpage.MenuList = ml

			h.Log.Debug("MenuList", *fflpage.MenuList)

			cisuc, cicont := h.ContentService.GetContent(shoppingCartContent3)
			if cisuc {
				fflpage.Content = cicont
			} else {
				var ct conts.Content
				fflpage.Content = &ct
			}
			h.Log.Debug("fflpage: ", fflpage)
			h.Templates.ExecuteTemplate(w, fflSearchPage, &fflpage)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//FindFFLID FindFFLID
func (h *Six910Handler) FindFFLID(w http.ResponseWriter, r *http.Request) {
	cfflss3, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc fflid", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cfflss3) {
			fflvars := mux.Vars(r)
			fflid := fflvars["id"]
			h.Log.Debug("fflid", fflid)

			ffl, code := h.FFLService.GetFFL(fflid)
			h.Log.Debug("code fflid: ", code)
			h.Log.Debug("ffl: ", *ffl)

			var fflpage FFLSearchPage
			fflpage.FFL = ffl

			hd := h.getHeader(cfflss3)

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			fflpage.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(cfflss3, ml, hd)
			fflpage.MenuList = ml

			h.Log.Debug("MenuList", *fflpage.MenuList)

			cisuc, cicont := h.ContentService.GetContent(shoppingCartContent3)
			if cisuc {
				fflpage.Content = cicont
			} else {
				var ct conts.Content
				fflpage.Content = &ct
			}

			h.Log.Debug("fflpage: ", fflpage)
			h.Templates.ExecuteTemplate(w, fflPage, &fflpage)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//AddFFL AddFFL
func (h *Six910Handler) AddFFL(w http.ResponseWriter, r *http.Request) {
	fflses, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc add ffl", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(fflses) {
			fflemail := fflses.Values["username"]
			h.Log.Debug("fflemail: ", fflemail)
			var uname string
			if fflemail != nil {
				uname = fflemail.(string)
			}

			hd := h.getHeader(fflses)
			//cc.CustomerAccount.Customer.ID
			//cc := h.getCustomerCart(ccuuss)
			if uname != "" {
				fflcus := h.API.GetCustomer(uname, hd)

				fflvars := mux.Vars(r)
				fflid := fflvars["id"]
				h.Log.Debug("fflid in add ffl", fflid)

				ffl, code := h.FFLService.GetFFL(fflid)
				h.Log.Debug("code fflid in add ffl to address: ", code)
				h.Log.Debug("ffl: ", *ffl)
				var success bool
				if fflcus != nil {
					// address := ffl.Address
					// city := ffl.City
					// state := ffl.State
					// zip := ffl.PremiseZip
					// // country :=
					// atype := "FFL"
					var nad sdbi.Address
					nad.Address = ffl.Address
					nad.City = ffl.City
					nad.State = ffl.State
					nad.Zip = ffl.PremiseZip
					nad.Country = ""
					nad.Type = "FFL"
					nad.CustomerID = fflcus.ID
					var fflname = ffl.BusName
					if fflname == "" {
						fflname = ffl.LicName
					}
					nad.Attr1 = fflname
					nad.Attr2 = ffl.Lic
					nad.Attr3 = ffl.ExpDate
					nad.Attr4 = ffl.Phone
					res := h.API.AddAddress(&nad, hd)
					success = res.Success
					h.Log.Debug("add ffl address suc: ", success)
					if success {
						cc := h.getCustomerCart(fflses)
						addLst := h.API.GetAddressList(fflcus.ID, hd)
						cc.CustomerAccount.Addresses = addLst
						h.storeCustomerCart(cc, fflses, w, r)
					}
				}

				http.Redirect(w, r, startCheckout, http.StatusFound)

				// if success {
				// 	http.Redirect(w, r, customerInfoView, http.StatusFound)
				// } else {
				// 	http.Redirect(w, r, customerInfoViewFail, http.StatusFound)
				// }

				// var fflpage FFLSearchPage
				// fflpage.FFL = ffl
				// h.Log.Debug("fflpage: ", fflpage)
				// h.Templates.ExecuteTemplate(w, fflPage, &fflpage)
			} else {
				http.Redirect(w, r, customerLoginView, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}

	}
}
