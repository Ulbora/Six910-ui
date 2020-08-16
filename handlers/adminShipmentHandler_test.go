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

func TestSix910Handler_StoreAdminAddShipmentPage(t *testing.T) {
	var sh Six910Handler

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	sh.API = &sapi

	//-----------start mocking------------------

	var odr sdbi.Order
	odr.ID = 1
	odr.OrderNumber = "O123"

	sapi.MockOrder = &odr

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 11
	oc.OrderID = 1
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	//-----------end mocking --------

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
	h.StoreAdminAddShipmentPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddShipmentPageLogin(t *testing.T) {
	var sh Six910Handler

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	sh.API = &sapi

	//-----------start mocking------------------

	var odr sdbi.Order
	odr.ID = 1
	odr.OrderNumber = "O123"

	sapi.MockOrder = &odr

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 11
	oc.OrderID = 1
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	//-----------end mocking --------

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
	h.StoreAdminAddShipmentPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddShipment(t *testing.T) {
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
	sapi.MockAddShipmentResp = &pr

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)

	var oi2 sdbi.OrderItem
	oi2.ID = 222
	oi2.OrderID = 1
	oi2.ProductName = "stuff2"
	oi2.ProductID = 2222
	oilst = append(oilst, oi2)

	var oi3 sdbi.OrderItem
	oi3.ID = 2223
	oi3.OrderID = 1
	oi3.ProductName = "stuff23"
	oi3.ProductID = 22223
	oilst = append(oilst, oi3)

	sapi.MockOrderItemList = &oilst

	var sir api.ResponseID
	sir.ID = 66
	sir.Success = true
	sapi.MockAddShipmentItemResp = &sir

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("status=shipped&orderId=5"))
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
	h.StoreAdminAddShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddShipmentItemInsertFail(t *testing.T) {
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
	sapi.MockAddShipmentResp = &pr

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)

	var oi2 sdbi.OrderItem
	oi2.ID = 222
	oi2.OrderID = 1
	oi2.ProductName = "stuff2"
	oi2.ProductID = 2222
	oilst = append(oilst, oi2)

	var oi3 sdbi.OrderItem
	oi3.ID = 2223
	oi3.OrderID = 1
	oi3.ProductName = "stuff23"
	oi3.ProductID = 22223
	oilst = append(oilst, oi3)

	sapi.MockOrderItemList = &oilst

	var sir api.ResponseID
	sir.ID = 66
	//sir.Success = true
	sapi.MockAddShipmentItemResp = &sir

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("status=shipped&orderId=5"))
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
	h.StoreAdminAddShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddShipmentLogin(t *testing.T) {
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
	sapi.MockAddShipmentResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("id=33&status=tester"))
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
	h.StoreAdminAddShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddShipmentFail(t *testing.T) {
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
	sapi.MockAddShipmentResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("id=33&status=tester"))
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
	h.StoreAdminAddShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditShipmentPage(t *testing.T) {
	var sh Six910Handler

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	sh.API = &sapi

	//-----------start mocking------------------

	var shp sdbi.Shipment
	shp.ID = 44
	shp.OrderID = 1
	sapi.MockShipment = &shp

	var sbx sdbi.ShipmentBox
	sbx.ID = 8
	sbx.ShipmentID = 44
	var sblst []sdbi.ShipmentBox
	sblst = append(sblst, sbx)
	sapi.MockShipmentBoxList = &sblst

	var si sdbi.ShipmentItem
	si.ID = 5
	si.ShipmentID = 44
	si.Quantity = 4

	var silst []sdbi.ShipmentItem
	silst = append(silst, si)
	sapi.MockShippingItemList = &silst

	var odr sdbi.Order
	odr.ID = 1
	odr.OrderNumber = "O123"

	sapi.MockOrder = &odr

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 11
	oc.OrderID = 1
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	//-----------end mocking --------

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
	h.StoreAdminEditShipmentPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditShipmentPageLogin(t *testing.T) {
	var sh Six910Handler

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	sh.API = &sapi

	//-----------start mocking------------------

	var shp sdbi.Shipment
	shp.ID = 44
	shp.OrderID = 1
	sapi.MockShipment = &shp

	var sbx sdbi.ShipmentBox
	sbx.ID = 8
	sbx.ShipmentID = 44
	var sblst []sdbi.ShipmentBox
	sblst = append(sblst, sbx)
	sapi.MockShipmentBoxList = &sblst

	var si sdbi.ShipmentItem
	si.ID = 5
	si.ShipmentID = 44
	si.Quantity = 4

	var silst []sdbi.ShipmentItem
	silst = append(silst, si)
	sapi.MockShippingItemList = &silst

	var odr sdbi.Order
	odr.ID = 1
	odr.OrderNumber = "O123"

	sapi.MockOrder = &odr

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 11
	oc.OrderID = 1
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	//-----------end mocking --------

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
	s.Values["storeAdminUser"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminEditShipmentPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditShipment(t *testing.T) {
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
	sapi.MockUpdateShipmentResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("id=3&status=shipped&orderId=5"))
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
	h.StoreAdminEditShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditShipmentLogin(t *testing.T) {
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
	sapi.MockUpdateShipmentResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("id=3&status=shipped&orderId=5"))
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
	h.StoreAdminEditShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminEditShipmentFail(t *testing.T) {
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
	sapi.MockUpdateShipmentResp = &pr

	//-----------end mocking --------

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("id=3&status=shipped&orderId=5"))
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
	h.StoreAdminEditShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminViewShipmentList(t *testing.T) {
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

	var odr sdbi.Order
	odr.ID = 1
	odr.OrderNumber = "O123"

	sapi.MockOrder = &odr

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 11
	oc.OrderID = 1
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	var shp sdbi.Shipment
	shp.ID = 3
	shp.OrderID = 1

	var shlst []sdbi.Shipment
	shlst = append(shlst, shp)

	sapi.MockShipmentList = &shlst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"oid": "4",
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
	h.StoreAdminViewShipmentList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminViewShipmentListLogin(t *testing.T) {
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

	var odr sdbi.Order
	odr.ID = 1
	odr.OrderNumber = "O123"

	sapi.MockOrder = &odr

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 11
	oc.OrderID = 1
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)
	sapi.MockCommentList = &oclst

	var oi sdbi.OrderItem
	oi.ID = 22
	oi.OrderID = 1
	oi.ProductName = "stuff"
	oi.ProductID = 222
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	var shp sdbi.Shipment
	shp.ID = 3
	shp.OrderID = 1

	var shlst []sdbi.Shipment
	shlst = append(shlst, shp)

	sapi.MockShipmentList = &shlst

	//-----------end mocking --------

	r, _ := http.NewRequest("GET", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
	vars := map[string]string{
		"oid": "4",
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
	h.StoreAdminViewShipmentList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteShipment(t *testing.T) {
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
	sapi.MockDeleteShipmentResp = &dres

	//-----------end mocking --------

	r, _ := http.NewRequest("DELETE", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
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
	h.StoreAdminDeleteShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteShipmentLogin(t *testing.T) {
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
	sapi.MockDeleteShipmentResp = &dres

	//-----------end mocking --------

	r, _ := http.NewRequest("DELETE", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
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
	h.StoreAdminDeleteShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteShipmentFail(t *testing.T) {
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
	sapi.MockDeleteShipmentResp = &dres

	//-----------end mocking --------

	r, _ := http.NewRequest("DELETE", "https://test.com", strings.NewReader("sku=tester123&name=tester"))
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
	h.StoreAdminDeleteShipment(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
