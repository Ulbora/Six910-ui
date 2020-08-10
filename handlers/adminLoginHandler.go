package handlers

import (
	"net/http"
	"strconv"

	b64 "encoding/base64"

	userv "github.com/Ulbora/Six910-ui/usersrv"
	api "github.com/Ulbora/Six910API-Go"
	oauth2 "github.com/Ulbora/go-oauth2-client"
)

// import (
// 	api "github.com/Ulbora/Six910API-Go"
// 	sdbi "github.com/Ulbora/six910-database-interface"
// )

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
func (h *Six910Handler) StoreAdminLogin(w http.ResponseWriter, r *http.Request) {
	if !h.OAuth2Enabled {
		loginErr := r.URL.Query().Get("error")
		var lge LoginError
		lge.Error = loginErr
		h.Log.Debug("in login----")
		h.AdminTemplates.ExecuteTemplate(w, adminloginPage, &lge)
	} else {
		h.authorize(w, r)
	}
}

//StoreAdminLoginNonOAuthUser StoreAdminLoginNonOAuthUser
func (h *Six910Handler) StoreAdminLoginNonOAuthUser(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		username := r.FormValue("username")
		password := r.FormValue("password")
		sEnccl := b64.StdEncoding.EncodeToString([]byte(username + ":" + password))
		h.Log.Debug("sEnc: ", sEnccl)

		var hd api.Headers
		hd.Set("Authorization", "Basic "+sEnccl)
		//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

		h.Log.Debug("username", username)
		h.Log.Debug("password", password)

		var u api.User
		u.Username = username

		var loginSuc bool

		usrcl := h.API.GetUser(&u, &hd)
		h.Log.Debug("usr: ", *usrcl)
		if usrcl.Enabled && usrcl.Username == u.Username && usrcl.Role == storeAdmin {
			loginSuc = true
			h.Log.Debug("loginSuc", loginSuc)
		}
		h.Log.Debug("login suc", loginSuc)
		if loginSuc {
			//if lari.ResponseType == codeRespType || lari.ResponseType == tokenRespType {
			s.Values["loggedIn"] = true
			//s.Values["user"] = username
			serr := s.Save(r, w)
			h.Log.Debug("serr", serr)
			//session, sserr := store.Get(r, "temp-name")
			//fmt.Println("sserr", sserr)
			//session.Store()
			//session.Options.Path = "/oauth/"
			//session.Values["loggedIn"] = true
			//fmt.Println("store", session.Store())
			//session.Save(r, w)

			//clintStr := strconv.FormatInt(lari.ClientID, 10)
			http.Redirect(w, r, adminIndex, http.StatusFound)
			//} else {
			//http.Redirect(w, r, invalidGrantErrorURL, http.StatusFound)
			//}
		} else {
			http.Redirect(w, r, adminLoginFailedURL, http.StatusFound)
		}

	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

//StoreAdminChangePassword StoreAdminChangePassword
func (h *Six910Handler) StoreAdminChangePassword(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Values["loggedIn"]
		h.Log.Debug("loggedIn in backups: ", loggedInAuth)
		if loggedInAuth == true {
			h.AdminTemplates.ExecuteTemplate(w, adminChangePwPage, nil)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminChangeUserPassword StoreAdminChangeUserPassword
func (h *Six910Handler) StoreAdminChangeUserPassword(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	if suc {
		loggedIn := s.Values["userLoggenIn"]
		token := h.token
		h.Log.Debug("user update pw Logged in: ", loggedIn)

		if loggedIn == nil || loggedIn.(bool) == false || token == nil {
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

//StoreAdminLogout StoreAdminLogout
func (h *Six910Handler) StoreAdminLogout(w http.ResponseWriter, r *http.Request) {
	h.token = nil
	cookie := &http.Cookie{
		Name:   "goauth2-ui",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	cookie2 := &http.Cookie{
		Name:   "goauth2",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie2)
	http.Redirect(w, r, "/", http.StatusFound)
}

//StoreAdminHandleToken StoreAdminHandleToken
func (h *Six910Handler) StoreAdminHandleToken(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	h.Log.Debug("handle token")
	if state == h.ClientCreds.AuthCodeState {

		h.Auth.SetOauthHost(h.OauthHost)
		h.Auth.SetClientID(h.ClientCreds.AuthCodeClient)
		h.Auth.SetSecret(h.ClientCreds.AuthCodeSecret)
		h.Auth.SetCode(code)
		h.Auth.SetRedirectURI(h.getRedirectURI(r, authCodeRedirectURI))

		h.Log.Debug("getting token")

		resp := h.Auth.AuthCodeToken()
		h.Log.Debug("token resp: ", *resp)

		h.Log.Debug("token len: ", len(resp.AccessToken))

		h.Log.Debug("token : ", resp.AccessToken)
		if resp != nil && resp.AccessToken != "" {

			s, suc := h.getSession(r)
			if suc {
				h.Log.Debug("userLoggenIn : ", true)
				s.Values["userLoggenIn"] = true

				h.token = resp

				err := s.Save(r, w)
				h.Log.Debug(err)
				http.Redirect(w, r, "/clients", http.StatusFound)
			}
		}
	}
}

func (h *Six910Handler) authorize(w http.ResponseWriter, r *http.Request) bool {
	h.Log.Debug("in authorize")

	var a oauth2.AuthCodeAuthorize
	a.ClientID = h.ClientCreds.AuthCodeClient // h.getAuthCodeClient()
	a.OauthHost = h.OauthHost                 // getOauthRedirectHost()
	a.RedirectURI = h.getRedirectURI(r, authCodeRedirectURI)
	a.Scope = "write"
	a.State = h.ClientCreds.AuthCodeState // authCodeState
	a.Res = w
	a.Req = r

	h.Log.Debug("a: ", a)
	resp := a.AuthCodeAuthorizeUser()

	h.Log.Debug("Resp: ", resp)
	return resp
}

func (h *Six910Handler) getRedirectURI(req *http.Request, path string) string {
	var scheme = req.URL.Scheme
	var serverHost string
	if scheme != "" {
		serverHost = req.URL.String() + path
	} else {
		serverHost = h.SchemeDefault + req.Host + path
	}
	h.Log.Debug("login redirect url: ", serverHost)
	return serverHost
}
