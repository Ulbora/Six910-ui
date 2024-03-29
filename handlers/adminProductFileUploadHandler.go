package handlers

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	ml "github.com/Ulbora/go-mail-sender"
	"io/ioutil"
	"net/http"
	"strconv"
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
		if h.isStoreAdminLoggedIn(s) {
			h.AdminTemplates.ExecuteTemplate(w, productFileUploadPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUploadProductFile StoreAdminUploadProductFile
func (h *Six910Handler) StoreAdminUploadProductFile(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {

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

			//h.Log.Debug("updata not zip: ", string(updata))
			dcupdata := h.extractTarGz(&updata)
			h.Log.Debug("updata file in handlers: ", string(dcupdata))

			hd := h.getHeader(s)
			h.Log.Debug("header: ", hd)
			h.Log.Debug("manager: ", h.Manager)
			icnt, notImported := h.Manager.UploadProductFile(dcupdata, hd)
			h.Log.Debug("Imported: ", icnt)
			h.Log.Debug("notImported: ", notImported)
			if h.MailSenderAddress != "" && icnt > 0 {
				var icntStr = strconv.Itoa(icnt)
				var m ml.Mailer
				m.Subject = h.MailSubject
				m.Body = mailMessageUploadComplete + icntStr
				str := h.API.GetStore(h.StoreName, h.LocalDomain, hd)
				m.Recipients = []string{str.Email}
				m.SenderAddress = h.MailSenderAddress

				sendSuc := h.MailSender.SendMail(&m)
				h.Log.Debug("sendSuc in contact: ", sendSuc)
			}

			//var pg PageValues
			if icnt != 0 {
				//pg.Suc = suc
				//pg.RecordsNotImported = notImported
				//h.AdminTemplates.ExecuteTemplate(w, productUploadResultPage, &pg)
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				h.Log.Debug("csv upload of " + handler.Filename + " failed")
				//h.AdminTemplates.ExecuteTemplate(w, productUploadResultPage, &pg)
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) extractTarGz(cf *[]byte) []byte {
	var rtn []byte
	r := bytes.NewReader(*cf)
	gzf, gzerr := gzip.NewReader(r)
	h.Log.Debug("gz file reader err : ", gzerr)
	if gzerr != nil {
		rtn = *cf
	} else {
		tr := tar.NewReader(gzf)
		tr.Next()
		rtn, _ = ioutil.ReadAll(tr)
	}
	h.Log.Debug("file reader data : ", string(rtn))
	return rtn
}
