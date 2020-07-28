package managers

import (
	b64 "encoding/base64"

	api "github.com/Ulbora/Six910API-Go"
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

//StoreAdminLogin StoreAdminLogin
func (m *Six910Manager) StoreAdminLogin(u *api.User, hd *api.Headers) (bool, *api.User) {
	var suc bool
	var rtn api.User

	sEnca := b64.StdEncoding.EncodeToString([]byte(u.Username + ":" + u.Password))
	m.Log.Debug("sEnc: ", sEnca)

	hd.Set("Authorization", "Basic "+sEnca)

	//YWRtaW46YWRtaW4=

	usra := m.API.GetUser(u, hd)
	m.Log.Debug("usr: ", *usra)
	if usra.Enabled && usra.Username == u.Username && usra.Role == storeAdmin {
		suc = true
		rtn.Enabled = usra.Enabled
		rtn.Role = usra.Role
		rtn.StoreID = usra.StoreID
		rtn.Username = usra.Username
	}
	m.Log.Debug("rtn: ", rtn)
	return suc, &rtn
}

//StoreAdminChangePassword StoreAdminChangePassword
func (m *Six910Manager) StoreAdminChangePassword(u *api.User, hd *api.Headers) (bool, *api.User) {
	var suc bool
	var rtnac api.User
	sEncac := b64.StdEncoding.EncodeToString([]byte(u.Username + ":" + u.OldPassword))
	m.Log.Debug("sEnc: ", sEncac)

	hd.Set("Authorization", "Basic "+sEncac)

	//YWRtaW46YWRtaW4=

	usrac := m.API.GetUser(u, hd)
	m.Log.Debug("usr: ", *usrac)
	if usrac.Enabled && usrac.Username == u.Username && usrac.Role == storeAdmin {
		res := m.API.UpdateUser(u, hd)
		suc = res.Success
		rtnac.Enabled = usrac.Enabled
		rtnac.Role = usrac.Role
		rtnac.StoreID = usrac.StoreID
		rtnac.Username = usrac.Username
	}
	m.Log.Debug("rtn: ", rtnac)
	return suc, &rtnac
}
