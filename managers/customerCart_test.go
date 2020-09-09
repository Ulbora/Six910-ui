package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Manager_AddProductToCart(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var ciares api.ResponseID
	ciares.Success = true
	ciares.ID = 5
	sapi.MockCartItemAddResp = &ciares

	var ci sdbi.CartItem
	ci.CartID = 3
	ci.ID = 7
	ci.ProductID = 7
	ci.Quantity = 3

	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sapi.MockCartItemList = &cilst

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
	cc := m.AddProductToCart(&cp, &head)

	fmt.Println("customer cc: ", cc)
	if cc.Cart == nil || cc.Cart.CustomerID != 18 || cc.Items == nil {
		t.Fail()
	}
}

func TestSix910Manager_AddProductToCartNoCID(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------

	var sapi mapi.MockAPI

	var acresp api.ResponseID
	acresp.Success = true
	acresp.ID = 3

	sapi.MockAddCartResp = &acresp

	var ciares api.ResponseID
	ciares.Success = true
	ciares.ID = 5
	sapi.MockCartItemAddResp = &ciares

	var ci sdbi.CartItem
	ci.CartID = 3
	ci.ID = 7
	ci.ProductID = 7
	ci.Quantity = 3

	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sapi.MockCartItemList = &cilst

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
	//cp.CustomerID = 18
	cp.ProductID = 7
	cp.Quantity = 3
	cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cc := m.AddProductToCart(&cp, &head)

	fmt.Println("customer cc: ", cc)
	if cc.Cart == nil || cc.Items == nil {
		t.Fail()
	}
}

func TestSix910Manager_AddProductToCartNoCIDExistingCartID(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------

	var sapi mapi.MockAPI

	var acresp api.ResponseID
	acresp.Success = true
	acresp.ID = 3

	sapi.MockAddCartResp = &acresp

	var ciares api.ResponseID
	ciares.Success = true
	ciares.ID = 5
	sapi.MockCartItemAddResp = &ciares

	var ci sdbi.CartItem
	ci.CartID = 3
	ci.ID = 7
	ci.ProductID = 7
	ci.Quantity = 3

	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sapi.MockCartItemList = &cilst

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
	//cp.CustomerID = 18
	cp.ProductID = 7
	var cart sdbi.Cart
	cart.ID = 21
	cp.Cart = &cart
	cp.Quantity = 3
	cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cc := m.AddProductToCart(&cp, &head)

	fmt.Println("customer cc: ", cc)
	if cc.Cart == nil || cc.Items == nil {
		t.Fail()
	}
}

func TestSix910Manager_UpdateProductToCart(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var ciures api.Response
	ciures.Success = true

	var ci sdbi.CartItem
	ci.CartID = 3
	ci.ID = 7
	ci.ProductID = 7
	ci.Quantity = 3

	var cilst []sdbi.CartItem
	cilst = append(cilst, ci)

	sapi.MockCartItemList = &cilst
	sapi.MockCartItemUpdateResp = &ciures

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

	var cart2 sdbi.Cart
	cart2.CustomerID = 18
	cart2.ID = 21
	cart2.StoreID = 59

	var ci2 sdbi.CartItem
	ci2.CartID = 21
	ci2.ID = 1
	ci2.ProductID = 7
	ci2.Quantity = 3

	var cp CustomerProductUpdate
	cp.CustomerID = 18
	cp.Cart = &cart2
	cp.CartItem = &ci2

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cc := m.UpdateProductToCart(&cp, &head)

	fmt.Println("customer cc: ", cc)
	if cc.Cart == nil || cc.Cart.CustomerID != 18 || cc.Items == nil {
		t.Fail()
	}
}

func TestSix910Manager_generateOrderNumber(t *testing.T) {
	var sm Six910Manager

	res := sm.generateOrderNumber()
	fmt.Println("Order #: ", res)
}

