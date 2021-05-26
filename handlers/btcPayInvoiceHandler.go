package handlers

import (
	"net/http"
	"strconv"

	cl "github.com/Ulbora/BTCPayClient"
	pi "github.com/Ulbora/Six910BTCPayServerPlugin"
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

//CompleteBTCPayTransaction CompleteBTCPayTransaction
func (h *Six910Handler) CompleteBTCPayTransaction(w http.ResponseWriter, r *http.Request) {
	cocodbc, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cocodbc) {
			codrbcvars := mux.Vars(r)
			//cartIdstr := codrvars["cartId"]
			total := codrbcvars["total"]
			taxIncluded := codrbcvars["tax"]
			firstName := codrbcvars["firstName"]
			lastName := codrbcvars["lastName"]
			email := codrbcvars["email"]
			hd := h.getHeader(cocodbc)
			pg := h.API.GetPaymentGatewayByName(btcPayServer, hd)
			h.checkBTCPayPlugin(pg)
			// if !h.BTCPlugin.IsPluginLoaded() {
			// 	var btc pi.BTCPay
			// 	btc.ClientID = pg.ClientID
			// 	btc.Host = pg.CheckoutURL
			// 	btc.PrivateKey = pg.ClientKey
			// 	btc.Token = pg.Token
			// 	h.BTCPlugin.NewClient(&btc)

			// }
			var inv cl.InvoiceReq
			price, err := strconv.ParseFloat(total, 64)
			if err == nil {
				tax, err := strconv.ParseFloat(taxIncluded, 64)
				if err == nil {
					inv.Price = price
					inv.TaxIncluded = tax
					inv.Currency = h.BTCPayCurrency
					inv.Buyer.Name = firstName + " " + lastName
					inv.Buyer.Email = email
					inv.TransactionSpeed = "medium"
					inv.Token = h.BTCPlugin.GetToken()
					inv.RedirectURL = pg.PostOrderURL + "{InvoiceId}" // h.Six910SiteURL + "/completeOrder/"
					inv.RedirectAutomatically = true
					res := h.BTCPlugin.CreateInvoice(&inv)
					h.Log.Debug("Invoice Res", *res)
					if res.Data.URL != "" {
						http.Redirect(w, r, res.Data.URL, http.StatusFound)
					}
				}
			}
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

func (h *Six910Handler) checkBTCPayPlugin(pg *sdbi.PaymentGateway) {
	if !h.BTCPlugin.IsPluginLoaded() {
		var btc pi.BTCPay
		btc.ClientID = pg.ClientID
		btc.Host = pg.CheckoutURL
		btc.PrivateKey = pg.ClientKey
		btc.Token = pg.Token
		h.BTCPlugin.NewClient(&btc)
	}
}
