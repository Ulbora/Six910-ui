package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	conts "github.com/Ulbora/Six910-ui/contentsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	ds "github.com/Ulbora/json-datastore"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
)

func TestSix910Handler_CreateCustomerAccountPage(t *testing.T) {
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

	var prod1 sdbi.Product
	prod1.ID = 2
	prod1.Desc = "test"
	sapi.MockProduct = &prod1

	var prod sdbi.Product
	prod.ID = 2
	prod.Desc = "test"

	var plst []sdbi.Product
	plst = append(plst, prod)
	sapi.MockProductList = &plst

	//-----------end mocking --------

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

	var c conts.CmsService
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	//ds.Delete("books1")
	c.Log = &l
	c.Store = ds.GetNew()

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

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"

	sh.Templates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.CreateCustomerAccountPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CreateCustomerAccount(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.ResponseID
	macres.Success = true
	macres.ID = 3

	sapi.MockAddCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CreateCustomerAccount(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CreateCustomerAccountFail(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.ResponseID
	//macres.Success = true
	macres.ID = 3

	sapi.MockAddCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CreateCustomerAccount(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CreateCustomerAccountExisting(t *testing.T) {
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

	var cusm sdbi.Customer
	cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.ResponseID
	macres.Success = true
	macres.ID = 3

	sapi.MockAddCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CreateCustomerAccount(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCustomerAccountPage(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

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
		"email": "bob@bob.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateCustomerAccountPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCustomerAccountPageLogin(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

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
		"email": "bob@bob.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateCustomerAccountPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCustomerAccount(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.Response
	macres.Success = true

	sapi.MockUpdateCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("email=bob@bob.com&firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateCustomerAccount(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCustomerAccountLogin(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.Response
	macres.Success = true

	sapi.MockUpdateCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("email=bob@bob.com&firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateCustomerAccount(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_UpdateCustomerAccountFail(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.Response
	//macres.Success = true

	sapi.MockUpdateCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("email=bob@bob.com&firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.UpdateCustomerAccount(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CustomerAddAddressPage(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

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
		"email": "bob@bob.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.CustomerAddAddressPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_CustomerAddAddressPageLogin(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

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
		"email": "bob@bob.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Save(r, w)
	h := sh.GetNew()
	h.CustomerAddAddressPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CustomerAddAddress(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.Response
	macres.Success = true

	sapi.MockUpdateCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("email=bob@bob.com&firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.CustomerAddAddress(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CustomerAddAddressLogin(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.Response
	macres.Success = true

	sapi.MockUpdateCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("email=bob@bob.com&firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CustomerAddAddress(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_CustomerAddAddressFail(t *testing.T) {
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

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.Response
	macres.Success = true

	sapi.MockUpdateCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	//aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("email=bob@bob.com&firstName=tester&"+
		"lastName=testertest&zip=12345&billAddress=123&billCity=dd&billState=rr&billZip=22&"+
		"billCountry=55&shipAddress=444&shipCity=444&shipState=dfg&shipZip=234&shipCountry=55"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	var cccs m.CustomerCart
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Values["customerCart"] = &cccs
	s.Save(r, w)
	h := sh.GetNew()
	h.CustomerAddAddress(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCustomerAddress(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.Response
	cires.Success = true

	sapi.MockDeleteAddressRes = &cires

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
		"id":  "3",
		"cid": "4",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Save(r, w)
	h := sh.GetNew()
	h.DeleteCustomerAddress(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCustomerAddressLogin(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.Response
	cires.Success = true

	sapi.MockDeleteAddressRes = &cires

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
		"id":  "3",
		"cid": "4",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Save(r, w)
	h := sh.GetNew()
	h.DeleteCustomerAddress(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_DeleteCustomerAddressFail(t *testing.T) {
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

	var prod sdbi.Customer
	prod.ID = 2
	prod.Email = "bob@bob.com"
	sapi.MockCustomer = &prod

	var crtres api.ResponseID
	crtres.ID = 3
	crtres.Success = true

	sapi.MockAddCartResp = &crtres

	var cires api.Response
	//cires.Success = true

	sapi.MockDeleteAddressRes = &cires

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
		"id":  "3",
		"cid": "4",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = 55
	s.Save(r, w)
	h := sh.GetNew()
	h.DeleteCustomerAddress(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
