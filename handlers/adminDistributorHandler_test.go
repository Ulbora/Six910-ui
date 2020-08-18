package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	m "github.com/Ulbora/Six910-ui/managers"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
)

func TestSix910Handler_StoreAdminAddDistributorPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	//-----------end mocking --------

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminAddDistributorPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddDistributorPageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	//-----------end mocking --------

	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"
	sh.OauthHost = "test.com"
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminAddDistributorPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddDistributor(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddDistributorResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("company=tester123&contactName=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminAddDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddDistributorLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddDistributorResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("company=tester123&contactName=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminAddDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddDistributorFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.ResponseID
	//pr.Success = true
	pr.ID = 5
	sapi.MockAddDistributorResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("company=tester123&contactName=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminAddDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditDistributorPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr sdbi.Distributor
	pr.Company = "test"
	pr.ID = 5
	sapi.MockDistributor = &pr

	// var clts []sdbi.Category
	// clts = append(clts, pr)
	// sapi.MockCategoryList = &clts

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))

	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditDistributorPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditDistributorPageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr sdbi.Distributor
	pr.Company = "test"
	pr.ID = 5
	sapi.MockDistributor = &pr

	// var clts []sdbi.Category
	// clts = append(clts, pr)
	// sapi.MockCategoryList = &clts

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))

	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditDistributorPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditDistributor(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.Response
	pr.Success = true
	sapi.MockUpdateDistributorResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("PUT", "https://test.com", strings.NewReader("id=3&company=tester123&contactName=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditDistributorLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.Response
	pr.Success = true
	sapi.MockUpdateDistributorResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("PUT", "https://test.com", strings.NewReader("id=3&company=tester123&contactName=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditDistributorFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr api.Response
	//pr.Success = true
	sapi.MockUpdateDistributorResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("PUT", "https://test.com", strings.NewReader("id=3&company=tester123&contactName=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminViewDistributorList(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr sdbi.Distributor
	pr.Company = "test"
	pr.ID = 5

	var flst []sdbi.Distributor
	flst = append(flst, pr)
	sapi.MockDistributorList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"start": "0",
		"end":   "100",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminViewDistributorList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminViewDistributorListLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var pr sdbi.Distributor
	pr.Company = "test"
	pr.ID = 5

	var flst []sdbi.Distributor
	flst = append(flst, pr)
	sapi.MockDistributorList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"start": "0",
		"end":   "100",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminViewDistributorList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteDistributor(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var dres api.Response
	dres.Success = true
	sapi.MockDeleteDistributorResp = &dres

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminDeleteDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteDistributorLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var dres api.Response
	dres.Success = true
	sapi.MockDeleteDistributorResp = &dres

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminDeleteDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteDistributorFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	sh.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	var dres api.Response
	//dres.Success = true
	sapi.MockDeleteDistributorResp = &dres

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"id": "1",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminDeleteDistributor(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
