package handlers

import (
	"net/http"
	"strconv"
	"sync"

	api "github.com/Ulbora/Six910API-Go"
	six910api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
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

//ShipPage ShipPage
type ShipPage struct {
	Error         string
	Shipment      *sdbi.Shipment
	ShipmentItems *[]sdbi.ShipmentItem
	ShipmentBoxes *[]sdbi.ShipmentBox
	Shipments     *[]sdbi.Shipment
	Order         *sdbi.Order
	OrderItems    *[]sdbi.OrderItem
	OrderComments *[]sdbi.OrderComment
}

//StoreAdminAddShipmentPage StoreAdminAddShipmentPage
func (h *Six910Handler) StoreAdminAddShipmentPage(w http.ResponseWriter, r *http.Request) {
	gss, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gss) {
			asvars := mux.Vars(r)
			asidstr := asvars["id"]
			asOIID, _ := strconv.ParseInt(asidstr, 10, 64)
			aspErr := r.URL.Query().Get("error")
			var page ShipPage
			page.Error = aspErr
			hd := h.getHeader(gss)
			var wg sync.WaitGroup
			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				page.Order = h.API.GetOrder(oid, header.DeepCopy())
			}(asOIID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				page.OrderComments = h.API.GetOrderCommentList(oid, header.DeepCopy())
			}(asOIID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				page.OrderItems = h.API.GetOrderItemList(oid, header.DeepCopy())
			}(asOIID, hd)

			wg.Wait()
			h.Log.Debug("shipment page", page)
			// h.Log.Debug("shipment order", *page.Order)
			// h.Log.Debug("shipment order notes", *page.OrderComments)
			// h.Log.Debug("shipment order items", *page.OrderItems)

			h.AdminTemplates.ExecuteTemplate(w, adminAddShipmentPage, &page)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminAddShipment StoreAdminAddShipment
func (h *Six910Handler) StoreAdminAddShipment(w http.ResponseWriter, r *http.Request) {
	as, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(as) {
			sh := h.processShipment(r)
			h.Log.Debug("shipment in add", *sh)
			hd1 := h.getHeader(as)
			shres := h.API.AddShipment(sh, hd1)
			h.Log.Debug("shipment add resp", *shres)
			var success = true
			if shres.Success {
				hd2 := h.getHeader(as)
				oil := h.API.GetOrderItemList(sh.OrderID, hd2)
				var oichan = make(chan *api.ResponseID, len(*oil))
				var wg sync.WaitGroup
				for i := range *oil {
					wg.Add(1)
					go func(oi *sdbi.OrderItem, ch chan *api.ResponseID) {
						//do deep copy here
						defer wg.Done()
						hd := h.getHeader(as)
						h.Log.Debug("order item in goroutine", *oi)
						var si sdbi.ShipmentItem
						si.OrderItemID = oi.ID
						si.Quantity = oi.Quantity
						si.ShipmentID = shres.ID
						h.Log.Debug("shipment item in goroutine", si)
						ires := h.API.AddShipmentItem(&si, hd)
						ch <- ires
					}(&(*oil)[i], oichan)
				}
				wg.Wait()
				close(oichan)
				for res := range oichan {
					if !res.Success {
						success = false
					}
				}
			} else {
				success = false
			}
			h.Log.Debug("shipment all add suc", success)
			if success {
				http.Redirect(w, r, adminOrderListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminAddShipmentViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditShipmentPage StoreAdminEditShipmentPage
func (h *Six910Handler) StoreAdminEditShipmentPage(w http.ResponseWriter, r *http.Request) {
	ess, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(ess) {
			var esparm ShipPage
			edErr := r.URL.Query().Get("error")
			esparm.Error = edErr

			hd := h.getHeader(ess)
			esvars := mux.Vars(r)
			esidstr := esvars["id"]
			esID, _ := strconv.ParseInt(esidstr, 10, 64)
			h.Log.Debug("shipment id in edit", esID)

			ship := h.API.GetShipment(esID, hd)
			esparm.Shipment = ship
			h.Log.Debug("shipment in edit", ship)

			var wg sync.WaitGroup
			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.Order = h.API.GetOrder(oid, header.DeepCopy())
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.OrderComments = h.API.GetOrderCommentList(oid, header.DeepCopy())
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.OrderItems = h.API.GetOrderItemList(oid, header.DeepCopy())
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.Shipments = h.API.GetShipmentList(oid, header.DeepCopy())
			}(ship.OrderID, hd)

			wg.Add(1)
			go func(spid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.ShipmentBoxes = h.API.GetShipmentBoxList(spid, header.DeepCopy())
			}(ship.ID, hd)

			wg.Add(1)
			go func(spid int64, header *six910api.Headers) {
				defer wg.Done()
				esparm.ShipmentItems = h.API.GetShipmentItemList(spid, header.DeepCopy())
			}(ship.ID, hd)

			wg.Wait()

			h.Log.Debug("shipment page", esparm)

			h.AdminTemplates.ExecuteTemplate(w, adminEditShipmentPage, &esparm)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminEditShipment StoreAdminEditShipment
func (h *Six910Handler) StoreAdminEditShipment(w http.ResponseWriter, r *http.Request) {
	esss, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(esss) {
			epp := h.processShipment(r)
			h.Log.Debug("shipment update", *epp)
			hd := h.getHeader(esss)
			res := h.API.UpdateShipment(epp, hd)
			h.Log.Debug("shipment update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminOrderListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminEditShipmentViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminViewShipmentList StoreAdminViewShipmentList
func (h *Six910Handler) StoreAdminViewShipmentList(w http.ResponseWriter, r *http.Request) {
	sls, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment list view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(sls) {
			hd := h.getHeader(sls)
			vpvars := mux.Vars(r)
			oidstr := vpvars["oid"]
			foid, _ := strconv.ParseInt(oidstr, 10, 64)
			//shps := h.API.GetShipmentList(oid, hd)
			plErr := r.URL.Query().Get("error")
			var slparm ShipPage
			slparm.Error = plErr
			//slparm.Shipments = shps

			var wg sync.WaitGroup
			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				slparm.Order = h.API.GetOrder(oid, header.DeepCopy())
			}(foid, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				slparm.OrderComments = h.API.GetOrderCommentList(oid, header.DeepCopy())
			}(foid, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				slparm.OrderItems = h.API.GetOrderItemList(oid, header.DeepCopy())
			}(foid, hd)

			wg.Add(1)
			go func(oid int64, header *six910api.Headers) {
				defer wg.Done()
				slparm.Shipments = h.API.GetShipmentList(oid, header.DeepCopy())
			}(foid, hd)

			wg.Wait()
			h.Log.Debug("shipments in list", slparm)
			h.AdminTemplates.ExecuteTemplate(w, adminShipmentListView, &slparm)
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

//StoreAdminDeleteShipment StoreAdminDeleteShipment
func (h *Six910Handler) StoreAdminDeleteShipment(w http.ResponseWriter, r *http.Request) {
	dss, suc := h.getSession(r)
	h.Log.Debug("session suc in shipment delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dss) {
			hd := h.getHeader(dss)
			dsvars := mux.Vars(r)
			idstrd := dsvars["id"]
			idd, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeleteShipment(idd, hd)
			h.Log.Debug("shipment delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminShipmentListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminShipmentListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminloginPage, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processShipment(r *http.Request) *sdbi.Shipment {
	var p sdbi.Shipment
	id := r.FormValue("id")
	p.ID, _ = strconv.ParseInt(id, 10, 64)
	p.Status = r.FormValue("status")
	boxes := r.FormValue("boxes")
	p.Boxes, _ = strconv.ParseInt(boxes, 10, 64)
	shippingHandling := r.FormValue("shippingHandling")
	p.ShippingHandling, _ = strconv.ParseFloat(shippingHandling, 64)
	insurance := r.FormValue("insurance")
	p.Insurance, _ = strconv.ParseFloat(insurance, 64)
	orderID := r.FormValue("orderId")
	p.OrderID, _ = strconv.ParseInt(orderID, 10, 64)

	return &p
}
