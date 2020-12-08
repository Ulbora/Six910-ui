package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	m "github.com/Ulbora/Six910-ui/managers"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	userv "github.com/Ulbora/Six910-ui/usersrv"
	usrv "github.com/Ulbora/Six910-ui/usersrv"
	api "github.com/Ulbora/Six910API-Go"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	"github.com/gorilla/mux"
)

func TestSix910Handler_StoreAdminChangePassword(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
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
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangePassword(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminChangeOauthPassword(t *testing.T) {
	var sh Six910Handler
	sh.OAuth2Enabled = true
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	sh.token = &mTkn

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
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangePassword(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminChangePasswordNotLoggedin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
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
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangePassword(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminChangeOauthPasswordNotLoggedin(t *testing.T) {
	var sh Six910Handler
	sh.OAuth2Enabled = true
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
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
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangePassword(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminChangeUserPassword(t *testing.T) {
	var sh Six910Handler
	var mockUserService userv.MockOauth2UserService
	var ures userv.UserResponse
	ures.Success = true

	mockUserService.MockUpdateUserResponse = &ures
	sh.UserService = &mockUserService
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	sh.token = &mTkn
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["storeAdminUser"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangeUserPassword(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminChangeUserPasswordNotLoggedIn(t *testing.T) {
	var sh Six910Handler
	var mockUserService userv.MockOauth2UserService
	var ures userv.UserResponse
	ures.Success = true

	mockUserService.MockUpdateUserResponse = &ures
	sh.UserService = &mockUserService
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	sh.token = &mTkn
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["userLoggenIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangeUserPassword(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminChangeUserPasswordUpdateFail(t *testing.T) {
	var sh Six910Handler
	var mockUserService userv.MockOauth2UserService
	var ures userv.UserResponse
	//ures.Success = true

	mockUserService.MockUpdateUserResponse = &ures
	sh.UserService = &mockUserService
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	sh.ClientCreds.AuthCodeClient = "1"

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	sh.token = &mTkn
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["storeAdminUser"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminChangeUserPassword(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAdminUserList(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	var flst []api.UserResponse
	flst = append(flst, usr)
	sapi.MockUserList = &flst

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
	h.StoreAdminAdminUserList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAdminUserListAuth(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	var flst []api.UserResponse
	flst = append(flst, usr)
	sapi.MockUserList = &flst

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
	h.StoreAdminAdminUserList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAdminUserListOauth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.OAuth2Enabled = true

	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"username":"tester", "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	var oa2us usrv.Oauth2UserService
	oa2us.ClientID = "44"
	oa2us.UserHost = "http://test"
	oa2us.Log = &l
	oa2us.Proxy = p.GetNewProxy()
	sh.UserService = &oa2us
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	sh.token = &mTkn

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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	var flst []api.UserResponse
	flst = append(flst, usr)
	sapi.MockUserList = &flst

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
	h.StoreAdminAdminUserList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminCustomerUserList(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	var flst []api.UserResponse
	flst = append(flst, usr)
	sapi.MockUserList = &flst

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
	h.StoreAdminCustomerUserList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminCustomerUserListAuth(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	var flst []api.UserResponse
	flst = append(flst, usr)
	sapi.MockUserList = &flst

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
	h.StoreAdminCustomerUserList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUserPage(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	sapi.MockUser = &usr

	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"username":"tester", "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	var oa2us usrv.Oauth2UserService
	oa2us.ClientID = "44"
	oa2us.UserHost = "http://test"
	oa2us.Log = &l
	oa2us.Proxy = p.GetNewProxy()
	sh.UserService = &oa2us
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	sh.token = &mTkn

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"username": "tester",
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
	h.StoreAdminEditUserPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUserPage2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.OAuth2Enabled = true
	var tkn oauth2.Token
	sh.token = &tkn

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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	sapi.MockUser = &usr

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"username":"tester", "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	var oa2us usrv.Oauth2UserService
	oa2us.ClientID = "44"
	oa2us.UserHost = "http://test"
	oa2us.Log = &l
	oa2us.Proxy = p.GetNewProxy()
	sh.UserService = &oa2us
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	sh.token = &mTkn

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"username": "tester",
		"role":     "StoreAdmin",
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
	h.StoreAdminEditUserPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUserPageAuth(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	sapi.MockUser = &usr

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"username": "tester",
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
	h.StoreAdminEditUserPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddAdminUserPage(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	sapi.MockUser = &usr

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"username": "tester",
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
	h.StoreAdminAddAdminUserPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddAdminUserPageAuth(t *testing.T) {
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

	var usr api.UserResponse
	usr.Enabled = true
	usr.Role = storeAdmin
	usr.StoreID = 5
	usr.Username = "testAdmin"

	sapi.MockUser = &usr

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"username": "tester",
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
	h.StoreAdminAddAdminUserPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddAdminUser(t *testing.T) {
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

	var res api.Response
	res.Success = true
	sapi.MockAddAdminUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=true"))
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
	h.StoreAdminAddAdminUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddAdminUser2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.OAuth2Enabled = true

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

	var res api.Response
	res.Success = true
	sapi.MockAddAdminUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"username":"tester", "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	var oa2us usrv.Oauth2UserService
	oa2us.ClientID = "44"
	oa2us.UserHost = "http://test"
	oa2us.Log = &l
	oa2us.Proxy = p.GetNewProxy()
	sh.UserService = &oa2us
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"
	sh.token = &mTkn

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=true"))
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
	h.StoreAdminAddAdminUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddAdminUserAuth(t *testing.T) {
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

	var res api.Response
	res.Success = true
	sapi.MockAddAdminUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=true"))
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
	h.StoreAdminAddAdminUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddAdminUserFail(t *testing.T) {
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

	var res api.Response
	//res.Success = true
	sapi.MockAddAdminUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=true"))
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
	h.StoreAdminAddAdminUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUser(t *testing.T) {
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

	var res api.Response
	res.Success = true
	sapi.MockAdminUpdateUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=on&role=customer"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester123"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUserAuth(t *testing.T) {
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

	var res api.Response
	res.Success = true
	sapi.MockAdminUpdateUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=true&role=customer"))
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
	h.StoreAdminEditUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUser2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.OAuth2Enabled = true
	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	sh.token = &mTkn

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

	var res api.Response
	res.Success = true
	sapi.MockAdminUpdateUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[{"username":"tester", "clientId": 2}]`))
	p.MockResp = &ress
	p.MockRespCode = 200
	var oa2us usrv.Oauth2UserService
	oa2us.ClientID = "44"
	oa2us.UserHost = "http://test"
	oa2us.Log = &l
	oa2us.Proxy = p.GetNewProxy()
	sh.UserService = &oa2us
	var cc ClientCreds
	cc.AuthCodeState = "123"
	sh.ClientCreds = &cc
	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"
	// sh.token = &mTkn

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=off&role=StoreAdmin"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester123"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditUser3(t *testing.T) {
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

	var res api.Response
	res.Success = true
	sapi.MockAdminUpdateUserResp = &res

	// var flst []api.UserResponse
	// flst = append(flst, usr)
	// sapi.MockUserList = &flst

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("username=tester123&password=tester&enabled=on&role=admin"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester123"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
