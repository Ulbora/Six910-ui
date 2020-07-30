package managers

import (
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
func (m *Six910Manager) AddProductToCart(cp *CustomerProduct, hd *api.Headers) *CustomerCart {
	var rtn CustomerCart
	var cart *sdbi.Cart
	if cp.CustomerID != 0 {
		cart = m.API.GetCart(cp.CustomerID, hd)
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
		var ci sdbi.CartItem
		ci.CartID = cart.ID
		ci.ProductID = cp.ProductID
		ci.Quantity = cp.Quantity
		res := m.API.AddCartItem(&ci, cp.CustomerID, hd)
		m.Log.Debug("cart add res: ", *res)
		if res.Success {
			rtn.Cart = cart
			rtn.Items = m.API.GetCartItemList(cart.ID, cp.CustomerID, hd)
		}
	}
	return &rtn
}
