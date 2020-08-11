package handlers

import (
	b64 "encoding/base64"
	"io/ioutil"
	"net/http"

	api "github.com/Ulbora/Six910API-Go"
)

/*
 Six910 is a shopping cart and E-commerce system.
 Copyright (C) 2020 Ulbora Labs LLC. (www.ulboralabs.com)
 All rights reserved.
 Copyright (C) 2020 Ken Williamson
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

//PageValues PageValues
type PageValues struct {
	Suc                bool
	RecordsNotImported int
}

//StoreAdminUploadProductFilePage StoreAdminUploadProductFilePage
func (h *Six910Handler) StoreAdminUploadProductFilePage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Values["loggedIn"]
		storeAdminUser := s.Values["storeAdminUser"]
		h.Log.Debug("loggedIn in backups: ", loggedInAuth)
		if loggedInAuth == true && storeAdminUser == true {
			h.AdminTemplates.ExecuteTemplate(w, productFileUploadPage, nil)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminUploadProductFile StoreAdminUploadProductFile
func (h *Six910Handler) StoreAdminUploadProductFile(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		loggedInAuth := s.Values["loggedIn"]
		storeAdminUser := s.Values["storeAdminUser"]
		h.Log.Debug("loggedIn in upload : ", loggedInAuth)
		h.Log.Debug("storeAdminUser in upload : ", storeAdminUser)
		if loggedInAuth == true && storeAdminUser == true {

			uplerr := r.ParseMultipartForm(50000000)
			h.Log.Debug("ParseMultipartForm err: ", uplerr)

			file, handler, ferr := r.FormFile("productupload")
			if ferr == nil {
				defer file.Close()
			}
			h.Log.Debug("file err: ", ferr)

			updata, rferr := ioutil.ReadAll(file)
			h.Log.Debug("read file  err: ", rferr)
			//h.Log.Debug("read file  bkdata: ", bkdata)

			h.Log.Debug("handler.Filename: ", handler.Filename)
			var hd api.Headers
			if !h.OAuth2Enabled {
				username := s.Values["username"]
				password := s.Values["password"]
				sEnccl := b64.StdEncoding.EncodeToString([]byte(username.(string) + ":" + password.(string)))
				h.Log.Debug("sEnc: ", sEnccl)
				hd.Set("Authorization", "Basic "+sEnccl)
			} else {
				hd.Set("Authorization", "Bearer "+h.token.AccessToken)
			}

			suc, notImported := h.Manager.UploadProductFile(updata, &hd)
			h.Log.Debug("notImported: ", notImported)

			var pg PageValues
			if suc {
				pg.Suc = suc
				pg.RecordsNotImported = notImported
				h.AdminTemplates.ExecuteTemplate(w, productUploadResultPage, &pg)
			} else {
				h.Log.Debug("csv upload of " + handler.Filename + " failed")
				h.AdminTemplates.ExecuteTemplate(w, productUploadResultPage, &pg)
			}
		} else {
			h.authorize(w, r)
		}
	}
}
