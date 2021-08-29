package handlers

import (
	"bytes"
	"compress/gzip"
	b64 "encoding/base64"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"time"

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
	fflsrv "github.com/Ulbora/Six910-ui/findfflsrv"
	imgs "github.com/Ulbora/Six910-ui/imgsrv"
	mails "github.com/Ulbora/Six910-ui/mailsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	stsrv "github.com/Ulbora/Six910-ui/statesrv"
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	users "github.com/Ulbora/Six910-ui/usersrv"
	api "github.com/Ulbora/Six910API-Go"
	btc "github.com/Ulbora/Six910BTCPayServerPlugin"
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

	BTCPlugin      btc.Plugin
	BTCPayCurrency string

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
	FFLService      fflsrv.FFLService

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
	MailBodyOrderProcessing    string

	MailSubjectOrderShipped string
	MailBodyOrderShipped    string

	MailSubjectOrderCanceled string
	MailBodyOrderCanceled    string

	MailSubjectPasswordReset string
	MailBodyPasswordReset    string

	ImagePath     string
	ThumbnailPath string

	ActiveTemplateName     string
	ActiveTemplateLocation string

	//domain of site www.somesite.com
	Six910SiteURL string

	CompanyName string

	SiteMapDomain string
	SiteMapDate   time.Time

	BackupFileName string
}

