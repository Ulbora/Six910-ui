package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func TestSix910Handler_StoreAdminGetPageCSS(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs csssrv.Six910CSSService
	cs.CSSStorePath = "../csssrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../csssrv/testFiles"
	cs.CSSStore = cds.GetNew()
	sh.CSSService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testPage",
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
	h.StoreAdminGetPageCSS(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminGetPageCSSLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs csssrv.Six910CSSService
	cs.CSSStorePath = "../csssrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../csssrv/testFiles"
	cs.CSSStore = cds.GetNew()
	sh.CSSService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testPage",
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
	h.StoreAdminGetPageCSS(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdatePageCSS(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs csssrv.Six910CSSService
	cs.CSSStorePath = "../csssrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../csssrv/testFiles"
	cs.CSSStore = cds.GetNew()
	sh.CSSService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=testPage&background=red&color=blue&pageTitle=white"))
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
	h.StoreAdminUpdatePageCSS(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdatePageCSSFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs csssrv.Six910CSSService
	cs.CSSStorePath = "../csssrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../csssrv/testFiles22"
	cs.CSSStore = cds.GetNew()
	sh.CSSService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=testPage&background=red&color=blue&pageTitle=white"))
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
	h.StoreAdminUpdatePageCSS(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdatePageCSSLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs csssrv.Six910CSSService
	cs.CSSStorePath = "../csssrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../csssrv/testFiles"
	cs.CSSStore = cds.GetNew()
	sh.CSSService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=testPage&background=red&color=blue&pageTitle=white"))
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
	h.StoreAdminUpdatePageCSS(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
