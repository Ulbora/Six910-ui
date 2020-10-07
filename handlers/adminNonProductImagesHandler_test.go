package handlers

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	isrv "github.com/Ulbora/Six910-ui/imgsrv"
)

func TestSix910Handler_StoreAdminAddImagePage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

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
	h.StoreAdminAddImagePage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddImagePageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

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
	h.StoreAdminAddImagePage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminAddImage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testImages"
	iss.ImageFullPath = "../imgsrv/testImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../imgsrv/testUploadImages/test22.jpg")
	if err != nil {
		fmt.Println("file test image open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("image fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", fi.Name())
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
	h.StoreAdminUploadImage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminAddImageFail(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv2/testUploadImages"
	iss.ImageFullPath = "../imgsrv2/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../imgsrv/testUploadImages/test22.jpg")
	if err != nil {
		fmt.Println("file test image open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("image fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", fi.Name())
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
	h.StoreAdminUploadImage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminAddImageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testUploadImages"
	iss.ImageFullPath = "../imgsrv/testUploadImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	file, err := os.Open("../imgsrv/testUploadImages/test22.jpg")
	if err != nil {
		fmt.Println("file test image open err: ", err)
	}
	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		fmt.Println("file.Stat err: ", err)
	}
	fmt.Println("image fi name : ", fi.Name())

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("image", fi.Name())
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
	h.StoreAdminUploadImage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}

}

func TestSix910Handler_StoreAdminImageList(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

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
	h.StoreAdminImageList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 200 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminImageListLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

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
	h.StoreAdminImageList(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteImage(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testImages"
	iss.ImageFullPath = "../imgsrv/testImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"name": "test22.jpg",
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
	h.StoreAdminDeleteImage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteImage2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testImages"
	iss.ImageFullPath = "../imgsrv/testImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"name": "test22.jpg",
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
	h.StoreAdminDeleteImage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}

func TestSix910Handler_StoreAdminDeleteImageLogin(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	//var ci sr.CmsService
	//ci.ImagePath = "../services/testImages"
	//ci.ImageFullPath = "../services/testImages"

	//ci.Log = &l

	var iss isrv.Six910ImageService
	iss.ImagePath = "../imgsrv/testImages"
	iss.ImageFullPath = "../imgsrv/testImages"
	iss.Log = &l

	sh.ImageService = iss.GetNew()
	sh.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	r, _ := http.NewRequest("POST", "https://test.com", nil)
	vars := map[string]string{
		"name": "test22.jpg",
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
	h.StoreAdminDeleteImage(w, r)
	fmt.Println("code: ", w.Code)

	if w.Code != 302 {
		t.Fail()
	}
}
