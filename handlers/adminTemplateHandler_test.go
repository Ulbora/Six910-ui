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
	"strings"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func TestSix910Handler_AdminAddTemplatePage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testBackup/templateStore"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testBackup/templateStore"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

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
	h.AdminAddTemplatePage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminAddTemplatePageAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testBackup/templateStore"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testBackup/templateStore"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

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
	h.AdminAddTemplatePage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUploadTemplate(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../templatesrv/testUploads/testTemplate.tar.gz")
	if err != nil {
		fmt.Println("file test template open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("template fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("tempFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("template upload file writer.FormDataContentType() : ", writer.FormDataContentType())
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
	h.AdminUploadTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUploadTemplateFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../templatesrv/testUploads/testTemplate.tar.gz")
	if err != nil {
		fmt.Println("file test template open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("template fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("tempFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("template upload file writer.FormDataContentType() : ", writer.FormDataContentType())
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
	h.AdminUploadTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUploadTemplateAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../templatesrv/testUploads/testTemplate.tar.gz")
	if err != nil {
		fmt.Println("file test template open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("template fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("tempFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("template upload file writer.FormDataContentType() : ", writer.FormDataContentType())
	err = writer.Close()
	if err != nil {
		fmt.Println(" writer.Close err: ", err)
	}
	w := httptest.NewRecorder()
	s, suc := sh.getSession(r)
	fmt.Println("suc: ", suc)
	//s.Values["loggedIn"] = true
	s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.AdminUploadTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminTemplateList(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

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
	h.AdminTemplateList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminTemplateListAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

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
	h.AdminTemplateList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminActivateTemplate(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.ActiveTemplateLocation = "../templatesrv/testDownloads"

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testTemplate",
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
	h.AdminActivateTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminActivateTemplateAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.ActiveTemplateLocation = "../templatesrv/testDownloads"

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testTemplate",
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
	h.AdminActivateTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminDeleteTemplate(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.ActiveTemplateLocation = "../templatesrv/testDownloads"

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testTemplate",
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
	h.AdminDeleteTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminDeleteTemplateAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.ActiveTemplateLocation = "../templatesrv/testDownloads"

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testTemplate",
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
	h.AdminDeleteTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminActivateTemplate2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.ActiveTemplateLocation = "../templatesrv/testDownloads"

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testTemplate2",
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
	h.AdminActivateTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminDeleteTemplate2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.ActiveTemplateLocation = "../templatesrv/testDownloads"

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", strings.NewReader("name=tester123"))
	vars := map[string]string{
		"name": "testTemplate",
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
	h.AdminDeleteTemplate(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
