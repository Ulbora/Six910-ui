package handlers

import (
	"net/http"
	"strconv"
	"sync"

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

//ProdPage ProdPage
type ProdPage struct {
	Error           string
	Product         *sdbi.Product
	Products        *[]sdbi.Product
	CategoryList    *[]sdbi.Category
	DistributorList *[]sdbi.Distributor
	Pagination      *Pagination
	HasProducts     bool
	ExistingCats    []int64
	LastCatSearch   int64
	// Pages    *[]Pageinate
	// PrevLink string
	// NextLink string
}

// //Pageinate Pageinate
// type Pageinate struct {
// 	PageCount int
// 	Active    string
// 	// Start     int
// 	// End       int
// 	// PrevStart int
// 	// PrevEnd   int
// 	// NextStart int
// 	// NextEnd   int
// 	//PrevLink string
// 	Link string
// 	// NextLink string
// 	First bool
// 	Last  bool
// }

//StoreAdminAddProductPage StoreAdminAddProductPage
func (h *Six910Handler) StoreAdminAddProductPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			hd := h.getHeader(s)
			loginErr := r.URL.Query().Get("error")
			var lge ProdPage
			lge.Error = loginErr

			var wg sync.WaitGroup
			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				cats := h.API.GetHierarchicalCategoryList(header)
				h.Log.Debug("prod  in edit", cats)
				lge.CategoryList = cats
			}(hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				dist := h.API.GetDistributorList(header)
				h.Log.Debug("prod  in edit", dist)
				lge.DistributorList = dist
			}(hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminAddProductPage, &lge)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddProduct StoreAdminAddProduct
