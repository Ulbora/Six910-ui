package handlers

import (
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strings"
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

//AdminAddTemplatePage AdminAddTemplatePage
func (h *Six910Handler) AdminAddTemplatePage(w http.ResponseWriter, r *http.Request) {
	ats, suc := h.getSession(r)
	h.Log.Debug("session suc in template add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ats) {
			//hd := h.getHeader(ats)
			h.AdminTemplates.ExecuteTemplate(w, adminTemplateUploadPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminUploadTemplate AdminUploadTemplate
func (h *Six910Handler) AdminUploadTemplate(w http.ResponseWriter, r *http.Request) {
	auts, suc := h.getSession(r)
	h.Log.Debug("session suc in template upload", suc)
	if suc {
		if h.isStoreAdminLoggedIn(auts) {
			mperr := r.ParseMultipartForm(10000000)
			h.Log.Debug("ParseMultipartForm err: ", mperr)

			tfile, handler, ferr := r.FormFile("tempFile")
			h.Log.Debug("template file err: ", ferr)
			defer tfile.Close()
			//h.Log.Debug("template file : ", *handler)

			data, rferr := ioutil.ReadAll(tfile)
			h.Log.Debug("read file  err: ", rferr)

			i := strings.Index(handler.Filename, ".")
			var tname = string(handler.Filename[:i])
			h.Log.Debug("tname: ", tname)

			suc := h.TemplateService.AddTemplateFile(tname, handler.Filename, data)
			var tasus bool
			h.Log.Debug("AddTemplateFile: ", suc)
			if suc {
				var tmp tmpsrv.Template
				tmp.Name = tname
				tasus = h.TemplateService.AddTemplate(&tmp)
			}
			if tasus {
				http.Redirect(w, r, adminTemplates, http.StatusFound)
			} else {
				h.Log.Debug("Template upload of " + handler.Filename + " failed")
				h.AdminTemplates.ExecuteTemplate(w, adminTemplateUploadPage, nil)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminTemplateList AdminTemplateList
func (h *Six910Handler) AdminTemplateList(w http.ResponseWriter, r *http.Request) {
	gtls, suc := h.getSession(r)
	h.Log.Debug("session suc in cats view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gtls) {
			h.Log.Debug("template: ", h.AdminTemplates)
			res := h.TemplateService.GetTemplateList()
			//h.Log.Debug("templates in admin template list: ", *res)

			h.AdminTemplates.ExecuteTemplate(w, adminTemplateList, &res)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminActivateTemplate AdminActivateTemplate
func (h *Six910Handler) AdminActivateTemplate(w http.ResponseWriter, r *http.Request) {
	gtls, suc := h.getSession(r)
	h.Log.Debug("session suc in cats view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gtls) {
			h.Log.Debug("template: ", h.AdminTemplates)
			vars := mux.Vars(r)
			name := vars["name"]
			res := h.TemplateService.ActivateTemplate(name)
			if res {
				h.LoadTemplate()
			}
			h.Log.Debug("activate templates in admin: ", res)
			http.Redirect(w, r, adminTemplates, http.StatusFound)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminDeleteTemplate AdminDeleteTemplate
func (h *Six910Handler) AdminDeleteTemplate(w http.ResponseWriter, r *http.Request) {
	dtls, suc := h.getSession(r)
	h.Log.Debug("session suc in cats view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dtls) {
			h.Log.Debug("template: ", h.AdminTemplates)
			vars := mux.Vars(r)
			name := vars["name"]
			suc := h.TemplateService.DeleteTemplate(name)
			h.Log.Debug("delete template in admin: ", suc)
			if suc {
				suc = h.TemplateService.DeleteTemplateFile(name)
				h.Log.Debug("delete templates files in admin: ", suc)
			}
			http.Redirect(w, r, adminTemplates, http.StatusFound)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
