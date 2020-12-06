package handlers

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"

	userv "github.com/Ulbora/Six910-ui/usersrv"
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

//StoreAdminChangePassword StoreAdminChangePassword-- Page
func (h *Six910Handler) StoreAdminChangePassword(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			h.AdminTemplates.ExecuteTemplate(w, adminChangePwPage, nil)
		} else {
			h.authorize(w, r)
		}

	}
}

//StoreAdminChangeUserPassword StoreAdminChangeUserPassword-- OAuth Pasword change
func (h *Six910Handler) StoreAdminChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		storeAdminUser := s.Values["storeAdminUser"]
		token := h.token
		h.Log.Debug("user update pw Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || storeAdminUser.(bool) == false || token == nil {
			h.authorize(w, r)
		} else {
			var uu userv.UserPW
			clientID := r.FormValue("clientId")
			h.Log.Debug("user update pw client: ", clientID)
			clientIDD, _ := strconv.ParseInt(clientID, 10, 0)
			uu.ClientID = clientIDD

			username := r.FormValue("username")
			h.Log.Debug("user update pw username: ", username)
			uu.Username = username

			password := r.FormValue("password")
			h.Log.Debug("user update pw password: ", password)
			uu.Password = password

			h.UserService.SetToken(h.token.AccessToken)

			res := h.UserService.UpdateUser(&uu)
			h.Log.Debug("user update pw res: ", *res)
			if res.Success {
				http.Redirect(w, r, adminIndex, http.StatusFound)
			} else {
				http.Redirect(w, r, adminChangePassword, http.StatusFound)
			}
		}
	}
}

