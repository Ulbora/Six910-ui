package managers

import (
	"fmt"
	"io/ioutil"
	"testing"

	lg "github.com/Ulbora/Level_Logger"
	mapi "github.com/Ulbora/Six910-ui/mockapi"
	api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
)

func TestSix910Manager_UploadProductFile(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dlst []sdbi.Distributor
	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

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

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	fmt.Println("readFile err: ", err)
	scnt, cnt := m.UploadProductFile(sourceFile, &head)
	fmt.Println("suc upload file: ", scnt)
	fmt.Println("not imported: ", cnt)

	if scnt == 0 {
		t.Fail()
	}
}

func TestSix910Manager_UploadProductFileExistingDist(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var dist sdbi.Distributor
	dist.Company = "bobs warehouse"
	dist.ID = 4
	var dlst []sdbi.Distributor
	dlst = append(dlst, dist)

	sapi.MockDistributorList = &dlst

	var sares api.ResponseID
	sares.Success = true
	sares.ID = 5

	sapi.MockAddDistributorResp = &sares

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

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

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	m := sm.GetNew()
	sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	fmt.Println("readFile err: ", err)
	scnt, cnt := m.UploadProductFile(sourceFile, &head)
	fmt.Println("suc upload file: ", scnt)
	fmt.Println("not imported: ", cnt)

	if scnt == 0 {
		t.Fail()
	}
}

func TestSix910Manager_processProductCategory(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

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
	//m := sm.GetNew()
	//sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	//fmt.Println("readFile err: ", err)
	var cat = "/cat1/cat2/cat2"
	var p Product
	suc := sm.processProductCategory(cat, &p, &head)
	fmt.Println("processProductCategory: ", suc)

	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_processProductCategoryFail(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	//cr1.Success = true
	//cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	//cr2.Success = true
	//cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	//cr3.Success = true
	//cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

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
	//m := sm.GetNew()
	//sourceFile, err := ioutil.ReadFile("../scripts/testUploadFile.csv")
	//fmt.Println("readFile err: ", err)
	var cat = "/cat1/cat2/cat2"
	var p Product
	suc := sm.processProductCategory(cat, &p, &head)
	fmt.Println("processProductCategory2: ", suc)

	if suc {
		t.Fail()
	}
}

func TestSix910Manager_createCategory(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var ctlist []sdbi.Category
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

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

	var catList = []string{"cat1", "cat2", "cat3"}
	catID := sm.createCategory(&catList, &head)
	fmt.Println("catID: ", catID)

	if catID == 0 {
		t.Fail()
	}
}

func TestSix910Manager_createCategoryExistingCat(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI
	var cart sdbi.Cart
	cart.CustomerID = 18
	cart.ID = 3
	cart.StoreID = 59
	sapi.MockCart = &cart

	var c1 sdbi.Category
	c1.ID = 3
	c1.Name = "cat1"

	var ctlist []sdbi.Category
	ctlist = append(ctlist, c1)
	sapi.MockCategoryList = &ctlist

	var cr1 api.ResponseID
	cr1.Success = true
	cr1.ID = 5
	sapi.MockAddCategoryResp1 = &cr1

	var cr2 api.ResponseID
	cr2.Success = true
	cr2.ID = 6
	sapi.MockAddCategoryResp2 = &cr2

	var cr3 api.ResponseID
	cr3.Success = true
	cr3.ID = 7
	sapi.MockAddCategoryResp3 = &cr3

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

	var catList = []string{"cat1", "cat2", "cat3"}
	catID := sm.createCategory(&catList, &head)
	fmt.Println("catID: ", catID)

	if catID == 0 {
		t.Fail()
	}
}

func TestSix910Manager_processParentProduct(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	prod.ID = 4

	sapi.MockProduct = &prod

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

	var sku = "12345"
	var p Product
	p.DistributorID = 4
	suc := sm.processParentProduct(sku, &p, &head)
	fmt.Println("suc: ", suc)

	if !suc {
		t.Fail()
	}
}

func TestSix910Manager_processParentProductFail(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

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

	var sku = "12345"
	var p Product
	p.DistributorID = 4
	suc := sm.processParentProduct(sku, &p, &head)
	fmt.Println("suc: ", suc)

	if suc {
		t.Fail()
	}
}

func TestSix910Manager_prepProducts(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	//prod.ID = 4

	sapi.MockProduct = &prod

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

	var csvRecs [][]string
	var hr = []string{"parent_product_sku"}
	var vr = []string{"12345"}
	csvRecs = append(csvRecs, hr)
	csvRecs = append(csvRecs, vr)

	pl, hl := sm.prepProducts(4, &csvRecs, &head)
	fmt.Println("pl: ", pl)
	fmt.Println("hl: ", hl)

	if len(*hl) == 0 {
		t.Fail()
	}
}

func TestSix910Manager_processHoldList(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	prod.ID = 4

	sapi.MockProduct = &prod

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

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var sku = "12345"
	var sku1 = "55555"
	var holdSkuList []string
	holdSkuList = append(holdSkuList, sku)
	holdSkuList = append(holdSkuList, sku1)

	var hp Product
	var hp2 Product

	var holdProdList []Product
	holdProdList = append(holdProdList, hp)
	holdProdList = append(holdProdList, hp2)

	var p Product
	p.DistributorID = 4

	var prodList []Product
	prodList = append(prodList, p)

	fmt.Println("prodList len before: ", len(prodList))

	hprods := sm.processHoldList(&prodList, &holdSkuList, &holdProdList, &head)

	fmt.Println("prodList len after: ", len(prodList))
	fmt.Println("hprods: ", hprods)

	if len(*hprods) > 0 {
		t.Fail()
	}
}

func TestSix910Manager_processHoldListWrongSize(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	prod.ID = 4

	sapi.MockProduct = &prod

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

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var sku = "12345"
	//var sku1 = "55555"
	var holdSkuList []string
	holdSkuList = append(holdSkuList, sku)
	//holdSkuList = append(holdSkuList, sku1)

	var hp Product
	var hp2 Product

	var holdProdList []Product
	holdProdList = append(holdProdList, hp)
	holdProdList = append(holdProdList, hp2)

	var p Product
	p.DistributorID = 4

	var prodList []Product
	prodList = append(prodList, p)

	hprods := sm.processHoldList(&prodList, &holdSkuList, &holdProdList, &head)
	fmt.Println("hprods: ", hprods)

	if len(*hprods) == 0 {
		t.Fail()
	}
}

func TestSix910Manager_processHoldListNotFound(t *testing.T) {
	var sm Six910Manager

	//var sapi api.Six910API

	//-----------start mocking------------------
	var sapi mapi.MockAPI

	var prod sdbi.Product
	//prod.ID = 4

	sapi.MockProduct = &prod

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

	var head api.Headers
	head.Set("Authorization", "Basic YWRtaW46YWRtaW4=")

	var sku = "12345"
	var sku1 = "55555"
	var holdSkuList []string
	holdSkuList = append(holdSkuList, sku)
	holdSkuList = append(holdSkuList, sku1)

	var hp Product
	var hp2 Product

	var holdProdList []Product
	holdProdList = append(holdProdList, hp)
	holdProdList = append(holdProdList, hp2)

	var p Product
	p.DistributorID = 4

	var prodList []Product
	prodList = append(prodList, p)

	hprods := sm.processHoldList(&prodList, &holdSkuList, &holdProdList, &head)
	fmt.Println("hprods: ", hprods)

	if len(*hprods) == 0 {
		t.Fail()
	}
}
