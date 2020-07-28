package managers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	api "github.com/Ulbora/Six910API-Go"
)

func TestSix910Manager_StoreAdminLogin(t *testing.T) {
	var sm Six910Manager
	var sapi api.Six910API
	//sapi.SetAPIKey("123")
	//sapi.storeID = 59
	sapi.SetStoreID(59)

	sapi.SetRestURL("http://localhost:3002")
	sapi.SetStore("defaultLocalStore", "defaultLocalStore.mydomain.com")
	sapi.SetAPIKey("GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5")

	//api := sapi.GetNew()

	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sm.API = sapi.GetNew()

	//---mock out the call
	var gp px.MockGoProxy
	var mres http.Response
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "admin", "enabled": true, "role": "StoreAdmin" }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var u api.User
	//u.CustomerID = 18
	//u.Enabled = true
	u.Password = "admin"
	//u.Role = "customer"
	u.Username = "admin"

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc, us := m.StoreAdminLogin(&u, &head)
	fmt.Println("suc: ", suc)
	fmt.Println("us: ", us)
	if !suc || us.Username != "admin" {
		t.Fail()
	}

}
