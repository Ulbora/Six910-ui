package managers

import (
	"fmt"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Manager_importProducts(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	// var ctlist []sdbi.Category
	// sapi.MockCategoryList = &ctlist

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = true

	sapi.MockAddProductCategoryResp = &cr

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

	// var cp CustomerProduct
	// cp.CustomerID = 18
	// cp.ProductID = 7
	// cp.Quantity = 3
	// cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//m := sm.GetNew()
	//sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	//fmt.Println("readFile err: ", err)
	// var cat = "/cat1/cat2/cat2"
	// var p sdbi.Product

	var plst []Product

	var p1 Product
	p1.CategoryID = 2
	p1.Sku = "12345"
	p1.Name = "test product1"
	p1.Desc = "test1"
	p1.Price = 10.00
	p1.Manufacturer = "testco"
	plst = append(plst, p1)

	var p2 Product
	p2.CategoryID = 2
	p2.Sku = "123456"
	p2.Name = "test product2"
	p2.Desc = "test2"
	p2.Price = 20.00
	p2.Manufacturer = "testco"
	plst = append(plst, p2)

	for i := range plst {
		fmt.Printf("add of slice elements in test: %p\n", &plst[i])
	}

	suc := sm.importProducts(&plst, &head)
	fmt.Println("importProducts: ", suc)

	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_importProductsFail(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	// var ctlist []sdbi.Category
	// sapi.MockCategoryList = &ctlist

	var pr api.ResponseID
	pr.Success = true
	pr.ID = 5
	sapi.MockAddProductResp = &pr

	var cr api.Response
	cr.Success = false

	sapi.MockAddProductCategoryResp = &cr

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

	// var cp CustomerProduct
	// cp.CustomerID = 18
	// cp.ProductID = 7
	// cp.Quantity = 3
	// cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//m := sm.GetNew()
	//sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	//fmt.Println("readFile err: ", err)
	// var cat = "/cat1/cat2/cat2"
	// var p sdbi.Product

	var plst []Product

	var p1 Product
	p1.CategoryID = 2
	p1.Sku = "12345"
	p1.Name = "test product1"
	p1.Desc = "test1"
	p1.Price = 10.00
	p1.Manufacturer = "testco"
	plst = append(plst, p1)

	var p2 Product
	p2.CategoryID = 2
	p2.Sku = "123456"
	p2.Name = "test product2"
	p2.Desc = "test2"
	p2.Price = 20.00
	p2.Manufacturer = "testco"
	plst = append(plst, p2)

	suc := sm.importProducts(&plst, &head)
	fmt.Println("importProducts: ", suc)

	if suc {
		t.Fail()
	}
}

func TestSix910Manager_importExistingProducts(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	// var ctlist []sdbi.Category
	// sapi.MockCategoryList = &ctlist

	var prod sdbi.Product
	prod.ID = 55
	prod.Sku = "12345"
	sapi.MockProduct = &prod

	var pr api.Response
	pr.Success = true
	sapi.MockUpdateProductResp = &pr

	var cr api.Response
	cr.Success = true

	sapi.MockAddProductCategoryResp = &cr

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

	// var cp CustomerProduct
	// cp.CustomerID = 18
	// cp.ProductID = 7
	// cp.Quantity = 3
	// cp.StoreID = 59

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	//m := sm.GetNew()
	//sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	//fmt.Println("readFile err: ", err)
	// var cat = "/cat1/cat2/cat2"
	// var p sdbi.Product

	var plst []Product

	var p1 Product
	p1.CategoryID = 2
	p1.Sku = "12345"
	p1.Name = "test product1"
	p1.Desc = "test1"
	p1.Price = 10.00
	p1.Manufacturer = "testco"
	plst = append(plst, p1)

	var p2 Product
	p2.CategoryID = 2
	p2.Sku = "123456"
	p2.Name = "test product2"
	p2.Desc = "test2"
	p2.Price = 20.00
	p2.Manufacturer = "testco"
	plst = append(plst, p2)

	for i := range plst {
		fmt.Printf("add of slice elements in test: %p\n", &plst[i])
	}

	suc := sm.importProducts(&plst, &head)
	fmt.Println("importProducts: ", suc)

	if !suc {
		t.Fail()
	}
}