func TestSix910Manager_processOrderItems(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	sapi.MockProduct = &prod

	var oires api.ResponseID
	oires.Success = true
	oires.ID = 3

	sapi.MockAddOrderItemResp = &oires

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

	var cart2 sdbi.Cart
	cart2.CustomerID = 18
	cart2.ID = 21
	cart2.StoreID = 59

	var ci2 sdbi.CartItem
	ci2.CartID = 21
	ci2.ID = 1
	ci2.ProductID = 7
	ci2.Quantity = 3

	var cp CustomerProductUpdate
	cp.CustomerID = 18
	cp.Cart = &cart2
	cp.CartItem = &ci2

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var cilst []sdbi.CartItem

	var ci1 sdbi.CartItem
	ci1.CartID = 1
	ci1.ProductID = 1
	ci1.Quantity = 1
	cilst = append(cilst, ci1)

	var ci22 sdbi.CartItem
	ci22.CartID = 2
	ci22.ProductID = 2
	ci22.Quantity = 2
	cilst = append(cilst, ci22)

	var ci3 sdbi.CartItem
	ci3.CartID = 3
	ci3.ProductID = 3
	ci3.Quantity = 3
	cilst = append(cilst, ci3)

	var ci4 sdbi.CartItem
	ci4.CartID = 4
	ci4.ProductID = 4
	ci4.Quantity = 4
	cilst = append(cilst, ci4)

	suc, res := sm.processOrderItems(&cilst, 4, &head)
	fmt.Println("suc from process cart items: ", suc)
	fmt.Println("res from process cart items: ", *res)
	if !suc || res == nil || len(*res) != 4 {
		t.Fail()
	}
}

func TestSix910Manager_processOrderItemsFail(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	sapi.MockProduct = &prod

	var oires api.ResponseID
	oires.Success = false
	oires.ID = 3

	sapi.MockAddOrderItemResp = &oires

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

	var cart2 sdbi.Cart
	cart2.CustomerID = 18
	cart2.ID = 21
	cart2.StoreID = 59

	var ci2 sdbi.CartItem
	ci2.CartID = 21
	ci2.ID = 1
	ci2.ProductID = 7
	ci2.Quantity = 3

	var cp CustomerProductUpdate
	cp.CustomerID = 18
	cp.Cart = &cart2
	cp.CartItem = &ci2

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var cilst []sdbi.CartItem

	var ci1 sdbi.CartItem
	ci1.CartID = 1
	ci1.ProductID = 1
	ci1.Quantity = 1
	cilst = append(cilst, ci1)

	var ci22 sdbi.CartItem
	ci22.CartID = 2
	ci22.ProductID = 2
	ci22.Quantity = 2
	cilst = append(cilst, ci22)

	var ci3 sdbi.CartItem
	ci3.CartID = 3
	ci3.ProductID = 3
	ci3.Quantity = 3
	cilst = append(cilst, ci3)

	var ci4 sdbi.CartItem
	ci4.CartID = 4
	ci4.ProductID = 4
	ci4.Quantity = 4
	cilst = append(cilst, ci4)

	suc, res := sm.processOrderItems(&cilst, 4, &head)
	fmt.Println("res from process cart items: ", *res)
	if suc {
		t.Fail()
	}
}

