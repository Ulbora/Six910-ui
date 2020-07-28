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

	sEnc := b64.StdEncoding.EncodeToString([]byte(u.Username + ":" + u.Password))
	m.Log.Debug("sEnc: ", sEnc)

	hd.Set("Authorization", "Basic "+sEnc)

	//YWRtaW46YWRtaW4=

	usr := m.API.GetUser(u, hd)
	m.Log.Debug("usr: ", *usr)
	if usr.Enabled && usr.Username == u.Username && usr.Role == storeAdmin {
		suc = true
		rtn.Enabled = usr.Enabled
		rtn.Role = usr.Role
		rtn.StoreID = usr.StoreID
		rtn.Username = usr.Username
	}
	m.Log.Debug("rtn: ", rtn)
	return suc, &rtn
}
