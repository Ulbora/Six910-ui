package handlers

import (
	"container/list"
	"net/http"
	"strconv"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	api "github.com/Ulbora/Six910API-Go"
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

//ViewProductList ViewProductList
func (h *Six910Handler) ViewProductList(w http.ResponseWriter, r *http.Request) {
	cpls, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		cplvars := mux.Vars(r)
		cplcatidstr := cplvars["catId"]
		cplststr := cplvars["start"]
		cplendstr := cplvars["end"]
		cplcatid, _ := strconv.ParseInt(cplcatidstr, 10, 64)
		cplstart, _ := strconv.ParseInt(cplststr, 10, 64)
		cplend, _ := strconv.ParseInt(cplendstr, 10, 64)
		if cplend == 0 {
			cplend = 100
		}
		h.Log.Debug("cplcatid: ", cplcatid)
		hd := h.getHeader(cpls)
		ppl := h.API.GetProductsByCaterory(cplcatid, cplstart, cplend, hd)
		cisuc, cicont := h.ContentService.GetContent(productListContent)

		var cplpage CustomerPage
		cplpage.ProductList = ppl
		if cisuc {
			cplpage.Content = cicont
		}
		h.Log.Debug("cplpage: ", cplpage)
		h.Templates.ExecuteTemplate(w, customerProductListPage, &cplpage)
	}
}

//ViewProductByCatList ViewProductByCatList
func (h *Six910Handler) ViewProductByCatList(w http.ResponseWriter, r *http.Request) {
	cpls, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		cplvars := mux.Vars(r)
		ccplcatidstr := cplvars["catId"]
		catName := cplvars["catName"]
		ccplststr := cplvars["start"]
		ccplendstr := cplvars["end"]
		cplcatid, _ := strconv.ParseInt(ccplcatidstr, 10, 64)
		cplstart, _ := strconv.ParseInt(ccplststr, 10, 64)
		cplend, _ := strconv.ParseInt(ccplendstr, 10, 64)
		if cplend == 0 {
			cplend = 100
		}
		h.Log.Debug("cplcatid: ", cplcatid)
		hd := h.getHeader(cpls)
		ppl := h.API.GetProductsByCaterory(cplcatid, cplstart, cplend, hd)
		cisuc, cicont := h.ContentService.GetContent(productCategoryListContent)

		var cplpage CustomerPage
		cplpage.ProductList = ppl
		if cisuc {
			cplpage.Content = cicont
		} else {
			var ct conts.Content
			cplpage.Content = &ct
		}
		cplpage.CategoryName = catName

		cplpage.MenuList = h.MenuService.GetMenuList()
		h.Log.Debug("MenuList", *cplpage.MenuList)

		var prowListc []*ProductRow
		var prowc *ProductRow
		var rc = 1
		for i, p := range *ppl {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				prowc = new(ProductRow)
				prowc.ProductLeft = p
				rc++
				if i == len(*ppl)-1 {
					prowListc = append(prowListc, prowc)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				prowc.ProductMiddle = p
				rc++
				if i == len(*ppl)-1 {
					prowListc = append(prowListc, prowc)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				prowc.ProductRight = p
				h.Log.Debug("prow", prowc)
				prowListc = append(prowListc, prowc)
				rc = 1
			}
		}

		cplpage.ProductListRowList = &prowListc

		h.Log.Debug("prowList", prowListc)
		cplpage.Pagination = h.doPagination(cplstart, len(*ppl), 100, "/productByCategoryList/"+ccplcatidstr+"/"+catName)
		h.Log.Debug("plparm.Pagination:", *cplpage.Pagination)
		h.Log.Debug("plparm.Pagination.Pages:", *cplpage.Pagination.Pages)
		h.Log.Debug("cplpage: ", cplpage)
		h.Templates.ExecuteTemplate(w, customerProductByCatPage, &cplpage)
	}
}

