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
		h.storeCustomerCart(cres, cpls, w, r)

		h.Log.Debug("cres: ", cres)
	}
}
