package handlers

import (
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

//VisitorData VisitorData
type VisitorData struct {
	Day      string
	Visitors int64
}

//SalesData SalesData
type SalesData struct {
	Day   string
	Sales float64
}

//Charts Charts
type Charts struct {
	VisitorData [][]interface{}
	SalesData   [][]interface{}
}

//StoreAdminIndex StoreAdminIndex
func (h *Six910Handler) StoreAdminIndex(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			hd := h.getHeader(s)
			vdl := h.API.GetVisitorData(hd)
			h.Log.Debug("vdl", *vdl)

			sdl := h.API.GetOrderSalesData(hd)
			h.Log.Debug("sdl", *sdl)
			// loggedInAuth := s.Values["loggedIn"]
			// storeAdminUser := s.Values["storeAdminUser"]
			// h.Log.Debug("loggedIn in backups: ", loggedInAuth)
			// if loggedInAuth == true && storeAdminUser == true {

			var arr [][]interface{}
			for _, vd := range *vdl {
				var e1 []interface{}
				m := vd.VisitDate.Month()
				d := vd.VisitDate.Day()
				e1 = append(e1, strconv.Itoa(int(m))+"/"+strconv.Itoa(d))
				e1 = append(e1, vd.VisitCount)
				arr = append(arr, e1)
			}

			var salesArr [][]interface{}
			for _, sd := range *sdl {
				var e1 []interface{}
				m := sd.OrderDate.Month()
				d := sd.OrderDate.Day()
				e1 = append(e1, strconv.Itoa(int(m))+"/"+strconv.Itoa(d))
				e1 = append(e1, sd.OrderTotal)
				salesArr = append(salesArr, e1)
			}

			var crts Charts
			crts.VisitorData = arr
			crts.SalesData = salesArr

			h.AdminTemplates.ExecuteTemplate(w, adminIndexPage, crts)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}
