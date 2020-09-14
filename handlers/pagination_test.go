package handlers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
)

func TestSix910Handler_doPagination1(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	res := sh.doPagination(100, 100, 100, "http://test")
	fmt.Println("res: ", *res)
	fmt.Println("res.pages: ", *res.Pages)
	fmt.Println("len(*res.Pages): ", len(*res.Pages))
	if len(*res.Pages) != 3 {
		t.Fail()
	}
}

func TestSix910Handler_doPagination2(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	res := sh.doPagination(100, 5, 100, "http://test")
	fmt.Println("res: ", *res)
	fmt.Println("res.pages: ", *res.Pages)
	fmt.Println("len(*res.Pages): ", len(*res.Pages))
	if len(*res.Pages) != 2 {
		t.Fail()
	}
}

func TestSix910Handler_doPagination3(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	res := sh.doPagination(300, 100, 100, "http://test")
	fmt.Println("res: ", *res)
	fmt.Println("res.pages: ", *res.Pages)
	fmt.Println("len(*res.Pages): ", len(*res.Pages))
	if len(*res.Pages) != 5 {
		t.Fail()
	}
}

func TestSix910Handler_doPagination4(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	res := sh.doPagination(400, 5, 100, "http://test")
	fmt.Println("res: ", *res)
	fmt.Println("res.pages: ", *res.Pages)
	fmt.Println("len(*res.Pages): ", len(*res.Pages))
	if len(*res.Pages) != 5 {
		t.Fail()
	}
}

func TestSix910Handler_doPagination5(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	res := sh.doPagination(0, 5, 100, "http://test")
	fmt.Println("res: ", *res)
	fmt.Println("res.pages: ", *res.Pages)
	fmt.Println("len(*res.Pages): ", len(*res.Pages))
	if len(*res.Pages) != 1 {
		t.Fail()
	}
}
