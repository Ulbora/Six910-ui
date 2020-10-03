package handlers

import (
	b64 "encoding/base64"
	"encoding/gob"
	"html/template"
	"net/http"

	lg "github.com/Ulbora/Level_Logger"
	bks "github.com/Ulbora/Six910-ui/bkupsrv"
	conts "github.com/Ulbora/Six910-ui/contsrv"
	imgs "github.com/Ulbora/Six910-ui/imgsrv"
	mails "github.com/Ulbora/Six910-ui/mailsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	users "github.com/Ulbora/Six910-ui/usersrv"
	api "github.com/Ulbora/Six910API-Go"
	ml "github.com/Ulbora/go-mail-sender"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	gs "github.com/Ulbora/go-sessions"
	"github.com/gorilla/sessions"
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

//Six910Handler Six910Handler
type Six910Handler struct {
	Log     *lg.Logger
	Manager m.Manager
	API     api.API

	Session        gs.GoSession
	Templates      *template.Template
	AdminTemplates *template.Template
	Store          *sessions.CookieStore

	//services
	BackupService  bks.BackupService
	ContentService conts.Service
	ImageService   imgs.ImageService
	MailService    mails.MailService
	UserService    users.UserService

	OauthHost     string
	UserHost      string
	SchemeDefault string // = "http://"
	Auth          oauth2.AuthToken
	token         *oauth2.Token
	ClientCreds   *ClientCreds

	BackendURL    string
	StoreName     string
	LocalDomain   string
	APIKey        string
	OAuth2Enabled bool

	MailSender        ml.Sender
	MailSenderAddress string
	MailSubject       string

	ImagePath     string
	ThumbnailPath string
}

//GetNew GetNew
func (h *Six910Handler) GetNew() Handler {
	return h
}

func (h *Six910Handler) getSession(r *http.Request) (*sessions.Session, bool) {
	//fmt.Println("getSession--------------------------------------------------")
	var suc bool
	var srtn *sessions.Session
	if h.Store == nil {
		h.Session.Name = "Six910-ui"
		h.Session.MaxAge = 3600
		h.Store = h.Session.InitSessionStore()
		h.Log.Debug("h.Store : ", h.Store)
		//errors without this
		gob.Register(&m.CustomerCart{})
		//-------gob.Register(&AuthorizeRequestInfo{})
	}
	if r != nil {
		// fmt.Println("secure in getSession", h.Session.Secure)
		// fmt.Println("name in getSession", h.Session.Name)
		// fmt.Println("MaxAge in getSession", h.Session.MaxAge)
		// fmt.Println("SessionKey in getSession", h.Session.SessionKey)

		//h.Session.HTTPOnly = true

		//h.Session.InitSessionStore()
		s, err := h.Store.Get(r, h.Session.Name)
		//s, err := store.Get(r, "temp-name")
		//s, err := store.Get(r, "goauth2")

		loggedInAuth := s.Values["userLoggenIn"]
		//userAuth := s.Values["user"]
		h.Log.Debug("userLoggenIn: ", loggedInAuth)
		//h.Log.Debug("user: ", userAuth)

		//larii := s.Values["authReqInfo"]
		//h.Log.Debug("arii-----login", larii)

		h.Log.Debug("session error in getSession: ", err)
		if err == nil {
			suc = true
			srtn = s
		}
	}
	//fmt.Println("exit getSession--------------------------------------------------")
	return srtn, suc
}

func (h *Six910Handler) getHeader(s *sessions.Session) *api.Headers {
	var hd api.Headers
	storeCustomerUserpa := s.Values["customerUser"]
	if !h.OAuth2Enabled || storeCustomerUserpa == true {
		var sEnccl string
		username := s.Values["username"]
		password := s.Values["password"]
		if username != nil && password != nil {
			sEnccl = b64.StdEncoding.EncodeToString([]byte(username.(string) + ":" + password.(string)))
		}
		h.Log.Debug("sEnc: ", sEnccl)
		hd.Set("Authorization", "Basic "+sEnccl)
	} else {
		hd.Set("Authorization", "Bearer "+h.token.AccessToken)
	}
	return &hd
}

func (h *Six910Handler) isStoreAdminLoggedIn(s *sessions.Session) bool {
	var rtn bool
	loggedInAuthpa := s.Values["loggedIn"]
	storeAdminUserpa := s.Values["storeAdminUser"]
	h.Log.Debug("loggedIn in backups: ", loggedInAuthpa)
	if loggedInAuthpa == true && storeAdminUserpa == true {
		rtn = true
	}
	return rtn
}

func (h *Six910Handler) isStoreCustomerLoggedIn(s *sessions.Session) bool {
	var rtn bool
	loggedInAuthpa := s.Values["loggedIn"]
	storeCustomerUserpa := s.Values["customerUser"]
	h.Log.Debug("loggedIn : ", loggedInAuthpa)
	if loggedInAuthpa == true && storeCustomerUserpa == true {
		rtn = true
	}
	return rtn
}

func (h *Six910Handler) getCustomerID(s *sessions.Session) int64 {
	var rtn int64
	cidstr := s.Values["customerId"]
	if cidstr != nil {
		cid := cidstr.(int)
		rtn = int64(cid)
	}
	return rtn
}

func (h *Six910Handler) storeCustomerCart(cc *m.CustomerCart, s *sessions.Session, w http.ResponseWriter, r *http.Request) bool {
	var rtn bool
	s.Values["customerCart"] = cc
	serr := s.Save(r, w)
	h.Log.Debug("serr", serr)
	if serr == nil {
		rtn = true
	}
	return rtn
}

func (h *Six910Handler) getCustomerCart(s *sessions.Session) *m.CustomerCart {
	var rtn *m.CustomerCart
	fc := s.Values["customerCart"]
	if fc != nil {
		rtn = fc.(*m.CustomerCart)
	}
	return rtn
}
