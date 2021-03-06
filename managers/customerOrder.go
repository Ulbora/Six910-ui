package managers

import (
	api "github.com/Ulbora/Six910API-Go"
	//sdbi "github.com/Ulbora/six910-database-interface"
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

//ViewCustomerOrder ViewCustomerOrder
func (m *Six910Manager) ViewCustomerOrder(orderID int64, cid int64, hd *api.Headers) *CustomerOrder {
	var rtn CustomerOrder
	rtn.Order = m.API.GetOrder(orderID, hd)
	rtn.Items = m.API.GetOrderItemList(orderID, hd)
	rtn.Comments = m.API.GetOrderCommentList(orderID, hd)
	var ca CustomerAccount
	ca.Customer = m.API.GetCustomerID(cid, hd)
	ca.Addresses = m.API.GetAddressList(cid, hd)
	rtn.CustomerAccount = &ca
	rtn.Success = true
	return &rtn
}

//ViewCustomerOrderList ViewCustomerOrderList
func (m *Six910Manager) ViewCustomerOrderList(cid int64, hd *api.Headers) *[]CustomerOrder {
	var rtn []CustomerOrder
	cus := m.API.GetCustomerID(cid, hd)
	cusal := m.API.GetAddressList(cid, hd)
	ol := m.API.GetOrderList(cid, hd)
	for i := range *ol {
		var o = (*ol)[i]
		var co CustomerOrder
		co.Order = &o
		co.Items = m.API.GetOrderItemList(o.ID, hd)
		co.Comments = m.API.GetOrderCommentList(o.ID, hd)
		var ca CustomerAccount
		ca.Customer = cus
		ca.Addresses = cusal
		co.CustomerAccount = &ca
		co.Success = true
		rtn = append(rtn, co)
	}
	return &rtn
}