func TestSix910Manager_completeOrder(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	// var ciures api.Response
	// ciures.Success = true

	// var ci sdbi.CartItem
	// ci.CartID = 3
	// ci.ID = 7
	// ci.ProductID = 7
	// ci.Quantity = 3

	// var mcilst []sdbi.CartItem
	// mcilst = append(mcilst, ci)

	//sapi.MockCartItemList = &mcilst
	//sapi.MockCartItemUpdateResp = &ciures

	var prod sdbi.Product
	sapi.MockProduct = &prod

	var ores api.ResponseID
	ores.Success = true
	ores.ID = 5
	sapi.MockAddOrderResp = &ores

	var oires api.ResponseID
	oires.Success = true
	oires.ID = 3

	sapi.MockAddOrderItemResp = &oires

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

	var cart2 sdbi.Cart
	cart2.CustomerID = 18
	cart2.ID = 21
	cart2.StoreID = 59

	var ci2 sdbi.CartItem
	ci2.CartID = 21
	ci2.ID = 1
	ci2.ProductID = 7
	ci2.Quantity = 3

	// var cp CustomerProductUpdate
	// cp.CustomerID = 18
	// cp.Cart = &cart2
	// cp.CartItem = &ci2

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var cilst []sdbi.CartItem

	var ci1 sdbi.CartItem
	ci1.CartID = 21
	ci1.ProductID = 7
	ci1.Quantity = 1
	cilst = append(cilst, ci1)

	var ci22 sdbi.CartItem
	ci22.CartID = 21
	ci22.ProductID = 8
	ci22.Quantity = 2
	cilst = append(cilst, ci22)

	var ci3 sdbi.CartItem
	ci3.CartID = 21
	ci3.ProductID = 9
	ci3.Quantity = 3
	cilst = append(cilst, ci3)

	// var ci4 sdbi.CartItem
	// ci4.CartID = 21
	// ci4.ProductID = 4
	// ci4.Quantity = 4
	// cilst = append(cilst, ci4)

	var usr api.User
	usr.Password = "tester"
	usr.Username = "tester"
	usr.Enabled = true

	var cus sdbi.Customer
	cus.ID = 18
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

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

	var ca CustomerAccount
	ca.Customer = &cus
	ca.Addresses = &alst
	ca.User = &usr

	var ccart CustomerCart
	ccart.Cart = &cart2
	ccart.CustomerAccount = &ca
	ccart.Items = &cilst
	ccart.InsuranceCost = 4.12
	ccart.OrderType = "Delivery"
	ccart.Pickup = false
	ccart.ShippingHandling = 12.52
	ccart.Subtotal = 52.20
	ccart.Taxes = 2.00
	ccart.Total = 54.20

	m := sm.GetNew()
	res := m.CheckOut(&ccart, &head)
	if !res.Success {
		t.Fail()
	}

}

func TestSix910Manager_completeOrderWithComment(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	// var ciures api.Response
	// ciures.Success = true

	// var ci sdbi.CartItem
	// ci.CartID = 3
	// ci.ID = 7
	// ci.ProductID = 7
	// ci.Quantity = 3

	// var mcilst []sdbi.CartItem
	// mcilst = append(mcilst, ci)

	//sapi.MockCartItemList = &mcilst
	//sapi.MockCartItemUpdateResp = &ciures

	var prod sdbi.Product
	sapi.MockProduct = &prod

	var ores api.ResponseID
	ores.Success = true
	ores.ID = 5
	sapi.MockAddOrderResp = &ores

	var oires api.ResponseID
	oires.Success = true
	oires.ID = 3

	sapi.MockAddOrderItemResp = &oires

	var cmtres api.ResponseID
	cmtres.Success = true
	cmtres.ID = 4

	sapi.MockAddCommentResp = &cmtres

	var cmt sdbi.OrderComment
	cmt.Comment = "Leave at door"
	cmt.OrderID = 3
	cmt.Username = "tester"
	cmt.ID = 2

	var cmtlst []sdbi.OrderComment
	cmtlst = append(cmtlst, cmt)
	sapi.MockCommentList = &cmtlst

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

	var cart2 sdbi.Cart
	cart2.CustomerID = 18
	cart2.ID = 21
	cart2.StoreID = 59

	var ci2 sdbi.CartItem
	ci2.CartID = 21
	ci2.ID = 1
	ci2.ProductID = 7
	ci2.Quantity = 3

	var cp CustomerProductUpdate
	cp.CustomerID = 18
	cp.Cart = &cart2
	cp.CartItem = &ci2

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var cilst []sdbi.CartItem

	var ci1 sdbi.CartItem
	ci1.CartID = 1
	ci1.ProductID = 1
	ci1.Quantity = 1
	cilst = append(cilst, ci1)

	var ci22 sdbi.CartItem
	ci22.CartID = 2
	ci22.ProductID = 2
	ci22.Quantity = 2
	cilst = append(cilst, ci22)

	var ci3 sdbi.CartItem
	ci3.CartID = 3
	ci3.ProductID = 3
	ci3.Quantity = 3
	cilst = append(cilst, ci3)

	var ci4 sdbi.CartItem
	ci4.CartID = 4
	ci4.ProductID = 4
	ci4.Quantity = 4
	cilst = append(cilst, ci4)

	var usr api.User
	usr.Password = "tester"
	usr.Username = "tester"
	usr.Enabled = true

	var cus sdbi.Customer
	cus.ID = 4
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

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

	var ca CustomerAccount
	ca.Customer = &cus
	ca.Addresses = &alst
	ca.User = &usr

	var ccart CustomerCart
	ccart.Cart = &cart
	ccart.CustomerAccount = &ca
	ccart.Items = &cilst
	ccart.InsuranceCost = 4.12
	ccart.OrderType = "Delivery"
	ccart.Pickup = false
	ccart.ShippingHandling = 12.52
	ccart.Subtotal = 52.20
	ccart.Taxes = 2.00
	ccart.Total = 54.20
	ccart.Comment = "Leave at door"

	m := sm.GetNew()
	res := m.CheckOut(&ccart, &head)
	if !res.Success {
		t.Fail()
	}

}

