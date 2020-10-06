package handlers

import (
	"net/http"
	"sort"
	"strings"
	"time"

	sr "github.com/Ulbora/Six910-ui/contentsrv"
	img "github.com/Ulbora/Six910-ui/imgsrv"
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

//ContPage ContPage
type ContPage struct {
	Error     string
	ImageList *[]img.Image
	Content   *sr.Content
}

//StoreAdminAddContentPage StoreAdminAddContentPage
func (h *Six910Handler) StoreAdminAddContentPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			var cp ContPage
			res := h.ImageService.GetImageList()
			cp.ImageList = res
			h.Log.Debug("image list in content add: ", *res)
			h.AdminTemplates.ExecuteTemplate(w, adminAddContentPage, &cp)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddContent StoreAdminAddContent
func (h *Six910Handler) StoreAdminAddContent(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			ct := h.processContent(r)
			ct.CreateDate = time.Now()
			res := h.ContentService.AddContent(ct)
			h.Log.Debug("add content res", *res)
			if res.Success {
				http.Redirect(w, r, adminContentList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddContentFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUpdateContent StoreAdminUpdateContent
func (h *Six910Handler) StoreAdminUpdateContent(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			uct := h.processContent(r)
			uct.ModifiedDate = time.Now()
			res := h.ContentService.UpdateContent(uct)
			h.Log.Debug("update content res", *res)
			if res.Success {
				http.Redirect(w, r, adminContentList, http.StatusFound)
			} else {
				//go back
				http.Redirect(w, r, adminGetContent+"/"+uct.Name, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminGetContent StoreAdminGetContent
func (h *Six910Handler) StoreAdminGetContent(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			vars := mux.Vars(r)
			name := vars["name"]
			ires := h.ImageService.GetImageList()
			h.Log.Debug("image list in content get: ", *ires)

			_, cres := h.ContentService.GetContent(name)
			h.Log.Debug("content in content get: ", *cres)
			var ci ContPage
			ci.ImageList = ires
			ci.Content = cres
			h.Log.Debug("content and image list in get content: ", ci)

			h.AdminTemplates.ExecuteTemplate(w, adminUpdateContent, &ci)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminContentList StoreAdminContentList
func (h *Six910Handler) StoreAdminContentList(w http.ResponseWriter, r *http.Request) {
	gcls, suc := h.getSession(r)
	h.Log.Debug("session suc in ins view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gcls) {
			csl := h.ContentService.GetContentList(false)
			sort.Slice(*csl, func(p, q int) bool {
				return (*csl)[p].Title < (*csl)[q].Title
			})
			h.Log.Debug("Content  in list", csl)
			h.AdminTemplates.ExecuteTemplate(w, adminContentListPage, &csl)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminDeleteContent StoreAdminDeleteContent
func (h *Six910Handler) StoreAdminDeleteContent(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			vars := mux.Vars(r)
			name := vars["name"]

			res := h.ContentService.DeleteContent(name)
			h.Log.Debug("content delete in content delete: ", *res)

			http.Redirect(w, r, adminContentList, http.StatusFound)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processContent(r *http.Request) *sr.Content {
	name := r.FormValue("name")
	name = strings.Replace(name, " ", "", -1)
	h.Log.Debug("name in new content: ", name)

	content := r.FormValue("content")
	h.Log.Debug("content in new content: ", content)

	visible := r.FormValue("visible")
	h.Log.Debug("visible in new content: ", visible)

	title := r.FormValue("title")
	h.Log.Debug("title in new content: ", title)

	subject := r.FormValue("subject")
	h.Log.Debug("subject in new content: ", subject)

	author := r.FormValue("author")
	h.Log.Debug("author in new content: ", author)

	metaKeyWords := r.FormValue("metaKeyWords")
	h.Log.Debug("metaKeyWords in new content: ", metaKeyWords)

	metaDesc := r.FormValue("desc")
	h.Log.Debug("metaDesc in new content: ", metaDesc)

	blogpost := r.FormValue("blogpost")
	h.Log.Debug("blogpost in new content: ", blogpost)

	archived := r.FormValue("archived")
	h.Log.Debug("archived in new content: ", archived)

	var ct sr.Content
	ct.Author = author
	ct.MetaDesc = metaDesc
	ct.MetaKeyWords = metaKeyWords
	ct.Name = name
	ct.Title = title
	ct.Subject = subject
	ct.Text = content

	if archived == "on" {
		ct.Archived = true
	} else {
		ct.Archived = false
	}

	if blogpost == "on" {
		ct.BlogPost = true
	} else {
		ct.BlogPost = false
	}
	if visible == "on" {
		ct.Visible = true
	} else {
		ct.Visible = false
	}

	return &ct
}
