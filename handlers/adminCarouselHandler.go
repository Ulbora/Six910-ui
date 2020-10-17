package handlers

import (
	"net/http"

	carsrv "github.com/Ulbora/Six910-ui/carouselsrv"
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

//StoreAdminGetCarousel StoreAdminGetCarousel
func (h *Six910Handler) StoreAdminGetCarousel(w http.ResponseWriter, r *http.Request) {
	cars, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(cars) {
			crvars := mux.Vars(r)
			crname := crvars["name"]
			_, cres := h.CarouselService.GetCarousel(crname)
			h.Log.Debug("Carousel in  get: ", *cres)
			var ci ContPage
			ci.Carousel = cres
			h.Log.Debug("Carousel in page: ", ci)

			h.AdminTemplates.ExecuteTemplate(w, adminEditCarousel, &ci)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminUpdateCarousel StoreAdminUpdateCarousel
func (h *Six910Handler) StoreAdminUpdateCarousel(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {

			var car carsrv.Carousel

			name := r.FormValue("name")
			h.Log.Debug("Carousel name in update: ", name)

			enabled := r.FormValue("enabled")
			h.Log.Debug("Carousel enabled in update: ", enabled)

			image1 := r.FormValue("image1")
			h.Log.Debug("Carousel image1 in update: ", image1)

			image2 := r.FormValue("image2")
			h.Log.Debug("Carousel image2 in update: ", image2)

			image3 := r.FormValue("image3")
			h.Log.Debug("Carousel image3 in update: ", image3)

			car.Name = name
			if enabled == "on" {
				car.Enabled = true
			} else {
				car.Enabled = false
			}
			car.Image1 = image1
			car.Image2 = image2
			car.Image3 = image3

			res := h.CarouselService.UpdateCarousel(&car)
			h.Log.Debug("update Carousel res", res)
			if res {
				http.Redirect(w, r, adminIndex, http.StatusFound)
			} else {
				//go back
				http.Redirect(w, r, adminEditCarousel+"/"+name, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
