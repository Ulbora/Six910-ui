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
	// var cat = "/cat1/cat2/cat2"
	// var p sdbi.Product

	var plst []Product

	suc := sm.importProducts(&plst, &head)
	fmt.Println("importProducts: ", suc)

	if !suc {
		t.Fail()
	}
}
