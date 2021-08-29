package findfflsrv

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"

	"testing"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

func TestSix910FFLService_GetFFLList(t *testing.T) {
	var c Six910FFLService
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`[
		{
			"id": "158223013K15918",
			"licenseName": "VAULT WORLDWIDE, LLC",
			"businessName": "THE VAULT WORLDWIDE, LLC",
			"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132"
		},
		{
			"id": "158223024A02897",
			"licenseName": "PATRIOT PAWN-N-SHOP, INC",
			"businessName": "",
			"premiseAddress": "562 HARDEE ST SUITE B\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F12198",
			"licenseName": "CITY PAWN OF DALLAS LLC",
			"businessName": "",
			"premiseAddress": "620 W MEMORIAL DR   STE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223013D10984",
			"licenseName": "WAHNER, BRIAN KEITH",
			"businessName": "WESTERN ARMS",
			"premiseAddress": "407 AMSTERDAM WAY\nDALLAS, GA 30132"
		},
		{
			"id": "158223074E11283",
			"licenseName": "RJL INC",
			"businessName": "PATRIOT ARMS AND SUPPLY",
			"premiseAddress": "105 VILLAGE WALK, SUITE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223074B07655",
			"licenseName": "U K PRECISION INC",
			"businessName": "",
			"premiseAddress": "2029 MARSHALL HUFF RD SUITE A\nDALLAS, GA 30132"
		},
		{
			"id": "158223071G08116",
			"licenseName": "SKY GUNS INTERNATIONAL LLC",
			"businessName": "SGI",
			"premiseAddress": "224 BRANCH VALLEY DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223013C00889",
			"licenseName": "STINSON, GLEN ELLIS",
			"businessName": "",
			"premiseAddress": "35 COURTHOUSE SQUARE\nDALLAS, GA 30132"
		},
		{
			"id": "158223073H15687",
			"licenseName": "SG3 GUNWORKS LLC",
			"businessName": "",
			"premiseAddress": "655 SHOALS TRAIL\nDALLAS, GA 30132"
		},
		{
			"id": "158223072M09852",
			"licenseName": "WHITTEMORE, HEATH BLAINE",
			"businessName": "HBW FIREARMS",
			"premiseAddress": "277 NEW FARM RD\nDALLAS, GA 30132"
		},
		{
			"id": "158223022F05648",
			"licenseName": "CATES, ROBERT CHRISTOPHER",
			"businessName": "C & C PAWN",
			"premiseAddress": "293 WI PARKWAY STE E\nDALLAS, GA 30132"
		},
		{
			"id": "158223014D11179",
			"licenseName": "ARMAGEDDON ARMS LLC",
			"businessName": "ARMAGEDDON ARMS",
			"premiseAddress": "350 BLACKBERRY RUN DR\nDALLAS, GA 30132"
		},
		{
			"id": "158223072F12092",
			"licenseName": "C PRECISION LLC",
			"businessName": "C PRECISION",
			"premiseAddress": "105 VILLAGE WALK STE 184\nDALLAS, GA 30132"
		},
		{
			"id": "158223011E11380",
			"licenseName": "CRABBE, TRENTON RAYMOND",
			"businessName": "",
			"premiseAddress": "2762 NARROWAY CHURCH CIRCLE\nDALLAS, GA 30132"
		}
	]`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "123456"
	s := c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	res, code := s.GetFFLList("30141")
	fmt.Print("res: ")
	fmt.Println(res)
	if (*res)[0].Key != "158223013K15918" || code != 200 {
		t.Fail()
	}

	// t.Fail()

}

func TestSix910FFLService_GetFFL(t *testing.T) {
	var c Six910FFLService
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	c.Log = &l
	var p px.MockGoProxy
	p.MockDoSuccess1 = true
	var ress http.Response
	ress.Body = ioutil.NopCloser(bytes.NewBufferString(`{
		"id": "158223013K15918",
		"license": "1-58-223-01-3K-15918",
		"expDate": "October 1, 2023",
		"licenseName": "VAULT WORLDWIDE, LLC",
		"businessName": "THE VAULT WORLDWIDE, LLC",
		"premiseAddress": "531 HARDEE STREET\nDALLAS, GA 30132",
		"address": "531 HARDEE STREET",
		"city": "DALLAS",
		"state": "GA",
		"premiseZip": "30132",
		"mailingAddress": "218 FIVE OAKS DRIVE\nHIRAM, GA 30141",
		"phone": "404-374-6970"
	}`))
	p.MockResp = &ress
	p.MockRespCode = 200

	c.Host = "http://api.findfflbyzip.com"

	c.APIKey = "3456789"
	s := c.New()
	c.SetProxy(&p)
	fmt.Println("f.Proxy in test: ", c.Proxy)

	res, code := s.GetFFL("158223013K15918")
	fmt.Print("res: ")
	fmt.Println(res)
	if res.Key != "158223013K15918" || code != 200 {
		t.Fail()
	}

	// t.Fail()
}
