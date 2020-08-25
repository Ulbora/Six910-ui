//Package managers ...
package managers

import (
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

const (
	storeAdmin   = "StoreAdmin"
	customerRole = "customer"

	billingAddressType  = "Billing"
	shippingAddressType = "Shipping"

	orderStatusProcessing = "processing"
)

//Product Product
type Product struct {
	ID                    int64   `json:"id"`
	Sku                   string  `json:"sku"`
	Gtin                  string  `json:"gtin"`
	Name                  string  `json:"name"`
	ShortDesc             string  `json:"shortDesc"`
	Desc                  string  `json:"desc"`
	Cost                  float64 `json:"cost"`
	Msrp                  float64 `json:"msrp"`
	Map                   float64 `json:"map"`
	Price                 float64 `json:"price"`
	SalePrice             float64 `json:"salePrice"`
	Currency              string  `json:"currency"`
	ManufacturerID        string  `json:"manufacturerId"`
	Manufacturer          string  `json:"manufacturer"`
	Stock                 int64   `json:"stock"`
	StockAlert            int64   `json:"stockAlert"`
	Weight                float64 `json:"weight"`
	Width                 float64 `json:"width"`
	Height                float64 `json:"height"`
	Depth                 float64 `json:"depth"`
	ShippingMarkup        float64 `json:"shippingMarkup"`
	Visible               bool    `json:"visible"`
	Searchable            bool    `json:"searchable"`
	MultiBox              bool    `json:"multibox"`
	ShipSeparately        bool    `json:"shipSeparately"`
	FreeShipping          bool    `json:"freeShipping"`
	Promoted              bool    `json:"promoted"`
	Dropship              bool    `json:"dropship"`
	SpecialProcessing     bool    `json:"specialProcessing"`
	SpecialProcessingType string  `json:"specialProcessingType"`
	Size                  string  `json:"size"`
	Color                 string  `json:"color"`
	Thumbnail             string  `json:"thumbnail"`
	Image1                string  `json:"image1"`
	Image2                string  `json:"image2"`
	Image3                string  `json:"image3"`
	Image4                string  `json:"image4"`
	DistributorID         int64   `json:"distributorId"`
	StoreID               int64   `json:"storeId"`
	ParentProductID       int64   `json:"parentProductId"`
	CategoryID            int64
}

//CustomerAccount CustomerAccount
type CustomerAccount struct {
	Customer  *sdbi.Customer
	Addresses *[]sdbi.Address
	User      *api.User
	//Cart      *CustomerCart
}

//CustomerProduct Product
type CustomerProduct struct {
	ProductID int64
	Quantity  int64
	//Cart      *CustomerCart
	CustomerID int64
	//CartID     int64
	Cart *sdbi.Cart
	//CartItem   *sdbi.CartItem
	//CustomerEmail string
	StoreID int64
}

//CustomerProductUpdate Product
type CustomerProductUpdate struct {
	//ProductID int64
	//Quantity  int64
	//Cart      *CustomerCart
	CustomerID int64
	Cart       *sdbi.Cart
	CartItem   *sdbi.CartItem
	//CustomerEmail string
	//StoreID int64
}

//CustomerCart CustomerCart
type CustomerCart struct {
	Cart             *sdbi.Cart
	Items            *[]sdbi.CartItem
	Comment          string
	CustomerAccount  *CustomerAccount
	InsuranceCost    float64
	OrderType        string
	Pickup           bool
	ShippingHandling float64
	Subtotal         float64
	Taxes            float64
	Total            float64
}

//CustomerOrder CustomerOrder
type CustomerOrder struct {
	Success         bool
	Order           *sdbi.Order
	Items           *[]sdbi.OrderItem
	Comments        *[]sdbi.OrderComment
	CustomerAccount *CustomerAccount
	Cart            *sdbi.Cart
}

//OrderItemResults OrderItemResults
type OrderItemResults struct {
	OrderItem *sdbi.OrderItem
	Resp      *api.ResponseID
}

//CartViewItem CartViewItem
type CartViewItem struct {
	ProductID int64
	Desc      string
	Image     string
	Quantity  int64
	Price     float64
	Total     float64
}

//CartView CartView
type CartView struct {
	Items *[]*CartViewItem
	Total float64
}

//Manager Manager
type Manager interface {
	// ------ Customer methods -----------

	//--------------------start----new------------

	AddProductToCart(cp *CustomerProduct, hd *api.Headers) *CustomerCart
	ViewCart(cc *CustomerCart, hd *api.Headers) *CartView
	UpdateProductToCart(cp *CustomerProductUpdate, hd *api.Headers) *CustomerCart
	CheckOut(cart *CustomerCart, hd *api.Headers) *CustomerOrder

	CreateCustomerAccount(cus *CustomerAccount, hd *api.Headers) (bool, *CustomerAccount)
	UpdateCustomerAccount(cus *CustomerAccount, hd *api.Headers) bool

	ViewCustomerOrder(orderID int64, cid int64, hd *api.Headers) *CustomerOrder
	ViewCustomerOrderList(cid int64, hd *api.Headers) *[]CustomerOrder

	CustomerLogin(u *api.User, hd *api.Headers) (bool, *api.User)
	CustomerChangePassword(u *api.User, hd *api.Headers) (bool, *api.User)

	//--------------------end ---new------------

	// ------some may change
	///////AddAddress(a *sdbi.Address, hd *api.Headers) *api.ResponseID
	///////UpdateAddress(a *sdbi.Address, hd *api.Headers) *api.Response
	// GetAddress(id int64, cid int64, hd *Headers) *sdbi.Address
	// GetAddressList(cid int64, hd *Headers) *[]sdbi.Address
	// DeleteAddress(id int64, cid int64, hd *Headers) *Response

	// //cart
	// AddCart(c *sdbi.Cart, hd *Headers) *ResponseID
	// UpdateCart(c *sdbi.Cart, hd *Headers) *Response
	// GetCart(cid int64, hd *Headers) *sdbi.Cart
	// DeleteCart(id int64, cid int64, hd *Headers) *Response

	// //cartItem
	// AddCartItem(ci *sdbi.CartItem, cid int64, hd *Headers) *ResponseID
	// UpdateCartItem(ci *sdbi.CartItem, cid int64, hd *Headers) *Response
	// GetCartItem(cid int64, prodID int64, hd *Headers) *sdbi.CartItem
	// GetCartItemList(cartID int64, cid int64, hd *Headers) *[]sdbi.CartItem
	// DeleteCartItem(id int64, prodID int64, cartID int64, hd *Headers) *Response

	// //customer
	// AddCustomer(c *sdbi.Customer, hd *Headers) *ResponseID
	// UpdateCustomer(c *sdbi.Customer, hd *Headers) *Response
	// GetCustomer(email string, hd *Headers) *sdbi.Customer
	// GetCustomerID(id int64, hd *Headers) *sdbi.Customer
	// GetCustomerList(hd *Headers) *[]sdbi.Customer
	// DeleteCustomer(id int64, hd *Headers) *Response

	// //order
	// AddOrder(o *sdbi.Order, hd *Headers) *ResponseID
	// UpdateOrder(o *sdbi.Order, hd *Headers) *Response
	// GetOrder(id int64, hd *Headers) *sdbi.Order
	// GetOrderList(cid int64, hd *Headers) *[]sdbi.Order
	// DeleteOrder(id int64, hd *Headers) *Response

	// //order comments
	// AddOrderComments(c *sdbi.OrderComment, hd *Headers) *ResponseID
	// GetOrderCommentList(orderID int64, hd *Headers) *[]sdbi.OrderComment

	// //order items
	// AddOrderItem(i *sdbi.OrderItem, hd *Headers) *ResponseID
	// UpdateOrderItem(i *sdbi.OrderItem, hd *Headers) *Response
	// GetOrderItem(id int64, hd *Headers) *sdbi.OrderItem
	// GetOrderItemList(orderID int64, hd *Headers) *[]sdbi.OrderItem
	// DeleteOrderItem(id int64, hd *Headers) *Response

	// //order transaction
	// AddOrderTransaction(t *sdbi.OrderTransaction, hd *Headers) *ResponseID
	// GetOrderTransactionList(orderID int64, hd *Headers) *[]sdbi.OrderTransaction

	// ------Super Admin------------

	// //store
	////////// AddStore(s *sdbi.Store, hd *Headers) *ResponseID
	// UpdateStore(s *sdbi.Store, hd *Headers) *Response
	// GetStore(sname string, localDomain string, hd *Headers) *sdbi.Store
	///////// DeleteStore(sname string, localDomain string, hd *Headers) *Response

	// //user
	//////// AddCustomerUser(u *User, hd *Headers) *Response
	// UpdateUser(u *User, hd *Headers) *Response
	// GetUser(u *User, hd *Headers) *UserResponse
	// GetAdminUsers(hd *Headers) *[]UserResponse
	// GetCustomerUsers(hd *Headers) *[]UserResponse

	//-------Store Admin-----

	StoreAdminLogin(u *api.User, hd *api.Headers) (bool, *api.User)
	StoreAdminChangePassword(u *api.User, hd *api.Headers) (bool, *api.User)
	UploadProductFile(file []byte, hd *api.Headers) (success bool, productNotImported int)

	// //category
	// AddCategory(c *sdbi.Category, hd *Headers) *ResponseID
	// UpdateCategory(c *sdbi.Category, hd *Headers) *Response
	// GetCategory(id int64, hd *Headers) *sdbi.Category
	// GetCategoryList(hd *Headers) *[]sdbi.Category
	// GetSubCategoryList(catID int64, hd *Headers) *[]sdbi.Category
	// DeleteCategory(id int64, hd *Headers) *Response

	// //distrubutor
	// AddDistributor(d *sdbi.Distributor, hd *Headers) *ResponseID
	// UpdateDistributor(d *sdbi.Distributor, hd *Headers) *Response
	// GetDistributor(id int64, hd *Headers) *sdbi.Distributor
	// GetDistributorList(hd *Headers) *[]sdbi.Distributor
	// DeleteDistributor(id int64, hd *Headers) *Response

	// // insurance
	// AddInsurance(i *sdbi.Insurance, hd *Headers) *ResponseID
	// UpdateInsurance(i *sdbi.Insurance, hd *Headers) *Response
	// GetInsurance(id int64, hd *Headers) *sdbi.Insurance
	// GetInsuranceList(hd *Headers) *[]sdbi.Insurance
	// DeleteInsurance(id int64, hd *Headers) *Response

	// //plugins
	// AddPlugin(p *sdbi.Plugins, hd *Headers) *ResponseID
	// UpdatePlugin(p *sdbi.Plugins, hd *Headers) *Response
	// GetPlugin(id int64, hd *Headers) *sdbi.Plugins
	// GetPluginList(start int64, end int64, hd *Headers) *[]sdbi.Plugins
	// DeletePlugin(id int64, hd *Headers) *Response

	// //store plugin
	// AddStorePlugin(sp *sdbi.StorePlugins, hd *Headers) *ResponseID
	// UpdateStorePlugin(sp *sdbi.StorePlugins, hd *Headers) *Response
	// GetStorePlugin(id int64, hd *Headers) *sdbi.StorePlugins
	// GetStorePluginList(hd *Headers) *[]sdbi.StorePlugins
	// DeleteStorePlugin(id int64, hd *Headers) *Response

	// //payment gateway
	// AddPaymentGateway(pgw *sdbi.PaymentGateway, hd *Headers) *ResponseID
	// UpdatePaymentGateway(pgw *sdbi.PaymentGateway, hd *Headers) *Response
	// GetPaymentGateway(id int64, hd *Headers) *sdbi.PaymentGateway
	// GetPaymentGateways(hd *Headers) *[]sdbi.PaymentGateway
	// DeletePaymentGateway(id int64, hd *Headers) *Response

	// //shipment carrier
	// AddShippingCarrier(c *sdbi.ShippingCarrier, hd *Headers) *ResponseID
	// UpdateShippingCarrier(c *sdbi.ShippingCarrier, hd *Headers) *Response
	// GetShippingCarrier(id int64, hd *Headers) *sdbi.ShippingCarrier
	// GetShippingCarrierList(hd *Headers) *[]sdbi.ShippingCarrier
	// DeleteShippingCarrier(id int64, hd *Headers) *Response

	// //shipment method
	// AddShippingMethod(s *sdbi.ShippingMethod, hd *Headers) *ResponseID
	// UpdateShippingMethod(s *sdbi.ShippingMethod, hd *Headers) *Response
	// GetShippingMethod(id int64, hd *Headers) *sdbi.ShippingMethod
	// GetShippingMethodList(hd *Headers) *[]sdbi.ShippingMethod
	// DeleteShippingMethod(id int64, hd *Headers) *Response

	// //region
	// AddRegion(r *sdbi.Region, hd *Headers) *ResponseID
	// UpdateRegion(r *sdbi.Region, hd *Headers) *Response
	// GetRegion(id int64, hd *Headers) *sdbi.Region
	// GetRegionList(hd *Headers) *[]sdbi.Region
	// DeleteRegion(id int64, hd *Headers) *Response

	// //sub region
	// AddSubRegion(s *sdbi.SubRegion, hd *Headers) *ResponseID
	// UpdateSubRegion(s *sdbi.SubRegion, hd *Headers) *Response
	// GetSubRegion(id int64, hd *Headers) *sdbi.SubRegion
	// GetSubRegionList(regionID int64, hd *Headers) *[]sdbi.SubRegion
	// DeleteSubRegion(id int64, hd *Headers) *Response

	// //excluded sub region
	// AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, hd *Headers) *ResponseID
	// GetExcludedSubRegionList(regionID int64, hd *Headers) *[]sdbi.ExcludedSubRegion
	// DeleteExcludedSubRegion(id int64, regionID int64, hd *Headers) *Response

	// //included sub region
	// AddIncludedSubRegion(e *sdbi.IncludedSubRegion, hd *Headers) *ResponseID
	// GetIncludedSubRegionList(regionID int64, hd *Headers) *[]sdbi.IncludedSubRegion
	// DeleteIncludedSubRegion(id int64, regionID int64, hd *Headers) *Response

	// //zip code zone
	// AddZoneZip(z *sdbi.ZoneZip, hd *Headers) *ResponseID
	// GetZoneZipListByExclusion(exID int64, hd *Headers) *[]sdbi.ZoneZip
	// GetZoneZipListByInclusion(incID int64, hd *Headers) *[]sdbi.ZoneZip
	// DeleteZoneZip(id int64, incID int64, exID int64, hd *Headers) *Response

	// //products
	// AddProduct(p *sdbi.Product, hd *Headers) *ResponseID
	// UpdateProduct(p *sdbi.Product, hd *Headers) *Response
	// GetProductByID(id int64, hd *Headers) *sdbi.Product
	// GetProductsByName(name string, start int64, end int64, hd *Headers) *[]sdbi.Product
	// GetProductsByCaterory(catID int64, start int64, end int64, hd *Headers) *[]sdbi.Product
	// GetProductList(start int64, end int64, hd *Headers) *[]sdbi.Product
	// DeleteProduct(id int64, hd *Headers) *Response

	// //product category
	// AddProductCategory(pc *sdbi.ProductCategory, hd *Headers) *Response
	// DeleteProductCategory(pc *sdbi.ProductCategory, hd *Headers) *Response

	// //shipment
	// AddShipment(s *sdbi.Shipment, hd *Headers) *ResponseID
	// UpdateShipment(s *sdbi.Shipment, hd *Headers) *Response
	// GetShipment(id int64, hd *Headers) *sdbi.Shipment
	// GetShipmentList(orderID int64, hd *Headers) *[]sdbi.Shipment
	// DeleteShipment(id int64, hd *Headers) *Response

	// //shipment box
	// AddShipmentBox(sb *sdbi.ShipmentBox, hd *Headers) *ResponseID
	// UpdateShipmentBox(sb *sdbi.ShipmentBox, hd *Headers) *Response
	// GetShipmentBox(id int64, hd *Headers) *sdbi.ShipmentBox
	// GetShipmentBoxList(shipmentID int64, hd *Headers) *[]sdbi.ShipmentBox
	// DeleteShipmentBox(id int64, hd *Headers) *Response

	// //shipment item
	// AddShipmentItem(si *sdbi.ShipmentItem, hd *Headers) *ResponseID
	// UpdateShipmentItem(si *sdbi.ShipmentItem, hd *Headers) *Response
	// GetShipmentItem(id int64, hd *Headers) *sdbi.ShipmentItem
	// GetShipmentItemList(shipmentID int64, hd *Headers) *[]sdbi.ShipmentItem
	// GetShipmentItemListByBox(boxNumber int64, shipmentID int64, hd *Headers) *[]sdbi.ShipmentItem
	// DeleteShipmentItem(id int64, hd *Headers) *Response

	// //cart
	// GetCart(cid int64, hd *Headers) *sdbi.Cart

	// //cartItem
	// GetCartItem(cid int64, prodID int64, hd *Headers) *sdbi.CartItem
	// GetCartItemList(cartID int64, cid int64, hd *Headers) *[]sdbi.CartItem

	// //customer
	// GetCustomer(email string, hd *Headers) *sdbi.Customer
	// GetCustomerID(id int64, hd *Headers) *sdbi.Customer
	// GetCustomerList(hd *Headers) *[]sdbi.Customer
	// DeleteCustomer(id int64, hd *Headers) *Response

	// //order
	// UpdateOrder(o *sdbi.Order, hd *Headers) *Response
	// GetOrder(id int64, hd *Headers) *sdbi.Order
	// GetOrderList(cid int64, hd *Headers) *[]sdbi.Order
	// DeleteOrder(id int64, hd *Headers) *Response

	// //order comments
	// AddOrderComments(c *sdbi.OrderComment, hd *Headers) *ResponseID
	// GetOrderCommentList(orderID int64, hd *Headers) *[]sdbi.OrderComment

	// //order items
	// GetOrderItem(id int64, hd *Headers) *sdbi.OrderItem
	// GetOrderItemList(orderID int64, hd *Headers) *[]sdbi.OrderItem
	// DeleteOrderItem(id int64, hd *Headers) *Response

	// //order transaction
	// GetOrderTransactionList(orderID int64, hd *Headers) *[]sdbi.OrderTransaction

	// //dataStore
	// AddLocalDatastore(d *sdbi.LocalDataStore, hd *Headers) *Response
	// UpdateLocalDatastore(d *sdbi.LocalDataStore, hd *Headers) *Response
	// GetLocalDatastore(dataStoreName string, hd *Headers) *sdbi.LocalDataStore

	// //instances
	// AddInstance(i *sdbi.Instances, hd *Headers) *Response
	// UpdateInstance(i *sdbi.Instances, hd *Headers) *Response
	// GetInstance(name string, dataStoreName string, hd *Headers) *sdbi.Instances
	// GetInstanceList(dataStoreName string, hd *Headers) *[]sdbi.Instances

	// //dataStoreWriteLock
	// AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock, hd *Headers) *Response
	// UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock, hd *Headers) *Response
	// GetDataStoreWriteLock(dataStore string, hd *Headers) *sdbi.DataStoreWriteLock

}
