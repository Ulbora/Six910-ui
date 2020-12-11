package managers

import (
	"strconv"
	"strings"

	frr "github.com/Ulbora/FileReader"
	api "github.com/Ulbora/Six910API-Go"

	//api "github.com/Ulbora/Six910API-Go"
	sdbi "github.com/Ulbora/six910-database-interface"
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

// -------File attributes used---------------------------------------
// Distributor           string //required
// Sku                   string //required
// Gtin                  string
// Name                  string  //required
// ShortDesc             string  //required
// Desc                  string  //required
// Cost                  float64 //required
// Msrp                  float64 //required
// Map                   float64
// Price                 float64 //required
// SalePrice             float64
// ManufacturerID        string
// Manufacturer          string
// Stock                 int64
// StockAlert            int64
// Weight                float64
// Width                 float64
// Height                float64
// Depth                 float64
// ShippingMarkup        float64
// Visible               bool
// Searchable            bool
// MultiBox              bool
// ShipSeparately        bool
// FreeShipping          bool
// Promoted              bool
// Dropship              bool
// SpecialProcessing     bool
// SpecialProcessingType string
// Size                  string
// Color                 string
// Thumbnail             string //required
// Image1                string //required
// Image2                string
// Image3                string
// Image4                string
// ParentProductName     string
// ParentProductSku      string
// Category              string //required cat/cat2/cat3/cat4

//UploadProductFile UploadProductFile
func (m *Six910Manager) UploadProductFile(file []byte, hd *api.Headers) (productsImported int, productNotImported int) {
	//var rtn bool
	var cr frr.CsvFileReader
	//var rprep fpp.RecordPrep
	var distName string
	rd := cr.GetNew()
	rec := rd.ReadCsvFile(file)
	if rec.CsvReadErr == nil && len(rec.CsvFileList) > 1 && len(rec.CsvFileList[0]) > 1 {
		if len(rec.CsvFileList) > 1 && len(rec.CsvFileList[0]) > 0 {
			distName = rec.CsvFileList[1][0]
		}
		m.Log.Debug("distributor: ", distName)

		m.Log.Debug("recs: ", *&rec.CsvFileList)

		var columns = rec.CsvFileList[0]

		var colMap = make(map[string]int)
		for i, v := range columns {
			colMap[v] = i
		}
		m.Log.Debug("colMap: ", colMap)

		var disID int64
		dlst := m.API.GetDistributorList(hd)
		for _, d := range *dlst {
			if d.Company == distName {
				disID = d.ID
				break
			}
		}
		if disID == 0 {
			var dist sdbi.Distributor
			dist.Company = distName
			dres := m.API.AddDistributor(&dist, hd)
			if dres.Success && dres.ID != 0 {
				disID = dres.ID
			}
		}

		prods, holdProds := m.prepProducts(disID, &rec.CsvFileList, hd)
		m.Log.Debug("prods: ", prods)
		m.Log.Debug("holdProds: ", holdProds)
		productNotImported = len(*holdProds)
		productsImported = m.importProducts(prods, hd)
	}

	return productsImported, productNotImported
}

func (m *Six910Manager) prepProducts(distributorID int64, csvRecs *[][]string, hd *api.Headers) (productList *[]Product, holdProductList *[]Product) {
	var prodList []Product
	var holdProdList []Product
	var col []string

	//var cnt = 0
	for i, row := range *csvRecs {
		if i == 0 {
			col = row
		} else {
			// cnt++
			// if cnt > 10 {
			// 	m.Log.Debug("sleeping :", 2000*time.Millisecond)
			// 	cnt = 0
			// 	m.Log.Debug("cnt :", cnt)
			// 	time.Sleep(2000 * time.Millisecond)
			// }
			var p Product
			p.DistributorID = distributorID
			var complete bool
			var partialUpload = true
			var holdParentSkuList []string
			for i, v := range row {
				var vname = col[i]
				m.Log.Debug("vname: ", vname)
				m.Log.Debug("val: ", v)
				if vname == "desc" || vname == "parent_product_sku" {
					partialUpload = false
				}
				complete = m.processValue(vname, v, &p, hd)
				if !complete && vname == "parent_product_sku" {
					holdParentSkuList = append(holdParentSkuList, v)
				}
				m.Log.Debug("complete: ", complete)
			}
			//m.Log.Debug("partialUpload: ", partialUpload)
			m.Log.Debug("product: ", p)
			if complete || partialUpload {
				prodList = append(prodList, p)
			} else {
				holdProdList = append(holdProdList, p)
			}
			holdProdList = *m.processHoldList(&prodList, &holdParentSkuList, &holdProdList, hd)

		}
	}

	return &prodList, &holdProdList
}

func (m *Six910Manager) processHoldList(prodList *[]Product, holdParentSkuList *[]string, holdProdList *[]Product, hd *api.Headers) *[]Product {
	var rtnHoldList *[]Product
	if len(*holdParentSkuList) == len(*holdProdList) && len(*holdParentSkuList) > 0 {
		var hlst []Product
		for i, sku := range *holdParentSkuList {
			p := (*holdProdList)[i]
			pp := m.API.GetProductBySku(sku, p.DistributorID, hd)
			if pp.ID != 0 {
				p.ParentProductID = pp.ID
				*prodList = append(*prodList, p)
			} else {
				hlst = append(hlst, p)
			}
		}
		rtnHoldList = &hlst
	} else {
		rtnHoldList = holdProdList
	}
	return rtnHoldList
}

func (m *Six910Manager) processValue(name string, val string, p *Product, hd *api.Headers) bool {
	var complete = false
	var parentProductFound = true
	var categoryFound = true
	switch name {
	case "sku":
		p.Sku = val
	case "gtin":
		p.Gtin = val
	case "name":
		p.Name = val
	case "short_desc":
		p.ShortDesc = val
	case "desc":
		p.Desc = val
	case "cost":
		cost, _ := strconv.ParseFloat(val, 64)
		p.Cost = cost
	case "msrp":
		msrp, _ := strconv.ParseFloat(val, 64)
		p.Msrp = msrp
	case "map":
		mapv, _ := strconv.ParseFloat(val, 64)
		p.Map = mapv
	case "price":
		price, _ := strconv.ParseFloat(val, 64)
		p.Price = price
	case "sale_price":
		salePrice, _ := strconv.ParseFloat(val, 64)
		p.SalePrice = salePrice
	case "currency":
		p.Currency = val
	case "manufacturer_id":
		p.ManufacturerID = val
	case "manufacturer":
		p.Manufacturer = val
	case "stock":
		stock, _ := strconv.ParseInt(val, 10, 64)
		p.Stock = stock
	case "stock_alert":
		stockAlert, _ := strconv.ParseInt(val, 10, 64)
		p.StockAlert = stockAlert
	case "weight":
		weight, _ := strconv.ParseFloat(val, 64)
		p.Weight = weight
	case "width":
		width, _ := strconv.ParseFloat(val, 64)
		p.Width = width
	case "height":
		height, _ := strconv.ParseFloat(val, 64)
		p.Height = height
	case "depth":
		depth, _ := strconv.ParseFloat(val, 64)
		p.Depth = depth
	case "shipping_markup":
		smarkup, _ := strconv.ParseFloat(val, 64)
		p.ShippingMarkup = smarkup
	case "visible":
		visible, _ := strconv.ParseBool(val)
		p.Visible = visible
	case "searchable":
		searchable, _ := strconv.ParseBool(val)
		p.Searchable = searchable
	case "multi_box":
		multiBox, _ := strconv.ParseBool(val)
		p.MultiBox = multiBox
	case "ship_separately":
		shipSep, _ := strconv.ParseBool(val)
		p.ShipSeparately = shipSep
	case "free_shipping":
		fship, _ := strconv.ParseBool(val)
		p.FreeShipping = fship
	case "promoted":
		promoted, _ := strconv.ParseBool(val)
		p.Promoted = promoted
	case "dropship":
		dropship, _ := strconv.ParseBool(val)
		p.Dropship = dropship
	case "special_processing":
		sproc, _ := strconv.ParseBool(val)
		p.SpecialProcessing = sproc
	case "special_processing_type":
		p.SpecialProcessingType = val
	case "size":
		p.Size = val
	case "color":
		p.Color = val
	case "thumbnail":
		p.Thumbnail = val
	case "image1":
		p.Image1 = val
	case "image2":
		p.Image2 = val
	case "image3":
		p.Image3 = val
	case "image4":
		p.Image4 = val
	case "parent_product_sku":
		parentProductFound = m.processParentProduct(val, p, hd)
	case "category":
		categoryFound = m.processProductCategory(val, p, hd)
	}
	m.Log.Debug("parentProductFound: ", parentProductFound)
	m.Log.Debug("categoryFound: ", categoryFound)
	if parentProductFound && categoryFound {
		complete = true
	}
	return complete
}

func (m *Six910Manager) processParentProduct(val string, p *Product, hd *api.Headers) bool {
	var found = true
	if val != "" {
		pprod := m.API.GetProductBySku(val, p.DistributorID, hd)
		m.Log.Debug("pprod lookup by sku: ", pprod)
		if pprod == nil || pprod.ID == 0 {
			found = false
		} else {
			p.ParentProductID = pprod.ID
		}
	}
	return found
}

func (m *Six910Manager) processProductCategory(val string, p *Product, hd *api.Headers) bool {
	var found = true
	m.Log.Debug("category: ", val)
	catList := strings.Split(val, "/")
	m.Log.Debug("catList: ", catList)
	if len(catList) > 0 {
		p.CategoryID = m.createCategory(&catList, hd)
		if p.CategoryID == 0 {
			found = false
		}
	}
	m.Log.Debug("product: ", *p)

	return found
}

func (m *Six910Manager) createCategory(catList *[]string, hd *api.Headers) int64 {
	//time.Sleep(100 * time.Millisecond)
	var catID int64
	m.Log.Debug("catList in createCategory: ", catList)
	var catList2 []string
	for _, c := range *catList {
		if c != "" {
			catList2 = append(catList2, c)
		}
	}
	fcs := m.API.GetCategoryList(hd)

	var clist []sdbi.Category
	for i, c := range catList2 {
		m.Log.Debug("looking for can: ", c)
		var found bool
		for _, fc := range *fcs {
			if fc.Name == c {
				clist = append(clist, fc)
				found = true
				break
			}
		}
		if !found {
			var nc sdbi.Category
			nc.Name = c
			nc.Description = c + " desc"
			if len(clist) > 0 {
				nc.ParentCategoryID = clist[i-1].ID
			}
			m.Log.Debug("Creating new category: ", nc)
			res := m.API.AddCategory(&nc, hd)
			if res.Success && res.ID != 0 {
				nc.ID = res.ID
				clist = append(clist, nc)
			}
		}
	}
	m.Log.Debug("cat list built: ", clist)
	var lstlen = len(clist)
	if lstlen > 0 {
		catID = clist[lstlen-1].ID
	}
	return catID
}
