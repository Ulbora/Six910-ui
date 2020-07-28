package managers

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
)

func TestSix910Manager_CustomerLogin(t *testing.T) {
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
	mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "tester123", "enabled": true, "role": "customer" }`))
	gp.MockResp = &mres
	gp.MockDoSuccess1 = true
	gp.MockRespCode = 200
	sapi.OverrideProxy(&gp)
	//---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var u api.User
	u.CustomerID = 18
	//u.Enabled = true
	u.Password = "tester"
	//u.Role = "customer"
	u.Username = "tester123"

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc, us := m.CustomerLogin(&u, &head)
	fmt.Println("suc customer: ", suc)
	fmt.Println("us customer: ", us)
	if !suc || us.Username != "tester123" {
		t.Fail()
	}
}

func TestSix910Manager_CustomerChangePassword(t *testing.T) {
	var sm Six910Manager
	var sapi mapi.MockAPI

	//-----------start mocking------------------
	var user api.UserResponse
	user.Username = "tester123"
	user.Role = customerRole
	user.Enabled = true

	sapi.MockUser = &user

	var ur api.Response
	ur.Success = true
	sapi.MockUpdateUserResp = &ur

	//-----------end mocking --------

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

	// //---mock out the call
	// var gp px.MockGoProxy
	// var mres http.Response
	// mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "admin", "enabled": true, "role": "StoreAdmin" }`))
	// gp.MockResp = &mres
	// gp.MockDoSuccess1 = true
	// gp.MockRespCode = 200
	// sapi.OverrideProxy(&gp)
	// //---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var u api.User
	u.CustomerID = 18
	//u.Enabled = true
	u.Password = "tester"
	//u.Role = "customer"
	u.Username = "tester123"

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc, us := m.CustomerChangePassword(&u, &head)
	fmt.Println("suc: ", suc)
	fmt.Println("us: ", us)
	if !suc || us.Username != "tester123" {
		t.Fail()
	}
}
