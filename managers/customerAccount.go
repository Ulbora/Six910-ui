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

//CreateCustomerAccount CreateCustomerAccount
func (m *Six910Manager) CreateCustomerAccount(cus *CustomerAccount, hd *api.Headers) (bool, *CustomerAccount) {
	var suc bool
	var rtn *CustomerAccount
	ecus := m.API.GetCustomer(cus.Customer.Email, hd)
	m.Log.Debug("existing customer: ", *ecus)
	if ecus != nil && ecus.ID != 0 {
		suc, rtn = m.processExistingCustomer(ecus, cus, hd)
	} else {
		suc, rtn = m.processNewCustomer(cus, hd)
	}
	return suc, rtn
}

func (m *Six910Manager) compareAddresses(na *sdbi.Address, eadd *[]sdbi.Address) bool {
	var rtn = true
	var found bool
	for i := range *eadd {
		ea := (*eadd)[i]
		if na.Address == ea.Address && na.City == ea.City && na.State == ea.State && na.Zip == ea.Zip &&
			na.County == ea.County && na.Country == ea.Country && na.Type == ea.Type {
			found = true
			break
		}
	}
	if !found {
		rtn = false
	}
	return rtn
}

func (m *Six910Manager) processExistingCustomer(ecus *sdbi.Customer, cus *CustomerAccount, hd *api.Headers) (bool, *CustomerAccount) {
	var suc bool
	var rtn *CustomerAccount
	cus.Customer = ecus
	eadds := m.API.GetAddressList(ecus.ID, hd)
	for i := range *cus.Addresses {
		a := (*cus.Addresses)[i]
		m.Log.Debug("existing customer address: ", a)
		if !m.compareAddresses(&a, eadds) {
			a.CustomerID = ecus.ID
			m.Log.Debug("existing customer adding address: ", a)
			res := m.API.AddAddress(&a, hd)
			m.Log.Debug("existing customer adding address res: ", *res)
		}
	}
	eadds = m.API.GetAddressList(ecus.ID, hd)
	cus.Addresses = eadds

	cus.User.CustomerID = ecus.ID
	fu := m.API.GetUser(cus.User, hd)
	m.Log.Debug("found user: ", *fu)
	if fu != nil && fu.Username != "" && fu.Enabled {
		suc = true
		rtn = cus
	} else if fu != nil && fu.Username == "" {
		cus.User.Enabled = true
		cus.User.Role = customerRole
		m.Log.Debug("adding user: ", *cus.User)
		auseRes := m.API.AddCustomerUser(cus.User, hd)
		if auseRes.Success {
			suc = true
			rtn = cus
		}
	}
	return suc, rtn
}

func (m *Six910Manager) processNewCustomer(cus *CustomerAccount, hd *api.Headers) (bool, *CustomerAccount) {
	var suc bool
	var rtn *CustomerAccount
	var addlst []sdbi.Address
	cres := m.API.AddCustomer(cus.Customer, hd)
	if cres.Success && cres.ID != 0 {
		cus.Customer.ID = cres.ID
		for i := range *cus.Addresses {
			var a = &(*cus.Addresses)[i]
			a.CustomerID = cres.ID
			ares := m.API.AddAddress(a, hd)
			if ares.Success && ares.ID != 0 {
				a.ID = ares.ID
				addlst = append(addlst, *a)
			}
		}
		cus.User.CustomerID = cres.ID
		cus.User.Enabled = true
		cus.User.Role = customerRole
		m.Log.Debug("adding user for new customer: ", *cus.User)
		ures := m.API.AddCustomerUser(cus.User, hd)
		m.Log.Debug("adding user for new customer res: ", *ures)
		if ures.Success {
			cus.Addresses = &addlst
			suc = true
			rtn = cus
		}
	}
	return suc, rtn
}

//UpdateCustomerAccount UpdateCustomerAccount
func (m *Six910Manager) UpdateCustomerAccount(cus *CustomerAccount, hd *api.Headers) bool {
	var rtn bool
	ecus := m.API.GetCustomer(cus.Customer.Email, hd)
	m.Log.Debug("existing customer in update: ", *ecus)
	if ecus.ID == cus.Customer.ID {
		ures := m.API.UpdateCustomer(cus.Customer, hd)
		if ures.Success {
			for i := range *cus.Addresses {
				var add = &(*cus.Addresses)[i]
				if add.ID != 0 {
					ures := m.API.UpdateAddress(add, hd)
					m.Log.Debug("updated address in update customer: ", *ures)
				} else {
					ares := m.API.AddAddress(add, hd)
					m.Log.Debug("add address in update customer: ", *ares)
				}
			}
			eus := m.API.GetUser(cus.User, hd)
			m.Log.Debug("user in update customer: ", *eus)
			if eus.Enabled {
				uusr := m.API.UpdateUser(cus.User, hd)
				m.Log.Debug("user res in update customer: ", *uusr)
				if uusr.Success {
					rtn = true
				}
			}
		}
	}
	return rtn
}
