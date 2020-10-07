package handlers

import (
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
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

//StoreAdminAddImagePage StoreAdminAddImagePage
func (h *Six910Handler) StoreAdminAddImagePage(w http.ResponseWriter, r *http.Request) {
	is, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(is) {
			h.AdminTemplates.ExecuteTemplate(w, adminImageUploadPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUploadImage StoreAdminUploadImage
func (h *Six910Handler) StoreAdminUploadImage(w http.ResponseWriter, r *http.Request) {
	ius, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ius) {

			imperr := r.ParseMultipartForm(2000000)
			h.Log.Debug("ParseMultipartForm err: ", imperr)

			file, handler, ferr := r.FormFile("image")
			h.Log.Debug("image file err: ", ferr)
			defer file.Close()
			//h.Log.Debug("image file : ", *handler)

			data, rferr := ioutil.ReadAll(file)
			h.Log.Debug("read file  err: ", rferr)

			h.Log.Debug("handler.Filename: ", handler.Filename)

			aisuc := h.ImageService.AddImage(handler.Filename, data)

			if aisuc {
				http.Redirect(w, r, adminImageList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminImageListFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminImageList StoreAdminImageList
func (h *Six910Handler) StoreAdminImageList(w http.ResponseWriter, r *http.Request) {
	gils, suc := h.getSession(r)
	h.Log.Debug("session suc in ins view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gils) {
			ires := h.ImageService.GetImageList()
			//h.Log.Debug("image list in images: ", *res)
			h.AdminTemplates.ExecuteTemplate(w, adminImageListPage, &ires)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminDeleteImage StoreAdminDeleteImage
func (h *Six910Handler) StoreAdminDeleteImage(w http.ResponseWriter, r *http.Request) {
	dis, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dis) {
			vars := mux.Vars(r)
			iname := vars["name"]

			suc := h.ImageService.DeleteImage(iname)
			h.Log.Debug("image delete: ", suc)

			http.Redirect(w, r, adminImageList, http.StatusFound)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
