package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http/httptest"

	"net/http"

	"testing"

	lg "github.com/Ulbora/Level_Logger"

	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Handler_getSession(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	ses, suc := h.getSession(r)
	if ses == nil || !suc {
		t.Fail()
	}
}

func TestSix910Handler_getUserSession(t *testing.T) {
	var h Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	h.Log = &l
	r, _ := http.NewRequest("POST", "https://test.com", nil)
	w := httptest.NewRecorder()
	ses, suc := h.getUserSession(w, r)
	if ses == nil || !suc {
		t.Fail()
	}
}

func TestSix910Handler_processProductMetaData(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.CompanyName = "Six910"
	sh.Six910SiteURL = "http://six910.com"

	var prod sdbi.Product
	prod.ID = 1234
	prod.Image1 = "http:/someimage/img.png"
	prod.ShortDesc = "16g161gf6156gf1d6f5g16d5f1g56df1g65df16g5df65g1d6f5g16d5fg165df1"
	prod.Desc = "16516d51fd5s1f65ds6f51ds65f1d6s5f6d5s1f65ds65ds1f65ds1665165f65ds165dfs165165ds1d56fs6d5s165fd56fs165dsf65dsf65ds6f5d1s65fds65f16sd51f65ds1f65ds16f51ds65fds651f65s1df651ds65f1d6s5f16d5sf65ds165ds6d5sf6"

	r, _ := http.NewRequest("POST", "https://test.com", nil)

	sitedata := sh.processProductMetaData(&prod, r)
	fmt.Println("sitedate:", sitedata)
	if sitedata.Title == "" {
		t.Fail()
	}
}

func TestSix910Handler_processProductMetaData2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.CompanyName = "Six910"
	sh.Six910SiteURL = "http://six910.com"

	var prod sdbi.Product
	prod.ID = 1234
	prod.Image1 = "/someimage/img.png"
	prod.ShortDesc = "16g161gf6156gf1d6f5g16d5f1g56df1g65df16g5df65g1d6f5g16d5fg165df1"
	prod.Desc = "165165ds1d56fs6dd1df651ds65f1d6s5f16d5sf65ds165ds6d5sf6"

	r, _ := http.NewRequest("POST", "/test.com", nil)

	sitedata := sh.processProductMetaData(&prod, r)
	fmt.Println("sitedate:", sitedata)
	if sitedata.Title == "" {
		t.Fail()
	}
}

func TestSix910Handler_processMetaData(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.CompanyName = "Six910"
	sh.Six910SiteURL = "http://six910.com"

	var prod sdbi.Product
	prod.ID = 1234
	prod.Image1 = "/someimage/img.png"
	prod.ShortDesc = "16g161gf6156gf1d6f5g16d5f1g56df1g65df16g5df65g1d6f5g16d5fg165df1"
	prod.Desc = "165165ds1d56fs6dd1df651ds65f1d6s5f16d5sf65ds165ds6d5sf6"

	r, _ := http.NewRequest("POST", "/test.com", nil)

	sitedata := sh.processMetaData("/test/test", "some category", r)
	fmt.Println("sitedate:", sitedata)
	if sitedata.Title == "" {
		t.Fail()
	}
}

func TestSix910Handler_processMetaData2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.CompanyName = "Six910"
	//sh.Six910SiteURL = "http://six910.com"

	var prod sdbi.Product
	prod.ID = 1234
	prod.Image1 = "/someimage/img.png"
	prod.ShortDesc = "16g161gf6156gf1d6f5g16d5f1g56df1g65df16g5df65g1d6f5g16d5fg165df1"
	prod.Desc = "165165ds1d56fs6dd1df651ds65f1d6s5f16d5sf65ds165ds6d5sf6"

	r, _ := http.NewRequest("POST", "http://test.com", nil)

	sitedata := sh.processMetaData("/test/test", "some category", r)
	fmt.Println("sitedate:", sitedata)
	if sitedata.Title == "" {
		t.Fail()
	}
}

type testObj struct {
	Valid bool   `json:"valid"`
	Code  string `json:"code"`
}

func TestSix910Handler_ProcessBodyBad(t *testing.T) {
	var oh Six910Handler
	var l lg.Logger
	oh.Log = &l
	var robj testObj
	robj.Valid = true
	robj.Code = "3"
	// var res http.Response
	// res.Body = ioutil.NopCloser(bytes.NewBufferString(`{"valid":true, "code":"1"}`))
	var sURL = "http://localhost/test"
	aJSON, _ := json.Marshal(robj)
	r, _ := http.NewRequest("POST", sURL, bytes.NewBuffer(aJSON))
	var obj testObj
	suc, _ := oh.ProcessBody(r, nil)
	if suc || obj.Valid != false || obj.Code != "" {
		t.Fail()
	}
}
