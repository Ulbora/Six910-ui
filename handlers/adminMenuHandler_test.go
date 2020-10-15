package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	sr "github.com/Ulbora/Six910-ui/contentsrv"
	isrv "github.com/Ulbora/Six910-ui/imgsrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func TestSix910Handler_StoreAdminMenuList(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	ci.Store = cds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
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

	var m2 musrv.Menu
	m2.Name = "menu2"
	m2.Active = true
	m2.Location = "top"
	m2.Shade = "light"
	m2.Background = "light"
	m2.Style = ""
	m2.ShadeList = &[]string{"light", "dark"}
	m2.BackgroundList = &[]string{"light", "dark"}

	msuc2 := ms.AddMenu(&m2)
	fmt.Println("menu save: ", msuc2)

	sh.MenuService = ms

	sh.ImageService = iss.GetNew()
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
	h.StoreAdminMenuList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminMenuListLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var cds ds.DataStore
	cds.Path = "../contentsrv/testFiles"
	ci.Store = cds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
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

	var m2 musrv.Menu
	m2.Name = "menu2"
	m2.Active = true
	m2.Location = "top"
	m2.Shade = "light"
	m2.Background = "light"
	m2.Style = ""
	m2.ShadeList = &[]string{"light", "dark"}
	m2.BackgroundList = &[]string{"light", "dark"}

	msuc2 := ms.AddMenu(&m2)
	fmt.Println("menu save: ", msuc2)

	sh.MenuService = ms

	sh.ImageService = iss.GetNew()
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
	h.StoreAdminMenuList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminGetMenu(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

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

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "menu1",
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
	h.StoreAdminGetMenu(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminGetMenuLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

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

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "menu1",
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
	h.StoreAdminGetMenu(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddMenuPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

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
	h.StoreAdminAddMenuPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddMenuPageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

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
	h.StoreAdminAddMenuPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddMenu(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com",
		strings.NewReader("name=tester123&location=123&active=on"))
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
	h.StoreAdminAddMenu(w, r)
	fmt.Println("code: ", w.Code)
	//sh.MenuService.DeleteMenu("tester123")

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddMenuFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFilesww"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com",
		strings.NewReader("name=tester123&location=123"))
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
	h.StoreAdminAddMenu(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddMenuLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com",
		strings.NewReader("name=tester123&location=123&active=on"))
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
	h.StoreAdminAddMenu(w, r)
	fmt.Println("code: ", w.Code)
	//sh.MenuService.DeleteMenu("tester123")

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateMenu(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com",
		strings.NewReader("name=tester123&location=123&active=on"))
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
	h.StoreAdminUpdateMenu(w, r)
	fmt.Println("code: ", w.Code)
	sh.MenuService.DeleteMenu("tester123")
	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateMenuFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFilesee"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com",
		strings.NewReader("name=tester123&location=123&active=on"))
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
	h.StoreAdminUpdateMenu(w, r)
	fmt.Println("code: ", w.Code)
	sh.MenuService.DeleteMenu("tester123")
	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateMenuLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sms musrv.Six910MenuService
	var mds ds.DataStore
	mds.Path = "../menusrv/testFiles"
	sms.MenuStore = mds.GetNew()
	sms.Log = &l
	ms := sms.GetNew()

	sh.MenuService = ms

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com",
		strings.NewReader("name=tester123&location=123&active=on"))
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
	h.StoreAdminUpdateMenu(w, r)
	fmt.Println("code: ", w.Code)
	sh.MenuService.DeleteMenu("tester123")
	if w.Code != 302 {
		t.Fail()
	}
}
