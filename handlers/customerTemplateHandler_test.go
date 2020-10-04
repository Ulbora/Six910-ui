package handlers

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
	"html/template"
	"testing"
	//"github.com/gorilla/sessions"
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
)

func TestSix910Handler_LoadTemplate(t *testing.T) {
	var ch Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	ch.Log = &l
	ch.AdminTemplates = template.Must(template.ParseFiles("testHtmls/test.html"))

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "../data/templateStore"

	ts.Log = &l
	var ds ds.DataStore
	ds.Path = "../data/templateStore"
	ts.TemplateStore = ds.GetNew()
	ch.TemplateService = ts.GetNew()

	ch.ActiveTemplateLocation = "../static/templates"

	h := ch.GetNew()

	h.LoadTemplate()

	fmt.Println("template in use: ", ch.ActiveTemplateName)
	if ch.ActiveTemplateName == "" {
		t.Fail()
	}
}