func TestSix910Manager_completeOrderCreateCustomer(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var cusm sdbi.Customer
	cusm.ID = 5
	cusm.Email = "test@tester.com"
	sapi.MockCustomer = &cusm

	var ma1 sdbi.Address
	ma1.Address = "123 Whitehead St"
	ma1.City = "Key West"
	ma1.Country = "CR"
	ma1.County = "Monroe"
	ma1.State = "FL"
	ma1.Type = "Shipping"

	var malst []sdbi.Address
	malst = append(malst, ma1)

	sapi.MockAddressList1 = &malst

	var ma2 sdbi.Address
	ma2.Address = "907 Whitehead St"
	ma2.City = "Key West"
	ma2.Country = "CR"
	ma2.County = "Monroe"
	ma2.State = "FL"
	ma2.Type = "Billing"

	var malst2 []sdbi.Address
	malst2 = append(malst2, ma1)
	malst2 = append(malst2, ma2)

	sapi.MockAddressList2 = &malst2

	var mu api.UserResponse
	mu.Enabled = true
	mu.Username = "tester"

	sapi.MockUser = &mu

	var aares api.ResponseID
	aares.Success = true
	aares.ID = 8
	sapi.MockAddAddressRes = &aares

	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	// var ciures api.Response
	// ciures.Success = true

	// var ci sdbi.CartItem
	// ci.CartID = 3
	// ci.ID = 7
	// ci.ProductID = 7
	// ci.Quantity = 3

	// var mcilst []sdbi.CartItem
	// mcilst = append(mcilst, ci)

	//sapi.MockCartItemList = &mcilst
	//sapi.MockCartItemUpdateResp = &ciures

	var prod sdbi.Product
	sapi.MockProduct = &prod

	var ores api.ResponseID
	ores.Success = true
	ores.ID = 5
	sapi.MockAddOrderResp = &ores

	var oires api.ResponseID
	oires.Success = true
	oires.ID = 3

	sapi.MockAddOrderItemResp = &oires

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

	var cart2 sdbi.Cart
	cart2.CustomerID = 18
	cart2.ID = 21
	cart2.StoreID = 59

	var ci2 sdbi.CartItem
	ci2.CartID = 21
	ci2.ID = 1
	ci2.ProductID = 7
	ci2.Quantity = 3

	var cp CustomerProductUpdate
	cp.CustomerID = 18
	cp.Cart = &cart2
	cp.CartItem = &ci2

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var cilst []sdbi.CartItem

	var ci1 sdbi.CartItem
	ci1.CartID = 1
	ci1.ProductID = 1
	ci1.Quantity = 1
	cilst = append(cilst, ci1)

	var ci22 sdbi.CartItem
	ci22.CartID = 2
	ci22.ProductID = 2
	ci22.Quantity = 2
	cilst = append(cilst, ci22)

	var ci3 sdbi.CartItem
	ci3.CartID = 3
	ci3.ProductID = 3
	ci3.Quantity = 3
	cilst = append(cilst, ci3)

	var ci4 sdbi.CartItem
	ci4.CartID = 4
	ci4.ProductID = 4
	ci4.Quantity = 4
	cilst = append(cilst, ci4)

	var usr api.User
	usr.Password = "tester"
	usr.Username = "tester"
	usr.Enabled = true

	var cus sdbi.Customer
	//cus.ID = 4
	cus.Email = "test@tester.com"
	cus.City = "Atlanta"
	cus.Company = "tester inc"
	cus.FirstName = "Tommy"
	cus.LastName = "Tutone"
	cus.Phone = "867-2309"

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

	var ca CustomerAccount
	ca.Customer = &cus
	ca.Addresses = &alst
	ca.User = &usr

	var ccart CustomerCart
	ccart.Cart = &cart
	ccart.CustomerAccount = &ca
	ccart.Items = &cilst
	ccart.InsuranceCost = 4.12
	ccart.OrderType = "Delivery"
	ccart.Pickup = false
	ccart.ShippingHandling = 12.52
	ccart.Subtotal = 52.20
	ccart.Taxes = 2.00
	ccart.Total = 54.20

	m := sm.GetNew()
	res := m.CheckOut(&ccart, &head)
	if !res.Success {
		t.Fail()
	}

}

