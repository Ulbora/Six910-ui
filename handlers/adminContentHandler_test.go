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
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func TestSix910Handler_StoreAdminAddContentPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

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
	h.StoreAdminAddContentPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminAddContentPageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

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
	h.StoreAdminAddContentPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
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
	h.StoreAdminAddContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddContent2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
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
	h.StoreAdminAddContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddContentLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
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
	h.StoreAdminAddContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123&archived=on&blogpost=on&visible=on"))
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
	h.StoreAdminUpdateContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateContentFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles2"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles2"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
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
	h.StoreAdminUpdateContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateContentLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
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
	h.StoreAdminUpdateContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminGetContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "tester123",
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
	h.StoreAdminGetContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminGetContentLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "tester123",
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
	h.StoreAdminGetContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminContentList(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

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
	h.StoreAdminContentList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminContentListLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

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
	h.StoreAdminContentList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminDeleteContent(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "tester123",
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
	h.StoreAdminDeleteContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteContentLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ci sr.CmsService
	ci.ContentStorePath = "../contentsrv/testFiles"

	ci.Log = &l
	var ds ds.DataStore
	ds.Path = "../contentsrv/testFiles"
	ci.Store = ds.GetNew()
	sh.ContentService = ci.GetNew()

	ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "tester123",
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
	h.StoreAdminDeleteContent(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
