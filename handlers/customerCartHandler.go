package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"
	"sync"
	"time"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	six910api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"

	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"

	mll "github.com/Ulbora/go-mail-sender"
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
	CheckoutURL    template.URL
}

//ShippingMethod ShippingMethod
type ShippingMethod struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Cost       float64 `json:"cost"`
	RegionName string  `json:"regionName"`
}

//CheckoutPage CheckoutPage
type CheckoutPage struct {
	CustomerCart           *m.CustomerCart
	PageBody               *csssrv.PageCSS
	MenuList               *[]musrv.Menu
	Content                *conts.Content
	PaymentMethodList      *[]PaymentMethod
	PaymentMethod          *PaymentMethod
	ShippingMethodList     *[]ShippingMethod
	ShippingMethod         *sdbi.ShippingMethod
	InsuranceList          *[]sdbi.Insurance
	ShowInsurance          bool
	CustomerAddressList    *[]sdbi.Address
	BillingAddress         *sdbi.Address
	ShippingAddress        *sdbi.Address
	ShowAddressList        bool
	Subtotal               string
	ShippingHandling       string
	InsuranceCost          string
	Taxes                  string
	Total                  string
	OrderInfo              string
	PayPalAuthorizePayment bool
	PayPalPayment          bool
	BillMeLaterPayment     bool
	BTCPayServerPayment    bool
	OrderNumber            string

	HeaderData *HeaderData
}

//CartPage CartPage
type CartPage struct {
	CustomerCart *m.CustomerCart
	PageBody     *csssrv.PageCSS
	MenuList     *[]musrv.Menu
	Content      *conts.Content
	//meta data
	HeaderData *HeaderData
	OrderList  *[]sdbi.Order
}

// //PayPalPayload PayPalPayload
// type PayPalPayload struct {
// 	Description string `json:"description"`
// }

//AddProductToCart AddProductToCart
func (h *Six910Handler) AddProductToCart(w http.ResponseWriter, r *http.Request) {
	cpls, suc := h.getUserSession(w, r)
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
		h.Log.Debug("cc: ", cc)
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

		cres := h.Manager.AddProductToCart(cc, &cpd, hd)
		acres := h.storeCustomerCart(cres, cpls, w, r)

		h.Log.Debug("cres: ", cres)
		h.Log.Debug("acres: ", acres)

		http.Redirect(w, r, customerShoppingCartView, http.StatusFound)
	}
}

//ViewCart ViewCart
func (h *Six910Handler) ViewCart(w http.ResponseWriter, r *http.Request) {
	ccvs, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		var cv *m.CartView
		cc := h.getCustomerCart(ccvs)
		h.Log.Debug("cc: ", cc)
		//h.Log.Debug("cc: ", *cc.Cart)
		//h.Log.Debug("cc items: ", *cc.Items)
		h.Log.Debug("cusId in viewCart: ", h.getCustomerID(ccvs))
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
	ucpls, suc := h.getUserSession(w, r)
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
		cc := h.getCustomerCart(ucpls)
		ucres := h.Manager.UpdateProductToCart(cc, &ucpd, hd)
		acres := h.storeCustomerCart(ucres, ucpls, w, r)

		h.Log.Debug("cres: ", ucres)
		h.Log.Debug("acres: ", acres)

		http.Redirect(w, r, customerShoppingCartView, http.StatusFound)
	}
}

