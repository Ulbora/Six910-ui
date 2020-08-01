package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Manager_CreateCustomerAccountExistingCustomer(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cusm sdbi.Customer
	cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "CR"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "CR"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

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
	// mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "tester123", "enabled": true, "role": "customer" }`))
	// gp.MockResp = &mres
	// gp.MockDoSuccess1 = true
	// gp.MockRespCode = 200
	// sapi.OverrideProxy(&gp)
	// //---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var cus sdbi.Customer
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

	var usr api.User
	usr.Password = "tester"
	usr.Username = "tester"

	var alst []sdbi.Address
	alst = append(alst, a1)
	alst = append(alst, a2)

	var cp CustomerAccount
	cp.Customer = &cus
	cp.Addresses = &alst
	cp.User = &usr

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc, ca := m.CreateCustomerAccount(&cp, &head)

	fmt.Println("customer ca: ", *ca.Customer)
	fmt.Println("customer address: ", *ca.Addresses)
	fmt.Println("customer user: ", *ca.User)
	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_CreateCustomerAccountExistingCustomer2NewUser(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cusm sdbi.Customer
	cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var mu api.UserResponse
	//mu.Enabled = true
	//mu.Username = "tester"

	sapi.MockUser = &mu

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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
	// mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "tester123", "enabled": true, "role": "customer" }`))
	// gp.MockResp = &mres
	// gp.MockDoSuccess1 = true
	// gp.MockRespCode = 200
	// sapi.OverrideProxy(&gp)
	// //---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var cus sdbi.Customer
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

	var usr api.User
	usr.Password = "tester"
	usr.Username = "tester"

	var alst []sdbi.Address
	alst = append(alst, a1)
	alst = append(alst, a2)

	var cp CustomerAccount
	cp.Customer = &cus
	cp.Addresses = &alst
	cp.User = &usr

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc, ca := m.CreateCustomerAccount(&cp, &head)

	fmt.Println("customer ca: ", *ca.Customer)
	fmt.Println("customer address: ", *ca.Addresses)
	fmt.Println("customer user: ", *ca.User)
	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_CreateCustomerAccountNewCustomer(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cusm sdbi.Customer
	//cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var macres api.ResponseID
	macres.Success = true
	macres.ID = 3

	sapi.MockAddCustomerResp = &macres

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "Conch Republic"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, a1)

	sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "Conch Republic"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var aures api.Response
	aures.Success = true

	sapi.MockAddCustomerUserRes = &aures

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
	// mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "tester123", "enabled": true, "role": "customer" }`))
	// gp.MockResp = &mres
	// gp.MockDoSuccess1 = true
	// gp.MockRespCode = 200
	// sapi.OverrideProxy(&gp)
	// //---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var cus sdbi.Customer
	cus.Email = "test2@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

	var usr api.User
	usr.Password = "tester"
	usr.Username = "test2@tester.com"

	var alst []sdbi.Address
	alst = append(alst, a1)
	alst = append(alst, a2)

	var cp CustomerAccount
	cp.Customer = &cus
	cp.Addresses = &alst
	cp.User = &usr

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc, ca := m.CreateCustomerAccount(&cp, &head)

	fmt.Println("customer ca: ", *ca.Customer)
	fmt.Println("customer address: ", *ca.Addresses)
	fmt.Println("customer user: ", *ca.User)
	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_UpdateCustomerAccount(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cusm sdbi.Customer
	cusm.ID = 18
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var cures api.Response
	cures.Success = true
	sapi.MockUpdateCustomerResp = &cures

	var a1 sdbi.Address
	a1.ID = 18
	a1.Address = "123 Whitehead Street"
	a1.City = "Key West"
	a1.Country = "CR"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"
	a1.CustomerID = 18

	// var malst []sdbi.Address
	// malst = append(malst, a1)

	// sapi.MockAddressList1 = &malst

	var a2 sdbi.Address
	//a2.ID = 19
	a2.Address = "907 Whitehead Street"
	a2.City = "Key West"
	a2.Country = "CR"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"
	a2.CustomerID = 18

	var malst2 []sdbi.Address
	malst2 = append(malst2, a1)
	malst2 = append(malst2, a2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var uares api.Response
	uares.Success = true
	sapi.MockUpdateAddressRes = &uares

	var uures api.Response
	uures.Success = true

	sapi.MockUpdateUserResp = &uures

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
	// mres.Body = ioutil.NopCloser(bytes.NewBufferString(`{"username": "tester123", "enabled": true, "role": "customer" }`))
	// gp.MockResp = &mres
	// gp.MockDoSuccess1 = true
	// gp.MockRespCode = 200
	// sapi.OverrideProxy(&gp)
	// //---end mock out the call

	sapi.SetLogLever(lg.AllLevel)
	sm.Log = &l

	var cus sdbi.Customer
	cus.ID = 18
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

	var usr api.User
	usr.Password = "tester"
	usr.OldPassword = "tester2"
	usr.Username = "tester"
	usr.Enabled = true
	usr.CustomerID = 18

	var alst []sdbi.Address
	alst = append(alst, a1)
	alst = append(alst, a2)

	var cp CustomerAccount
	cp.Customer = &cus
	cp.Addresses = &alst
	cp.User = &usr

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	suc := m.UpdateCustomerAccount(&cp, &head)

	if !suc {
		t.Fail()
	}
}
