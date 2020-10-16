package csssrv

import (
	"fmt"

	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

func TestSix910CSSService_GetPageCSS(t *testing.T) {

	var cs Six910CSSService
	var cds ds.DataStore
	cds.Path = "./testFiles"
	cs.CSSStore = cds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	suc, pg := s.GetPageCSS("testPage")
	fmt.Println("found page css", *pg)
	if !suc {
		t.Fail()
	}
}

func TestSix910CSSService_UpdatePageCSS(t *testing.T) {

	var cs Six910CSSService
	var cds ds.DataStore
	cds.Path = "./testFiles"
	cs.CSSStore = cds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	var p Page
	p.Name = "testPage"
	p.Background = "red"
	p.Color = "green"
	p.PageTitle = "black"

	suc := s.UpdatePage(&p)
	if !suc {
		t.Fail()
	}

}

func TestSix910CSSService_UpdatePageCSS2(t *testing.T) {

	var cs Six910CSSService
	var cds ds.DataStore
	cds.Path = "./testFiles"
	cs.CSSStore = cds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	var p Page
	p.Name = "testPage"
	p.Background = "grey"
	p.Color = "white"
	p.PageTitle = "blue"

	suc := s.UpdatePage(&p)
	if !suc {
		t.Fail()
	}

}

func TestSix910CSSService_GetPage(t *testing.T) {

	var cs Six910CSSService
	var cds ds.DataStore
	cds.Path = "./testFiles"
	cs.CSSStore = cds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	cs.Log = &l

	s := cs.GetNew()

	suc, pg := s.GetPage("testPage")
	fmt.Println("found page css", *pg)
	if !suc {
		t.Fail()
	}
}
