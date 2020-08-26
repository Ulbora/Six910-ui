package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	conts "github.com/Ulbora/Six910-ui/contsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
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
	var ds ds.DataStore
	ds.Path = "../contsrv/testFiles"
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
		"prodId":   "12",
		"quantity": "2",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
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
	var ds ds.DataStore
	ds.Path = "../contsrv/testFiles"
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
		"prodId":   "12",
		"quantity": "2",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
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
	var ds ds.DataStore
	ds.Path = "../contsrv/testFiles"
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

	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewCart(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
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
	ds.Path = "../contsrv/testFiles"
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
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
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
	ds.Path = "../contsrv/testFiles"
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
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
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
	ds.Path = "../contsrv/testFiles"
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
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
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
	ds.Path = "../contsrv/testFiles"
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
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CheckOutView(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
