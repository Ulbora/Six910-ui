package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	m "github.com/Ulbora/Six910-ui/managers"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	ml "github.com/Ulbora/go-mail-sender"
	ds "github.com/Ulbora/json-datastore"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
)

func TestSix910Handler_AddProductToCart(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.ResponseID
	cires.ID = 4
	cires.Success = true

	sapi.MockCartItemAddResp = &cires

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

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

	var mm musrv.Menu
	mm.Name = "menu1"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
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
		"prodId": "12",
		//"quantity": "2",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 12
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	var crt sdbi.Cart
	cccs.Cart = &crt
	cccs.Items = &cilstp

	var cus sdbi.Customer
	cus.ID = 3

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cccs.CustomerAccount = &cusa

	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.AddProductToCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AddProductToCart2(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.ResponseID
	cires.ID = 4
	cires.Success = true

	sapi.MockCartItemAddResp = &cires

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

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

	var mm musrv.Menu
	mm.Name = "menu1"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	//-----------end mocking --------

	// var c conts.CmsService
	// var ds ds.DataStore
	// ds.Path = "../contentsrv/testFiles"
	// //ds.Delete("books1")
	// c.Log = &l
	// c.Store = ds.GetNew()

	// var ct conts.Content
	// ct.Name = "product"
	// ct.Author = "ken"
	// ct.MetaAuthorName = "ken"
	// ct.MetaDesc = "shopping cart index"
	// ct.Text = "some book text"
	// ct.Title = "the best book ever"
	// ct.Visible = true
	// res := c.AddContent(&ct)
	// fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com?id=123&qty=3", nil)
	// vars := map[string]string{
	// 	"prodId": "12",
	// 	//"quantity": "2",
	// }
	// r = mux.SetURLVars(r, vars)

	var cccs m.CustomerCart
	var crt sdbi.Cart
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cccs.CustomerAccount = &cusa

	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.AddProductToCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_ViewCart(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.ResponseID
	cires.ID = 4
	cires.Success = true

	sapi.MockCartItemAddResp = &cires

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCart"
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
		"prodId":   "12",
		"quantity": "2",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewCartCartSession(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pd sdbi.Product
	pd.ID = 4
	pd.SalePrice = 28.95
	pd.ShortDesc = "test one"
	pd.Thumbnail = "/test/"
	pd.SpecialProcessing = true
	pd.SpecialProcessingType = "FFL"
	sapi.MockProduct = &pd

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.ResponseID
	cires.ID = 4
	cires.Success = true

	sapi.MockCartItemAddResp = &cires

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCart"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	//res := c.AddContent(&ct)
	res := c.DeleteContent("shoppingCart")
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
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
	// vars := map[string]string{
	// 	"prodId":   "12",
	// 	"quantity": "2",
	// }
	// r = mux.SetURLVars(r, vars)

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var cus sdbi.Customer
	cus.ID = 3

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cccs.CustomerAccount = &cusa

	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
	// t.Fail()
}

func TestSix910Handler_UpdateProductToCart(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	// var crtres api.ResponseID
	// crtres.ID = 3
	// crtres.Success = true

	// sapi.MockAddCartResp = &crtres

	var cires api.Response
	cires.Success = true

	sapi.MockCartItemUpdateResp = &cires

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

	//-----------end mocking --------

	var c conts.CmsService
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = ds.GetNew()

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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("GET", "https://test.com?id=9&quantity=2", nil)
	vars := map[string]string{
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var cus sdbi.Customer
	cus.ID = 3

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cccs.CustomerAccount = &cusa

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateProductToCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateProductToCartFail(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	// var crtres api.ResponseID
	// crtres.ID = 3
	// crtres.Success = true

	// sapi.MockAddCartResp = &crtres

	var cires api.Response
	//cires.Success = true

	sapi.MockCartItemUpdateResp = &cires

	// var plst []sdbi.Product
	// plst = append(plst, prod)
	// sapi.MockProductList = &plst

	//-----------end mocking --------

	var c conts.CmsService
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = ds.GetNew()

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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var cus sdbi.Customer
	cus.ID = 3

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cccs.CustomerAccount = &cusa

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateProductToCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutView(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	prod.SpecialProcessing = true
	prod.SpecialProcessingType = "FFL"
	sapi.MockProduct = &prod

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var rgn sdbi.Region
	rgn.ID = 1
	rgn.Name = "test"
	sapi.MockRegion = &rgn

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var cadd sdbi.Address
	cadd.ID = 3
	var caddlst []sdbi.Address
	caddlst = append(caddlst, cadd)
	var ffladd sdbi.Address
	ffladd.ID = 4
	ffladd.Type = "FFL"
	caddlst = append(caddlst, ffladd)
	sapi.MockAddressList1 = &caddlst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCart2"
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

	var mm musrv.Menu
	mm.Name = "menu1"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
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
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp
	var cus sdbi.Customer
	cus.ID = 3
	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cccs.CustomerAccount = &cusa

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutView(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
	// t.Fail()
}

func TestSix910Handler_CheckOutView2(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var rgn sdbi.Region
	rgn.ID = 1
	rgn.Name = "test"
	sapi.MockRegion = &rgn

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var add sdbi.Address
	add.ID = 1

	var adlst []sdbi.Address
	adlst = append(adlst, add)
	sapi.MockAddressList1 = &adlst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCart2")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCart2"
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

	var mm musrv.Menu
	mm.Name = "menu1"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
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
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutView(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutView3(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var rgn sdbi.Region
	rgn.ID = 1
	rgn.Name = "test"
	sapi.MockRegion = &rgn

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var add sdbi.Address
	add.ID = 1

	var adlst []sdbi.Address
	adlst = append(adlst, add)
	sapi.MockAddressList1 = &adlst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCart2")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCart2"
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

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
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
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	//cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	//cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutView(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutViewBtcPay(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"
	sapi.MockProduct = &prod

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "BTCPayServer"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var rgn sdbi.Region
	rgn.ID = 1
	rgn.Name = "test"
	sapi.MockRegion = &rgn

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var add sdbi.Address
	add.ID = 1

	var adlst []sdbi.Address
	adlst = append(adlst, add)
	sapi.MockAddressList1 = &adlst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCart2")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCart2"
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

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
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
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	//cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	//cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutView(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutViewLogin(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	//-----------end mocking --------

	var c conts.CmsService
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = ds.GetNew()

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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"prodId":   "9",
		"quantity": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutView(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutContinue(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2,&shippingAddressID=4"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var usr api.User
	usr.Enabled = true

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	fmt.Println("cccs.Items: ", cccs.Items)
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutContinue(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutContinue2(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL authorize"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCartContinue")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	//res := c.AddContent(&ct)
	//fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var usr api.User
	usr.Enabled = true
	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutContinue(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutContinue3(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "Bill Me Later"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCartContinue")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	//res := c.AddContent(&ct)
	//fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var usr api.User
	usr.Enabled = true
	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutContinue(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutContinueBtcPayServer(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	pg.Name = "BTCPayServer"
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "BTCPayServer"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCartContinue")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	//res := c.AddContent(&ct)
	//fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var usr api.User
	usr.Enabled = true
	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutContinue(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
	// t.Fail()
}

func TestSix910Handler_CheckOutContinueLogin(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	//-----------end mocking --------

	var c conts.CmsService
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = ds.GetNew()

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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"PaymentGatewayID":  "9",
		"ShippingMethodID":  "2",
		"InsuranceID":       "2",
		"BillingAddressID":  "2",
		"ShippingAddressID": "2",
	}

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutContinue(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutComplateOrder(t *testing.T) {
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

	var mss ml.MockSecureSender
	mss.MockSuccess = true
	sh.MailSender = mss.GetNew()

	sh.MailSenderAddress = "test@test.com"

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	var tranres api.ResponseID
	tranres.ID = 3
	tranres.Success = true
	sapi.MockOrderTransactionRes = &tranres

	var pupres api.Response
	pupres.Success = true
	sapi.MockUpdateProductQuantityResp = &pupres

	var str sdbi.Store
	str.ID = 5
	sapi.MockStore = &str

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	vars := map[string]string{
		"transactionCode": "billMeLaterTransaction",
	}
	r = mux.SetURLVars(r, vars)

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var cilstpv []*m.CartViewItem

	var ctit1v m.CartViewItem
	ctit1v.Quantity = 3
	ctit1v.ProductID = 7
	cilstpv = append(cilstpv, &ctit1v)

	var ctit2v m.CartViewItem
	ctit2v.Quantity = 4
	ctit2v.ProductID = 9
	cilstpv = append(cilstpv, &ctit2v)

	//var cccs m.CustomerCart
	var cv m.CartView
	cv.Items = &cilstpv
	cccs.CartView = &cv

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var usr api.User
	usr.Enabled = true

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	fmt.Println("cccs.Items: ", cccs.Items)
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutComplateOrder(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutComplateOrder2(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	var tranres api.ResponseID
	tranres.ID = 3
	tranres.Success = true
	sapi.MockOrderTransactionRes = &tranres

	var pupres api.Response
	pupres.Success = true
	sapi.MockUpdateProductQuantityResp = &pupres

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("shoppingCartContinue")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	//res := c.AddContent(&ct)
	//fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var cilstpv []*m.CartViewItem

	var ctit1v m.CartViewItem
	ctit1v.Quantity = 3
	ctit1v.ProductID = 7
	cilstpv = append(cilstpv, &ctit1v)

	var ctit2v m.CartViewItem
	ctit2v.Quantity = 4
	ctit2v.ProductID = 9
	cilstpv = append(cilstpv, &ctit2v)

	//var cccs m.CustomerCart
	var cv m.CartView
	cv.Items = &cilstpv
	cccs.CartView = &cv

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var usr api.User
	usr.Enabled = true

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	fmt.Println("cccs.Items: ", cccs.Items)
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutComplateOrder(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutComplateOrder3(t *testing.T) {
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

	var mss ml.MockSecureSender
	mss.MockSuccess = true
	sh.MailSender = mss.GetNew()

	sh.MailSenderAddress = "test@test.com"

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	var tranres api.ResponseID
	tranres.ID = 3
	tranres.Success = true
	sapi.MockOrderTransactionRes = &tranres

	var str sdbi.Store
	str.ID = 5
	sapi.MockStore = &str

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	//cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var usr api.User
	usr.Enabled = true

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	fmt.Println("cccs.Items: ", cccs.Items)
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutComplateOrder(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CheckOutComplateOrderLogin(t *testing.T) {
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

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	//-----------start mocking------------------

	var pg sdbi.PaymentGateway
	pg.ID = 2
	pg.StorePluginsID = 4
	var pglst []sdbi.PaymentGateway
	pglst = append(pglst, pg)
	sapi.MockPaymentGatewayList = &pglst

	sapi.MockPaymentGateway = &pg

	var spi sdbi.StorePlugins
	spi.ID = 4
	spi.PluginName = "PAYPAL"
	sapi.MockStorePlugin = &spi

	var sm sdbi.ShippingMethod
	sm.ID = 5
	sm.Name = "USP"
	var smslt []sdbi.ShippingMethod
	smslt = append(smslt, sm)
	sapi.MockShippingMethodList = &smslt

	var ins sdbi.Insurance
	ins.ID = 4
	ins.Cost = 4.55
	var inslst []sdbi.Insurance
	inslst = append(inslst, ins)
	sapi.MockInsuranceList = &inslst

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true
	sapi.MockAddOrderResp = &crtres

	var prod sdbi.Product
	prod.ID = 3
	sapi.MockProduct = &prod

	var oid sdbi.OrderItem
	oid.ID = 3
	sapi.MockOrderItem = &oid

	var aoires api.ResponseID
	aoires.ID = 3
	aoires.Success = true
	sapi.MockAddOrderItemResp = &aoires

	var tranres api.ResponseID
	tranres.ID = 3
	tranres.Success = true
	sapi.MockOrderTransactionRes = &tranres

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "shoppingCartContinue"
	ct.Author = "ken"
	ct.MetaAuthorName = "ken"
	ct.MetaDesc = "shopping cart index"
	ct.Text = "some book text"
	ct.Title = "the best book ever"
	ct.Visible = true
	res := c.AddContent(&ct)
	fmt.Println("content save: ", res)

	sh.ContentService = c.GetNew()

	var css csssrv.Six910CSSService
	var csds ds.DataStore
	csds.Path = "./testFiles"
	css.CSSStore = csds.GetNew()
	css.Log = &l
	sh.CSSService = css.GetNew()

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "./testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	var mm musrv.Menu
	mm.Name = "navBar"
	mm.Active = true
	mm.Location = "top"
	mm.Shade = "light"
	mm.Background = "light"
	mm.Style = ""
	mm.ShadeList = &[]string{"light", "dark"}
	mm.BackgroundList = &[]string{"light", "dark"}

	msuc := ms.AddMenu(&mm)
	fmt.Println("menu save: ", msuc)

	sh.MenuService = ms

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("paymentGatewayID=9&"+
		"shippingMethodID=22&insuranceID=2&billingAddressID=23&shippingAddressID=2"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	// vars := map[string]string{
	// 	"PaymentGatewayID":  "9",
	// 	"ShippingMethodID":  "2",
	// 	"InsuranceID":       "2",
	// 	"BillingAddressID":  "2",
	// 	"ShippingAddressID": "2",
	// }

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cccs m.CustomerCart
	cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	var usr api.User
	usr.Enabled = true

	var ad sdbi.Address
	ad.ID = 4
	var adlst []sdbi.Address
	adlst = append(adlst, ad)

	var cusa m.CustomerAccount
	cusa.Customer = &cus
	cusa.User = &usr
	cusa.Addresses = &adlst
	cccs.CustomerAccount = &cusa

	//var itmList []sdbi.CartItem
	//cccs.Items = &itmList

	//r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	//s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	//s.Values["customerCart"] = &cccs
	fmt.Println("cccs.Items: ", cccs.Items)
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutComplateOrder(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
