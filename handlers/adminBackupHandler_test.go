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
	bkupsrv "github.com/Ulbora/Six910-ui/bkupsrv"
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func TestSix910Handler_AdminBackupMainPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l
	var bds ds.DataStore
	bds.Path = "./bkupsrv/testBackup/contentStore"
	bs.TemplateStore = bds.GetNew()
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
	h.AdminBackupMainPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminBackupMainPageAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l
	var bds ds.DataStore
	bds.Path = "./bkupsrv/testBackup/contentStore"
	bs.TemplateStore = bds.GetNew()
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
	//s.Values["storeAdminUser"] = true
	s.Values["username"] = "tester"
	s.Values["password"] = "tester"
	s.Save(r, w)
	h := sh.GetNew()
	h.AdminBackupMainPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminBackupUploadPage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l
	var bds ds.DataStore
	bds.Path = "./bkupsrv/testBackup/contentStore"
	bs.TemplateStore = bds.GetNew()
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
	h.AdminBackupUploadPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminBackupUploadPageAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l
	var bds ds.DataStore
	bds.Path = "./bkupsrv/testBackup/contentStore"
	bs.TemplateStore = bds.GetNew()
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
	h.AdminBackupUploadPage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUploadBackups(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackupRestore/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackupRestore/contentStore"
	bs.CarouselStorePath = "../bkupsrv/testBackupRestore/carouselStore"
	bs.CountryStorePath = "../bkupsrv/testBackupRestore/countryStore"
	bs.CSSStorePath = "../bkupsrv/testBackupRestore/cssStore"
	bs.MenuStorePath = "../bkupsrv/testBackupRestore/menuStore"
	bs.StateStorePath = "../bkupsrv/testBackupRestore/stateStore"
	bs.ImagePath = "../bkupsrv/testBackupRestore/images"
	bs.TemplateFilePath = "../bkupsrv/testBackupRestore/templates"
	bs.Log = &l

	// var bds ds.DataStore
	// bds.Path = "./bkupsrv/testBackup/contentStore"
	// bs.TemplateStore = bds.GetNew()
	// //ch.Service = ci.GetNew()

	var cds ds.DataStore
	cds.Path = "../bkupsrv/testBackupRestore/contentStore"
	bs.Store = cds.GetNew()

	var tds ds.DataStore
	tds.Path = "../bkupsrv/testBackupRestore/templateStore"
	bs.TemplateStore = tds.GetNew()

	var crds ds.DataStore
	crds.Path = "../bkupsrv/testBackupRestore/carouselStore"
	bs.CarouselStore = crds.GetNew()

	var cyds ds.DataStore
	cyds.Path = "../bkupsrv/testBackupRestore/countryStore"
	bs.CountryStore = cyds.GetNew()

	var cssds ds.DataStore
	cssds.Path = "../bkupsrv/testBackupRestore/cssStore"
	bs.CSSStore = cssds.GetNew()

	var mds ds.DataStore
	mds.Path = "../bkupsrv/testBackupRestore/menuStore"
	bs.MenuStore = mds.GetNew()

	var sds ds.DataStore
	sds.Path = "../bkupsrv/testBackupRestore/stateStore"
	bs.StateStore = sds.GetNew()

	sh.BackupService = bs.GetNew()

	// sh.ActiveTemplateLocation = "../templatesrv/testDownloads"
	sh.ActiveTemplateLocation = "../bkupsrv/testBackupRestore/templates"

	var ts tmpsrv.Six910TemplateService
	// ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	// ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.TemplateStorePath = "../bkupsrv/testBackupRestore/templateStore"
	ts.TemplateFilePath = "../bkupsrv/testBackupRestore/templates"
	ts.Log = &l
	//var ttds ds.DataStore
	//ttds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../bkupsrv/testBackupZips/compress.dat")
	if err != nil {
		fmt.Println("file test backup open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("backup fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("backupFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("image upload file writer.FormDataContentType() : ", writer.FormDataContentType())

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
	h.AdminUploadBackups(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUploadBackupsAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.CarouselStorePath = "../bkupsrv/testBackup/carouselStore"
	bs.CountryStorePath = "../bkupsrv/testBackup/countryStore"
	bs.CSSStorePath = "../bkupsrv/testBackup/cssStore"
	bs.MenuStorePath = "../bkupsrv/testBackup/menuStore"
	bs.StateStorePath = "../bkupsrv/testBackup/stateStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l

	// var bds ds.DataStore
	// bds.Path = "./bkupsrv/testBackup/contentStore"
	// bs.TemplateStore = bds.GetNew()
	// //ch.Service = ci.GetNew()

	var cds ds.DataStore
	cds.Path = "../bkupsrv/testBackup/contentStore"
	bs.Store = cds.GetNew()

	var tds ds.DataStore
	tds.Path = "../bkupsrv/testBackup/templateStore"
	bs.TemplateStore = tds.GetNew()

	var crds ds.DataStore
	crds.Path = "../bkupsrv/testBackup/carouselStore"
	bs.CarouselStore = crds.GetNew()

	var cyds ds.DataStore
	cyds.Path = "../bkupsrv/testBackup/countryStore"
	bs.CountryStore = cyds.GetNew()

	var cssds ds.DataStore
	cssds.Path = "../bkupsrv/testBackup/cssStore"
	bs.CSSStore = cssds.GetNew()

	var mds ds.DataStore
	mds.Path = "../bkupsrv/testBackup/menuStore"
	bs.MenuStore = mds.GetNew()

	var sds ds.DataStore
	sds.Path = "../bkupsrv/testBackup/stateStore"
	bs.StateStore = sds.GetNew()

	sh.BackupService = bs.GetNew()

	// sh.ActiveTemplateLocation = "../templatesrv/testDownloads"
	sh.ActiveTemplateLocation = "../bkupsrv/testBackup/templates"

	var ts tmpsrv.Six910TemplateService
	// ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	// ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	ts.TemplateFilePath = "../bkupsrv/testBackup/templates"
	ts.Log = &l
	//var ttds ds.DataStore
	//ttds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../bkupsrv/testBackupZips/compress.dat")
	if err != nil {
		fmt.Println("file test backup open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("backup fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("backupFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("image upload file writer.FormDataContentType() : ", writer.FormDataContentType())

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
	h.AdminUploadBackups(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminUploadBackupsFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackupRestore/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackupRestore/contentStore"
	bs.CarouselStorePath = "../bkupsrv/testBackupRestore/carouselStore"
	bs.CountryStorePath = "../bkupsrv/testBackupRestore/countryStore"
	bs.CSSStorePath = "../bkupsrv/testBackupRestore/cssStore"
	bs.MenuStorePath = "../bkupsrv/testBackupRestore/menuStore"
	bs.StateStorePath = "../bkupsrv/testBackupRestore/stateStore"
	bs.ImagePath = "../bkupsrv/testBackupRestore/images"
	bs.TemplateFilePath = "../bkupsrv/testBackupRestore/templates"
	bs.Log = &l

	// var bds ds.DataStore
	// bds.Path = "./bkupsrv/testBackup/contentStore"
	// bs.TemplateStore = bds.GetNew()
	// //ch.Service = ci.GetNew()

	var cds ds.DataStore
	cds.Path = "../bkupsrv/testBackupRestore/contentStore"
	bs.Store = cds.GetNew()

	var tds ds.DataStore
	tds.Path = "../bkupsrv/testBackupRestore/templateStore"
	bs.TemplateStore = tds.GetNew()

	var crds ds.DataStore
	crds.Path = "../bkupsrv/testBackupRestore/carouselStore"
	bs.CarouselStore = crds.GetNew()

	var cyds ds.DataStore
	cyds.Path = "../bkupsrv/testBackupRestore/countryStore"
	bs.CountryStore = cyds.GetNew()

	var cssds ds.DataStore
	cssds.Path = "../bkupsrv/testBackupRestore/cssStore"
	bs.CSSStore = cssds.GetNew()

	var mds ds.DataStore
	mds.Path = "../bkupsrv/testBackupRestore/menuStore"
	bs.MenuStore = mds.GetNew()

	var sds ds.DataStore
	sds.Path = "../bkupsrv/testBackupRestore/stateStore"
	bs.StateStore = sds.GetNew()

	sh.BackupService = bs.GetNew()

	// sh.ActiveTemplateLocation = "../templatesrv/testDownloads"
	sh.ActiveTemplateLocation = "../bkupsrv/testBackupRestore/templates"

	var ts tmpsrv.Six910TemplateService
	// ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	// ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.TemplateStorePath = "../bkupsrv/testBackupRestore/templateStore"
	ts.TemplateFilePath = "../bkupsrv/testBackupRestore/templates"
	ts.Log = &l
	//var ttds ds.DataStore
	//ttds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../bkupsrv/testBackupZips/test.jpg")
	if err != nil {
		fmt.Println("file test backup open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("backup fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("backupFile", fi.Name())
	if err != nil {
		fmt.Println("create form err: ", err)
	}

	_, err = io.Copy(part, file)
	fmt.Println("io.Copy err: ", err)

	writer.WriteField("name", fi.Name())

	r, _ := http.NewRequest("POST", "/test", body)
	r.Header.Set("Content-Type", writer.FormDataContentType())
	fmt.Println("image upload file writer.FormDataContentType() : ", writer.FormDataContentType())

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
	h.AdminUploadBackups(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_AdminDownloadBackups(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.CarouselStorePath = "../bkupsrv/testBackup/carouselStore"
	bs.CountryStorePath = "../bkupsrv/testBackup/countryStore"
	bs.CSSStorePath = "../bkupsrv/testBackup/cssStore"
	bs.MenuStorePath = "../bkupsrv/testBackup/menuStore"
	bs.StateStorePath = "../bkupsrv/testBackup/stateStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l

	// var bds ds.DataStore
	// bds.Path = "./bkupsrv/testBackup/contentStore"
	// bs.TemplateStore = bds.GetNew()
	// //ch.Service = ci.GetNew()

	var cds ds.DataStore
	cds.Path = "../bkupsrv/testBackup/contentStore"
	bs.Store = cds.GetNew()

	var tds ds.DataStore
	tds.Path = "../bkupsrv/testBackup/templateStore"
	bs.TemplateStore = tds.GetNew()

	var crds ds.DataStore
	crds.Path = "../bkupsrv/testBackup/carouselStore"
	bs.CarouselStore = crds.GetNew()

	var cyds ds.DataStore
	cyds.Path = "../bkupsrv/testBackup/countryStore"
	bs.CountryStore = cyds.GetNew()

	var cssds ds.DataStore
	cssds.Path = "../bkupsrv/testBackup/cssStore"
	bs.CSSStore = cssds.GetNew()

	var mds ds.DataStore
	mds.Path = "../bkupsrv/testBackup/menuStore"
	bs.MenuStore = mds.GetNew()

	var sds ds.DataStore
	sds.Path = "../bkupsrv/testBackup/stateStore"
	bs.StateStore = sds.GetNew()

	sh.BackupService = bs.GetNew()

	// sh.ActiveTemplateLocation = "../templatesrv/testDownloads"
	sh.ActiveTemplateLocation = "../bkupsrv/testBackup/templates"

	var ts tmpsrv.Six910TemplateService
	// ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	// ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	ts.TemplateFilePath = "../bkupsrv/testBackup/templates"
	ts.Log = &l
	//var ttds ds.DataStore
	//ttds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.BackupFileName = "backup.dat"

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
	h.AdminDownloadBackups(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_AdminDownloadBackupsAuth(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var bs bkupsrv.Six910BackupService
	bs.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	bs.ContentStorePath = "../bkupsrv/testBackup/contentStore"
	bs.CarouselStorePath = "../bkupsrv/testBackup/carouselStore"
	bs.CountryStorePath = "../bkupsrv/testBackup/countryStore"
	bs.CSSStorePath = "../bkupsrv/testBackup/cssStore"
	bs.MenuStorePath = "../bkupsrv/testBackup/menuStore"
	bs.StateStorePath = "../bkupsrv/testBackup/stateStore"
	bs.ImagePath = "../bkupsrv/testBackup/images"
	bs.TemplateFilePath = "../bkupsrv/testBackup/templates"
	bs.Log = &l

	// var bds ds.DataStore
	// bds.Path = "./bkupsrv/testBackup/contentStore"
	// bs.TemplateStore = bds.GetNew()
	// //ch.Service = ci.GetNew()

	var cds ds.DataStore
	cds.Path = "../bkupsrv/testBackup/contentStore"
	bs.Store = cds.GetNew()

	var tds ds.DataStore
	tds.Path = "../bkupsrv/testBackup/templateStore"
	bs.TemplateStore = tds.GetNew()

	var crds ds.DataStore
	crds.Path = "../bkupsrv/testBackup/carouselStore"
	bs.CarouselStore = crds.GetNew()

	var cyds ds.DataStore
	cyds.Path = "../bkupsrv/testBackup/countryStore"
	bs.CountryStore = cyds.GetNew()

	var cssds ds.DataStore
	cssds.Path = "../bkupsrv/testBackup/cssStore"
	bs.CSSStore = cssds.GetNew()

	var mds ds.DataStore
	mds.Path = "../bkupsrv/testBackup/menuStore"
	bs.MenuStore = mds.GetNew()

	var sds ds.DataStore
	sds.Path = "../bkupsrv/testBackup/stateStore"
	bs.StateStore = sds.GetNew()

	sh.BackupService = bs.GetNew()

	// sh.ActiveTemplateLocation = "../templatesrv/testDownloads"
	sh.ActiveTemplateLocation = "../bkupsrv/testBackup/templates"

	var ts tmpsrv.Six910TemplateService
	// ts.TemplateStorePath = "../templatesrv/testTemplateFiles"
	// ts.TemplateFilePath = "../templatesrv/testDownloads"
	ts.TemplateStorePath = "../bkupsrv/testBackup/templateStore"
	ts.TemplateFilePath = "../bkupsrv/testBackup/templates"
	ts.Log = &l
	//var ttds ds.DataStore
	//ttds.Path = "../templatesrv/testTemplateFiles"
	ts.TemplateStore = tds.GetNew()
	//ch.Service = ci.GetNew()

	sh.TemplateService = ts.GetNew()

	sh.BackupFileName = "backup.dat"

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
	h.AdminDownloadBackups(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
