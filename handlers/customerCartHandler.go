package handlers

import (
	"net/http"
	"strconv"
	"sync"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	six910api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"

	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
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

//PaymentMethod PaymentMethod
type PaymentMethod struct {
	Name           string
	PaymentGateway *sdbi.PaymentGateway
}

//CheckoutPage CheckoutPage
type CheckoutPage struct {
	CustomerCart       *m.CustomerCart
	PaymentMethodList  *[]PaymentMethod
	ShippingMethodList *[]sdbi.ShippingMethod
	InsuranceList      *[]sdbi.Insurance
}

//CartPage CartPage
type CartPage struct {
	CustomerCart *m.CustomerCart
	PageBody     *csssrv.PageCSS
	MenuList     *[]musrv.Menu
	Content      *conts.Content
	//meta data
	HeaderData *HeaderData
}

//AddProductToCart AddProductToCart
func (h *Six910Handler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	cpls, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		var cppid int64
		var cppqty int64

		query := r.URL.Query()
		id := query.Get("id")
		qty := query.Get("qty")
		if id != "" && qty != "" {
			cppid, _ = strconv.ParseInt(id, 10, 64)
			cppqty, _ = strconv.ParseInt(qty, 10, 64)
		} else {
			appvars := mux.Vars(r)
			appidstr := appvars["prodId"]
			// appqtystr := appvars["quantity"]
			cppid, _ = strconv.ParseInt(appidstr, 10, 64)
			// cppqty, _ = strconv.ParseInt(appqtystr, 10, 64)
			cppqty = 1
		}

		var cpd m.CustomerProduct
		cpd.ProductID = cppid
		cpd.Quantity = cppqty
		if h.isStoreCustomerLoggedIn(cpls) {
			cpd.CustomerID = h.getCustomerID(cpls)
		}
		h.Log.Debug("cusid: ", cpd.CustomerID)

		hd := h.getHeader(cpls)
		cc := h.getCustomerCart(cpls)
		if cc != nil {
			cpd.Cart = cc.Cart
			if cc.Items != nil {
				for i := range *cc.Items {
					if (*cc.Items)[i].ProductID == cppid {
						cpd.CartItem = &(*cc.Items)[i]
					}
				}
			}
		}
		cres := h.Manager.AddProductToCart(&cpd, hd)
		acres := h.storeCustomerCart(cres, cpls, w, r)

		h.Log.Debug("cres: ", cres)
		h.Log.Debug("acres: ", acres)

		http.Redirect(w, r, customerShoppingCartView, http.StatusFound)
	}
}

//ViewCart ViewCart
func (h *Six910Handler) ViewCart(w http.ResponseWriter, r *http.Request) {
	ccvs, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		var cv *m.CartView
		cc := h.getCustomerCart(ccvs)
		h.Log.Debug("cc: ", cc)
		hd := h.getHeader(ccvs)
		if cc != nil && cc.Items != nil && len(*cc.Items) > 0 {
			cv = h.Manager.ViewCart(cc, hd)
			cc.CartView = cv
			h.storeCustomerCart(cc, ccvs, w, r)
		} else {
			var ncv m.CartView
			var ncil []*m.CartViewItem
			ncv.Items = &ncil
			cv = &ncv
			cc = new(m.CustomerCart)
		}

		var cpage CartPage
		cpage.CustomerCart = cc

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		h.Log.Debug("PageBody: ", *csspg)
		cpage.PageBody = csspg

		ml := h.MenuService.GetMenuList()
		h.getCartTotal(ccvs, ml, hd)
		cpage.MenuList = ml

		h.Log.Debug("MenuList", *cpage.MenuList)

		cisuc, cicont := h.ContentService.GetContent(shoppingCartContent)
		if cisuc {
			cpage.Content = cicont
		} else {
			var ct conts.Content
			cpage.Content = &ct
		}

		h.Log.Debug("CartView: ", *cv)
		h.Templates.ExecuteTemplate(w, customerShoppingCartPage, &cpage)
	}
}

//UpdateProductToCart UpdateProductToCart
func (h *Six910Handler) UpdateProductToCart(w http.ResponseWriter, r *http.Request) {
	ucpls, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {

		query := r.URL.Query()
		uappidstr := query.Get("id")
		uappqtystr := query.Get("qty")
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
				(*ccart.Items)[i].Quantity = ucppqty
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

//CheckOutView CheckOutView
func (h *Six910Handler) CheckOutView(w http.ResponseWriter, r *http.Request) {
	cocvs, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cocvs) {
			var cop CheckoutPage
			cocc := h.getCustomerCart(cocvs)
			cop.CustomerCart = cocc
			var wg sync.WaitGroup
			hd := h.getHeader(cocvs)
			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				var mplist []PaymentMethod
				pgs := h.API.GetPaymentGateways(header)
				for i := range *pgs {
					var pg = (*pgs)[i]
					sp := h.API.GetStorePlugin(pg.StorePluginsID, header)
					var pm PaymentMethod
					pm.Name = sp.PluginName
					pm.PaymentGateway = &pg
					mplist = append(mplist, pm)
				}
				cop.PaymentMethodList = &mplist
			}(hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				cop.ShippingMethodList = h.API.GetShippingMethodList(header)
			}(hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				cop.InsuranceList = h.API.GetInsuranceList(header)
			}(hd)

			wg.Wait()

			h.Log.Debug("CheckoutPage: ", cop)
			h.Templates.ExecuteTemplate(w, customerShoppingCartPage, &cop)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//CheckOutContinue CheckOutContinue
func (h *Six910Handler) CheckOutContinue(w http.ResponseWriter, r *http.Request) {
	// this is where insurance, shipping and taxes are calculated
	// returns results to user before final checkout
	//items to get:
	//1. PaymentGatewayID
	//2. ShippingMethodID
	//3. InsuranceID
	//4. BillingAddressID
	//5. ShippingAddressID

	//tax calc: country, state, zipstart zipend, %, prod category, inc handling, inc shipping, tax type(sales, vat)
	cocccs, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cocccs) {
			//uappvars := mux.Vars(r)
			pidStr := r.FormValue("paymentGatewayID")
			smidStr := r.FormValue("shippingMethodID")
			insidStr := r.FormValue("insuranceID")
			baidStr := r.FormValue("billingAddressID")
			saidStr := r.FormValue("shippingAddressID")

			pgwid, _ := strconv.ParseInt(pidStr, 10, 64)
			smid, _ := strconv.ParseInt(smidStr, 10, 64)
			insid, _ := strconv.ParseInt(insidStr, 10, 64)
			baid, _ := strconv.ParseInt(baidStr, 10, 64)
			said, _ := strconv.ParseInt(saidStr, 10, 64)

			ccoart := h.getCustomerCart(cocccs)
			ccoart.PaymentGatewayID = pgwid
			ccoart.ShippingMethodID = smid
			ccoart.InsuranceID = insid
			ccoart.BillingAddressID = baid
			ccoart.ShippingAddressID = said
			h.Log.Debug("ccoart: ", *ccoart)

			hd := h.getHeader(cocccs)
			ccotres := h.Manager.CalculateCartTotals(ccoart, hd)

			acres := h.storeCustomerCart(ccotres, cocccs, w, r)

			h.Log.Debug("ccotres: ", ccotres)
			h.Log.Debug("acres: ", acres)

			h.Templates.ExecuteTemplate(w, customerShoppingCartContinuePage, &ccotres)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}
