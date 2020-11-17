package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"net/http/httptest"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Handler_StoreAdminIndex(t *testing.T) {
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

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi
	//mock
	var vd sdbi.VisitorData
	var vdlst []sdbi.VisitorData
	vdlst = append(vdlst, vd)
	sapi.MockVisitorData = &vdlst

	var sd sdbi.OrderSalesData
	var sdlst []sdbi.OrderSalesData
	sdlst = append(sdlst, sd)
	sapi.MockOrderSalesData = &sdlst

	//end mock

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminIndex(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminIndexLoggedIn(t *testing.T) {
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

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi

	//mock
	var vd sdbi.VisitorData
	var vdlst []sdbi.VisitorData
	vdlst = append(vdlst, vd)
	sapi.MockVisitorData = &vdlst

	var sd sdbi.OrderSalesData
	var sdlst []sdbi.OrderSalesData
	sdlst = append(sdlst, sd)
	sapi.MockOrderSalesData = &sdlst

	//end mock

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminIndex(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}
