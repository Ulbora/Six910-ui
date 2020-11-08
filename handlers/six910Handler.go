package handlers

import (
	"bytes"
	"compress/gzip"
	b64 "encoding/base64"
	"encoding/gob"
	"encoding/json"
	"io/ioutil"

	"html/template"
	"net/http"
	"strconv"
	"strings"

	lg "github.com/Ulbora/Level_Logger"
	bks "github.com/Ulbora/Six910-ui/bkupsrv"
	carsrv "github.com/Ulbora/Six910-ui/carouselsrv"
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	cntrysrv "github.com/Ulbora/Six910-ui/countrysrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	imgs "github.com/Ulbora/Six910-ui/imgsrv"
	mails "github.com/Ulbora/Six910-ui/mailsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	stsrv "github.com/Ulbora/Six910-ui/statesrv"
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	users "github.com/Ulbora/Six910-ui/usersrv"
	api "github.com/Ulbora/Six910API-Go"
	ml "github.com/Ulbora/go-mail-sender"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	gs "github.com/Ulbora/go-sessions"
	sdbi "github.com/Ulbora/six910-database-interface"
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
	UserSession    gs.GoSession
	Templates      *template.Template
	AdminTemplates *template.Template
	Store          *sessions.CookieStore
	UserStore      *sessions.CookieStore

	//services
	BackupService   bks.BackupService
	ContentService  conts.Service
	ImageService    imgs.ImageService
	MailService     mails.MailService
	UserService     users.UserService
	TemplateService tmpsrv.TemplateService
	MenuService     musrv.MenuService
	CSSService      csssrv.CSSService
	CarouselService carsrv.CarouselService
	StateService    stsrv.StateService
	CountryService  cntrysrv.CountryService

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

	//order Emails
	MailSubjectOrderReceived string
	MailBodyOrderReceived    string

	MailSubjectOrderProcessing string
	MailBodyOrderProcession    string

	MailSubjectOrderShipped string
	MailBodyOrderShipped    string

	ImagePath     string
	ThumbnailPath string

	ActiveTemplateName     string
	ActiveTemplateLocation string

	//domain of site www.somesite.com
	Six910SiteURL string

	CompanyName string
}

//HeaderData HeaderData
type HeaderData struct {
	Title         string
	SiteData      *SiteData
	BasicSiteData *BasicSiteData
}

//SiteData SiteData
type SiteData struct {
	Canonical   template.URL
	OgImage     template.URL
	OgType      string
	OgSiteName  string
	OgTitle     string
	OgURL       template.URL
	Description string
}

//BasicSiteData BasicSiteData
type BasicSiteData struct {
	Canonical   template.URL
	Description string
}

func (h *Six910Handler) processProductMetaData(prod *sdbi.Product, r *http.Request) *HeaderData {
	var rtn HeaderData
	var scheme = r.URL.Scheme
	// fmt.Println("scheme: ", scheme)
	// fmt.Println("scheme: ", len(scheme))
	var serverHost string
	if scheme != "" {
		serverHost = r.URL.String()
	} else {
		serverHost = h.SchemeDefault + r.Host
	}
	var sd SiteData
	rtn.Title = prod.ShortDesc
	pidstr := strconv.FormatInt(prod.ID, 10)
	if h.Six910SiteURL != "" {
		sd.Canonical = template.URL(h.Six910SiteURL + "/viewProduct/" + pidstr)
	} else {
		sd.Canonical = template.URL(serverHost + "/viewProduct/" + pidstr)
	}

	if strings.Contains(prod.Image1, "http") {
		sd.OgImage = template.URL(prod.Image1)
	} else if h.Six910SiteURL != "" {
		sd.OgImage = template.URL(h.Six910SiteURL + prod.Image1)
	} else {
		sd.OgImage = template.URL(serverHost + prod.Image1)
	}

	sd.OgType = "product"

	sd.OgSiteName = h.CompanyName

	sd.OgTitle = prod.ShortDesc

	if h.Six910SiteURL != "" {
		sd.OgURL = template.URL(h.Six910SiteURL + "/viewProduct/" + pidstr)
	} else {
		sd.OgURL = template.URL(serverHost + "/viewProduct/" + pidstr)
	}

	if len(prod.Desc) > 159 {
		sd.Description = prod.Desc[0:158]
	} else {
		sd.Description = prod.Desc
	}
	rtn.SiteData = &sd

	return &rtn
}

