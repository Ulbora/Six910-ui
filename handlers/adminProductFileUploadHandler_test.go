package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	m "github.com/Ulbora/Six910-ui/managers"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	oauth2 "github.com/Ulbora/go-oauth2-client"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Handler_StoreAdminUploadProductFilePage(t *testing.T) {
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
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFilePage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUploadProductFilePageLoggedIn(t *testing.T) {
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
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFilePage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUploadProductFile(t *testing.T) {
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
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------
	// var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dlst []sdbi.Distributor
	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = true

	sapi.MockAddProductCategoryResp = &cr

	//-----------end mocking --------

	//file, err := os.Open("../../testUploadFile.csv")
	file, err := os.Open("../../testUploadFile.tar.gz")
	if err != nil {
		fmt.Println("file open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("csv fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("productupload", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	writer.WriteField("sleeptime", "20")

	r, _ := http.NewRequest("POST", "/test", body)
	//r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("csv upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}

	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFile(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUploadProductFileNotGzfile(t *testing.T) {
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
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------
	// var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dlst []sdbi.Distributor
	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = true

	sapi.MockAddProductCategoryResp = &cr

	//-----------end mocking --------

	file, err := os.Open("../../testUploadFile.csv")
	//file, err := os.Open("../../testUploadFile.tar.gz")
	if err != nil {
		fmt.Println("file open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("csv fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("productupload", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	writer.WriteField("sleeptime", "20")

	r, _ := http.NewRequest("POST", "/test", body)
	//r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("csv upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}

	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFile(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUploadProductFileLogin(t *testing.T) {
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
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------
	// var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dlst []sdbi.Distributor
	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = true

	sapi.MockAddProductCategoryResp = &cr

	//-----------end mocking --------

	//file, err := os.Open("../../testUploadFile.csv")
	file, err := os.Open("../../testUploadFile.tar.gz")
	if err != nil {
		fmt.Println("file test backup open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("csv fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("productupload", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	writer.WriteField("sleeptime", "20")

	r, _ := http.NewRequest("POST", "/test", body)
	//r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("csv upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}

	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = false
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFile(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUploadProductFileFail(t *testing.T) {
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
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------
	// var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dlst []sdbi.Distributor
	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

	var pr api.ResponseID
	//pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = false

	sapi.MockAddProductCategoryResp = &cr

	//-----------end mocking --------

	//file, err := os.Open("../../testUploadFile.csv")
	file, err := os.Open("../../testUploadFile.tar.gz")
	if err != nil {
		fmt.Println("file test backup open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("csv fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("productupload", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	writer.WriteField("sleeptime", "20")

	r, _ := http.NewRequest("POST", "/test", body)
	//r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("csv upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}

	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFile(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminUploadProductFileOauth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.OAuth2Enabled = true

	var mTkn oauth2.Token
	mTkn.AccessToken = "45ffffff"

	var mockAcTkn oauth2.MockAuthCodeToken
	mockAcTkn.MockToken = &mTkn

	sh.Auth = &mockAcTkn
	sh.token = &mTkn

	var sapi mapi.MockAPI
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	var man m.Six910Manager
	man.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	//-----------start mocking------------------
	// var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dlst []sdbi.Distributor
	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = true

	sapi.MockAddProductCategoryResp = &cr

	//-----------end mocking --------

	//file, err := os.Open("../../testUploadFile.csv")
	file, err := os.Open("../../testUploadFile.tar.gz")
	if err != nil {
		fmt.Println("file test backup open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("csv fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("productupload", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	writer.WriteField("sleeptime", "20")

	r, _ := http.NewRequest("POST", "/test", body)
	//r.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	//r.Header.Set("Content-Type", "application/json")

	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("csv upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}

	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.StoreAdminUploadProductFile(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