//CheckOutView CheckOutView
func (h *Six910Handler) CheckOutView(w http.ResponseWriter, r *http.Request) {
	cocvs, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cocvs) {
			var cop CheckoutPage
			cocc := h.getCustomerCart(cocvs)
			h.Log.Debug("Customer cart: ", *cocc)
			h.Log.Debug("Customer Account: ", cocc.CustomerAccount)
			cop.CustomerCart = cocc
			var wg sync.WaitGroup
			hd := h.getHeader(cocvs)
			cid := h.getCustomerID(cocvs)
			h.Log.Debug("Customer ID: ", cid)
			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				var mplist []PaymentMethod
				pgs := h.API.GetPaymentGateways(header)
				h.Log.Debug("Payment Gateways: ", *pgs)
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
				slst := h.API.GetShippingMethodList(header)
				var smlst []ShippingMethod
				for _, sm := range *slst {
					var nsm ShippingMethod
					nsm.ID = sm.ID
					nsm.Cost = sm.Cost
					nsm.Name = sm.Name
					rg := h.API.GetRegion(sm.RegionID, hd)
					nsm.RegionName = rg.Name
					smlst = append(smlst, nsm)
				}
				cop.ShippingMethodList = &smlst
			}(hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				cop.InsuranceList = h.API.GetInsuranceList(header)
				if len(*cop.InsuranceList) > 0 {
					cop.ShowInsurance = true
				}
			}(hd)

			wg.Add(1)
			// go func(cart *m.CustomerCart, header *six910api.Headers) {
			// 	defer wg.Done()
			// 	//if cart.CustomerAccount != nil && cart.CustomerAccount.Customer != nil {

			// 	if h.getCustomerID != 0 {
			// 		cop.CustomerAddressList = h.API.GetAddressList(h.getCustomerID, header)
			// 		if len(*cop.CustomerAddressList) > 0 {
			// 			cop.ShowAddressList = true
			// 		}
			// 	}
			// }(cocc, hd)

			go func(cusId int64, header *six910api.Headers) {
				defer wg.Done()
				//if cart.CustomerAccount != nil && cart.CustomerAccount.Customer != nil {

				if cusId != 0 {
					cop.CustomerAddressList = h.API.GetAddressList(cusId, header)
					if len(*cop.CustomerAddressList) > 0 {
						cop.ShowAddressList = true
					}
				}
			}(cid, hd)

			wg.Wait()

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			cop.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(cocvs, ml, hd)
			cop.MenuList = ml

			h.Log.Debug("MenuList", *cop.MenuList)

			cisuc, cicont := h.ContentService.GetContent(shoppingCartContent2)
			if cisuc {
				cop.Content = cicont
			} else {
				var ct conts.Content
				cop.Content = &ct
			}

			h.Log.Debug("CheckoutPage: ", cop)
			h.Templates.ExecuteTemplate(w, customerShoppingCartPage2, &cop)
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
	cocccs, suc := h.getUserSession(w, r)
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
			h.Log.Debug("ccoart.InsuranceID: ", ccoart.InsuranceID)
			h.Log.Debug("ccoart.Items: ", ccoart.Items)

			hd := h.getHeader(cocccs)
			ccotres := h.Manager.CalculateCartTotals(ccoart, hd)
			// h.Log.Debug("ccotres.OrderID: ", ccotres.OrderID)
			// if ccotres.OrderID == 0 {
			// 	odrRes := h.Manager.CheckOut(ccotres, hd)
			// 	h.Log.Debug("odrRes after CheckOut: ", *odrRes.Order)
			// 	ccotres.OrderID = odrRes.Order.ID
			// 	h.Log.Debug("ccotres.OrderID after create: ", ccotres.OrderID)
			// }

			pgw := h.API.GetPaymentGateway(pgwid, hd)
			sp := h.API.GetStorePlugin(pgw.StorePluginsID, hd)
			var pm PaymentMethod
			pm.Name = sp.PluginName
			pm.PaymentGateway = pgw
			pm.CheckoutURL = template.URL(pgw.CheckoutURL)

			sm := h.API.GetShippingMethod(smid, hd)

			acres := h.storeCustomerCart(ccotres, cocccs, w, r)

			h.Log.Debug("ccotres: ", ccotres)
			h.Log.Debug("acres: ", acres)
			h.Log.Debug("pgw: ", *pgw)

			// var wg sync.WaitGroup
			var ccop CheckoutPage
			if strings.Contains(strings.ToLower(pm.Name), "paypal authorize") {
				h.Log.Debug("Using PayPay Authorize Gateway")
				ccop.PayPalAuthorizePayment = true
			} else if strings.Contains(strings.ToLower(pm.Name), "paypal") {
				h.Log.Debug("Using PayPay Regualr Gateway")
				ccop.PayPalPayment = true
			} else if strings.Contains(strings.ToLower(pm.Name), "bill me later") {
				h.Log.Debug("Using Bill Me Later Gateway")
				ccop.BillMeLaterPayment = true
			} else if strings.Contains(strings.ToLower(pgw.Name), "btcpayserver") {
				h.Log.Debug("Using BTCPayServer Gateway")
				ccop.BTCPayServerPayment = true
			}
			ccop.OrderInfo = h.CompanyName
			ccop.CustomerCart = ccotres
			ccop.PaymentMethod = &pm
			ccop.ShippingMethod = sm
			ccop.BillingAddress = h.API.GetAddress(baid, ccotres.Cart.CustomerID, hd)
			ccop.ShippingAddress = h.API.GetAddress(said, ccotres.Cart.CustomerID, hd)
			ccop.Subtotal = fmt.Sprintf("%.2f", ccotres.Subtotal)
			ccop.ShippingHandling = fmt.Sprintf("%.2f", ccotres.ShippingHandling)
			ccop.InsuranceCost = fmt.Sprintf("%.2f", ccotres.InsuranceCost)
			ccop.Taxes = fmt.Sprintf("%.2f", ccotres.Taxes)
			ccop.Total = fmt.Sprintf("%.2f", ccotres.Total)

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			ccop.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(cocccs, ml, hd)
			ccop.MenuList = ml

			h.Log.Debug("MenuList", *ccop.MenuList)

			cisuc, cicont := h.ContentService.GetContent(shoppingCartContent3)
			if cisuc {
				ccop.Content = cicont
			} else {
				var ct conts.Content
				ccop.Content = &ct
			}

			h.Templates.ExecuteTemplate(w, customerShoppingCartContinuePage, &ccop)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//CheckOutComplateOrder CheckOutComplateOrder
