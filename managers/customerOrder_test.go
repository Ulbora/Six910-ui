package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Manager_ViewCustomerOrder(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var odr sdbi.Order
	odr.ID = 4
	odr.CustomerID = 18

	sapi.MockOrder = &odr

	var oi sdbi.OrderItem
	oi.ID = 2
	oi.OrderID = 4
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 2
	oc.OrderID = 4
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)

	sapi.MockCommentList = &oclst

	var cus sdbi.Customer
	cus.ID = 18
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

	sapi.MockCustomer = &cus

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "CR"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "CR"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var alst []sdbi.Address
	alst = append(alst, a1)
	alst = append(alst, a2)

	sapi.MockAddressList1 = &alst

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

	var cp CustomerProduct
	cp.CustomerID = 18
	cp.ProductID = 7
	cp.Quantity = 3
	cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cc := m.ViewCustomerOrder(58, 18, &head)

	fmt.Println("customer cc: ", cc)
	if cc.CustomerAccount.Customer.ID != 18 || cc.Items == nil {
		t.Fail()
	}
}

func TestSix910Manager_ViewCustomerOrderList(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var odr sdbi.Order
	odr.ID = 4
	odr.CustomerID = 18

	var odr2 sdbi.Order
	odr2.ID = 44
	odr2.CustomerID = 18

	var olst []sdbi.Order
	olst = append(olst, odr)
	olst = append(olst, odr2)

	sapi.MockOrderList = &olst

	var oi sdbi.OrderItem
	oi.ID = 2
	oi.OrderID = 4
	var oilst []sdbi.OrderItem
	oilst = append(oilst, oi)
	sapi.MockOrderItemList = &oilst

	var oc sdbi.OrderComment
	oc.Comment = "test"
	oc.ID = 2
	oc.OrderID = 4
	var oclst []sdbi.OrderComment
	oclst = append(oclst, oc)

	sapi.MockCommentList = &oclst

	var cus sdbi.Customer
	cus.ID = 18
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

	sapi.MockCustomer = &cus

	var a1 sdbi.Address
	a1.Address = "123 Whitehead St"
	a1.City = "Key West"
	a1.Country = "CR"
	a1.County = "Monroe"
	a1.State = "FL"
	a1.Type = "Shipping"

	var a2 sdbi.Address
	a2.Address = "907 Whitehead St"
	a2.City = "Key West"
	a2.Country = "CR"
	a2.County = "Monroe"
	a2.State = "FL"
	a2.Type = "Billing"

	var alst []sdbi.Address
	alst = append(alst, a1)
	alst = append(alst, a2)

	sapi.MockAddressList1 = &alst

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

	var cp CustomerProduct
	cp.CustomerID = 18
	cp.ProductID = 7
	cp.Quantity = 3
	cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cc := m.ViewCustomerOrderList(18, &head)

	fmt.Println("customer cc in list: ", cc)
	if (*cc)[0].CustomerAccount.Customer.ID != 18 || (*cc)[0].Items == nil {
		t.Fail()
	}
}
