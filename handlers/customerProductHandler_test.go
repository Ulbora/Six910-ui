package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	//api "github.com/Ulbora/Six910API-Go"
	ds "github.com/Ulbora/json-datastore"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
	//"github.com/gorilla/sessions"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
)

func TestSix910Handler_ViewProductList(t *testing.T) {
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

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"

	var plst []sdbi.Product
	plst = append(plst, prod)
	sapi.MockProductList = &plst

	//-----------end mocking --------

	var c conts.CmsService
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = ds.GetNew()

	var ct conts.Content
	ct.Name = "productList"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"catId": "11",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewProductList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_SearchProductList(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "productList"
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"search": "test",
		"start":  "0",
		"end":    "0",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.SearchProductList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_SearchProductList2(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("productList")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "productList"
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"search": "test",
		"start":  "0",
		"end":    "0",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.SearchProductList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_SearchProductByManufacturerList(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "productList"
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"manf":   "test",
		"search": "test",
		"start":  "0",
		"end":    "0",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.SearchProductByManufacturerList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_SearchProductByManufacturerList2(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("productList")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "productList"
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"manf":   "test",
		"search": "test",
		"start":  "0",
		"end":    "0",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.SearchProductByManufacturerList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewProduct(t *testing.T) {
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

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

	var pclst = []int64{2, 4}
	sapi.MockProductCategoryIDList = pclst

	var cat sdbi.Category
	cat.ID = 2
	cat.ParentCategoryID = 0
	sapi.MockCategory = &cat

	//-----------end mocking --------

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

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"id": "12",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewProduct(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewProductByCatList(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "productCategoryList"
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewProductByCatList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewProductByCatList2(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("productCategoryList")
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewProductByCatList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewProductByCatAndManufacturerList(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "productCategoryList"
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewProductByCatAndManufacturerList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewProductByCatAndManufacturerList2(t *testing.T) {
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

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("productCategoryList")
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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewProductByCatAndManufacturerList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}
