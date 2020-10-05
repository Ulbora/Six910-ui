package menusrv

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	ds "github.com/Ulbora/json-datastore"
)

func TestSix910MenuService_AddMenu(t *testing.T) {

	var sms Six910MenuService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	ds.Delete("menu1")
	sms.MenuStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sms.Log = &l

	ms := sms.GetNew()

	var m Menu
	m.Name = "menu1"
	m.Active = true
	m.Location = "top"
	m.Shade = "light"
	m.Background = "light"
	m.Style = ""
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	suc := ms.AddMenu(&m)
	if !suc {
		t.Fail()
	}
}

func TestSix910MenuService_UpdateMenu(t *testing.T) {

	var sms Six910MenuService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	sms.MenuStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sms.Log = &l

	ms := sms.GetNew()

	var m Menu
	m.Name = "menu1"
	m.Active = true
	m.Shade = "dark"
	m.Background = "light"
	m.Style = "background: blue;"
	m.ShadeList = &[]string{"light", "dark"}
	m.BackgroundList = &[]string{"light", "dark"}

	suc := ms.UpdateMenu(&m)
	if !suc {
		t.Fail()
	}

}

func TestSix910MenuService_GetMenu(t *testing.T) {

	var sms Six910MenuService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	sms.MenuStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sms.Log = &l

	ms := sms.GetNew()

	suc, m := ms.GetMenu("menu1")
	fmt.Println("found menu", *m)
	if !suc {
		t.Fail()
	}
}

func TestSix910MenuService_GetMenuList(t *testing.T) {

	var sms Six910MenuService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	sms.MenuStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sms.Log = &l

	ms := sms.GetNew()

	mlst := ms.GetMenuList()
	fmt.Println("found menu list", *mlst)
	if len(*mlst) == 0 {
		t.Fail()
	}
}

func TestSix910MenuService_DeleteMenu(t *testing.T) {

	var sms Six910MenuService
	var ds ds.DataStore
	ds.Path = "./testFiles"
	sms.MenuStore = ds.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sms.Log = &l

	ms := sms.GetNew()

	suc := ms.DeleteMenu("menu1")
	if !suc {
		t.Fail()
	}
}
