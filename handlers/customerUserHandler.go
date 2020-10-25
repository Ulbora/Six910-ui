package handlers

import (
	"net/http"

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

//CustomerLoginPage CustomerLoginPage
func (h *Six910Handler) CustomerLoginPage(w http.ResponseWriter, r *http.Request) {
	clinss, suc := h.getUserSession(r)
	h.Log.Debug("session suc", clinss)
	if suc {
		var locp CustomerPage
		h.Templates.ExecuteTemplate(w, customerLoginPage, locp)
	}
}

//CustomerLogin CustomerLogin
func (h *Six910Handler) CustomerLogin(w http.ResponseWriter, r *http.Request) {
	slin, suc := h.getUserSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		cusername := r.FormValue("username")
		cpassword := r.FormValue("password")

		hd := h.getHeader(slin)

		h.Log.Debug("username", cusername)
		h.Log.Debug("password", cpassword)

		var u api.User
		u.Username = cusername
		u.Password = cpassword

		//var loginSuc bool

		loginSuc, uu := h.Manager.CustomerLogin(&u, hd)
		h.Log.Debug("uu: ", uu)

		h.Log.Debug("login suc", loginSuc)
		if loginSuc {
			//if lari.ResponseType == codeRespType || lari.ResponseType == tokenRespType {
			slin.Values["loggedIn"] = true
			slin.Values["customerUser"] = true
			slin.Values["username"] = cusername
			slin.Values["password"] = cpassword
			slin.Values["customerId"] = uu.CustomerID

			serr := slin.Save(r, w)
			h.Log.Debug("serr", serr)

			http.Redirect(w, r, customerIndexView, http.StatusFound)

		} else {
			http.Redirect(w, r, customerLoginViewFail, http.StatusFound)
		}
	}
}

//CustomerChangePasswordPage CustomerChangePasswordPage
func (h *Six910Handler) CustomerChangePasswordPage(w http.ResponseWriter, r *http.Request) {
	ccpwpuvss, suc := h.getUserSession(r)
	h.Log.Debug("session suc", ccpwpuvss)
	if suc {
		if h.isStoreCustomerLoggedIn(ccpwpuvss) {
			var ucpname string
			ccaemail := ccpwpuvss.Values["username"]
			if ccaemail != nil {
				ucpname = ccaemail.(string)
			}
			hd := h.getHeader(ccpwpuvss)
			h.Log.Debug("ucpname: ", ucpname)
			var u api.User
			u.Username = ucpname
			uu := h.API.GetUser(&u, hd)
			h.Log.Debug("uu: ", *uu)
			h.Templates.ExecuteTemplate(w, customerCreatePage, &uu)
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//CustomerChangePassword CustomerChangePassword
func (h *Six910Handler) CustomerChangePassword(w http.ResponseWriter, r *http.Request) {
	ccpwuvss, suc := h.getUserSession(r)
	h.Log.Debug("session suc", ccpwuvss)
	if suc {
		if h.isStoreCustomerLoggedIn(ccpwuvss) {
			ccpwpwnpw := r.FormValue("password")
			ccpwpwopw := r.FormValue("oldPassword")

			hd := h.getHeader(ccpwuvss)

			var ccpwuemail string
			ccpwaemail := ccpwuvss.Values["username"]
			if ccpwaemail != nil {
				ccpwuemail = ccpwaemail.(string)
			}
			h.Log.Debug("ccpwuemail: ", ccpwuemail)

			var u api.User
			u.Username = ccpwuemail
			u.Password = ccpwpwnpw
			u.OldPassword = ccpwpwopw
			h.Log.Debug("user in change pw: ", u)
			suc, uu := h.Manager.CustomerChangePassword(&u, hd)
			h.Log.Debug("uu: ", uu)
			if suc {
				http.Redirect(w, r, customerInfoView, http.StatusFound)
			} else {
				http.Redirect(w, r, customerInfoViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, customerLoginView, http.StatusFound)
		}
	}
}

//CustomerLogout CustomerLogout
func (h *Six910Handler) CustomerLogout(w http.ResponseWriter, r *http.Request) {
	h.token = nil
	ccookie := &http.Cookie{
		Name:   "Six910-ui",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, ccookie)

	ccookie2 := &http.Cookie{
		Name:   "Six910",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, ccookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}
