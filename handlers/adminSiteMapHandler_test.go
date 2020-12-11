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
)

func TestSix910Handler_GenerateSiteMap(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	// var man m.Six910Manager
	// man.API = &sapi
	sh.API = &sapi
	// man.Log = &l
	//sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	// var pr sdbi.ShippingMethod
	// pr.Cost = 56.25
	// pr.ID = 5

	// //var flst []sdbi.ShippingMethod
	// flst = append(flst, pr)
	// sapi.MockShippingMethodList = &flst

	var idlst []int64
	idlst = append(idlst, 5)
	idlst = append(idlst, 53)
	idlst = append(idlst, 55)
	sapi.MockProductIDList = &idlst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	// vars := map[string]string{
	// 	"start": "0",
	// 	"end":   "100",
	// }
	// r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.GenerateSiteMap(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_GenerateSiteMapAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	// var man m.Six910Manager
	// man.API = &sapi
	sh.API = &sapi
	// man.Log = &l
	//sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------

	// var pr sdbi.ShippingMethod
	// pr.Cost = 56.25
	// pr.ID = 5

	// //var flst []sdbi.ShippingMethod
	// flst = append(flst, pr)
	// sapi.MockShippingMethodList = &flst

	var idlst []int64
	idlst = append(idlst, 5)
	idlst = append(idlst, 53)
	idlst = append(idlst, 55)
	sapi.MockProductIDList = &idlst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	// vars := map[string]string{
	// 	"start": "0",
	// 	"end":   "100",
	// }
	// r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.GenerateSiteMap(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
