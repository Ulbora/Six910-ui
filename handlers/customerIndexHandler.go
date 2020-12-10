package handlers

import (
	"os"
	"time"
	//"html/template"
	"net/http"
	//"strconv"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	sdbi "github.com/Ulbora/six910-database-interface"
	//"github.com/gorilla/mux"
	carsrv "github.com/Ulbora/Six910-ui/carouselsrv"
	cntrysrv "github.com/Ulbora/Six910-ui/countrysrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	stsrv "github.com/Ulbora/Six910-ui/statesrv"
	six910api "github.com/Ulbora/Six910API-Go"
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

//ProductRow ProductRow
type ProductRow struct {
	ProductLeft   sdbi.Product
	ProductMiddle sdbi.Product
	ProductRight  sdbi.Product
}

// //PageBody PageBody
// type PageBody struct {
// 	Background template.CSS
// 	Color      template.CSS
// 	PageTitle  template.CSS
// }

//CustomerPage CustomerPage
type CustomerPage struct {
	ProductListRowList *[]*ProductRow
	ProductList        *[]sdbi.Product
	Product            *sdbi.Product
	Content            *conts.Content
	MenuList           *[]musrv.Menu
	CategoryList       *[]sdbi.Category
	Pagination         *Pagination
	CategoryID         string
	CategoryName       string
	SearchName         string
	ManufacturerList   *[]string
	Manufacturer       string
	PageBody           *csssrv.PageCSS
	Carousel           *carsrv.Carousel
	Customer           *sdbi.Customer
	AddressList        *[]sdbi.Address
	StateList          *[]stsrv.State
	CountryList        *[]cntrysrv.Country

	//meta data
	HeaderData *HeaderData
	Error      string
}

//Index Index
func (h *Six910Handler) Index(w http.ResponseWriter, r *http.Request) {
	cis, suc := h.getUserSession(w, r)
	h.Log.Debug("session suc", suc)
	if suc {
		//var pagebdy PageBody
		// pagebdy.Background = "background: grey !important;"
		// pagebdy.Color = "" //"color: white;"
		hd := h.getHeader(cis)

		origin := r.Header.Get("Origin")
		host := r.Host
		h.Log.Debug("X-Forwarded-For :" + r.Header.Get("X-FORWARDED-FOR"))
		//h.Log.Debug("headers", headers)
		//h.Log.Debug("request", *r)
		h.Log.Debug("origin", origin)
		h.Log.Debug("host", host)
		h.Log.Debug("scheme", r.URL.Scheme)
		//h.Log.Debug("request", *r)
		var v sdbi.Visitor
		v.Origin = r.Header.Get("Origin")
		v.Host = r.Host
		v.IPAddress = r.Header.Get("X-FORWARDED-FOR")
		today := time.Now()
		//ptwo := today.Add(2 * time.Hour)
		lastHit := h.getLastHit(cis, w, r)
		h.Log.Debug("lastHit", lastHit)
		ptwo := lastHit.Add(2 * time.Hour)
		if today.After(ptwo) {
			h.Log.Debug("add new hit")
			h.setLastHit(cis, w, r)
			go func(vis *sdbi.Visitor, header *six910api.Headers) {
				h.API.AddVisit(vis, header)
			}(&v, hd)
		}

		ppl := h.API.GetProductsByPromoted(0, 100, hd)
		h.Log.Debug("promoted products", *ppl)

		var cipage CustomerPage

		_, csspg := h.CSSService.GetPageCSS("pageCss")
		h.Log.Debug("PageBody: ", *csspg)
		cipage.PageBody = csspg

		_, carpg := h.CarouselService.GetCarousel("carousel")
		h.Log.Debug("Carousel: ", *carpg)
		cipage.Carousel = carpg

		var prowList []*ProductRow
		var prow *ProductRow
		var rc = 1
		for i, p := range *ppl {
			if rc == 1 {
				h.Log.Debug("sku1", p.Sku)
				prow = new(ProductRow)
				prow.ProductLeft = p
				rc++
				if i == len(*ppl)-1 {
					prowList = append(prowList, prow)
				}
				continue
			} else if rc == 2 {
				h.Log.Debug("sku2", p.Sku)
				prow.ProductMiddle = p
				rc++
				if i == len(*ppl)-1 {
					prowList = append(prowList, prow)
				}
				continue
			} else if rc == 3 {
				h.Log.Debug("sku3", p.Sku)
				prow.ProductRight = p
				h.Log.Debug("prow", prow)
				prowList = append(prowList, prow)
				rc = 1
			}
		}

		cipage.ProductListRowList = &prowList

		h.Log.Debug("prowList", prowList)

		//cipage.MenuList = h.MenuService.GetMenuList()
		ml := h.MenuService.GetMenuList()
		h.getCartTotal(cis, ml, hd)
		cipage.MenuList = ml
		h.Log.Debug("MenuList", *cipage.MenuList)
		cisuc, cont := h.ContentService.GetContent(indexContent)
		if cisuc {
			cipage.Content = cont
		} else {
			var ct conts.Content
			cipage.Content = &ct
		}
		//h.ContentService.HitCheck()
		//headers := r.Header

		//site map
		smtime := h.SiteMapDate
		h.Log.Debug("smtime: ", smtime)
		if today.After(smtime) {
			h.Log.Debug("Saving new Site Map")
			smtime = today
			smtime = smtime.Add(720 * time.Hour)
			h.SiteMapDate = smtime
			idlst := h.API.GetProductIDList(hd)
			path := "./static"
			// h.ActiveTemplateLocation + "/" + h.ActiveTemplateName
			h.saveSiteMap(idlst, path)
		}

		h.Log.Debug("cipage: ", cipage)
		h.Templates.ExecuteTemplate(w, customerIndexPage, &cipage)
	}
}

func (h *Six910Handler) saveSiteMap(ids *[]int64, path string) {
	var vp SiteMapValues
	vp.Domain = h.getSiteMapDomain()
	vp.ProductIDList = ids
	smbs := h.generateSiteMap(&vp)
	//path := h.ActiveTemplateLocation + "/" + h.ActiveTemplateName
	f, err := os.Create(path + "/sitemap.xml")
	h.Log.Debug("os.Create err: ", err)
	defer f.Close()
	_, err2 := f.Write(smbs)
	h.Log.Debug("f.Write err: ", err2)
}