func TestSix910Manager_ViewCart(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var pd sdbi.Product
	pd.ID = 4
	pd.Price = 28.95
	pd.ShortDesc = "test one"
	pd.Thumbnail = "/test/"
	sapi.MockProduct = &pd

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

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cc CustomerCart
	cc.Items = &cilstp

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cv := m.ViewCart(&cc, &head)

	fmt.Println("customer cv: ", cv)
	if cv.Total != 202.65 {
		t.Fail()
	}
}

func TestSix910Manager_ViewCartSaleprice(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var pd sdbi.Product
	pd.ID = 4
	pd.SalePrice = 28.95
	pd.ShortDesc = "test one"
	pd.Thumbnail = "/test/"
	sapi.MockProduct = &pd

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

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cc CustomerCart
	cc.Items = &cilstp

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cv := m.ViewCart(&cc, &head)

	fmt.Println("customer cv: ", cv)
	if cv.Total != 202.65 {
		t.Fail()
	}
}

func TestSix910Manager_CalculateCartTotals(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var pd sdbi.Product
	pd.ID = 4
	pd.SalePrice = 28.95
	pd.ShortDesc = "test one"
	pd.Thumbnail = "/test/"
	sapi.MockProduct = &pd

	var smth sdbi.ShippingMethod
	smth.Cost = 4.55
	smth.Handling = 2.22
	sapi.MockShippingMethod = &smth

	var insc sdbi.Insurance
	insc.Cost = 2.34

	sapi.MockInsurance = &insc

	var uadd sdbi.Address
	uadd.Country = "USA"
	uadd.State = "GA"
	uadd.Zip = "12345"

	sapi.MockAddress = &uadd

	var trate1 sdbi.TaxRate
	trate1.Country = "USA"
	trate1.State = "GA"
	trate1.PercentRate = 5
	trate1.ZipStart = "12344"
	trate1.ZipEnd = "12346"

	var trate2 sdbi.TaxRate
	trate2.Country = "USA"
	trate2.State = "GA"
	trate2.PercentRate = 7
	trate2.ZipStart = ""
	trate2.ZipEnd = ""
	trate2.IncludeHandling = true
	trate2.IncludeShipping = true

	var trlst []sdbi.TaxRate
	trlst = append(trlst, trate1)
	trlst = append(trlst, trate2)

	sapi.MockTaxRateList = &trlst

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

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cvv CartView
	cvv.Total = 100.25

	var crt sdbi.Cart
	crt.CustomerID = 4

	var cc CustomerCart
	cc.Items = &cilstp
	cc.CartView = &cvv
	cc.Cart = &crt
	cc.InsuranceID = 5

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cv := m.CalculateCartTotals(&cc, &head)

	fmt.Println("customer cv: ", cv)
	fmt.Println("Subtotal: ", cv.Subtotal)
	fmt.Println("ShippingHandling: ", cv.ShippingHandling)
	fmt.Println("InsuranceCost: ", cv.InsuranceCost)
	fmt.Println("Taxes: ", cv.Taxes)
	fmt.Println("Total: ", cv.Total)
	if cv.Total != 114.37 {
		t.Fail()
	}
}

