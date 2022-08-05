package handlers

import (
	"container/list"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	cpls, suc := h.getUserSession(w, r)
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
		} else {
			var ct conts.Content
			cplpage.Content = &ct
		}
		h.Log.Debug("cplpage: ", cplpage)
		h.Templates.ExecuteTemplate(w, customerProductListPage, &cplpage)
	}
}

//ViewProductByCatList ViewProductByCatList
func (h *Six910Handler) ViewProductByCatList(w http.ResponseWriter, r *http.Request) {
	cpls, suc := h.getUserSession(w, r)
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

		//make call to get manufact by cat
		mlst := h.API.GetProductManufacturerListByCatID(cplcatid, hd)

		var cplpage CustomerPage
		var turl = "/productByCategoryList/" + ccplcatidstr + "/" + catName + "/0/100"
		cplpage.HeaderData = h.processMetaData(turl, catName, r)

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		h.Log.Debug("PageBody: ", *csspg)
		cplpage.PageBody = csspg

		cplpage.ProductList = ppl
		cplpage.ManufacturerList = mlst
		if cisuc {
			cplpage.Content = cicont
		} else {
			var ct conts.Content
			cplpage.Content = &ct
		}
		cplpage.CategoryName = catName
		cplpage.CategoryID = ccplcatidstr

		//cplpage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(cpls, ml, hd)
		cplpage.MenuList = ml
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

//ViewProductByCatAndManufacturerList ViewProductByCatAndManufacturerList
func (h *Six910Handler) ViewProductByCatAndManufacturerList(w http.ResponseWriter, r *http.Request) {
	mcpls, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		mcplvars := mux.Vars(r)
		mccplcatidstr := mcplvars["catId"]
		mcatName := mcplvars["catName"]
		manf := mcplvars["manf"]
		mccplststr := mcplvars["start"]
		mccplendstr := mcplvars["end"]
		mcplcatid, _ := strconv.ParseInt(mccplcatidstr, 10, 64)
		mcplstart, _ := strconv.ParseInt(mccplststr, 10, 64)
		mcplend, _ := strconv.ParseInt(mccplendstr, 10, 64)
		if mcplend == 0 {
			mcplend = 100
		}
		h.Log.Debug("cplcatid: ", mcplcatid)
		hd := h.getHeader(mcpls)
		ppl := h.API.GetProductByCatAndManufacturer(mcplcatid, manf, mcplstart, mcplend, hd)
		cisuc, cicont := h.ContentService.GetContent(productCategoryListContent)

		//make call to get manufact by cat
		mlst := h.API.GetProductManufacturerListByCatID(mcplcatid, hd)

		var mcplpage CustomerPage
		var turl = "/productByCategoryAndManufacturerList/" + mccplcatidstr + "/" + mcatName + "/" + manf + "/0/100"
		mcplpage.HeaderData = h.processMetaData(turl, mcatName, r)

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		h.Log.Debug("PageBody: ", *csspg)
		mcplpage.PageBody = csspg

		mcplpage.ProductList = ppl
		mcplpage.ManufacturerList = mlst
		if cisuc {
			mcplpage.Content = cicont
		} else {
			var ct conts.Content
			mcplpage.Content = &ct
		}
		mcplpage.CategoryName = mcatName
		mcplpage.CategoryID = mccplcatidstr
		mcplpage.Manufacturer = manf

		//mcplpage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(mcpls, ml, hd)
		mcplpage.MenuList = ml
		h.Log.Debug("MenuList", *mcplpage.MenuList)

		var mprowListc []*ProductRow
		var mprowc *ProductRow
		var rc = 1
		for i, p := range *ppl {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				mprowc = new(ProductRow)
				mprowc.ProductLeft = p
				rc++
				if i == len(*ppl)-1 {
					mprowListc = append(mprowListc, mprowc)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				mprowc.ProductMiddle = p
				rc++
				if i == len(*ppl)-1 {
					mprowListc = append(mprowListc, mprowc)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				mprowc.ProductRight = p
				h.Log.Debug("prow", mprowc)
				mprowListc = append(mprowListc, mprowc)
				rc = 1
			}
		}

		mcplpage.ProductListRowList = &mprowListc

		h.Log.Debug("prowList", mprowListc)
		mcplpage.Pagination = h.doPagination(mcplstart, len(*ppl), 100, "/productByCategoryAndManufacturerList/"+mccplcatidstr+"/"+mcatName+"/"+manf)
		h.Log.Debug("plparm.Pagination:", *mcplpage.Pagination)
		h.Log.Debug("plparm.Pagination.Pages:", *mcplpage.Pagination.Pages)
		h.Log.Debug("cplpage: ", mcplpage)
		h.Templates.ExecuteTemplate(w, customerProductByCatPage, &mcplpage)
	}
}

//SearchProductList SearchProductList
func (h *Six910Handler) SearchProductList(w http.ResponseWriter, r *http.Request) {
	cspls, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		//var pagebdy PageBody
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

		mlst := h.API.GetProductManufacturerListByProductName(csplsearch, hd)
		h.Log.Debug("mlst: ", mlst)

		//make call to get manufact by name of product

		cisuc, cscont := h.ContentService.GetContent(productListContent)

		var csplpage CustomerPage
		var turl = "/searchProductsByName/" + csplsearch + "/0/100"
		csplpage.HeaderData = h.processMetaData(turl, csplsearch, r)

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		csplpage.PageBody = csspg

		csplpage.ProductList = ppl
		csplpage.ManufacturerList = mlst
		if cisuc {
			csplpage.Content = cscont
		} else {
			var ct conts.Content
			csplpage.Content = &ct
		}
		csplpage.SearchName = csplsearch
		csplpage.Manufacturer = ""

		//csplpage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(cspls, ml, hd)
		csplpage.MenuList = ml
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

//SearchProductByManufacturerList SearchProductByManufacturerList
func (h *Six910Handler) SearchProductByManufacturerList(w http.ResponseWriter, r *http.Request) {
	mcspls, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {

		mcsplsearch := r.FormValue("search")
		var manf string
		var csplstart int64
		var csplend int64

		if mcsplsearch == "" {
			mcsplvars := mux.Vars(r)
			manf = mcsplvars["manf"]
			mcsplsearch = mcsplvars["search"]
			mcsplststr := mcsplvars["start"]
			mcsplendstr := mcsplvars["end"]

			csplstart, _ = strconv.ParseInt(mcsplststr, 10, 64)
			csplend, _ = strconv.ParseInt(mcsplendstr, 10, 64)
		}

		if csplend == 0 {
			csplend = 100
		}
		h.Log.Debug("csplsearch: ", mcsplsearch)
		hd := h.getHeader(mcspls)
		ppl := h.API.GetProductByNameAndManufacturerName(manf, mcsplsearch, csplstart, csplend, hd)

		mlst := h.API.GetProductManufacturerListByProductName(mcsplsearch, hd)
		h.Log.Debug("mlst: ", mlst)

		//make call to get manufact by name of product

		cisuc, cscont := h.ContentService.GetContent(productListContent)

		var mcsplpage CustomerPage
		var turl = "/searchProductsByManufacturerAndName/" + manf + "/" + mcsplsearch + "/0/100"
		mcsplpage.HeaderData = h.processMetaData(turl, mcsplsearch, r)

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		mcsplpage.PageBody = csspg

		mcsplpage.ProductList = ppl
		mcsplpage.ManufacturerList = mlst
		if cisuc {
			mcsplpage.Content = cscont
		} else {
			var ct conts.Content
			mcsplpage.Content = &ct
		}
		mcsplpage.SearchName = mcsplsearch
		mcsplpage.Manufacturer = manf

		//mcsplpage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(mcspls, ml, hd)
		mcsplpage.MenuList = ml
		h.Log.Debug("MenuList", *mcsplpage.MenuList)

		var msprowListc []*ProductRow
		var sprowc *ProductRow
		var rc = 1
		for i, p := range *ppl {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				sprowc = new(ProductRow)
				sprowc.ProductLeft = p
				rc++
				if i == len(*ppl)-1 {
					msprowListc = append(msprowListc, sprowc)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				sprowc.ProductMiddle = p
				rc++
				if i == len(*ppl)-1 {
					msprowListc = append(msprowListc, sprowc)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				sprowc.ProductRight = p
				h.Log.Debug("prow", sprowc)
				msprowListc = append(msprowListc, sprowc)
				rc = 1
			}
		}

		mcsplpage.ProductListRowList = &msprowListc

		h.Log.Debug("prowList", msprowListc)
		mcsplpage.Pagination = h.doPagination(csplstart, len(*ppl), 100, "/searchProductsByManufacturerAndName/"+manf+"/"+mcsplsearch)
		h.Log.Debug("plparm.Pagination:", *mcsplpage.Pagination)
		h.Log.Debug("plparm.Pagination.Pages:", *mcsplpage.Pagination.Pages)
		h.Log.Debug("csplpage: ", mcsplpage)
		h.Templates.ExecuteTemplate(w, customerProductsSearchListPage, &mcsplpage)
	}
}

//ProductSearchByDescAttributes ProductSearchByDescAttributes
func (h *Six910Handler) ProductSearchByDescAttributes(w http.ResponseWriter, r *http.Request) {
	cspls, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		//var pagebdy PageBody
		var color string
		var size string
		var gender string

		csrplsearch := r.FormValue("search")
		color = r.FormValue("color")
		size = r.FormValue("size")
		gender = r.FormValue("gender")

		var acsplstart int64
		var acsplend int64

		fmt.Println("Search: ", csrplsearch)
		fmt.Println("Gender: ", gender)

		if csrplsearch == "" {
			acsplvars := mux.Vars(r)
			csrplsearch = acsplvars["search"]
			acsplststr := acsplvars["start"]
			acsplendstr := acsplvars["end"]

			// color = acsplvars["color"]
			// size = acsplvars["size"]
			// gender = acsplvars["gender"]

			//add addition params for gender color and size
			//and do productSearch to temp product list and remove products that
			//don't meet new parameters like : male size 12, color black
			//assign new product list to sppl

			acsplstart, _ = strconv.ParseInt(acsplststr, 10, 64)
			acsplend, _ = strconv.ParseInt(acsplendstr, 10, 64)
		}

		var attrbs = strings.Split(csrplsearch, " ")

		if acsplend == 0 {
			acsplend = 100
		}
		h.Log.Debug("csrplsearch: ", csrplsearch)
		hd := h.getHeader(cspls)

		var psratt sdbi.ProductSearch
		psratt.DescAttributes = &attrbs
		psratt.End = acsplend

		var sppl *[]sdbi.Product

		tempsppl := h.API.ProductSearch(&psratt, hd)

		sppl = h.filterProduct(color, size, gender, tempsppl)

		h.Log.Debug("sppl: ", sppl)

		// may need to modify this search too

		smlst := h.API.GetProductManufacturerListByProductSearch(csrplsearch, hd)
		h.Log.Debug("smlst: ", smlst)

		//make call to get manufact by name of product

		cisuc, cscont := h.ContentService.GetContent(productListContent)

		var acsplpage CustomerPage
		var turl = "/searchProductsByDesc/" + csrplsearch + "/0/100"
		acsplpage.HeaderData = h.processMetaData(turl, csrplsearch, r)

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		acsplpage.PageBody = csspg

		acsplpage.ProductList = sppl
		acsplpage.ManufacturerList = smlst
		if cisuc {
			acsplpage.Content = cscont
		} else {
			var ct conts.Content
			acsplpage.Content = &ct
		}
		acsplpage.SearchName = csrplsearch
		acsplpage.Color = color
		acsplpage.Size = size
		acsplpage.Gender = gender
		acsplpage.Manufacturer = ""

		//csplpage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(cspls, ml, hd)
		acsplpage.MenuList = ml
		h.Log.Debug("MenuList", *acsplpage.MenuList)

		var asprowListc []*ProductRow
		var sprowc *ProductRow
		var rc = 1
		for i, p := range *sppl {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				sprowc = new(ProductRow)
				sprowc.ProductLeft = p
				rc++
				if i == len(*sppl)-1 {
					asprowListc = append(asprowListc, sprowc)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				sprowc.ProductMiddle = p
				rc++
				if i == len(*sppl)-1 {
					asprowListc = append(asprowListc, sprowc)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				sprowc.ProductRight = p
				h.Log.Debug("prow", sprowc)
				asprowListc = append(asprowListc, sprowc)
				rc = 1
			}
		}

		acsplpage.ProductListRowList = &asprowListc

		h.Log.Debug("prowList", asprowListc)
		acsplpage.Pagination = h.doPagination(acsplstart, len(*sppl), 100, "/searchProductsByDesc/"+csrplsearch)
		h.Log.Debug("plparm.Pagination:", *acsplpage.Pagination)
		h.Log.Debug("plparm.Pagination.Pages:", *acsplpage.Pagination.Pages)
		h.Log.Debug("acsplpage: ", acsplpage)
		h.Templates.ExecuteTemplate(w, customerProductsSearchListPage, &acsplpage)
	}
}

//ViewProduct ViewProduct
func (h *Six910Handler) ViewProduct(w http.ResponseWriter, r *http.Request) {
	cvps, suc := h.getUserSession(w, r)
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

		var likeProdCatID int64
		if len(catList) > 0 {
			likeProdCatID = catList[len(catList)-1].ID
		}

		likeProdList := h.API.GetProductsByCaterory(likeProdCatID, 0, 9, hd)

		cisuc, cicont := h.ContentService.GetContent(productContent)

		var cplpage CustomerPage
		cplpage.HeaderData = h.processProductMetaData(pp, r)
		cplpage.ProductList = likeProdList

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		h.Log.Debug("PageBody: ", *csspg)
		cplpage.PageBody = csspg

		cplpage.CategoryList = &catList

		//cplpage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(cvps, ml, hd)
		cplpage.MenuList = ml
		h.Log.Debug("MenuList", *cplpage.MenuList)
		//_, cont := h.ContentService.GetContent("product")
		//cplpage.Content = cont

		cplpage.Product = pp
		if cisuc {
			cplpage.Content = cicont
		} else {
			var ct conts.Content
			cplpage.Content = &ct
		}

		var pmsprowListc []*ProductRow
		var psprowc *ProductRow
		var rc = 1
		for i, p := range *likeProdList {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				psprowc = new(ProductRow)
				psprowc.ProductLeft = p
				rc++
				if i == len(*likeProdList)-1 {
					pmsprowListc = append(pmsprowListc, psprowc)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				psprowc.ProductMiddle = p
				rc++
				if i == len(*likeProdList)-1 {
					pmsprowListc = append(pmsprowListc, psprowc)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				psprowc.ProductRight = p
				h.Log.Debug("prow", psprowc)
				pmsprowListc = append(pmsprowListc, psprowc)
				rc = 1
			}
		}

		cplpage.ProductListRowList = &pmsprowListc

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

func (h *Six910Handler) filterProduct(color string, size string, gender string, plst *[]sdbi.Product) *[]sdbi.Product {
	var rtn []sdbi.Product

	for _, p := range *plst {
		if color != "" && strings.ToLower(p.Color) != strings.ToLower(color) {
			continue
		}
		if size != "" && strings.ToLower(p.Size) != strings.ToLower(size) {
			continue
		}
		if gender != "" && strings.ToLower(p.Gender) != strings.ToLower(gender) {
			continue
		}
		rtn = append(rtn, p)
	}
	return &rtn
}
