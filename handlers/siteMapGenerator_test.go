package handlers

import (
	"fmt"
	"testing"
	"time"

	lg "github.com/Ulbora/Level_Logger"
)

func TestSix910Handler_generateSiteMap(t *testing.T) {
	var sh Six910Handler
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l

	var vp SiteMapValues
	vp.Domain = "https://test.com"
	var idlst []int64
	idlst = append(idlst, 2)
	idlst = append(idlst, 3)
	idlst = append(idlst, 5)
	idlst = append(idlst, 77777)
	vp.ProductIDList = &idlst
	today := time.Date(2021, 1, 1, 20, 34, 58, 651387237, time.UTC)

	sm := sh.generateSiteMap(today, &vp)
	fmt.Println("site map: \n", string(sm))
	if sm == nil {
		t.Fail()
	}
	// t.Fail()
}
