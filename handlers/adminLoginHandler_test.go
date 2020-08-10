package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	userv "github.com/Ulbora/Six910-ui/usersrv"
	api "github.com/Ulbora/Six910API-Go"
	oauth2 "github.com/Ulbora/go-oauth2-client"
)

func TestSix910Handler_getRedirectURI(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	uri := h.getRedirectURI(r, "/dosometesting")
	if uri != "https://test.com/dosometesting" {
		t.Fail()
	}
}

func TestSix910Handler_getRedirectURI2(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	h.OauthHost = "test.com"
	h.SchemeDefault = "https://"
	r, _ := http.NewRequest("POST", "/test.com", nil)
	r.Host = "test.com"
	fmt.Println("req host: ", r.Host)
	uri := h.getRedirectURI(r, "/dosometesting")
	if uri != "https://test.com/dosometesting" {
		t.Fail()
	}
}

func TestSix910Handler_authorize(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	var cc ClientCreds
	cc.AuthCodeState = "123"
	h.ClientCreds = &cc
	h.ClientCreds.AuthCodeClient = "1"
	h.OauthHost = "test.com"

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	suc := h.authorize(w, r)
	if !suc {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminLogin(t *testing.T) {
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
	h := sh.GetNew()
	h.StoreAdminLogin(w, r)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminLoginOauth(t *testing.T) {
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
	h := sh.GetNew()
	h.StoreAdminLogin(w, r)

	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminHandleToken(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	var mockAcTkn oauth2.MockAuthCodeToken
	mockAcTkn.MockToken = &mTkn

	h.Auth = &mockAcTkn

	var cc ClientCreds
	cc.AuthCodeState = "123"
	cc.AuthCodeClient = "2"
	cc.AuthCodeSecret = "12345"
	h.ClientCreds = &cc
	h.OauthHost = "http://test12.com"
	r, _ := http.NewRequest("POST", "https://test.com?code=555&state=123", nil)
	w := httptest.NewRecorder()
	h.StoreAdminHandleToken(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_HandleLogout(t *testing.T) {
	var sh Six910Handler
	//h.TokenMap = make(map[string]*oauth2.Token)

	// var mTkn oauth2.Token
	// mTkn.AccessToken = "45ffffff"

	// var mockAcTkn oauth2.MockAuthCodeToken
	// mockAcTkn.MockToken = &mTkn

	// h.AuthToken = &mockAcTkn

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)

	s.Values["accessTokenKey"] = "123"

	w := httptest.NewRecorder()
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminLogout(w, r)
	fmt.Println("code: ", w.Code)
	if w.Code != 302 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminLoginNonOAuthUser(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	//-----------start mocking------------------
	var user api.UserResponse
	user.Username = "tester123"
	user.Role = storeAdmin
	user.Enabled = true

	sapi.MockUser = &user

	var ur api.Response
	ur.Success = true
	sapi.MockUpdateUserResp = &ur

	sh.API = &sapi

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester123&password=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := sh.GetNew()
	h.StoreAdminLoginNonOAuthUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminLoginNonOAuthUserNotLoggedIn(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	//-----------start mocking------------------
	var user api.UserResponse
	user.Username = "tester123"
	user.Role = storeAdmin
	//user.Enabled = true

	sapi.MockUser = &user

	var ur api.Response
	ur.Success = true
	sapi.MockUpdateUserResp = &ur

	sh.API = &sapi

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester123&password=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := sh.GetNew()
	h.StoreAdminLoginNonOAuthUser(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminLoginNonOAuthUserBadSession(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	//-----------start mocking------------------
	var user api.UserResponse
	user.Username = "tester123"
	user.Role = storeAdmin
	user.Enabled = true

	sapi.MockUser = &user

	var ur api.Response
	ur.Success = true
	sapi.MockUpdateUserResp = &ur

	sh.API = &sapi

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "/test", strings.NewReader("username=tester123&password=tester"))
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	h := sh.GetNew()
	h.StoreAdminLoginNonOAuthUser(w, nil)
	fmt.Println("code: ", w.Code)

	if w.Code != 500 {
		t.Fail()
	}
}

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
