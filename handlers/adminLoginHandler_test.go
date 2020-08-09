package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
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
}
