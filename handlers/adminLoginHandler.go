package handlers

import (
	"net/http"

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
