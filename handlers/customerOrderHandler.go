package handlers

import (
	"net/http"
	"strconv"
	"sync"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
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

//OrderViewPage OrderViewPage
type OrderViewPage struct {
	Order    *sdbi.Order
	Items    *[]sdbi.OrderItem
	Comments *[]sdbi.OrderComment
	PageBody *csssrv.PageCSS
	MenuList *[]musrv.Menu
	Content  *conts.Content
	//meta data
	HeaderData      *HeaderData
	CustomerName    string
	PaymentMethod   string
	ShippingMethod  string
	ShippingAddress string
	BillingAddress  string
}

//CustomerOrderItem CustomerOrderItem
type CustomerOrderItem struct {
	ID               int64  `json:"id"`
	Quantity         int64  `json:"quantity"`
	BackOrdered      bool   `json:"backOrdered"`
	Dropship         bool   `json:"dropship"`
	ProductName      string `json:"productName"`
	ProductShortDesc string `json:"productShortDesc"`
	ProductID        int64  `json:"productId"`
	OrderID          int64  `json:"orderId"`
	Desc             string `json:"desc"`
	Image            string `json:"image"`
}

//ViewCustomerOrder ViewCustomerOrder
func (h *Six910Handler) ViewCustomerOrder(w http.ResponseWriter, r *http.Request) {
	vorps, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(vorps) {
			vodrvars := mux.Vars(r)
			vodridstr := vodrvars["id"]
			vodrid, _ := strconv.ParseInt(vodridstr, 10, 64)

			h.Log.Debug("vodrid: ", vodrid)
			hd := h.getHeader(vorps)
			var ovpage OrderViewPage

			var wg sync.WaitGroup

			wg.Add(1)
			func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				ovpage.Order = h.API.GetOrder(oid, header)
			}(vodrid, hd)

			wg.Add(1)
			func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				ovpage.Items = h.API.GetOrderItemList(oid, header)
			}(vodrid, hd)

			wg.Add(1)
			func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				ovpage.Comments = h.API.GetOrderCommentList(oid, header)
			}(vodrid, hd)

			wg.Wait()

			ptranList := h.API.GetOrderTransactionList(ovpage.Order.ID, hd)
			if len(*ptranList) > 0 {
				ovpage.PaymentMethod = (*ptranList)[0].Method
			}

			ovpage.CustomerName = ovpage.Order.CustomerName

			ovpage.ShippingMethod = ovpage.Order.ShippingMethodName
			ovpage.BillingAddress = ovpage.Order.BillingAddress
			ovpage.ShippingAddress = ovpage.Order.ShippingAddress

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			ovpage.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(vorps, ml, hd)
			ovpage.MenuList = ml

			h.Log.Debug("MenuList", *ovpage.MenuList)

			cisuc, cicont := h.ContentService.GetContent(orderContent)
			if cisuc {
				ovpage.Content = cicont
			} else {
				var ct conts.Content
				ovpage.Content = &ct
			}

			h.Log.Debug("ovpage: ", ovpage)
			h.Templates.ExecuteTemplate(w, customerOrderPage, &ovpage)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//ViewCustomerOrderList ViewCustomerOrderList
func (h *Six910Handler) ViewCustomerOrderList(w http.ResponseWriter, r *http.Request) {
	vorrls, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreCustomerLoggedIn(vorrls) {
			// var cid int64
			// //var cidi int
			// fcid := vorrls.Values["customerId"]
			// if fcid != nil {
			// 	cid = fcid.(int64)
			// 	//cid = int64(cidi)
			// }
			comccotres := h.getCustomerCart(vorrls)
			cid := comccotres.Cart.CustomerID
			h.Log.Debug("cid: ", cid)
			hd := h.getHeader(vorrls)

			odlst := h.API.GetOrderList(cid, hd)
			var opage CartPage
			var newOdrLst []sdbi.Order

			for i := len(*odlst) - 1; i >= 0; i-- {
				newOdrLst = append(newOdrLst, (*odlst)[i])
			}
			opage.OrderList = &newOdrLst

			_, csspg := h.CSSService.GetPageCSS("pageCss")
			h.Log.Debug("PageBody: ", *csspg)
			opage.PageBody = csspg

			ml := h.MenuService.GetMenuList()
			h.getCartTotal(vorrls, ml, hd)
			opage.MenuList = ml

			h.Log.Debug("MenuList", *opage.MenuList)

			cisuc, cicont := h.ContentService.GetContent(orderListContent)
			if cisuc {
				opage.Content = cicont
			} else {
				var ct conts.Content
				opage.Content = &ct
			}

			h.Log.Debug("odlst: ", odlst)
			h.Templates.ExecuteTemplate(w, customerOrderListPage, &opage)

		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}