//SearchProductList SearchProductList
func (h *Six910Handler) SearchProductList(w http.ResponseWriter, r *http.Request) {
	cspls, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {

		csplsearch := r.FormValue("search")
		var csplstart int64
		var csplend int64

		if csplsearch == "" {
			csplvars := mux.Vars(r)
			csplsearch = csplvars["search"]
			csplststr := csplvars["start"]
			csplendstr := csplvars["end"]

			csplstart, _ = strconv.ParseInt(csplststr, 10, 64)
			csplend, _ = strconv.ParseInt(csplendstr, 10, 64)
		}

		if csplend == 0 {
			csplend = 100
		}
		h.Log.Debug("csplsearch: ", csplsearch)
		hd := h.getHeader(cspls)
		ppl := h.API.GetProductsByName(csplsearch, csplstart, csplend, hd)

		cisuc, cscont := h.ContentService.GetContent(productListContent)

		var csplpage CustomerPage
		csplpage.ProductList = ppl
		if cisuc {
			csplpage.Content = cscont
		} else {
			var ct conts.Content
			csplpage.Content = &ct
		}
		csplpage.SearchName = csplsearch

		csplpage.MenuList = h.MenuService.GetMenuList()
		h.Log.Debug("MenuList", *csplpage.MenuList)

		var sprowListc []*ProductRow
		var sprowc *ProductRow
		var rc = 1
		for i, p := range *ppl {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				sprowc = new(ProductRow)
				sprowc.ProductLeft = p
				rc++
				if i == len(*ppl)-1 {
					sprowListc = append(sprowListc, sprowc)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				sprowc.ProductMiddle = p
				rc++
				if i == len(*ppl)-1 {
					sprowListc = append(sprowListc, sprowc)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				sprowc.ProductRight = p
				h.Log.Debug("prow", sprowc)
				sprowListc = append(sprowListc, sprowc)
				rc = 1
			}
		}

		csplpage.ProductListRowList = &sprowListc

		h.Log.Debug("prowList", sprowListc)
		csplpage.Pagination = h.doPagination(csplstart, len(*ppl), 100, "/searchProductsByName/"+csplsearch)
		h.Log.Debug("plparm.Pagination:", *csplpage.Pagination)
		h.Log.Debug("plparm.Pagination.Pages:", *csplpage.Pagination.Pages)
		h.Log.Debug("csplpage: ", csplpage)
		h.Templates.ExecuteTemplate(w, customerProductsSearchListPage, &csplpage)
	}
}

//ViewProduct ViewProduct
func (h *Six910Handler) ViewProduct(w http.ResponseWriter, r *http.Request) {
	cvps, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		cvpvars := mux.Vars(r)
		cpidstr := cvpvars["id"]
		cplcatid, _ := strconv.ParseInt(cpidstr, 10, 64)

		h.Log.Debug("cplcatid: ", cplcatid)
		hd := h.getHeader(cvps)
		pp := h.API.GetProductByID(cplcatid, hd)

		prodCat := h.API.GetProductCategoryList(cplcatid, hd)
		var catList []sdbi.Category
		for _, pcc := range prodCat {
			h.Log.Debug("pcc: ", pcc)
			l := list.New()
			h.getProductCatList(l, pcc, hd)
			h.Log.Debug("l len: ", l.Len())
			h.Log.Debug("l vals: ", *l)
			for e := l.Front(); e != nil; e = e.Next() {

				h.Log.Debug("l e vals: ", *e.Value.(*sdbi.Category))
				catList = append(catList, *e.Value.(*sdbi.Category))
			}
			break

		}

		cisuc, cicont := h.ContentService.GetContent(productContent)

		var cplpage CustomerPage
		cplpage.CategoryList = &catList
		cplpage.MenuList = h.MenuService.GetMenuList()
		h.Log.Debug("MenuList", *cplpage.MenuList)
		_, cont := h.ContentService.GetContent("product")
		cplpage.Content = cont

		cplpage.Product = pp
		if cisuc {
			cplpage.Content = cicont
		}
		h.Log.Debug("cplpage: ", cplpage)
		h.Templates.ExecuteTemplate(w, customerProductPage, &cplpage)
	}
}

func (h *Six910Handler) getProductCatList(l *list.List, cid int64, hd *api.Headers) int64 {
	if cid == 0 {
		return 0
	}
	pc := h.API.GetCategory(cid, hd)
	h.Log.Debug("pc: ", *pc)
	l.PushFront(pc)
	return h.getProductCatList(l, pc.ParentCategoryID, hd)

}
