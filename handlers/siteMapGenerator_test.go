package handlers

import (
	"fmt"
	lg "github.com/Ulbora/Level_Logger"
	"testing"
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
	sm := sh.generateSiteMap(&vp)
	fmt.Println("site map: \n", string(sm))
	if sm == nil {
		t.Fail()
	}
}
