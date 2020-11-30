package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	m "github.com/Ulbora/Six910-ui/managers"
	api "github.com/Ulbora/Six910API-Go"
	mll "github.com/Ulbora/go-mail-sender"
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
		pageErr := r.URL.Query().Get("error")
		var locp CustomerPage
		locp.Error = pageErr
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
			slin.Values["userLoggenIn"] = true
			slin.Values["customerUser"] = true
			slin.Values["username"] = cusername
			slin.Values["password"] = cpassword
			slin.Values["customerId"] = uu.CustomerID

			serr := slin.Save(r, w)
			h.Log.Debug("serr", serr)
			cc := h.getCustomerCart(slin)
			h.Log.Debug("cusID uu in login", uu.CustomerID)
			h.Log.Debug("cusID h.getCustomerID in login", h.getCustomerID(slin))
			h.Log.Debug("cc in login", cc)
			if cc != nil {
				h.Log.Debug("cart in login", cc.Cart)
				cc.CustomerAccount = new(m.CustomerAccount)
				cc.CustomerAccount.User = uu
				cc.CustomerAccount.Customer = h.API.GetCustomerID(uu.CustomerID, hd)
				cc.CustomerAccount.Addresses = h.API.GetAddressList(uu.CustomerID, hd)
				cc.Cart = h.API.GetCart(uu.CustomerID, hd)
				if cc.Cart != nil {
					h.Log.Debug("cart in login", *cc.Cart)
					h.Log.Debug("cart item in login", cc.Items)
					itmlst := h.API.GetCartItemList(cc.Cart.ID, uu.CustomerID, hd)
					if cc.Items != nil {
						h.Log.Debug("cart item len in login", len(*cc.Items))
						for _, ci := range *cc.Items {
							var delcartID = ci.CartID
							h.Log.Debug("existing cart item in login", ci)
							h.Log.Debug("existing cart item id in login", ci.ID)
							h.Log.Debug("existing cart item cart id in login", ci.CartID)
							ci.ID = 0
							ci.CartID = cc.Cart.ID
							ciadd := h.API.AddCartItem(&ci, uu.CustomerID, hd)
							h.Log.Debug("add cart item in login", ciadd)
							h.API.DeleteCart(delcartID, 0, hd)
						}
						for i := range *itmlst {
							*cc.Items = append(*cc.Items, (*itmlst)[i])
						}
					} else {
						cc.Items = itmlst
					}
					h.Log.Debug("cart items in login", cc.Items)
				}

				h.storeCustomerCart(cc, slin, w, r)
			}

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
			ccpwpwocid := r.FormValue("customerId")
			cid, _ := strconv.ParseInt(ccpwpwocid, 10, 64)

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
			u.CustomerID = cid
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
	cloutss, suc := h.getUserSession(r)
	h.Log.Debug("session suc", cloutss)
	if suc {
		if h.isStoreCustomerLoggedIn(cloutss) {
			cloutss.Values["userLoggenIn"] = false
			serr := cloutss.Save(r, w)
			h.Log.Debug("serr", serr)
			cc := h.getCustomerCart(cloutss)
			cc.CustomerAccount = nil
			cc.Items = nil
			cc.Cart = nil
			h.storeCustomerCart(cc, cloutss, w, r)

		}
	}
	// h.token = nil
	// ccookie := &http.Cookie{
	// 	Name:   "Six910-ui-user",
	// 	Value:  "",
	// 	Path:   "/",
	// 	MaxAge: -1,
	// }
	// http.SetCookie(w, ccookie)

	// ccookie2 := &http.Cookie{
	// 	Name:   "Six910",
	// 	Value:  "",
	// 	Path:   "/",
	// 	MaxAge: -1,
	// }
	// http.SetCookie(w, ccookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}

//CustomerResetPasswordPage CustomerResetPasswordPage
func (h *Six910Handler) CustomerResetPasswordPage(w http.ResponseWriter, r *http.Request) {
	rpclinss, suc := h.getUserSession(r)
	h.Log.Debug("session suc", rpclinss)
	if suc {
		h.Templates.ExecuteTemplate(w, customerResetPasswordPage, nil)
	}
}

//CustomerResetPassword CustomerResetPassword
func (h *Six910Handler) CustomerResetPassword(w http.ResponseWriter, r *http.Request) {
	rpwpuvss, suc := h.getUserSession(r)
	h.Log.Debug("session suc", rpwpuvss)
	if suc {

		cusername := r.FormValue("username")

		hd := h.getHeader(rpwpuvss)

		h.Log.Debug("username", cusername)

		var u api.User
		u.Username = cusername

		upr := h.API.ResetCustomerUserPassword(&u, hd)

		h.Log.Debug("upr: ", *upr)
		if upr.Success && h.MailSenderAddress != "" {
			var buyerMail mll.Mailer
			buyerMail.Subject = fmt.Sprintf(h.MailSubjectPasswordReset)
			buyerMail.Body = fmt.Sprintf(h.MailBodyPasswordReset, h.Six910SiteURL, upr.Password)
			buyerMail.Recipients = []string{upr.Username}
			buyerMail.SenderAddress = h.MailSenderAddress

			buyerSendSuc := h.MailSender.SendMail(&buyerMail)
			h.Log.Debug("reset pw sendSuc to buyer: ", buyerSendSuc)
			h.Log.Debug("reset pw from: ", h.MailSenderAddress)

		}

		http.Redirect(w, r, customerIndexView, http.StatusFound)

	}
}
