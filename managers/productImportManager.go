package managers

import ( //api "github.com/Ulbora/Six910API-Go"
	//api "github.com/Ulbora/Six910API-Go"

	"sync"

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

func (m *Six910Manager) importProducts(prodList *[]Product, hd *api.Headers) bool {
	var rtn = true
	var wg sync.WaitGroup
	var pchan = make(chan *api.ResponseID, len(*prodList))

	for i := range *prodList {
		var cp = &(*prodList)[i]
		m.Log.Debug("before goroutine :", cp.Sku)
		wg.Add(1)
		go func(product *Product, header *api.Headers, prodchan chan *api.ResponseID) {
			m.Log.Debug("in goroutine product address:", product)
			m.Log.Debug("in goroutine product.Sku:", product.Sku)
			m.Log.Debug("in goroutine product.Name:", product.Name)
			defer wg.Done()
			// need to search for product before adding
			fpd := m.API.GetProductBySku(product.Sku, product.DistributorID, hd)
			if fpd != nil && fpd.ID != 0 {
				ep := m.parseExistingProduct(fpd)
				m.Log.Debug("in goroutine parsed existing product:", *ep)
				pres := m.API.UpdateProduct(ep, hd)
				m.Log.Debug("in goroutine pres existing:", pres)
				var npres api.ResponseID
				npres.Success = pres.Success
				prodchan <- &npres
			} else {
				np := m.parseProduct(product)
				m.Log.Debug("in goroutine parsed product:", *np)
				pres := m.API.AddProduct(np, header)
				m.Log.Debug("in goroutine pres:", pres)
				if pres.Success && pres.ID != 0 {
					if product.CategoryID != 0 {
						var pc sdbi.ProductCategory
						pc.CategoryID = product.CategoryID
						pc.ProductID = pres.ID
						cres := m.API.AddProductCategory(&pc, header)
						pres.Success = cres.Success
					}
				}
				prodchan <- pres
			}
		}(cp, hd, pchan)
	}
	m.Log.Debug("before wait")
	wg.Wait()
	m.Log.Debug("after wait")
	close(pchan)
	m.Log.Debug("chan len :", len(pchan))
	for res := range pchan {
		if !res.Success {
			rtn = false
		}
	}
	return rtn
}

func (m *Six910Manager) parseProduct(p *Product) *sdbi.Product {
	var rtn sdbi.Product
	rtn.Color = p.Color
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
	return &rtn
}

func (m *Six910Manager) parseExistingProduct(p *sdbi.Product) *sdbi.Product {
	var rtn sdbi.Product
	rtn.ID = p.ID
	rtn.Color = p.Color
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
	return &rtn
}
