package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	carsrv "github.com/Ulbora/Six910-ui/carouselsrv"
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func TestSix910Handler_StoreAdminGetCarousel(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs carsrv.Six910CarouselService
	cs.StorePath = "../carouselsrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../carouselsrv/testFiles"
	cs.Store = cds.GetNew()
	sh.CarouselService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testCar",
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
	h.StoreAdminGetCarousel(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminGetCarouselLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs carsrv.Six910CarouselService
	cs.StorePath = "../carouselsrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../carouselsrv/testFiles"
	cs.Store = cds.GetNew()
	sh.CarouselService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testCar",
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
	h.StoreAdminGetCarousel(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateCarousel(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs carsrv.Six910CarouselService
	cs.StorePath = "../carouselsrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../carouselsrv/testFiles"
	cs.Store = cds.GetNew()
	sh.CarouselService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=testCar&enabled=on&image1=/img11&image2=/img22&image3=/img33"))
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
	h.StoreAdminUpdateCarousel(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateCarouselFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs carsrv.Six910CarouselService
	cs.StorePath = "../carouselsrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../carouselsrv/testFiles22"
	cs.Store = cds.GetNew()
	sh.CarouselService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=testCar&enabled=off&image1=/img11&image2=/img22&image3=/img33"))
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
	h.StoreAdminUpdateCarousel(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUpdateCarouselLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var cs carsrv.Six910CarouselService
	cs.StorePath = "../carouselsrv/testFiles"
	cs.Log = &l
	var cds ds.DataStore
	cds.Path = "../carouselsrv/testFiles"
	cs.Store = cds.GetNew()
	sh.CarouselService = cs.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=testCar&enabled=on&image1=/img11&image2=/img22&image3=/img33"))
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
	h.StoreAdminUpdateCarousel(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
