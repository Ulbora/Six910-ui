package findfflsrv

import (
	"bytes"
	"encoding/json"
	"net/http"

	px "github.com/Ulbora/GoProxy"
	lg "github.com/Ulbora/Level_Logger"
)

/*
 Copyright (C) 2019 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.

 Copyright (C) 2019 Ken Williamson
 All rights reserved.

 This program is free software: you can redistribute it and/or modify
 it under the terms of the GNU General Public License as published by
 the Free Software Foundation, either version 3 of the License, or
 (at your option) any later version.
 This program is distributed in the hope that it will be useful,
 but WITHOUT ANY WARRANTY; without even the implied warranty of
 MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 GNU General Public License for more details.
 You should have received a copy of the GNU General Public License
 along with this program.  If not, see <http://www.gnu.org/licenses/>.

*/

//FFLZip FFLZip
type FFLZip struct {
	Zip string `json:"zip"`
}

//FFLKey FFLKey
type FFLKey struct {
	Key string `json:"id"`
}

//FFLList FFLList
type FFLList struct {
	Key            string `json:"id"`
	LicName        string `json:"licenseName"`
	BusName        string `json:"businessName"`
	PremiseAddress string `json:"premiseAddress"`
}

//FFL FFL
type FFL struct {
	Key            string `json:"id"`
	Lic            string `json:"license"`
	ExpDate        string `json:"expDate"`
	LicName        string `json:"licenseName"`
	BusName        string `json:"businessName"`
	PremiseAddress string `json:"premiseAddress"`
	Address        string `json:"address"`
	City           string `json:"city"`
	State          string `json:"state"`
	PremiseZip     string `json:"premiseZip"`
	MailingAddress string `json:"mailingAddress"`
	Phone          string `json:"phone"`
}

//FFLService FFLService
type FFLService interface {
	GetFFLList(zip string) (*[]FFLList, int)
	GetFFL(key string) (*FFL, int)
}

//Six910FFLService Six910FFLService
type Six910FFLService struct {
	Host   string
	Proxy  px.Proxy
	Log    *lg.Logger
	APIKey string
}

//New New
func (f *Six910FFLService) New() FFLService {
	var prxy px.GoProxy
	f.Proxy = prxy.GetNewProxy()
	return f
}

//SetProxy SetProxy
func (f *Six910FFLService) SetProxy(prxy px.Proxy) {
	f.Proxy = prxy
}

//GetFFLList GetFFLList
func (f *Six910FFLService) GetFFLList(zip string) (*[]FFLList, int) {
	var rtn []FFLList
	var code int
	var gURL = f.Host + "/rs/findByZip"
	var zipObj FFLZip
	zipObj.Zip = zip
	aJSON, err := json.Marshal(zipObj)
	// fmt.Println("err: ", err)
	f.Log.Debug("find ffl err: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", gURL, bytes.NewBuffer(aJSON))
		f.Log.Debug("get ffl list url: ", gURL)
		f.Log.Debug("get ffl list req err: ", rErr)
		// fmt.Println("url: ", gURL)
		// fmt.Println("list error: ", rErr)
		if rErr == nil {
			req.Header.Set("api-key", f.APIKey)
			req.Header.Set("Content-Type", "application/json")
			// var de bool
			_, code = f.Proxy.Do(req, &rtn)
			// fmt.Println("de: ", de)
			// fmt.Println("code: ", code)
			f.Log.Debug("GetFFLList code: ", code)
		}
	}

	return &rtn, code
}

//GetFFL GetFFL
func (f *Six910FFLService) GetFFL(key string) (*FFL, int) {
	var rtn FFL
	var code int
	var gURL = f.Host + "/rs/findById"
	var keyObj FFLKey
	keyObj.Key = key
	aJSON, err := json.Marshal(keyObj)
	// fmt.Println("err: ", err)
	f.Log.Debug("find ffl by key err: ", err)
	if err == nil {
		req, rErr := http.NewRequest("POST", gURL, bytes.NewBuffer(aJSON))
		f.Log.Debug("get ffl by id url: ", gURL)
		f.Log.Debug("get ffl by id req err: ", rErr)
		// fmt.Println("url: ", gURL)
		// fmt.Println("list error: ", rErr)
		if rErr == nil {
			req.Header.Set("api-key", f.APIKey)
			req.Header.Set("Content-Type", "application/json")
			// var de bool
			_, code = f.Proxy.Do(req, &rtn)
			// fmt.Println("de: ", de)
			// fmt.Println("code: ", code)
			f.Log.Debug("GetFFL code: ", code)
		}
	}
	return &rtn, code
}
