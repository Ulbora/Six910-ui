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
