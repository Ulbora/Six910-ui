package handlers

import (
	"net/http"
	"strconv"

	m "github.com/Ulbora/Six910-ui/managers"
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

//AddProductToCart AddProductToCart
func (h *Six910Handler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	cpls, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		appvars := mux.Vars(r)
		appidstr := appvars["prodId"]
		appqtystr := appvars["quantity"]
		cppid, _ := strconv.ParseInt(appidstr, 10, 64)
		cppqty, _ := strconv.ParseInt(appqtystr, 10, 64)
		var cpd m.CustomerProduct
		cpd.ProductID = cppid
		cpd.Quantity = cppqty
		if h.isStoreCustomerLoggedIn(cpls) {
			cpd.CustomerID = h.getCustomerID(cpls)
		}
		h.Log.Debug("cusid: ", cpd.CustomerID)

		hd := h.getHeader(cpls)
		cres := h.Manager.AddProductToCart(&cpd, hd)
		acres := h.storeCustomerCart(cres, cpls, w, r)

		h.Log.Debug("cres: ", cres)
		h.Log.Debug("acres: ", acres)

		http.Redirect(w, r, customerShoppingCartView, http.StatusFound)
	}
}

//ViewCart ViewCart
func (h *Six910Handler) ViewCart(w http.ResponseWriter, r *http.Request) {
	ccvs, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		var cv *m.CartView
		cc := h.getCustomerCart(ccvs)
		h.Log.Debug("cc: ", cc)
		if cc != nil {
			hd := h.getHeader(ccvs)
			cv = h.Manager.ViewCart(cc, hd)
		} else {
			var ncv m.CartView
			var ncil []*m.CartViewItem
			ncv.Items = &ncil
			cv = &ncv
		}
		h.Log.Debug("CartView: ", *cv)
		h.Templates.ExecuteTemplate(w, customerShoppingCartView, &cv)
	}
}

//UpdateProductToCart UpdateProductToCart
func (h *Six910Handler) UpdateProductToCart(w http.ResponseWriter, r *http.Request) {
	ucpls, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		uappvars := mux.Vars(r)
		uappidstr := uappvars["prodId"]
		uappqtystr := uappvars["quantity"]
		ucppid, _ := strconv.ParseInt(uappidstr, 10, 64)
		ucppqty, _ := strconv.ParseInt(uappqtystr, 10, 64)
		var ucpd m.CustomerProductUpdate

		if h.isStoreCustomerLoggedIn(ucpls) {
			ucpd.CustomerID = h.getCustomerID(ucpls)
		}
		ccart := h.getCustomerCart(ucpls)
		ucpd.Cart = ccart.Cart

		for i := range *ccart.Items {
			h.Log.Debug("(*ccart.Items)[i]: ", (*ccart.Items)[i])
			if (*ccart.Items)[i].ProductID == ucppid {
				(*ccart.Items)[i].Quantity += ucppqty
				ucpd.CartItem = &(*ccart.Items)[i]
				break
			}
		}

		h.Log.Debug("cusid: ", ucpd.CustomerID)
		h.Log.Debug("CustomerProductUpdate: ", ucpd)
		//h.Log.Debug("CustomerProductUpdate item: ", *ucpd.CartItem)

		hd := h.getHeader(ucpls)
		ucres := h.Manager.UpdateProductToCart(&ucpd, hd)
		acres := h.storeCustomerCart(ucres, ucpls, w, r)

		h.Log.Debug("cres: ", ucres)
		h.Log.Debug("acres: ", acres)

		http.Redirect(w, r, customerShoppingCartView, http.StatusFound)
	}
}