func (h *Six910Handler) CheckOutComplateOrder(w http.ResponseWriter, r *http.Request) {
	cocod, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(cocod) {
			codrvars := mux.Vars(r)
			//cartIdstr := codrvars["cartId"]
			transactionCode := codrvars["transactionCode"]
			// appqtystr := appvars["quantity"]
			//cartID, _ = strconv.ParseInt(cartIdstr, 10, 64)

			hd := h.getHeader(cocod)
			comccotres := h.getCustomerCart(cocod)
			if transactionCode == "billMeLaterTransaction" {
				comccotres.BillMeLater = true
			}

			h.Log.Debug("comccotres: ", *comccotres.CustomerAccount)
			if comccotres.Items != nil {
				odrRes := h.Manager.CheckOut(comccotres, hd)
				h.Log.Debug("odrRes after CheckOut: ", *odrRes.Order)
				//ccotres.OrderID = odrRes.Order.ID
				h.Log.Debug("comccotres.OrderID after create: ", comccotres.OrderID)

				var ccopc CheckoutPage
				var wg1 sync.WaitGroup

				wg1.Add(1)
				go func(smid int64, header *six910api.Headers) {
					defer wg1.Done()
					ccopc.ShippingMethod = h.API.GetShippingMethod(smid, header)
				}(comccotres.ShippingMethodID, hd)

				wg1.Add(1)
				go func(baid int64, cid int64, header *six910api.Headers) {
					defer wg1.Done()
					ccopc.BillingAddress = h.API.GetAddress(baid, cid, header)
				}(comccotres.BillingAddressID, comccotres.Cart.CustomerID, hd)

				wg1.Add(1)
				go func(said int64, cid int64, header *six910api.Headers) {
					defer wg1.Done()
					ccopc.ShippingAddress = h.API.GetAddress(said, cid, header)
				}(comccotres.ShippingAddressID, comccotres.Cart.CustomerID, hd)

				wg1.Add(1)
				go func(pgwid int64, oid int64, tcode string, amount float64, header *six910api.Headers) {
					defer wg1.Done()
					pgw := h.API.GetPaymentGateway(pgwid, header)
					sp := h.API.GetStorePlugin(pgw.StorePluginsID, header)
					var pm PaymentMethod
					pm.Name = sp.PluginName
					pm.PaymentGateway = pgw
					ccopc.PaymentMethod = &pm
					var trans sdbi.OrderTransaction
					trans.Gwid = pgw.ID
					trans.Method = sp.PluginName
					trans.OrderID = oid
					trans.DateEntered = time.Now()
					trans.ReferenceNumber = tcode
					trans.Amount = amount
					trans.ResponseCode = "200"
					trans.ResponseMessage = "success"
					trans.Success = true
					trans.TransactionID = tcode
					trans.Type = sp.PluginName
					tres := h.API.AddOrderTransaction(&trans, header)
					h.Log.Debug("transaction res: ", *tres)
				}(comccotres.PaymentGatewayID, odrRes.Order.ID, transactionCode, odrRes.Order.Total, hd)

				//----

				// if strings.Contains(strings.ToLower(pm.Name), "paypal") {
				// 	h.Log.Debug("Using PayPay Gateway")
				// 	ccop.PayPalPayment = true
				// }
				ccopc.OrderNumber = odrRes.Order.OrderNumber
				ccopc.OrderInfo = h.CompanyName
				ccopc.CustomerCart = comccotres
				// ccopc.PaymentMethod = &pm
				//ccopc.ShippingMethod = sm
				//ccopc.BillingAddress = h.API.GetAddress(comccotres.BillingAddressID, comccotres.Cart.CustomerID, hd)
				//ccopc.ShippingAddress = h.API.GetAddress(comccotres.ShippingAddressID, comccotres.Cart.CustomerID, hd)
				ccopc.Subtotal = fmt.Sprintf("%.2f", comccotres.Subtotal)
				ccopc.ShippingHandling = fmt.Sprintf("%.2f", comccotres.ShippingHandling)
				ccopc.InsuranceCost = fmt.Sprintf("%.2f", comccotres.InsuranceCost)
				ccopc.Taxes = fmt.Sprintf("%.2f", comccotres.Taxes)
				ccopc.Total = fmt.Sprintf("%.2f", comccotres.Total)

				_, csspg := h.CSSService.GetPageCSS("pageCss")
				h.Log.Debug("PageBody: ", *csspg)
				ccopc.PageBody = csspg

				ml := h.MenuService.GetMenuList()
				h.getCartTotal(cocod, ml, hd)
				ccopc.MenuList = ml

				h.Log.Debug("MenuList", *ccopc.MenuList)

				cisuc, cicont := h.ContentService.GetContent(shoppingCartContent3)
				if cisuc {
					ccopc.Content = cicont
				} else {
					var ct conts.Content
					ccopc.Content = &ct
				}

				if h.MailSenderAddress != "" {
					var sellerMail mll.Mailer
					sellerMail.Subject = h.MailSubjectOrderReceived
					sellerMail.Body = fmt.Sprintf(h.MailBodyOrderReceived, odrRes.Order.OrderNumber, odrRes.Order.CustomerName)
					str := h.API.GetStore(h.StoreName, h.LocalDomain, hd)
					sellerMail.Recipients = []string{str.Email}
					sellerMail.SenderAddress = h.MailSenderAddress

					sellerSendSuc := h.MailSender.SendMail(&sellerMail)
					h.Log.Debug("sendSuc  to seller: ", sellerSendSuc)

					var buyerMail mll.Mailer
					buyerMail.Subject = fmt.Sprintf(h.MailSubjectOrderProcessing, h.CompanyName, odrRes.Order.OrderNumber)
					odridstr := strconv.FormatInt(odrRes.Order.ID, 10)
					var olnk = "<a href='/viewCustomerOrder/" + odridstr + ">" + odrRes.Order.OrderNumber + "</a>"
					buyerMail.Body = fmt.Sprintf(h.MailBodyOrderProcessing, odrRes.Order.CustomerName, olnk)
					//buystr := h.API.GetStore(h.StoreName, h.LocalDomain, hd)
					buyerMail.Recipients = []string{comccotres.CustomerAccount.User.Username}
					buyerMail.SenderAddress = h.MailSenderAddress

					buyerSendSuc := h.MailSender.SendMail(&buyerMail)
					h.Log.Debug("sendSuc to buyer: ", buyerSendSuc)
				}

				wg1.Wait()

				ecc := h.getCustomerCart(cocod)
				var wg sync.WaitGroup
				for _, ci := range *ecc.Items {
					wg.Add(1)
					go func(id int64, pid int64, cid int64, qty int64, header *six910api.Headers) {
						defer wg.Done()
						prd := h.API.GetProductByID(pid, header)
						prd.Stock -= qty
						h.Log.Debug("product after -=: ", *prd)
						upres := h.API.UpdateProductQuantity(prd, header)
						h.Log.Debug("update product after -=: ", *upres)
						h.API.DeleteCartItem(id, pid, cid, header)
					}(ci.ID, ci.ProductID, ci.CartID, ci.Quantity, hd)
				}
				wg.Wait()
				ecc.CartView = nil
				//ecc.CustomerAccount = nil
				ecc.Items = nil
				//ecc.Cart = nil
				h.storeCustomerCart(ecc, cocod, w, r)
				h.Templates.ExecuteTemplate(w, checkoutReceiptPage, &ccopc)
			} else {
				http.Redirect(w, r, customerOrderListView, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}
