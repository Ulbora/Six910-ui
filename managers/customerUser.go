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

//CustomerLogin CustomerLogin
func (m *Six910Manager) CustomerLogin(u *api.User, hd *api.Headers) (bool, *api.User) {
	var succl bool
	var rtncl api.User
	sEnccl := b64.StdEncoding.EncodeToString([]byte(u.Username + ":" + u.Password))
	m.Log.Debug("sEnc: ", sEnccl)

	hd.Set("Authorization", "Basic "+sEnccl)

	usrcl := m.API.GetUser(u, hd)
	m.Log.Debug("usr: ", *usrcl)
	if usrcl.Enabled && usrcl.Username == u.Username && usrcl.Role == customerRole {
		succl = true
		rtncl.Enabled = usrcl.Enabled
		rtncl.Role = usrcl.Role
		rtncl.StoreID = usrcl.StoreID
		rtncl.Username = usrcl.Username
		rtncl.CustomerID = usrcl.CustomerID
	}
	m.Log.Debug("rtn: ", rtncl)
	return succl, &rtncl
}

//CustomerChangePassword CustomerChangePassword
func (m *Six910Manager) CustomerChangePassword(u *api.User, hd *api.Headers) (bool, *api.User) {
	var succc bool
	var rtncc api.User
	sEnccc := b64.StdEncoding.EncodeToString([]byte(u.Username + ":" + u.OldPassword))
	m.Log.Debug("sEnc: ", sEnccc)

	hd.Set("Authorization", "Basic "+sEnccc)

	usrcc := m.API.GetUser(u, hd)
	m.Log.Debug("usr: ", *usrcc)
	if usrcc.Enabled && usrcc.Username == u.Username && usrcc.Role == customerRole {
		u.Enabled = usrcc.Enabled
		res := m.API.UpdateUser(u, hd)
		m.Log.Debug("usrcc: ", *usrcc)
		succc = res.Success
		rtncc.Enabled = usrcc.Enabled
		rtncc.Role = usrcc.Role
		rtncc.StoreID = usrcc.StoreID
		rtncc.Username = usrcc.Username
	}
	m.Log.Debug("rtn: ", rtncc)
	return succc, &rtncc
}