func TestSix910Manager_CalculateCartTotals2(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var pd sdbi.Product
	pd.ID = 4
	pd.SalePrice = 28.95
	pd.ShortDesc = "test one"
	pd.Thumbnail = "/test/"
	sapi.MockProduct = &pd

	var smth sdbi.ShippingMethod
	smth.Cost = 4.55
	smth.Handling = 2.22
	sapi.MockShippingMethod = &smth

	var insc sdbi.Insurance
	insc.Cost = 2.34

	sapi.MockInsurance = &insc

	var uadd sdbi.Address
	uadd.Country = "USA"
	uadd.State = "GA"
	uadd.Zip = "12345"

	sapi.MockAddress = &uadd

	var trate1 sdbi.TaxRate
	trate1.Country = "USA"
	trate1.State = "GA"
	trate1.PercentRate = 5
	trate1.ZipStart = "12355"
	trate1.ZipEnd = "12356"

	var trate2 sdbi.TaxRate
	trate2.Country = "USA"
	trate2.State = "GA"
	trate2.PercentRate = 7
	trate2.ZipStart = ""
	trate2.ZipEnd = ""
	trate2.IncludeHandling = true
	trate2.IncludeShipping = true

	var trlst []sdbi.TaxRate
	trlst = append(trlst, trate1)
	trlst = append(trlst, trate2)

	sapi.MockTaxRateList = &trlst

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

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cvv CartView
	cvv.Total = 100.25

	var crt sdbi.Cart
	crt.CustomerID = 4

	var cc CustomerCart
	cc.Items = &cilstp
	cc.CartView = &cvv
	cc.Cart = &crt
	cc.InsuranceID = 5

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cv := m.CalculateCartTotals(&cc, &head)

	fmt.Println("customer cv: ", cv)
	fmt.Println("Subtotal: ", cv.Subtotal)
	fmt.Println("ShippingHandling: ", cv.ShippingHandling)
	fmt.Println("InsuranceCost: ", cv.InsuranceCost)
	fmt.Println("Taxes: ", cv.Taxes)
	fmt.Println("Total: ", cv.Total)
	if cv.Total != 116.86 {
		t.Fail()
	}
}

func TestSix910Manager_CalculateCartTotals3(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var pd sdbi.Product
	pd.ID = 4
	pd.SalePrice = 28.95
	pd.ShortDesc = "test one"
	pd.Thumbnail = "/test/"
	sapi.MockProduct = &pd

	var smth sdbi.ShippingMethod
	smth.Cost = 4.55
	smth.Handling = 2.22
	sapi.MockShippingMethod = &smth

	var insc sdbi.Insurance
	insc.Cost = 2.34

	sapi.MockInsurance = &insc

	var uadd sdbi.Address
	uadd.Country = "USA"
	uadd.State = "GA"
	uadd.Zip = "12345"

	sapi.MockAddress = &uadd

	var trate1 sdbi.TaxRate
	trate1.Country = "USA"
	trate1.State = "GA"
	trate1.PercentRate = 5
	trate1.ZipStart = "12355"
	trate1.ZipEnd = "12356"

	var trate2 sdbi.TaxRate
	trate2.Country = "USA"
	trate2.State = "GA"
	trate2.PercentRate = 7
	trate2.ZipStart = ""
	trate2.ZipEnd = ""
	trate2.IncludeHandling = true
	trate2.IncludeShipping = true

	var trlst []sdbi.TaxRate
	//trlst = append(trlst, trate1)
	trlst = append(trlst, trate2)

	sapi.MockTaxRateList = &trlst

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

	var cilstp []sdbi.CartItem

	var ctit1 sdbi.CartItem
	ctit1.Quantity = 3
	ctit1.ProductID = 7
	cilstp = append(cilstp, ctit1)

	var ctit2 sdbi.CartItem
	ctit2.Quantity = 4
	ctit2.ProductID = 9
	cilstp = append(cilstp, ctit2)

	var cvv CartView
	cvv.Total = 100.25

	var crt sdbi.Cart
	crt.CustomerID = 4

	var cc CustomerCart
	cc.Items = &cilstp
	cc.CartView = &cvv
	cc.Cart = &crt
	cc.InsuranceID = 5

	var head api.Headers
	//head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	cv := m.CalculateCartTotals(&cc, &head)

	fmt.Println("customer cv: ", cv)
	fmt.Println("Subtotal: ", cv.Subtotal)
	fmt.Println("ShippingHandling: ", cv.ShippingHandling)
	fmt.Println("InsuranceCost: ", cv.InsuranceCost)
	fmt.Println("Taxes: ", cv.Taxes)
	fmt.Println("Total: ", cv.Total)
	if cv.Total != 116.86 {
		t.Fail()
	}
}
