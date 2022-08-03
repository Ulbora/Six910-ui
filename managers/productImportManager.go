package managers

import ( //api "github.com/Ulbora/Six910API-Go"
	"sync"

	//api "github.com/Ulbora/Six910API-Go"

	api "github.com/Ulbora/Six910API-Go"
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

// func (m *Six910Manager) importProducts(prodList *[]Product, hd *api.Headers) bool {
// 	var rtn = true
// 	var wg sync.WaitGroup
// 	var pchan = make(chan *api.ResponseID, len(*prodList))

// 	var cnt = 0
// 	for i := range *prodList {
// 		m.Log.Debug("cnt :", cnt)
// 		//time.Sleep(2000 * time.Millisecond)
// 		cnt++
// 		if cnt > 10 {
// 			m.Log.Debug("sleeping :", 5000*time.Millisecond)
// 			cnt = 0
// 			m.Log.Debug("cnt :", cnt)
// 			time.Sleep(5000 * time.Millisecond)
// 		}
// 		var cp = &(*prodList)[i]
// 		m.Log.Debug("before goroutine :", cp.Sku)
// 		wg.Add(1)
// 		go func(product *Product, header *api.Headers, prodchan chan *api.ResponseID) {
// 			m.Log.Debug("in goroutine product address:", product)
// 			m.Log.Debug("in goroutine product.Sku:", product.Sku)
// 			m.Log.Debug("in goroutine product.Name:", product.Name)
// 			defer wg.Done()
// 			// need to search for product before adding
// 			fpd := m.API.GetProductBySku(product.Sku, product.DistributorID, hd)
// 			if fpd != nil && fpd.ID != 0 {
// 				ep := m.parseExistingProduct(fpd, product)
// 				m.Log.Debug("in goroutine parsed existing product:", *ep)
// 				pres := m.API.UpdateProduct(ep, hd)
// 				m.Log.Debug("in goroutine pres existing:", pres)
// 				var npres api.ResponseID
// 				npres.Success = pres.Success
// 				prodchan <- &npres
// 			} else {
// 				np := m.parseProduct(product)
// 				m.Log.Debug("in goroutine parsed product:", *np)
// 				pres := m.API.AddProduct(np, header)
// 				m.Log.Debug("in goroutine pres:", pres)
// 				if pres.Success && pres.ID != 0 {
// 					if product.CategoryID != 0 {
// 						var pc sdbi.ProductCategory
// 						pc.CategoryID = product.CategoryID
// 						pc.ProductID = pres.ID
// 						cres := m.API.AddProductCategory(&pc, header)
// 						pres.Success = cres.Success
// 					}
// 				}
// 				prodchan <- pres
// 			}
// 		}(cp, hd, pchan)
// 	}
// 	m.Log.Debug("before wait")
// 	wg.Wait()
// 	m.Log.Debug("after wait")
// 	close(pchan)
// 	m.Log.Debug("chan len :", len(pchan))
// 	for res := range pchan {
// 		if !res.Success {
// 			rtn = false
// 		}
// 	}
// 	return rtn
// }

func (m *Six910Manager) importProducts(prodList *[]Product, hd *api.Headers) int {
	var rtn int
	var wg sync.WaitGroup
	var pchan = make(chan *api.ResponseID, len(*prodList))

	wg.Add(1)
	go func(pl *[]Product, header *api.Headers, prodchan chan *api.ResponseID) {
		defer wg.Done()

		for i := range *pl {
			var cp = &(*pl)[i]
			m.Log.Debug("before goroutine :", cp.Sku)
			//wg.Add(1)
			//go func(product *Product, header *api.Headers, prodchan chan *api.ResponseID) {
			//m.Log.Debug("in goroutine product address:", product)
			//m.Log.Debug("in goroutine product.Sku:", product.Sku)
			//m.Log.Debug("in goroutine product.Name:", product.Name)
			//defer wg.Done()
			// need to search for product before adding
			fpd := m.API.GetProductBySku(cp.Sku, cp.DistributorID, header)
			if fpd != nil && fpd.ID != 0 {
				ep := m.parseExistingProduct(fpd, cp)
				m.Log.Debug("in goroutine parsed existing product:", *ep)
				pres := m.API.UpdateProduct(ep, header)
				m.Log.Debug("in goroutine pres existing:", pres)
				var npres api.ResponseID
				if pres.Success {
					npres.Success = pres.Success
					prodchan <- &npres
				}
			} else if cp.Desc != "" && cp.Name != "" {
				np := m.parseProduct(cp)
				m.Log.Debug("in goroutine parsed product:", *np)
				pres := m.API.AddProduct(np, header)
				m.Log.Debug("in goroutine pres:", pres)
				if pres.Success && pres.ID != 0 {
					if cp.CategoryID != 0 {
						var pc sdbi.ProductCategory
						pc.CategoryID = cp.CategoryID
						pc.ProductID = pres.ID
						cres := m.API.AddProductCategory(&pc, header)
						pres.Success = cres.Success
					}
				}
				if pres.Success {
					prodchan <- pres
				}
			}
			//}(cp, hd, pchan)
		}
	}(prodList, hd, pchan)
	m.Log.Debug("before wait")
	wg.Wait()
	m.Log.Debug("after wait")
	close(pchan)
	m.Log.Debug("chan len :", len(pchan))
	rtn = len(pchan)
	// for res := range pchan {
	// 	if !res.Success {
	// 		rtn = false
	// 	}
	// }
	return rtn
}

func (m *Six910Manager) parseProduct(p *Product) *sdbi.Product {
	var rtn sdbi.Product
	rtn.Color = p.Color
	rtn.Gender = p.Gender
	rtn.Cost = p.Cost
	rtn.Currency = p.Currency
	rtn.Depth = p.Depth
	rtn.Desc = p.Desc
	rtn.DistributorID = p.DistributorID
	rtn.Dropship = p.Dropship
	rtn.FreeShipping = p.FreeShipping
	rtn.Gtin = p.Gtin
	rtn.Height = p.Height
	rtn.Image1 = p.Image1
	rtn.Image2 = p.Image2
	rtn.Image3 = p.Image3
	rtn.Image4 = p.Image4
	rtn.Manufacturer = p.Manufacturer
	rtn.ManufacturerID = p.ManufacturerID
	rtn.Map = p.Map
	rtn.Msrp = p.Msrp
	rtn.MultiBox = p.MultiBox
	rtn.Name = p.Name
	rtn.ParentProductID = p.ParentProductID
	rtn.Price = p.Price
	rtn.Promoted = p.Promoted
	rtn.SalePrice = p.SalePrice
	rtn.Searchable = p.Searchable
	rtn.ShipSeparately = p.ShipSeparately
	rtn.ShippingMarkup = p.ShippingMarkup
	rtn.ShortDesc = p.ShortDesc
	rtn.Size = p.Size
	rtn.Sku = p.Sku
	rtn.SpecialProcessing = p.SpecialProcessing
	rtn.SpecialProcessingType = p.SpecialProcessingType
	rtn.Stock = p.Stock
	rtn.StockAlert = p.StockAlert
	rtn.Thumbnail = p.Thumbnail
	rtn.Visible = p.Visible
	rtn.Weight = p.Weight
	rtn.Width = p.Width
	rtn.Visible = true
	rtn.Searchable = true
	return &rtn
}

func (m *Six910Manager) parseExistingProduct(ep *sdbi.Product, up *Product) *sdbi.Product {
	var rtn = ep
	//rtn.ID = ep.ID
	if up.Color != "" {
		rtn.Color = up.Color
	}
	if up.Gender != "" {
		rtn.Gender = up.Gender
	}
	if up.Cost != 0 {
		rtn.Cost = up.Cost
	}
	if up.Currency != "" {
		rtn.Currency = up.Currency
	}
	if up.Depth != 0 {
		rtn.Depth = up.Depth
	}
	if up.Desc != "" {
		rtn.Desc = up.Desc
	}
	//rtn.DistributorID = ep.DistributorID
	//rtn.Dropship = ep.Dropship
	//rtn.FreeShipping = ep.FreeShipping
	if up.Gtin != "" {
		rtn.Gtin = up.Gtin
	}
	if up.Height != 0 {
		rtn.Height = up.Height
	}
	if up.Image1 != "" {
		rtn.Image1 = up.Image1
	}
	if up.Image2 != "" {
		rtn.Image2 = up.Image2
	}
	if up.Image3 != "" {
		rtn.Image3 = up.Image3
	}
	if up.Image4 != "" {
		rtn.Image4 = up.Image4
	}

	//rtn.Image2 = ep.Image2
	//rtn.Image3 = ep.Image3
	//rtn.Image4 = ep.Image4
	if up.Manufacturer != "" {
		rtn.Manufacturer = up.Manufacturer
	}

	if up.ManufacturerID != "" {
		rtn.ManufacturerID = up.ManufacturerID
	}
	if up.Map != 0 {
		rtn.Map = up.Map
	}
	if up.Msrp != 0 {
		rtn.Msrp = up.Msrp
	}

	//rtn.MultiBox = ep.MultiBox
	if up.Name != "" {
		rtn.Name = up.Name
	}

	//rtn.ParentProductID = ep.ParentProductID
	if up.Price != 0 {
		rtn.Price = up.Price
	}

	//rtn.Promoted = ep.Promoted
	if up.SalePrice != 0 {
		rtn.SalePrice = up.SalePrice
	}

	//rtn.Searchable = ep.Searchable
	//rtn.ShipSeparately = ep.ShipSeparately
	//rtn.ShippingMarkup = ep.ShippingMarkup
	if up.ShortDesc != "" {
		rtn.ShortDesc = up.ShortDesc
	}
	if up.Size != "" {
		rtn.Size = up.Size
	}

	//rtn.Sku = ep.Sku
	//rtn.SpecialProcessing = ep.SpecialProcessing
	//rtn.SpecialProcessingType = ep.SpecialProcessingType
	//rtn.Stock = ep.Stock
	//rtn.StockAlert = ep.StockAlert
	if up.Thumbnail != "" {
		rtn.Thumbnail = up.Thumbnail
	}

	//rtn.Visible = ep.Visible
	if up.Weight != 0 {
		rtn.Weight = up.Weight
	}
	if up.Width != 0 {
		rtn.Width = up.Width
	}
	rtn.Stock = up.Stock

	return rtn
}