//StoreAdminAddAdminUserPage StoreAdminAddAdminUserPage
func (h *Six910Handler) StoreAdminAddAdminUserPage(w http.ResponseWriter, r *http.Request) {
	auss, suc := h.getSession(r)
	h.Log.Debug("session suc in user add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(auss) {
			h.AdminTemplates.ExecuteTemplate(w, addAdminUserPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddAdminUser StoreAdminAddAdminUser
func (h *Six910Handler) StoreAdminAddAdminUser(w http.ResponseWriter, r *http.Request) {
	aauss, suc := h.getSession(r)
	h.Log.Debug("session suc in user add user", suc)
	if suc {
		if h.isStoreAdminLoggedIn(aauss) {

			var suc bool
			if !h.OAuth2Enabled {
				au := h.processUser(r)
				au.Role = storeAdmin
				hd := h.getHeader(aauss)
				res := h.API.AddAdminUser(au, hd)
				suc = res.Success
			} else {
				h.UserService.SetToken(h.token.AccessToken)
				//usrs, _ := h.UserService.GetUser(username, h.ClientCreds.AuthCodeClient)
				var user userv.User
				user.Username = r.FormValue("username")
				user.Password = r.FormValue("password")
				user.ClientID, _ = strconv.ParseInt(h.ClientCreds.AuthCodeClient, 10, 0)
				user.Enabled = true
				user.RoleID = 1
				user.FirstName = r.FormValue("username")    //r.FormValue("fname")
				user.LastName = r.FormValue("username")     //r.FormValue("lname")
				user.EmailAddress = r.FormValue("username") //r.FormValue("email")
				h.Log.Debug("add oauth user:", user)
				oures := h.UserService.AddUser(user)
				suc = oures.Success
				h.Log.Debug("add oauth user suc:", suc)
			}
			if suc {
				http.Redirect(w, r, adminUserList, http.StatusFound)
			} else {
				http.Redirect(w, r, addAdminUserFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAdminUserList StoreAdminAdminUserList
func (h *Six910Handler) StoreAdminAdminUserList(w http.ResponseWriter, r *http.Request) {
	gauls, suc := h.getSession(r)
	h.Log.Debug("session suc in Admin User list view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gauls) {
			if !h.OAuth2Enabled {
				hd := h.getHeader(gauls)
				usl := h.API.GetAdminUsers(hd)
				h.Log.Debug("Admin User  in list", usl)
				h.AdminTemplates.ExecuteTemplate(w, adminUserListPage, &usl)
			} else {
				h.UserService.SetToken(h.token.AccessToken)
				usrs, _ := h.UserService.GetAdminUserList(h.ClientCreds.AuthCodeClient)
				h.Log.Debug("OAuth2 Admin User  in list", *usrs)
				h.AdminTemplates.ExecuteTemplate(w, oauthAdminUserListPage, &usrs)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminCustomerUserList StoreAdminCustomerUserList
func (h *Six910Handler) StoreAdminCustomerUserList(w http.ResponseWriter, r *http.Request) {
	gculs, suc := h.getSession(r)
	h.Log.Debug("session suc in customer User list view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gculs) {
			hd := h.getHeader(gculs)
			usl := h.API.GetCustomerUsers(hd)
			h.Log.Debug("Customer User  in list", usl)
			h.AdminTemplates.ExecuteTemplate(w, customerUserListPage, &usl)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditUserPage StoreAdminEditUserPage
func (h *Six910Handler) StoreAdminEditUserPage(w http.ResponseWriter, r *http.Request) {
	euss, suc := h.getSession(r)
	h.Log.Debug("session suc in user edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(euss) {
			edvars := mux.Vars(r)
			usernm := edvars["username"]
			//role := edvars["role"]
			var useOauth bool
			if h.OAuth2Enabled {
				useOauth = true
			}
			if !useOauth {
				hd := h.getHeader(euss)
				var us api.User
				us.Username = usernm
				user := h.API.GetUser(&us, hd)
				h.AdminTemplates.ExecuteTemplate(w, adminEditUserPage, &user)
			} else {
				h.UserService.SetToken(h.token.AccessToken)
				usrs, _ := h.UserService.GetUser(usernm, h.ClientCreds.AuthCodeClient)
				h.Log.Debug("OAuth2 Admin User ", *usrs)
				h.AdminTemplates.ExecuteTemplate(w, adminEditOAuth2UserPage, &usrs)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditUser StoreAdminEditUser
func (h *Six910Handler) StoreAdminEditUser(w http.ResponseWriter, r *http.Request) {
	edus, suc := h.getSession(r)
	h.Log.Debug("session suc in user edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(edus) {
			username := r.FormValue("username")
			var useOauth bool
			if h.OAuth2Enabled {
				useOauth = true
			}
			var suc bool
			if !useOauth {
				edd := h.processUser(r)
				h.Log.Debug("User update", *edd)
				hd := h.getHeader(edus)
				//add new update for admin use------------------------------
				res := h.API.AdminUpdateUser(edd, hd)
				h.Log.Debug("User update resp", *res)
				h.Log.Debug("edd", *edd)
				h.Log.Debug("edus.Values username", edus.Values["username"])
				suc = res.Success
				if res.Success && edus.Values["username"] != nil && edd.Username == edus.Values["username"].(string) && edd.Password != "" {
					edus.Values["password"] = edd.Password
					serr := edus.Save(r, w)
					h.Log.Debug("serr", serr)
				}
			} else {
				h.UserService.SetToken(h.token.AccessToken)
				//usrs, _ := h.UserService.GetUser(username, h.ClientCreds.AuthCodeClient)
				var user userv.UserPW
				user.Password = r.FormValue("password")
				user.Username = username
				user.ClientID, _ = strconv.ParseInt(h.ClientCreds.AuthCodeClient, 10, 0)
				user.Enabled = true
				oures := h.UserService.UpdateUser(&user)
				suc = oures.Success
				h.Log.Debug("update oauth user suc:", suc)
			}
			if suc {
				http.Redirect(w, r, adminIndex, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditUser+"/"+username, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processUser(r *http.Request) *api.User {
	var u api.User
	u.Username = r.FormValue("username")
	cID := r.FormValue("cid")
	u.CustomerID, _ = strconv.ParseInt(cID, 10, 64)
	u.Password = r.FormValue("password")
	u.Role = r.FormValue("role")
	// enabled := r.FormValue("enabled")
	// u.Enabled, _ = strconv.ParseBool(enabled)
	enabled := r.FormValue("enabled")
	if enabled == "on" {
		u.Enabled = true
	} else {
		u.Enabled = false
	}

	return &u
}