func (h *Six910Handler) StoreAdminAddProduct(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			p := h.processProduct(r)
			h.Log.Debug("prod add", *p)
			hd := h.getHeader(s)
			prres := h.API.AddProduct(p, hd)
			r.ParseForm()
			acats := r.Form["catIds"]

			var aformCats []int64
			for _, c := range acats {
				cid, _ := strconv.ParseInt(c, 10, 64)
				aformCats = append(aformCats, cid)
			}

			//adding new
			for _, c := range aformCats {
				h.Log.Debug("adding new cat to prodcat", c)
				var pc sdbi.ProductCategory
				pc.CategoryID = c
				pc.ProductID = prres.ID
				go func(pc *sdbi.ProductCategory, header *six910api.Headers) {
					h.API.AddProductCategory(pc, header)
				}(&pc, hd)
			}

			h.Log.Debug("prod add resp", *prres)
			if prres.Success {
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditProductPage StoreAdminEditProductPage
func (h *Six910Handler) StoreAdminEditProductPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			hd := h.getHeader(s)
			epvars := mux.Vars(r)
			idstr := epvars["id"]
			prodID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("prod id in edit", prodID)

			edErr := r.URL.Query().Get("error")
			var epparm ProdPage
			epparm.Error = edErr
			var wg sync.WaitGroup

			wg.Add(1)
			go func(id int64, header *six910api.Headers) {
				defer wg.Done()
				prod := h.API.GetProductByID(id, header)
				h.Log.Debug("prod  in edit", prod)
				epparm.Product = prod
			}(prodID, hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				cats := h.API.GetHierarchicalCategoryList(header)
				h.Log.Debug("prod  in edit", cats)
				epparm.CategoryList = cats
			}(hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				dist := h.API.GetDistributorList(header)
				h.Log.Debug("prod  in edit", dist)
				epparm.DistributorList = dist
			}(hd)

			wg.Add(1)
			go func(pid int64, header *six910api.Headers) {
				defer wg.Done()
				dist := h.API.GetProductCategoryList(pid, header)
				h.Log.Debug("prod category in edit", dist)
				epparm.ExistingCats = dist
			}(prodID, hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminEditProductPage, &epparm)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditProduct StoreAdminEditProduct
func (h *Six910Handler) StoreAdminEditProduct(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			epp := h.processProduct(r)
			h.Log.Debug("prod update", *epp)
			h.Log.Debug("image4: ", epp.Image4)
			r.ParseForm()
			cats := r.Form["catIds"]
			h.Log.Debug("cats in prod cat update", cats)
			hd := h.getHeader(s)

			var formCats []int64
			for _, c := range cats {
				cid, _ := strconv.ParseInt(c, 10, 64)
				formCats = append(formCats, cid)
			}

			exstCats := h.API.GetProductCategoryList(epp.ID, hd)

			//remove not used
			for _, c := range exstCats {
				var found = false
				for _, ec := range formCats {
					if c == ec {
						found = true
						break
					}
				}
				if !found {
					//remove
					h.Log.Debug("removing cat from prodcat", c)
					var pc sdbi.ProductCategory
					pc.CategoryID = c
					pc.ProductID = epp.ID
					go func(pc *sdbi.ProductCategory, header *six910api.Headers) {
						h.API.DeleteProductCategory(pc, header)
					}(&pc, hd)
				}
			}

			//adding new
			for _, c := range formCats {
				var found = false
				for _, ec := range exstCats {
					if c == ec {
						found = true
						break
					}
				}
				if !found {
					//add
					h.Log.Debug("adding new cat to prodcat", c)
					var pc sdbi.ProductCategory
					pc.CategoryID = c
					pc.ProductID = epp.ID
					go func(pc *sdbi.ProductCategory, header *six910api.Headers) {
						h.API.AddProductCategory(pc, header)
					}(&pc, hd)
				}
			}

			res := h.API.UpdateProduct(epp, hd)
			h.Log.Debug("prod update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminSearchProductBySkuPage StoreAdminSearchProductBySkuPage
func (h *Six910Handler) StoreAdminSearchProductBySkuPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			edErr := r.URL.Query().Get("error")
			var epparm ProdPage
			var plst []sdbi.Product
			sku := r.FormValue("sku")
			h.Log.Debug("sku", sku)
			if sku != "" {
				hd := h.getHeader(s)
				dist := h.API.GetDistributorList(hd)
				h.Log.Debug("dist", dist)
				for _, d := range *dist {
					prod := h.API.GetProductBySku(sku, d.ID, hd)
					if prod != nil && prod.ID != 0 {
						epparm.HasProducts = true
						plst = append(plst, *prod)
					}
				}
			}
			epparm.Products = &plst
			epparm.Error = edErr
			h.AdminTemplates.ExecuteTemplate(w, adminProductSkuSearchPage, &epparm)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminSearchProductByNamePage StoreAdminSearchProductByNamePage
func (h *Six910Handler) StoreAdminSearchProductByNamePage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			edErr := r.URL.Query().Get("error")
			var epparm ProdPage
			name := r.FormValue("name")
			h.Log.Debug("name", name)
			if name != "" {
				hd := h.getHeader(s)
				prods := h.API.GetProductsByName(name, 0, 100, hd)
				h.Log.Debug("prods by name", *prods)
				epparm.Products = prods
				if len(*prods) > 0 {
					epparm.HasProducts = true
				}
			} else {
				var plst []sdbi.Product
				epparm.Products = &plst
			}
			epparm.Error = edErr
			h.AdminTemplates.ExecuteTemplate(w, adminProductNameSearchPage, &epparm)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminSearchProductByCategoryPage StoreAdminSearchProductByCategoryPage
func (h *Six910Handler) StoreAdminSearchProductByCategoryPage(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod search by cat view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			ecdErr := r.URL.Query().Get("error")
			var cepparm ProdPage
			cidStr := r.FormValue("cid")
			h.Log.Debug("cidStr", cidStr)
			hd := h.getHeader(s)
			cats := h.API.GetCategoryList(hd)
			h.Log.Debug("prod  in edit", cats)
			cepparm.CategoryList = cats
			if cidStr != "" {
				cid, _ := strconv.ParseInt(cidStr, 10, 64)
				cepparm.LastCatSearch = cid
				prods := h.API.GetProductsByCaterory(cid, 0, 100, hd)
				h.Log.Debug("prods by name", *prods)
				cepparm.Products = prods
				if len(*prods) > 0 {
					cepparm.HasProducts = true
				}
			} else {
				var plst []sdbi.Product
				cepparm.Products = &plst
			}
			cepparm.Error = ecdErr
			h.AdminTemplates.ExecuteTemplate(w, adminProductCatSearchPage, &cepparm)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminViewProductList StoreAdminViewProductList
func (h *Six910Handler) StoreAdminViewProductList(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prods view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			hd := h.getHeader(s)
			vpvars := mux.Vars(r)
			ststr := vpvars["start"]
			endstr := vpvars["end"]
			vpstart, _ := strconv.ParseInt(ststr, 10, 64)
			vpend, _ := strconv.ParseInt(endstr, 10, 64)
			prods := h.API.GetProductList(vpstart, vpend, hd)
			plErr := r.URL.Query().Get("error")
			var plparm ProdPage
			plparm.Pagination = h.doPagination(vpstart, len(*prods), 100, "/admin/productList")
			h.Log.Debug("plparm.Pagination:", *plparm.Pagination)
			h.Log.Debug("plparm.Pagination.Pages:", *plparm.Pagination.Pages)
			plparm.Error = plErr
			plparm.Products = prods
			h.Log.Debug("prods  in edit", prods)
			h.AdminTemplates.ExecuteTemplate(w, adminProductListPage, &plparm)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminDeleteProduct StoreAdminDeleteProduct
func (h *Six910Handler) StoreAdminDeleteProduct(w http.ResponseWriter, r *http.Request) {
	s, suc := h.getSession(r)
	h.Log.Debug("session suc in prod list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(s) {
			hd := h.getHeader(s)
			dpvars := mux.Vars(r)
			idstrd := dpvars["id"]
			idd, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeleteProduct(idd, hd)
			h.Log.Debug("prod delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminProductList, http.StatusFound)
			} else {
				http.Redirect(w, r, adminProductListError, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processProduct(r *http.Request) *sdbi.Product {
	var p sdbi.Product
	id := r.FormValue("id")
	p.ID, _ = strconv.ParseInt(id, 10, 64)
	sku := r.FormValue("sku")
	p.Sku = sku
	p.Gtin = r.FormValue("gtin")
	p.Name = r.FormValue("name")
	p.ShortDesc = r.FormValue("shortDest")
	p.Desc = r.FormValue("desc")
	cost := r.FormValue("cost")
	p.Cost, _ = strconv.ParseFloat(cost, 64)
	msrp := r.FormValue("msrp")
	p.Msrp, _ = strconv.ParseFloat(msrp, 64)
	mapp := r.FormValue("map")
	p.Map, _ = strconv.ParseFloat(mapp, 64)
	price := r.FormValue("price")
	p.Price, _ = strconv.ParseFloat(price, 64)
	salePrice := r.FormValue("salePrice")
	p.SalePrice, _ = strconv.ParseFloat(salePrice, 64)
	p.Currency = r.FormValue("currency")
	p.ManufacturerID = r.FormValue("manfId")
	p.Manufacturer = r.FormValue("manf")
	stock := r.FormValue("stock")
	p.Stock, _ = strconv.ParseInt(stock, 10, 64)
	stockAlrt := r.FormValue("stockAlrt")
	p.StockAlert, _ = strconv.ParseInt(stockAlrt, 10, 64)
	weight := r.FormValue("weight")
	p.Weight, _ = strconv.ParseFloat(weight, 64)
	width := r.FormValue("width")
	p.Width, _ = strconv.ParseFloat(width, 64)
	height := r.FormValue("height")
	p.Height, _ = strconv.ParseFloat(height, 64)
	depth := r.FormValue("depth")
	p.Depth, _ = strconv.ParseFloat(depth, 64)
	shipMarkup := r.FormValue("shipMarkup")
	p.ShippingMarkup, _ = strconv.ParseFloat(shipMarkup, 64)

	visible := r.FormValue("visible")
	if visible == "on" {
		p.Visible = true
	} else {
		p.Visible = false
	}

	searchable := r.FormValue("searchable")
	if searchable == "on" {
		p.Searchable = true
	} else {
		p.Searchable = false
	}

	multibox := r.FormValue("multibox")
	if multibox == "on" {
		p.MultiBox = true
	} else {
		p.MultiBox = false
	}

	shipSep := r.FormValue("shipSep")
	if shipSep == "on" {
		p.ShipSeparately = true
	} else {
		p.ShipSeparately = false
	}

	freeShipping := r.FormValue("freeShipping")
	if freeShipping == "on" {
		p.FreeShipping = true
	} else {
		p.FreeShipping = false
	}

	promoted := r.FormValue("promoted")
	if promoted == "on" {
		p.Promoted = true
	} else {
		p.Promoted = false
	}

	dropship := r.FormValue("dropship")
	if dropship == "on" {
		p.Dropship = true
	} else {
		p.Dropship = false
	}

	specialProc := r.FormValue("specialProc")
	if specialProc == "on" {
		p.SpecialProcessing = true
	} else {
		p.SpecialProcessing = false
	}

	specialProcType := r.FormValue("specialProcType")
	p.SpecialProcessingType = specialProcType
	p.Size = r.FormValue("size")
	p.Color = r.FormValue("color")
	p.Gender = r.FormValue("gender")
	p.Thumbnail = r.FormValue("thumbnail")
	p.Image1 = r.FormValue("image1")
	p.Image2 = r.FormValue("image2")
	p.Image3 = r.FormValue("image3")
	p.Image4 = r.FormValue("image4")
	distributorID := r.FormValue("distributorId")
	p.DistributorID, _ = strconv.ParseInt(distributorID, 10, 64)
	storeID := r.FormValue("storeId")
	p.StoreID, _ = strconv.ParseInt(storeID, 10, 64)
	parentProductID := r.FormValue("parentProductId")
	p.ParentProductID, _ = strconv.ParseInt(parentProductID, 10, 64)

	return &p
}
