package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	fflsrv "github.com/Ulbora/Six910-ui/findfflsrv"
	man "github.com/Ulbora/Six910-ui/managers"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	ds "github.com/Ulbora/json-datastore"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
)

func TestSix910Handler_FindFFLZipPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "product"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("GET", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZipPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_FindFFLZipPageContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	cds.Delete("shoppingCartContinue")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "product"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	// res := c.AddContent(&ct)
	// fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("GET", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZipPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_FindFFLZipPageNotLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("GET", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = false
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZipPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_FindFFLZip(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		{
			"id": "158223013K15918",
			"licenseName": "VAULT WORLDWIDE, LLC",
			"businessName": "THE VAULT WORLDWIDE, LLC",
			"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132"
		},
		{
			"id": "158223024A02897",
			"licenseName": "PATRIOT PAWN-N-SHOP, INC",
			"businessName": "",
			"premiseAddress": "562 HARDEE ST SUITE B\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F12198",
			"licenseName": "CITY PAWN OF DALLAS LLC",
			"businessName": "",
			"premiseAddress": "620 W MEMORIAL DR   STE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223013D10984",
			"licenseName": "WAHNER, BRIAN KEITH",
			"businessName": "WESTERN ARMS",
			"premiseAddress": "407 AMSTERDAM WAY\nDALLAS, GA 30132"
		},
		{
			"id": "158223074E11283",
			"licenseName": "RJL INC",
			"businessName": "PATRIOT ARMS AND SUPPLY",
			"premiseAddress": "105 VILLAGE WALK, SUITE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223074B07655",
			"licenseName": "U K PRECISION INC",
			"businessName": "",
			"premiseAddress": "2029 MARSHALL HUFF RD SUITE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223071G08116",
			"licenseName": "SKY GUNS INTERNATIONAL LLC",
			"businessName": "SGI",
			"premiseAddress": "224 BRANCH VALLEY DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223013C00889",
			"licenseName": "STINSON, GLEN ELLIS",
			"businessName": "",
			"premiseAddress": "35 COURTHOUSE SQUARE\nDALLAS, GA 30132"
		},
		{
			"id": "158223073H15687",
			"licenseName": "SG3 GUNWORKS LLC",
			"businessName": "",
			"premiseAddress": "655 SHOALS TRAIL\nDALLAS, GA 30132"
		},
		{
			"id": "158223072M09852",
			"licenseName": "WHITTEMORE, HEATH BLAINE",
			"businessName": "HBW FIREARMS",
			"premiseAddress": "277 NEW FARM RD\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F05648",
			"licenseName": "CATES, ROBERT CHRISTOPHER",
			"businessName": "C & C PAWN",
			"premiseAddress": "293 WI PARKWAY STE E\nDALLAS, GA 30132"
		},
		{
			"id": "158223014D11179",
			"licenseName": "ARMAGEDDON ARMS LLC",
			"businessName": "ARMAGEDDON ARMS",
			"premiseAddress": "350 BLACKBERRY RUN DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223072F12092",
			"licenseName": "C PRECISION LLC",
			"businessName": "C PRECISION",
			"premiseAddress": "105 VILLAGE WALK STE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223011E11380",
			"licenseName": "CRABBE, TRENTON RAYMOND",
			"businessName": "",
			"premiseAddress": "2762 NARROWAY CHURCH CIRCLE\nDALLAS, GA 30132"
		}
	]`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var cts conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	cts.Log = &l
	cts.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := cts.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = cts.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZip(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_FindFFLZipContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		{
			"id": "158223013K15918",
			"licenseName": "VAULT WORLDWIDE, LLC",
			"businessName": "THE VAULT WORLDWIDE, LLC",
			"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132"
		},
		{
			"id": "158223024A02897",
			"licenseName": "PATRIOT PAWN-N-SHOP, INC",
			"businessName": "",
			"premiseAddress": "562 HARDEE ST SUITE B\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F12198",
			"licenseName": "CITY PAWN OF DALLAS LLC",
			"businessName": "",
			"premiseAddress": "620 W MEMORIAL DR   STE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223013D10984",
			"licenseName": "WAHNER, BRIAN KEITH",
			"businessName": "WESTERN ARMS",
			"premiseAddress": "407 AMSTERDAM WAY\nDALLAS, GA 30132"
		},
		{
			"id": "158223074E11283",
			"licenseName": "RJL INC",
			"businessName": "PATRIOT ARMS AND SUPPLY",
			"premiseAddress": "105 VILLAGE WALK, SUITE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223074B07655",
			"licenseName": "U K PRECISION INC",
			"businessName": "",
			"premiseAddress": "2029 MARSHALL HUFF RD SUITE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223071G08116",
			"licenseName": "SKY GUNS INTERNATIONAL LLC",
			"businessName": "SGI",
			"premiseAddress": "224 BRANCH VALLEY DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223013C00889",
			"licenseName": "STINSON, GLEN ELLIS",
			"businessName": "",
			"premiseAddress": "35 COURTHOUSE SQUARE\nDALLAS, GA 30132"
		},
		{
			"id": "158223073H15687",
			"licenseName": "SG3 GUNWORKS LLC",
			"businessName": "",
			"premiseAddress": "655 SHOALS TRAIL\nDALLAS, GA 30132"
		},
		{
			"id": "158223072M09852",
			"licenseName": "WHITTEMORE, HEATH BLAINE",
			"businessName": "HBW FIREARMS",
			"premiseAddress": "277 NEW FARM RD\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F05648",
			"licenseName": "CATES, ROBERT CHRISTOPHER",
			"businessName": "C & C PAWN",
			"premiseAddress": "293 WI PARKWAY STE E\nDALLAS, GA 30132"
		},
		{
			"id": "158223014D11179",
			"licenseName": "ARMAGEDDON ARMS LLC",
			"businessName": "ARMAGEDDON ARMS",
			"premiseAddress": "350 BLACKBERRY RUN DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223072F12092",
			"licenseName": "C PRECISION LLC",
			"businessName": "C PRECISION",
			"premiseAddress": "105 VILLAGE WALK STE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223011E11380",
			"licenseName": "CRABBE, TRENTON RAYMOND",
			"businessName": "",
			"premiseAddress": "2762 NARROWAY CHURCH CIRCLE\nDALLAS, GA 30132"
		}
	]`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var cts conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	// ds.Delete("books1")
	cds.Delete("shoppingCartContinue")
	cts.Log = &l
	cts.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "product"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	// res := cts.AddContent(&ct)
	// fmt.Println("content save: ", res)

	sh.ContentService = cts.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZip(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_FindFFLZipLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		{
			"id": "158223013K15918",
			"licenseName": "VAULT WORLDWIDE, LLC",
			"businessName": "THE VAULT WORLDWIDE, LLC",
			"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132"
		},
		{
			"id": "158223024A02897",
			"licenseName": "PATRIOT PAWN-N-SHOP, INC",
			"businessName": "",
			"premiseAddress": "562 HARDEE ST SUITE B\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F12198",
			"licenseName": "CITY PAWN OF DALLAS LLC",
			"businessName": "",
			"premiseAddress": "620 W MEMORIAL DR   STE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223013D10984",
			"licenseName": "WAHNER, BRIAN KEITH",
			"businessName": "WESTERN ARMS",
			"premiseAddress": "407 AMSTERDAM WAY\nDALLAS, GA 30132"
		},
		{
			"id": "158223074E11283",
			"licenseName": "RJL INC",
			"businessName": "PATRIOT ARMS AND SUPPLY",
			"premiseAddress": "105 VILLAGE WALK, SUITE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223074B07655",
			"licenseName": "U K PRECISION INC",
			"businessName": "",
			"premiseAddress": "2029 MARSHALL HUFF RD SUITE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223071G08116",
			"licenseName": "SKY GUNS INTERNATIONAL LLC",
			"businessName": "SGI",
			"premiseAddress": "224 BRANCH VALLEY DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223013C00889",
			"licenseName": "STINSON, GLEN ELLIS",
			"businessName": "",
			"premiseAddress": "35 COURTHOUSE SQUARE\nDALLAS, GA 30132"
		},
		{
			"id": "158223073H15687",
			"licenseName": "SG3 GUNWORKS LLC",
			"businessName": "",
			"premiseAddress": "655 SHOALS TRAIL\nDALLAS, GA 30132"
		},
		{
			"id": "158223072M09852",
			"licenseName": "WHITTEMORE, HEATH BLAINE",
			"businessName": "HBW FIREARMS",
			"premiseAddress": "277 NEW FARM RD\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F05648",
			"licenseName": "CATES, ROBERT CHRISTOPHER",
			"businessName": "C & C PAWN",
			"premiseAddress": "293 WI PARKWAY STE E\nDALLAS, GA 30132"
		},
		{
			"id": "158223014D11179",
			"licenseName": "ARMAGEDDON ARMS LLC",
			"businessName": "ARMAGEDDON ARMS",
			"premiseAddress": "350 BLACKBERRY RUN DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223072F12092",
			"licenseName": "C PRECISION LLC",
			"businessName": "C PRECISION",
			"premiseAddress": "105 VILLAGE WALK STE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223011E11380",
			"licenseName": "CRABBE, TRENTON RAYMOND",
			"businessName": "",
			"premiseAddress": "2762 NARROWAY CHURCH CIRCLE\nDALLAS, GA 30132"
		}
	]`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = false
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZip(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_FindFFLZipNoResults(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		
	]`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var cts conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	cts.Log = &l
	cts.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "product"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := cts.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = cts.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLZip(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_FindFFLID(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var cts conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	cts.Log = &l
	cts.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := cts.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = cts.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLID(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_FindFFLIDContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var cts conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	cds.Delete("shoppingCartContinue")
	cts.Log = &l
	cts.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "product"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := cts.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = cts.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLID(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_FindFFLIDLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["userLoggenIn"] = false
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.FindFFLID(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_AddFFL(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------
	var cusm sdbi.Customer
	cusm.ID = 5
	// cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	var cccs man.CustomerCart
	var cct sdbi.Cart
	cct.ID = 1
	cccs.Cart = &cct
	var ca man.CustomerAccount
	cccs.CustomerAccount = &ca

	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	//s.Values["loggedIn"] = true
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.AddFFL(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_AddFFLBadFFLBusName(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------
	var cusm sdbi.Customer
	cusm.ID = 5
	// cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	//s.Values["loggedIn"] = true

	var cccs man.CustomerCart
	var cct sdbi.Cart
	cct.ID = 1
	cccs.Cart = &cct
	var ca man.CustomerAccount
	cccs.CustomerAccount = &ca
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb

	s.Save(r, w)
	h := sh.GetNew()
	h.AddFFL(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_AddFFLBadEmail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------
	var cusm sdbi.Customer
	cusm.ID = 5
	// cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = ""
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.AddFFL(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()
}

func TestSix910Handler_AddFFLNotLoggedIn(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	sh.ActiveTemplateLocation = "./testsitemap"
	sh.ActiveTemplateName = "test"

	var c fflsrv.Six910FFLService
	// //var l lg.Logger
	// //l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	// s := c.New()
	sh.FFLService = c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	// sh.FFLService = c.New()

	//-----------start mocking------------------
	var cusm sdbi.Customer
	cusm.ID = 5
	// cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	//-----------end mocking --------

	//-----------start mocking------------------

	//-----------end mocking --------

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var m musrv.Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&m)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	// r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("zip=12345"))
	// r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "158223013K15918",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = false
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.AddFFL(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()
}
