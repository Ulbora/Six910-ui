package handlers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	ds "github.com/Ulbora/json-datastore"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
	//conts "github.com/Ulbora/Six910-ui/contentsrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
)

func TestSix910Handler_ViewCustomerOrder(t *testing.T) {
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

	var prod sdbi.Order
	prod.ID = 2
	prod.CustomerName = "tester"
	sapi.MockOrder = &prod

	var oi sdbi.OrderItem
	var plst []sdbi.OrderItem
	plst = append(plst, oi)
	sapi.MockOrderItemList = &plst

	var oc sdbi.OrderComment
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

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
		"id": "12",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewCustomerOrder(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewCustomerOrderLogin(t *testing.T) {
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

	var prod sdbi.Order
	prod.ID = 2
	prod.CustomerName = "tester"
	sapi.MockOrder = &prod

	var oi sdbi.OrderItem
	var plst []sdbi.OrderItem
	plst = append(plst, oi)
	sapi.MockOrderItemList = &plst

	var oc sdbi.OrderComment
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

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
		"id": "12",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewCustomerOrder(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_ViewCustomerOrderList(t *testing.T) {
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

	var odr sdbi.Order
	odr.ID = 2
	odr.CustomerName = "tester"
	var odrlst []sdbi.Order
	odrlst = append(odrlst, odr)
	sapi.MockOrderList = &odrlst

	var oi sdbi.OrderItem
	var plst []sdbi.OrderItem
	plst = append(plst, oi)
	sapi.MockOrderItemList = &plst

	var oc sdbi.OrderComment
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("orderList")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "orderList"
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
	// vars := map[string]string{
	// 	"id": "12",
	// }
	// r = mux.SetURLVars(r, vars)

	var cccs m.CustomerCart
	//cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	crt.CustomerID = 5
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	//s.Save(r, w)
	h := sh.GetNew()
	h.ViewCustomerOrderList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewCustomerOrderList2(t *testing.T) {
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

	var odr sdbi.Order
	odr.ID = 2
	odr.CustomerName = "tester"
	var odrlst []sdbi.Order
	odrlst = append(odrlst, odr)
	sapi.MockOrderList = &odrlst

	var oi sdbi.OrderItem
	var plst []sdbi.OrderItem
	plst = append(plst, oi)
	sapi.MockOrderItemList = &plst

	var oc sdbi.OrderComment
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	//-----------end mocking --------

	var c conts.CmsService
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	cds.Delete("orderList")
	c.Log = &l
	c.Store = cds.GetNew()

	var ct conts.Content
	ct.Name = "orderList"
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
	// vars := map[string]string{
	// 	"id": "12",
	// }
	// r = mux.SetURLVars(r, vars)

	var cccs m.CustomerCart
	//cccs.Items = &cilstp

	var crt sdbi.Cart
	crt.ID = 3
	crt.CustomerID = 5
	cccs.Cart = &crt

	var cus sdbi.Customer
	cus.ID = 3

	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	b, _ := json.Marshal(cccs)
	bb := sh.compressObj(b)
	s.Values["customerCart"] = bb
	s.Save(r, w)
	//s.Save(r, w)
	h := sh.GetNew()
	h.ViewCustomerOrderList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_ViewCustomerOrderListLogin(t *testing.T) {
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

	var odr sdbi.Order
	odr.ID = 2
	odr.CustomerName = "tester"
	var odrlst []sdbi.Order
	odrlst = append(odrlst, odr)
	sapi.MockOrderList = &odrlst

	var oi sdbi.OrderItem
	var plst []sdbi.OrderItem
	plst = append(plst, oi)
	sapi.MockOrderItemList = &plst

	var oc sdbi.OrderComment
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

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
		"id": "12",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.ViewCustomerOrderList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