//HeaderData HeaderData
type HeaderData struct {
	Title           string
	SiteData        *SiteData
	RichResultsData *RichResultsData
	BasicSiteData   *BasicSiteData
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

//RichResultsData RichResultsData
type RichResultsData struct {
	Type     string
	Sku      string
	Name     string
	Desc     string
	Image    template.JS
	Mpn      string
	Brand    string
	Price    string
	Currency string
}

//BasicSiteData BasicSiteData
type BasicSiteData struct {
	Canonical   template.URL
	Description string
}

func (h *Six910Handler) getSiteMapDomain() string {
	var rtn string
	if h.SiteMapDomain == "" {
		rtn = "http://localhost:8080"
	} else {
		rtn = h.SiteMapDomain
	}
	return rtn
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
	rtn.RichResultsData = h.processRichResultsData(prod, serverHost)

	return &rtn
}

func (h *Six910Handler) processRichResultsData(prod *sdbi.Product, serverHost string) *RichResultsData {
	var rsd RichResultsData
	rsd.Type = "Product"
	rsd.Sku = prod.Sku
	rsd.Name = prod.Name
	rsd.Desc = prod.ShortDesc
	rsd.Mpn = prod.ManufacturerID
	rsd.Brand = prod.Manufacturer
	if strings.Contains(prod.Image1, "http") {
		rsd.Image = template.JS(prod.Image1)
	} else if h.Six910SiteURL != "" {
		rsd.Image = template.JS(h.Six910SiteURL + prod.Image1)
	} else {
		rsd.Image = template.JS(serverHost + prod.Image1)
	}
	rsd.Price = fmt.Sprintf("%.2f", prod.Price)
	cur := prod.Currency
	if cur == "" {
		rsd.Currency = "USA"
	} else {
		rsd.Currency = prod.Currency
	}

	return &rsd
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
		h.Session.MaxAge = 3000
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

func (h *Six910Handler) getUserSession(w http.ResponseWriter, r *http.Request) (*sessions.Session, bool) {
	//fmt.Println("getSession--------------------------------------------------")
	var suc bool
	var srtn *sessions.Session
	if h.UserStore == nil {
		h.UserSession.Name = "Six910-ui-user"
		h.UserSession.MaxAge = 36000000
		h.createUserSession()
	}
	if r != nil {
		// fmt.Println("secure in getSession", h.Session.Secure)
		// fmt.Println("name in getSession", h.Session.Name)
		// fmt.Println("MaxAge in getSession", h.Session.MaxAge)
		// fmt.Println("SessionKey in getSession", h.Session.SessionKey)

		//h.Session.HTTPOnly = true

		//h.Session.InitSessionStore()
		//fmt.Println("h.UserSession.Name: ", h.UserSession.Name)

		s, err := h.UserStore.Get(r, h.UserSession.Name)

		//fmt.Println("session s: ", s)

		//s, err := store.Get(r, "temp-name")
		//s, err := store.Get(r, "goauth2")
		if s != nil {
			loggedInAuth := s.Values["userLoggenIn"]
			//userAuth := s.Values["user"]
			h.Log.Debug("userLoggenIn: ", loggedInAuth)
		}
		// loggedInAuth := s.Values["userLoggenIn"]
		// //userAuth := s.Values["user"]
		// h.Log.Debug("userLoggenIn: ", loggedInAuth)
		//h.Log.Debug("user: ", userAuth)

		//larii := s.Values["authReqInfo"]
		//h.Log.Debug("arii-----login", larii)

		h.Log.Debug("session error in getSession: ", err)
		if err == nil {
			suc = true
			srtn = s
			srtn.Values["customerUser"] = true
			serr := srtn.Save(r, w)
			h.Log.Debug("serr", serr)
		} else {
			h.UserStore = nil
			h.UserSession.Name = "Six910-ui-user"
			h.UserSession.MaxAge = 36000000
			h.createUserSession()
			s, _ := h.UserStore.Get(r, h.UserSession.Name)
			suc = true
			srtn = s
			srtn.Values["customerUser"] = true
			serr := srtn.Save(r, w)
			h.Log.Debug("serr", serr)
		}
	}
	//fmt.Println("exit getSession--------------------------------------------------")
	return srtn, suc
}

func (h *Six910Handler) createUserSession() {
	h.UserStore = h.UserSession.InitSessionStore()
	h.Log.Debug("h.UserStore : ", h.UserStore)
	//errors without this
	gob.Register(&m.CustomerCart{})
	//-------gob.Register(&AuthorizeRequestInfo{})
}

func (h *Six910Handler) getHeader(s *sessions.Session) *api.Headers {
	var hd api.Headers
	storeCustomerUserpa := s.Values["customerUser"]
	h.Log.Debug("storeCustomerUserpa: ", storeCustomerUserpa)
	if !h.OAuth2Enabled || storeCustomerUserpa == true {
		var sEnccl string
		username := s.Values["username"]
		password := s.Values["password"]
		if username != nil && password != nil {
			sEnccl = b64.StdEncoding.EncodeToString([]byte(username.(string) + ":" + password.(string)))
		}
		h.Log.Debug("sEnc: ", sEnccl)
		hd.Set("Authorization", "Basic "+sEnccl)
		if storeCustomerUserpa != true {
			hd.Set("clientId", "none")
		}
	} else {
		hd.Set("Authorization", "Bearer "+h.token.AccessToken)
		hd.Set("clientId", h.ClientCreds.AuthCodeClient)
	}
	return &hd
}

func (h *Six910Handler) isStoreAdminLoggedIn(s *sessions.Session) bool {
	var rtn bool
	loggedInAuthpa := s.Values["loggedIn"]
	storeAdminUserpa := s.Values["storeAdminUser"]
	h.Log.Debug("loggedIn in: ", loggedInAuthpa)
	if !h.OAuth2Enabled && loggedInAuthpa == true && storeAdminUserpa == true {
		rtn = true
	} else if loggedInAuthpa == true && storeAdminUserpa == true && h.token != nil {
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

func (h *Six910Handler) setLastHit(s *sessions.Session, w http.ResponseWriter, r *http.Request) {
	s.Values["lastHitTime"] = time.Now().UnixNano() / int64(time.Millisecond)
	serr := s.Save(r, w)
	h.Log.Debug("serr in setLastHit", serr)
}
func (h *Six910Handler) getLastHit(s *sessions.Session, w http.ResponseWriter, r *http.Request) time.Time {
	var rtn time.Time
	hv := s.Values["lastHitTime"]
	h.Log.Debug("hv", hv)
	if hv != nil {
		mill := hv.(int64)
		rtn = time.Unix(0, mill*int64(time.Millisecond))
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

func (h *Six910Handler) getCartTotal(s *sessions.Session, ml *[]musrv.Menu, hd *api.Headers) bool {
	var fflRtn bool
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
			if itm.SpecialProcessing && itm.SpecialProcessingType == "FFL" {
				fflRtn = true
			}
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
	return fflRtn
}

//CheckContent CheckContent
func (h *Six910Handler) CheckContent(r *http.Request) bool {
	var rtn bool
	cType := r.Header.Get("Content-Type")
	if cType == "application/json" {
		// http.Error(w, "json required", http.StatusUnsupportedMediaType)
		rtn = true
	}
	return rtn
}

//SetContentType SetContentType
func (h *Six910Handler) SetContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//ProcessBody ProcessBody
func (h *Six910Handler) ProcessBody(r *http.Request, obj interface{}) (bool, error) {
	var suc bool
	var err error
	//fmt.Println("r.Body: ", r.Body)
	if r.Body != nil {
		decoder := json.NewDecoder(r.Body)
		//fmt.Println("decoder: ", decoder)
		err = decoder.Decode(obj)
		//fmt.Println("decoder: ", decoder)
		if err != nil {
			//log.Println("Decode Error: ", err.Error())
			h.Log.Error("Decode Error: ", err.Error())
		} else {
			suc = true
		}
	} else {
		err = errors.New("Bad Body")
	}

	return suc, err
}
