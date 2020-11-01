package managers

import (
	"math"
	"strconv"
	"sync"
	"time"

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

//AddProductToCart AddProductToCart
func (m *Six910Manager) AddProductToCart(cc *CustomerCart, cp *CustomerProduct, hd *api.Headers) *CustomerCart {
	m.Log.Debug("cc : ", cc)
	var rtn *CustomerCart = cc
	if rtn == nil {
		rtn = new(CustomerCart)
	}
	var cart = rtn.Cart //*sdbi.Cart
	m.Log.Debug("customerId : ", cp.CustomerID)
	m.Log.Debug("cp cart : ", cp.Cart)
	m.Log.Debug("cp quantity : ", cp.Quantity)
	if cp.CustomerID != 0 && cp.Cart == nil {
		cart = m.API.GetCart(cp.CustomerID, hd)
	} else if cp.Cart != nil {
		cart = cp.Cart
	}
	m.Log.Debug("cart in add prod to cart: ", cart)

	if cart == nil || cart.ID == 0 {
		m.Log.Debug("cart nil or cartId = 0: ")
		var nc sdbi.Cart
		nc.StoreID = cp.StoreID
		nc.CustomerID = cp.CustomerID
		res := m.API.AddCart(&nc, hd)
		if res.Success {
			nc.ID = res.ID
			cart = &nc
		}
	}
	if cart != nil && cart.ID != 0 {
		m.Log.Debug("cart in add prod to cart: ", *cart)
		m.Log.Debug("cart.ID: ", cart.ID)
		prod := m.API.GetProductByID(cp.ProductID, hd)
		//citm := m.API.GetCartItem(cart.ID, prod.ID, hd)
		m.Log.Debug("prod stock : ", prod.Stock)
		m.Log.Debug("cp.Quantity : ", cp.Quantity)
		m.Log.Debug("cart.ID : ", cart.ID)
		m.Log.Debug("prod.ID : ", prod.ID)
		//m.Log.Debug("citm : ", cp.CartItem.Quantity)

		m.Log.Debug("cp.Quantity before adjust : ", cp.Quantity)
		if cp.CartItem != nil && (cp.Quantity+cp.CartItem.Quantity) > prod.Stock {
			cp.Quantity = 0
		}
		m.Log.Debug("cp.Quantity : ", cp.Quantity)
		var ci sdbi.CartItem
		ci.CartID = cart.ID
		ci.ProductID = cp.ProductID
		ci.Quantity = cp.Quantity
		m.Log.Debug("CartItem: ", ci)
		res := m.API.AddCartItem(&ci, cp.CustomerID, hd)
		m.Log.Debug("cart add res: ", *res)
		if res.Success {
			rtn.Cart = cart
			rtn.Items = m.API.GetCartItemList(cart.ID, cp.CustomerID, hd)
			//m.Log.Debug("cart Items: ", *rtn.Items)
		}
	}
	return rtn
}

//ViewCart ViewCart
func (m *Six910Manager) ViewCart(cc *CustomerCart, hd *api.Headers) *CartView {
	var rtn CartView
	var wg sync.WaitGroup
	var itemchan = make(chan *CartViewItem, len(*cc.Items))
	for i := range *cc.Items {
		wg.Add(1)
		go func(cItem *sdbi.CartItem, ichan chan *CartViewItem, header *api.Headers) {
			m.Log.Debug("in goroutine :", cItem.ProductID)
			defer wg.Done()
			var cvi CartViewItem
			prod := m.API.GetProductByID(cItem.ProductID, header)
			//m.Log.Debug("in goroutine prod:", *prod)
			cvi.ProductName = prod.Name
			cvi.ProductID = prod.ID
			cvi.Desc = prod.ShortDesc
			cvi.Image = prod.Thumbnail
			cvi.Stock = prod.Stock
			cvi.Quantity = cItem.Quantity
			if prod.SalePrice != 0 {
				cvi.Price = prod.SalePrice
			} else {
				cvi.Price = prod.Price
			}
			cvi.Total = math.Round((cvi.Price*float64(cvi.Quantity))*100) / 100
			m.Log.Debug("in goroutine cvi.Total:", cvi.Total)
			ichan <- &cvi
		}(&(*cc.Items)[i], itemchan, hd)
	}
	wg.Wait()
	close(itemchan)
	var cviList []*CartViewItem
	for ci := range itemchan {
		cviList = append(cviList, ci)
		rtn.Total = math.Round((rtn.Total+ci.Total)*100) / 100
		m.Log.Debug("ci:", *ci)
	}
	m.Log.Debug("rtn.Total:", rtn.Total)
	rtn.Items = &cviList
	m.Log.Debug("rtn:", rtn)

	return &rtn
}

//UpdateProductToCart UpdateCart
func (m *Six910Manager) UpdateProductToCart(cc *CustomerCart, cp *CustomerProductUpdate, hd *api.Headers) *CustomerCart {
	var rtn *CustomerCart = cc
	if cp.Cart != nil && cp.CartItem != nil {
		var res *api.Response
		if cp.CartItem.Quantity == 0 {
			res = m.API.DeleteCartItem(cp.CartItem.ID, cp.CartItem.ProductID, cp.CartItem.CartID, hd)
		} else {
			res = m.API.UpdateCartItem(cp.CartItem, cp.CustomerID, hd)
		}
		if res.Success {
			rtn.Cart = cp.Cart
			rtn.Items = m.API.GetCartItemList(cp.Cart.ID, cp.CustomerID, hd)
		}
	}
	return rtn
}

//CheckOut CheckOut
func (m *Six910Manager) CheckOut(cart *CustomerCart, hd *api.Headers) *CustomerOrder {
	var rtn *CustomerOrder
	if cart.CustomerAccount.Customer.ID != 0 && cart.CustomerAccount.User.Enabled {
		// check out with logged in user
		rtn = m.completeOrder(cart, hd)
	} else {
		//user not logged in
		suc, cc := m.CreateCustomerAccount(cart.CustomerAccount, hd)
		if suc && cc != nil && cc.Customer != nil {
			rtn = m.completeOrder(cart, hd)
		}
	}
	return rtn
}

//CalculateCartTotals CalculateCartTotals
func (m *Six910Manager) CalculateCartTotals(cart *CustomerCart, hd *api.Headers) *CustomerCart {
	// do all calculation lookups and calculations here
	cv := cart.CartView
	if cv != nil {
		cart.Subtotal = cv.Total
		sm := m.API.GetShippingMethod(cart.ShippingMethodID, hd)
		cart.ShippingHandling = math.Round((sm.Cost+sm.Handling)*100) / 100 //sm.Cost + sm.Handling
		if cart.InsuranceID != 0 {
			ins := m.API.GetInsurance(cart.InsuranceID, hd)
			cart.InsuranceCost = ins.Cost
		}
		sad := m.API.GetAddress(cart.ShippingAddressID, cart.Cart.CustomerID, hd)
		trs := m.API.GetTaxRate(sad.Country, sad.State, hd)
		var tr *sdbi.TaxRate
		if len(*trs) > 1 {
			for i := range *trs {
				if sad.Zip >= (*trs)[i].ZipStart && sad.Zip <= (*trs)[i].ZipEnd {
					tr = &(*trs)[i]
				}
			}
		} else if len(*trs) == 1 {
			tr = &(*trs)[0]
		}

		if tr == nil && len(*trs) > 1 {
			for i := range *trs {
				if (*trs)[i].ZipStart == "" && (*trs)[i].ZipEnd == "" {
					tr = &(*trs)[i]
					break
				}
			}
		}
		if tr.PercentRate != 0 {
			var hTax float64
			var sTax float64
			if tr.IncludeHandling {
				htrate := sm.Handling * (float64(tr.PercentRate) / float64(100))
				hTax = math.Round((htrate)*100) / 100
			}
			if tr.IncludeShipping {
				strate := sm.Cost * (float64(tr.PercentRate) / float64(100))
				sTax = math.Round((strate)*100) / 100
			}
			tax := cart.Subtotal * (float64(tr.PercentRate) / float64(100))
			cart.Taxes = math.Round((hTax+sTax+tax)*100) / 100
		}
		cart.Total = math.Round((cart.InsuranceCost+cart.ShippingHandling+cart.Subtotal+cart.Taxes)*100) / 100
	}
	return cart
}

func (m *Six910Manager) completeOrder(cart *CustomerCart, hd *api.Headers) *CustomerOrder {
	var rtn CustomerOrder
	var badd sdbi.Address
	var sadd sdbi.Address
	for _, a := range *cart.CustomerAccount.Addresses {
		if a.Type == billingAddressType {
			badd = a
		} else if a.Type == shippingAddressType {
			sadd = a
		}
	}
	var odr sdbi.Order
	odr.BillingAddress = badd.Address + ", " + badd.City + " " + badd.State + " " + badd.Zip
	odr.BillingAddressID = badd.ID
	odr.CustomerID = cart.CustomerAccount.Customer.ID
	odr.CustomerName = cart.CustomerAccount.Customer.FirstName + " " + cart.CustomerAccount.Customer.LastName
	odr.Insurance = cart.InsuranceCost
	odr.OrderDate = time.Now()
	odr.OrderNumber = m.generateOrderNumber()
	odr.OrderType = cart.OrderType
	odr.Pickup = cart.Pickup
	odr.ShippingAddress = sadd.Address + ", " + sadd.City + " " + sadd.State + " " + sadd.Zip
	odr.ShippingAddressID = sadd.ID
	odr.ShippingHandling = cart.ShippingHandling
	odr.Status = orderStatusProcessing
	odr.Subtotal = cart.Subtotal
	odr.Taxes = cart.Taxes
	odr.Total = cart.Total
	odr.Username = cart.CustomerAccount.User.Username

	ores := m.API.AddOrder(&odr, hd)
	if ores.Success && ores.ID != 0 {
		rtn.Order = &odr
		rtn.Cart = cart.Cart
		rtn.CustomerAccount = cart.CustomerAccount
		oisuc, oires := m.processOrderItems(cart.Items, ores.ID, hd)
		rtn.Items = oires
		if oisuc && cart.Comment != "" {
			var ocmt sdbi.OrderComment
			ocmt.Comment = cart.Comment
			ocmt.OrderID = ores.ID
			ocmt.Username = cart.CustomerAccount.User.Username
			cres := m.API.AddOrderComments(&ocmt, hd)
			if cres.Success && cres.ID != 0 {
				rtn.Comments = m.API.GetOrderCommentList(ores.ID, hd)
				rtn.Success = true
			}
		} else if oisuc {
			rtn.Success = true
		}
	}
	return &rtn
}

func (m *Six910Manager) processOrderItems(ois *[]sdbi.CartItem, orderID int64, hd *api.Headers) (bool, *[]sdbi.OrderItem) {
	m.Log.Debug("in processOrderItems")
	var rtn = true
	var rtnoi []sdbi.OrderItem
	var wg sync.WaitGroup
	oiresults := make(chan *OrderItemResults, len(*ois))
	for _, ci := range *ois {
		wg.Add(1)
		var oi sdbi.OrderItem
		oi.OrderID = orderID
		oi.ProductID = ci.ProductID
		oi.Quantity = ci.Quantity
		//make call to product to get rest of details
		prod := m.API.GetProductByID(ci.ProductID, hd)
		if prod.Stock == 0 {
			oi.BackOrdered = true
		}
		oi.Dropship = prod.Dropship
		oi.ProductName = prod.Name
		oi.ProductShortDesc = prod.ShortDesc
		go func(ioi *sdbi.OrderItem, ihd *api.Headers, reslt chan *OrderItemResults) {
			m.Log.Debug("in goroutine :", ioi.ProductID)
			defer wg.Done()
			oires := m.API.AddOrderItem(ioi, ihd)
			var oir OrderItemResults
			ioi.ID = oires.ID
			oir.OrderItem = ioi
			oir.Resp = oires
			reslt <- &oir
		}(&oi, hd, oiresults)
	}
	m.Log.Debug("before wait")
	wg.Wait()
	m.Log.Debug("after wait")
	close(oiresults)
	for oir := range oiresults {
		if !oir.Resp.Success || oir.OrderItem.ID == 0 {
			rtn = false
		} else {
			rtnoi = append(rtnoi, *oir.OrderItem)
		}
	}
	return rtn, &rtnoi
}

func (m *Six910Manager) generateOrderNumber() string {
	m.mu.Lock()
	defer m.mu.Unlock()
	var rtn string

	unixNano := time.Now().UnixNano()
	umillisec := unixNano / 1000000

	rtn = "OD-" + strconv.FormatInt(umillisec, 10)

	return rtn
}
