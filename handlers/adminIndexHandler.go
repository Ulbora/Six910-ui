package handlers

import (
	"net/http"
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

//StoreAdminIndex StoreAdminIndex
func (h *Six910Handler) StoreAdminIndex(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		// loggedInAuth := s.Values["loggedIn"]
		// storeAdminUser := s.Values["storeAdminUser"]
		// h.Log.Debug("loggedIn in backups: ", loggedInAuth)
		// if loggedInAuth == true && storeAdminUser == true {
		if h.isStoreAdminLoggedIn(s) {
			h.AdminTemplates.ExecuteTemplate(w, adminIndexPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
