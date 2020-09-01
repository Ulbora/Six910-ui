package handlers

import (
	"net/http"

	m "github.com/Ulbora/Six910-ui/managers"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
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

//CreateCustomerAccountPage CreateCustomerAccountPage
func (h *Six910Handler) CreateCustomerAccountPage(w http.ResponseWriter, r *http.Request) {
	ccuss, suc := h.getSession(r)
	h.Log.Debug("session suc", ccuss)
	if suc {
		h.Templates.ExecuteTemplate(w, customerCreatePage, nil)
	}
}

//CreateCustomerAccount CreateCustomerAccount
func (h *Six910Handler) CreateCustomerAccount(w http.ResponseWriter, r *http.Request) {
	crcuss, suc := h.getSession(r)
	h.Log.Debug("session suc", crcuss)
	if suc {
		hd := h.getHeader(crcuss)
		//appvars := mux.Vars(r)
		email := r.FormValue("email")
		fcus := h.API.GetCustomer(email, hd)
		if fcus != nil && fcus.ID != 0 {
			http.Redirect(w, r, createCustomerViewFail, http.StatusFound)
		} else {
			var ca m.CustomerAccount
			var cus sdbi.Customer
			var adlst []sdbi.Address
			var usr api.User

			firstName := r.FormValue("firstName")
			lastName := r.FormValue("lastName")
			company := r.FormValue("company")
			city := r.FormValue("city")
			state := r.FormValue("state")
			zip := r.FormValue("zip")
			phone := r.FormValue("phone")
			password := r.FormValue("password")

			billAddress := r.FormValue("billAddress")
			billCity := r.FormValue("billCity")
			billState := r.FormValue("billState")
			billZip := r.FormValue("billZip")
			billCountry := r.FormValue("billCountry")
			if billAddress != "" && billCity != "" && billState != "" && billZip != "" {
				var ba sdbi.Address
				ba.Address = billAddress
				ba.City = billCity
				ba.State = billState
				ba.Zip = billZip
				ba.Country = billCountry
				ba.Type = billingAddressType
				adlst = append(adlst, ba)
			}

			shipAddress := r.FormValue("shipAddress")
			shipCity := r.FormValue("shipCity")
			shipState := r.FormValue("shipState")
			shipZip := r.FormValue("shipZip")
			shipCountry := r.FormValue("shipCountry")
			if shipAddress != "" && shipCity != "" && shipState != "" && shipZip != "" {
				var sa sdbi.Address
				sa.Address = shipAddress
				sa.City = shipCity
				sa.State = shipState
				sa.Zip = shipZip
				sa.Country = shipCountry
				sa.Type = shippingAddressType
				adlst = append(adlst, sa)
			}

			cus.Email = email
			cus.FirstName = firstName
			cus.LastName = lastName
			cus.Company = company
			cus.City = city
			cus.State = state
			cus.Zip = zip
			cus.Phone = phone

			usr.Username = email
			usr.Password = password
			usr.Role = customerRole
			usr.Enabled = true

			ca.Customer = &cus
			ca.Addresses = &adlst
			ca.User = &usr

			suc, cres := h.Manager.CreateCustomerAccount(&ca, hd)
			h.Log.Debug("cres: ", suc)
			h.Log.Debug("acres: ", cres)
			//acres := h.storeCustomerCart(cres, cpls, w, r)
			if suc {
				crcuss.Values["username"] = email
				crcuss.Values["password"] = password
				crcuss.Values["loggedIn"] = true
				crcuss.Values["customerUser"] = true
				serr := crcuss.Save(r, w)
				h.Log.Debug("serr", serr)
				cc := h.getCustomerCart(crcuss)
				cc.CustomerAccount = cres
				h.storeCustomerCart(cc, crcuss, w, r)

				http.Redirect(w, r, customerShoppingCartView, http.StatusFound)
			} else {
				http.Redirect(w, r, createCustomerViewError, http.StatusFound)
			}
		}
	}
}
