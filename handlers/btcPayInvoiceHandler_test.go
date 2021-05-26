package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	cl "github.com/Ulbora/BTCPayClient"
	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	pi "github.com/Ulbora/Six910BTCPayServerPlugin"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
)

func TestSix910Handler_CompleteBTCPayTransaction(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ppi pi.PayPlugin

	var mc pi.MockBTCPayClient
	mc.MockClientID = "eeeddd"
	var tknr cl.TokenResponse
	var tkn cl.TokenData
	tkn.Token = "1123aaa"
	tkn.ParingCode = "pa111"
	tknr.Data = []cl.TokenData{tkn}
	mc.MockTokenResponse = &tknr
	mc.MockPairingCodeURL = "http://test.com/pair/123"

	var invres cl.InvoiceResponse
	invres.Data.URL = "http://test"
	mc.MockInvoiceResponse = &invres

	ppi.SetClient(mc.New())

	sh.BTCPlugin = ppi.New()
	sh.BTCPayCurrency = "USD"

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi

	//-----------start mocking------------------

	var pr sdbi.PaymentGateway
	pr.CheckoutURL = "http://test/test"
	pr.ID = 5
	pr.ClientID = "Tf4UFY3a3XyZX8PV4wj7zDaehQam4SkeTzq"
	pr.ClientKey = "74f522c6704e39d102db6ae98dfb286e11c678a76fa32c93cc50244e436d936f"
	pr.Token = "12345555"
	pr.PostOrderURL = "http://test.com"
	sapi.MockPaymentGateway = &pr

	//-----------end mocking

	r, _ := http.NewRequest("GET", "https://test.com", nil)
	vars := map[string]string{
		"total":     "21.99",
		"tax":       "1.54",
		"firstName": "Ricky",
		"lastName":  "Bobby",
		"email":     "ricky2@bobby.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = true
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.CompleteBTCPayTransaction(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()

}

func TestSix910Handler_CompleteBTCPayTransactionAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ppi pi.PayPlugin

	var mc pi.MockBTCPayClient
	mc.MockClientID = "eeeddd"
	var tknr cl.TokenResponse
	var tkn cl.TokenData
	tkn.Token = "1123aaa"
	tkn.ParingCode = "pa111"
	tknr.Data = []cl.TokenData{tkn}
	mc.MockTokenResponse = &tknr
	mc.MockPairingCodeURL = "http://test.com/pair/123"

	var invres cl.InvoiceResponse
	invres.Data.URL = "http://test"
	mc.MockInvoiceResponse = &invres

	ppi.SetClient(mc.New())

	sh.BTCPlugin = ppi.New()
	sh.BTCPayCurrency = "USD"

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")
	sh.API = &sapi

	//-----------start mocking------------------

	var pr sdbi.PaymentGateway
	pr.CheckoutURL = "http://test/test"
	pr.ID = 5
	pr.ClientID = "Tf4UFY3a3XyZX8PV4wj7zDaehQam4SkeTzq"
	pr.ClientKey = "74f522c6704e39d102db6ae98dfb286e11c678a76fa32c93cc50244e436d936f"
	pr.Token = "12345555"
	pr.PostOrderURL = "http://test.com"
	sapi.MockPaymentGateway = &pr

	//-----------end mocking

	r, _ := http.NewRequest("GET", "https://test.com", nil)
	vars := map[string]string{
		"total":     "21.99",
		"tax":       "1.54",
		"firstName": "Ricky",
		"lastName":  "Bobby",
		"email":     "ricky2@bobby.com",
	}
	r = mux.SetURLVars(r, vars)
	w := httptest.NewRecorder()
	s, suc := sh.getUserSession(w, r)
	fmt.Println("suc: ", suc)
	s.Values["userLoggenIn"] = false
	s.Values["customerUser"] = true
	s.Values["customerId"] = int64(55)
	s.Values["username"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.CompleteBTCPayTransaction(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

	// t.Fail()

}

func TestSix910Handler_checkBTCPayPlugin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ppi pi.PayPlugin
	sh.BTCPlugin = ppi.New()

	var pg sdbi.PaymentGateway
	pg.CheckoutURL = "http://test/test"
	pg.ID = 5
	pg.ClientID = "Tf4UFY3a3XyZX8PV4wj7zDaehQam4SkeTzq"
	pg.ClientKey = "74f522c6704e39d102db6ae98dfb286e11c678a76fa32c93cc50244e436d936f"
	pg.Token = "12345555"
	pg.PostOrderURL = "http://test.com"

	sh.checkBTCPayPlugin(&pg)

	if !sh.BTCPlugin.IsPluginLoaded() {
		t.Fail()
	}

}
