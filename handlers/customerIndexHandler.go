package handlers

import (
	"net/http"
	//"strconv"

	conts "github.com/Ulbora/Six910-ui/contentsrv"
	sdbi "github.com/Ulbora/six910-database-interface"
	//"github.com/gorilla/mux"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
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

//CustomerPage CustomerPage
type CustomerPage struct {
	//ProductListLift   *[]sdbi.Product
	//ProductListMiddle *[]sdbi.Product
	//ProductListRight  *[]sdbi.Product
	ProductListRowList *[]*ProductRow
	ProductList        *[]sdbi.Product
	Product            *sdbi.Product
	Content            *conts.Content
	MenuList           *[]musrv.Menu
	CategoryList       *[]sdbi.Category
}

//Index Index
func (h *Six910Handler) Index(w http.ResponseWriter, r *http.Request) {
	cis, suc := h.getSession(r)
	h.Log.Debug("session suc", suc)
	if suc {
		// civars := mux.Vars(r)
		// //ciststr := civars["start"]
		// ciendstr := civars["end"]
		// //cistart, _ := strconv.ParseInt(ciststr, 10, 64)
		// ciend, _ := strconv.ParseInt(ciendstr, 10, 64)
		// if ciend == 0 {
		// 	ciend = 100
		// }
		hd := h.getHeader(cis)
		ppl := h.API.GetProductsByPromoted(0, 100, hd)
		h.Log.Debug("promoted products", *ppl)
		//cisuc, cicont := h.ContentService.GetContent(indexContent)

		var cipage CustomerPage
		//var lp []sdbi.Product
		//var mp []sdbi.Product
		//var rp []sdbi.Product
		//cipage.ProductList = ppl
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
		//cipage.ProductListLift = &lp
		//cipage.ProductListMiddle = &mp
		cipage.ProductListRowList = &prowList

		h.Log.Debug("prowList", prowList)

		cipage.MenuList = h.MenuService.GetMenuList()
		h.Log.Debug("MenuList", *cipage.MenuList)
		_, cont := h.ContentService.GetContent("home")
		cipage.Content = cont
		// if cisuc {
		// 	cipage.Content = cicont
		// }
		h.Log.Debug("cipage: ", cipage)
		h.Templates.ExecuteTemplate(w, customerIndexPage, &cipage)
	}
}