func (h *Six910Handler) processMetaData(url string, name string, r *http.Request) *HeaderData {
	var rtn HeaderData
	var scheme = r.URL.Scheme
	// fmt.Println("scheme: ", scheme)
	// fmt.Println("scheme: ", len(scheme))
	var serverHost string
	if scheme != "" {
		serverHost = r.URL.String()
	} else {
		serverHost = h.SchemeDefault + r.Host
	}
	var bsd BasicSiteData
	rtn.Title = name

	if h.Six910SiteURL != "" {
		bsd.Canonical = template.URL(h.Six910SiteURL + url)
	} else {
		bsd.Canonical = template.URL(serverHost + url)
	}

	bsd.Description = name
	rtn.BasicSiteData = &bsd

	return &rtn
}

// func (h *Six910Handler) getSiteURL(r *http.Request) string {
// 	if h.Six910SiteURL != ""
// }

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

func (h *Six910Handler) getUserSession(r *http.Request) (*sessions.Session, bool) {
	//fmt.Println("getSession--------------------------------------------------")
	var suc bool
	var srtn *sessions.Session
	if h.UserStore == nil {
		h.UserSession.Name = "Six910-ui-user"
		h.UserSession.MaxAge = 36000000
		h.UserStore = h.UserSession.InitSessionStore()
		h.Log.Debug("h.UserStore : ", h.UserStore)
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
		s, err := h.UserStore.Get(r, h.UserSession.Name)
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
	loggedInAuthpa := s.Values["userLoggenIn"]
	storeCustomerUserpa := s.Values["customerUser"]
	h.Log.Debug("userLoggenIn : ", loggedInAuthpa)
	if loggedInAuthpa == true && storeCustomerUserpa == true {
		rtn = true
	}
	return rtn
}

func (h *Six910Handler) getCustomerID(s *sessions.Session) int64 {
	var rtn int64
	cidstr := s.Values["customerId"]
	h.Log.Debug("cidstr", cidstr)
	if cidstr != nil {
		cid := cidstr.(int64)
		rtn = cid
	}
	return rtn
}

func (h *Six910Handler) storeCustomerCart(cc *m.CustomerCart, s *sessions.Session, w http.ResponseWriter, r *http.Request) bool {
	var rtn bool
	//h.Log.Debug("cc items in save session: ", *cc.Items)
	b, _ := json.Marshal(cc)
	bb := h.compressObj(b)
	s.Values["customerCart"] = bb
	if cc.CustomerAccount != nil && cc.CustomerAccount.Customer != nil {
		s.Values["customerId"] = cc.CustomerAccount.Customer.ID
	}

	serr := s.Save(r, w)
	h.Log.Debug("serr", serr)
	if serr == nil {
		rtn = true
	}
	return rtn
}

func (h *Six910Handler) getCustomerCart(s *sessions.Session) *m.CustomerCart {
	var rtn m.CustomerCart
	fc := s.Values["customerCart"]
	if fc != nil {
		b := fc.([]byte)
		bb := h.decompressObj(b)
		json.Unmarshal(bb, &rtn)
		//rtn = fc.(*m.CustomerCart)
	}
	return &rtn
}

func (h *Six910Handler) compressObj(s []byte) []byte {

	zipbuf := bytes.Buffer{}
	zipped, _ := gzip.NewWriterLevel(&zipbuf, gzip.BestCompression)
	zipped.Write(s)
	zipped.Close()
	h.Log.Debug("compressed size (bytes): ", len(zipbuf.Bytes()))
	return zipbuf.Bytes()
}

func (h *Six910Handler) decompressObj(s []byte) []byte {

	rdr, _ := gzip.NewReader(bytes.NewReader(s))
	data, err := ioutil.ReadAll(rdr)
	h.Log.Debug("decompress err ", err)
	rdr.Close()
	h.Log.Debug("uncompressed size (bytes): ", len(data))
	return data
}

func (h *Six910Handler) getCartTotal(s *sessions.Session, ml *[]musrv.Menu, hd *api.Headers) {
	var rtn int64
	var isLoggedIn = h.isStoreCustomerLoggedIn(s)
	h.Log.Debug("isLoggedIn in carttotal: ", isLoggedIn)
	cc := h.getCustomerCart(s)
	h.Log.Debug("cc: ", cc)
	h.Log.Debug("cc.Items: ", cc.Items)
	if cc != nil && cc.Items != nil && len(*cc.Items) > 0 {
		cv := h.Manager.ViewCart(cc, hd)
		for _, itm := range *cv.Items {
			rtn += itm.Quantity
		}
		for i := range *ml {
			if (*ml)[i].Name == "navBar" && (*ml)[i].Location == "top" {
				(*ml)[i].CartCount = rtn
				(*ml)[i].LoggedIn = isLoggedIn
			}
		}
	} else if cc != nil {
		for i := range *ml {
			if (*ml)[i].Name == "navBar" && (*ml)[i].Location == "top" {
				(*ml)[i].LoggedIn = isLoggedIn
			}
		}
	}
}
