package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	carsrv "github.com/Ulbora/Six910-ui/carouselsrv"
	conts "github.com/Ulbora/Six910-ui/contentsrv"

	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	mapi "github.com/Ulbora/Six910-ui/mockapi"

	ds "github.com/Ulbora/json-datastore"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Handler_Index(t *testing.T) {
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

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 1
	prod.Sku = "1"
	prod.Desc = "test"

	var prod2 sdbi.Product
	prod2.ID = 2
	prod2.Sku = "2"
	prod2.Desc = "test"

	var prod3 sdbi.Product
	prod3.ID = 3
	prod3.Sku = "3"
	prod3.Desc = "test"

	var prod4 sdbi.Product
	prod4.ID = 4
	prod4.Sku = "4"
	prod4.Desc = "test"

	var plst []sdbi.Product
	plst = append(plst, prod)
	plst = append(plst, prod2)
	plst = append(plst, prod3)
	plst = append(plst, prod4)
	sapi.MockProductList = &plst

	var idlst []int64
	idlst = append(idlst, 5)
	idlst = append(idlst, 53)
	idlst = append(idlst, 555)
	sapi.MockProductIDList = &idlst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "index"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

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

	var cars carsrv.Six910CarouselService
	cars.StorePath = "../carouselsrv/testFiles"
	cars.Log = &l
	var cards ds.DataStore
	cards.Path = "../carouselsrv/testFiles"
	cars.Store = cds.GetNew()
	sh.CarouselService = cars.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.Index(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_Index2(t *testing.T) {
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

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 1
	prod.Sku = "1"
	prod.Desc = "test"

	var prod2 sdbi.Product
	prod2.ID = 2
	prod2.Sku = "2"
	prod2.Desc = "test"

	var prod3 sdbi.Product
	prod3.ID = 3
	prod3.Sku = "3"
	prod3.Desc = "test"

	var prod4 sdbi.Product
	prod4.ID = 4
	prod4.Sku = "4"
	prod4.Desc = "test"

	var prod5 sdbi.Product
	prod5.ID = 5
	prod5.Sku = "5"
	prod5.Desc = "test"

	var plst []sdbi.Product
	plst = append(plst, prod)
	plst = append(plst, prod2)
	plst = append(plst, prod3)
	plst = append(plst, prod4)
	plst = append(plst, prod5)
	sapi.MockProductList = &plst

	var idlst []int64
	idlst = append(idlst, 5)
	idlst = append(idlst, 53)
	idlst = append(idlst, 555)
	sapi.MockProductIDList = &idlst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("index")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "index"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	//res := c.AddContent(&ct)
	//fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

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

	var cars carsrv.Six910CarouselService
	cars.StorePath = "../carouselsrv/testFiles"
	cars.Log = &l
	var cards ds.DataStore
	cards.Path = "../carouselsrv/testFiles"
	cars.Store = cds.GetNew()
	sh.CarouselService = cars.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	sh.setLastHit(s, w, r)
	serr := s.Save(r, w)
	sh.Log.Debug("serr in test", serr)
	h := sh.GetNew()
	h.Index(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_saveSiteMap(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var idlst []int64
	idlst = append(idlst, 5)
	idlst = append(idlst, 53)
	idlst = append(idlst, 55)
	var path = "./testsitemap"
	sh.saveSiteMap(&idlst, path)

}

func TestSix910Handler_saveSiteMap2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.SiteMapDomain = "http://www.six910.com"

	var idlst []int64
	idlst = append(idlst, 5)
	idlst = append(idlst, 53)
	idlst = append(idlst, 55)
	var path = "./testsitemap"
	sh.saveSiteMap(&idlst, path)

}
