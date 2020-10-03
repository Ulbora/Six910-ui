package handlers

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"os"
	//ml "github.com/Ulbora/go-mail-sender"
	"io/ioutil"
	"net/http"
	//"strconv"
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

//StoreAdminUploadImageFilesPage StoreAdminUploadImageFilesPage
func (h *Six910Handler) StoreAdminUploadImageFilesPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session image files suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			h.AdminTemplates.ExecuteTemplate(w, imageFilesUploadPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUploadImageFiles StoreAdminUploadImageFiles
func (h *Six910Handler) StoreAdminUploadImageFiles(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {

			iuplerr := r.ParseMultipartForm(50000000)
			h.Log.Debug("ParseMultipartForm err: ", iuplerr)

			file, handler, ferr := r.FormFile("imageFilesUpload")
			if ferr == nil {
				defer file.Close()
			}
			h.Log.Debug("file err: ", ferr)

			iupdata, rferr := ioutil.ReadAll(file)
			h.Log.Debug("read file  err: ", rferr)
			//h.Log.Debug("read file  bkdata: ", bkdata)

			h.Log.Debug("handler.Filename: ", handler.Filename)

			//h.Log.Debug("updata not zip: ", string(updata))

			rd := bytes.NewReader(iupdata)
			gzf, gzerr := gzip.NewReader(rd)
			h.Log.Debug("gz file reader err : ", gzerr)
			var suc bool
			if gzerr == nil {
				tr := tar.NewReader(gzf)
				for {
					header, err := tr.Next()
					h.Log.Debug("tar.NewReader err : ", gzerr)
					if err == io.EOF {
						break
					}

					info := header.FileInfo()
					h.Log.Debug("tar.NewReader info isdir : ", info.IsDir())
					h.Log.Debug("tar.NewReader info name: ", info.Name())
					if !info.IsDir() {
						suc = true
						file, err := os.OpenFile(h.ImagePath+string(os.PathSeparator)+info.Name(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, info.Mode())
						h.Log.Debug("os.OpenFile err : ", err)
						h.Log.Debug("os.OpenFile : ", h.ImagePath+info.Name())
						if err == nil {
							defer file.Close()
							_, err = io.Copy(file, tr)
						}
					}
				}
			}
			if suc {
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				h.Log.Debug("image files upload of " + handler.Filename + " failed")
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUploadThumbnailFilesPage StoreAdminUploadThumbnailFilesPage
func (h *Six910Handler) StoreAdminUploadThumbnailFilesPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session thumbnail files suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			h.AdminTemplates.ExecuteTemplate(w, thumbnailFilesUploadPage, nil)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUploadThumbnailFiles StoreAdminUploadThumbnailFiles
func (h *Six910Handler) StoreAdminUploadThumbnailFiles(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session thumbnail suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {

			ituplerr := r.ParseMultipartForm(50000000)
			h.Log.Debug("ParseMultipartForm err: ", ituplerr)

			tfile, handler, ferr := r.FormFile("thumbnailFilesUpload")
			if ferr == nil {
				defer tfile.Close()
			}
			h.Log.Debug("file err: ", ferr)

			itupdata, rferr := ioutil.ReadAll(tfile)
			h.Log.Debug("read file  err: ", rferr)
			//h.Log.Debug("read file  bkdata: ", bkdata)

			h.Log.Debug("thumbnail handler.Filename: ", handler.Filename)

			//h.Log.Debug("updata not zip: ", string(updata))

			rd := bytes.NewReader(itupdata)
			gzf, gzerr := gzip.NewReader(rd)
			h.Log.Debug("gz thumbnail file reader err : ", gzerr)
			var tsuc bool
			if gzerr == nil {
				tr := tar.NewReader(gzf)
				for {
					header, err := tr.Next()
					h.Log.Debug("tar.NewReader err : ", gzerr)
					if err == io.EOF {
						break
					}

					tinfo := header.FileInfo()
					h.Log.Debug("tar.NewReader info isdir : ", tinfo.IsDir())
					h.Log.Debug("tar.NewReader info name: ", tinfo.Name())
					if !tinfo.IsDir() {
						tsuc = true
						file, err := os.OpenFile(h.ThumbnailPath+string(os.PathSeparator)+tinfo.Name(), os.O_CREATE|os.O_TRUNC|os.O_WRONLY, tinfo.Mode())
						h.Log.Debug("os.OpenFile err : ", err)
						h.Log.Debug("os.OpenFile : ", h.ImagePath+tinfo.Name())
						if err == nil {
							defer file.Close()
							_, err = io.Copy(file, tr)
						}
					}
				}
			}
			if tsuc {
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				h.Log.Debug("thumbnail files upload of " + handler.Filename + " failed")
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
