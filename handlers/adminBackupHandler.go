package handlers

import (
	// tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	// "github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	// "strings"
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

//AdminBackupMainPage AdminBackupMainPage
func (h *Six910Handler) AdminBackupMainPage(w http.ResponseWriter, r *http.Request) {
	bkms, suc := h.getSession(r)
	h.Log.Debug("session suc in template add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(bkms) {
			//hd := h.getHeader(ats)
			h.AdminTemplates.ExecuteTemplate(w, adminBackupsPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminBackupUploadPage AdminBackupUploadPage
func (h *Six910Handler) AdminBackupUploadPage(w http.ResponseWriter, r *http.Request) {
	bks, suc := h.getSession(r)
	h.Log.Debug("session suc in template add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(bks) {
			//hd := h.getHeader(ats)
			h.AdminTemplates.ExecuteTemplate(w, adminBackupUploadPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminUploadBackups AdminUploadBackups
func (h *Six910Handler) AdminUploadBackups(w http.ResponseWriter, r *http.Request) {
	aubks, suc := h.getSession(r)
	h.Log.Debug("session suc in template upload", suc)
	if suc {
		if h.isStoreAdminLoggedIn(aubks) {
			bkerr := r.ParseMultipartForm(50000000)
			h.Log.Debug("ParseMultipartForm err: ", bkerr)

			file, handler, ferr := r.FormFile("backupFile")
			if ferr == nil {
				defer file.Close()
			}
			h.Log.Debug("backup file err: ", ferr)

			//h.Log.Debug("image file : ", *handler)

			bkdata, rferr := ioutil.ReadAll(file)
			h.Log.Debug("read file  err: ", rferr)

			h.Log.Debug("handler.Filename: ", handler.Filename)
			var suc bool
			if ferr == nil && rferr == nil {
				suc = h.BackupService.UploadBackups(&bkdata)
			}

			if suc {
				h.LoadTemplate()
				http.Redirect(w, r, adminBackups, http.StatusFound)
			} else {
				h.Log.Debug("backup upload of " + handler.Filename + " failed")
				http.Redirect(w, r, adminBackups, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//AdminDownloadBackups AdminDownloadBackups
func (h *Six910Handler) AdminDownloadBackups(w http.ResponseWriter, r *http.Request) {
	adbks, suc := h.getSession(r)
	h.Log.Debug("session suc in template upload", suc)
	if suc {
		if h.isStoreAdminLoggedIn(adbks) {
			suc, file := h.BackupService.DownloadBackups()
			h.Log.Debug("download backup suc: ", suc)
			if suc {
				w.Header().Set("Content-Disposition", "attachment; filename="+h.BackupFileName)
				w.Header().Set("Content-Type", r.Header.Get("Content-Type"))
				w.Write(*file)
				//var buf = bytes.NewBuffer(*file)
				//io.Copy(w, buf)
				//w.WriteHeader(http.StatusOK)
				// out, err := os.Create(h.BackupFileName)
				// if err == nil {
				// 	defer out.Close()
				// 	// Write the body to file
				// _, err = io.Copy(out, resp.Body)
				// }
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
