package mockapi

import (
	px "github.com/Ulbora/GoProxy"
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

//MockAPI MockAPI
type MockAPI struct {
	MockUser               *api.UserResponse
	MockAddCustomerUserRes *api.Response
	MockUpdateUserResp     *api.Response

	MockCart        *sdbi.Cart
	MockAddCartResp *api.ResponseID

	MockCartItemAddResp    *api.ResponseID
	MockCartItemUpdateResp *api.Response
	MockCartItemList       *[]sdbi.CartItem

	MockCustomer           *sdbi.Customer
	MockAddCustomerResp    *api.ResponseID
	MockUpdateCustomerResp *api.Response
	MockCustomerList       *[]sdbi.Customer

	MockAddAddressRes    *api.ResponseID
	MockUpdateAddressRes *api.Response
	mockAddressList1Used bool
	MockAddressList1     *[]sdbi.Address
	MockAddressList2     *[]sdbi.Address
	MockAddress          *sdbi.Address
	MockDeleteAddressRes *api.Response

	MockProduct           *sdbi.Product
	MockAddProductResp    *api.ResponseID
	MockUpdateProductResp *api.Response
	MockProductList       *[]sdbi.Product
	MockDeleteProductResp *api.Response

	MockAddOrderResp    *api.ResponseID
	MockUpdateOrderResp *api.Response
	MockOrder           *sdbi.Order
	MockOrderList       *[]sdbi.Order

	MockAddOrderItemResp *api.ResponseID
	MockOrderItemList    *[]sdbi.OrderItem

	MockAddCommentResp *api.ResponseID
	MockCommentList    *[]sdbi.OrderComment

	mockAddCategory1User bool
	mockAddCategory2User bool
	mockAddCategory3User bool
	mockAddCategory4User bool
	MockAddCategoryResp1 *api.ResponseID
	MockAddCategoryResp2 *api.ResponseID
	MockAddCategoryResp3 *api.ResponseID
	MockAddCategoryResp4 *api.ResponseID

	MockUpdateCategoryResp *api.Response

	MockCategory           *sdbi.Category
	MockDeleteCategoryResp *api.Response

	MockCategoryList *[]sdbi.Category

	MockAddDistributorResp    *api.ResponseID
	MockUpdateDistributorResp *api.Response
	MockDistributor           *sdbi.Distributor
	MockDistributorList       *[]sdbi.Distributor
	MockDeleteDistributorResp *api.Response

	MockAddProductCategoryResp *api.Response

	MockProductCategoryIDList []int64

	MockAddShipmentResp    *api.ResponseID
	MockUpdateShipmentResp *api.Response
	MockShipment           *sdbi.Shipment
	MockShipmentList       *[]sdbi.Shipment
	MockDeleteShipmentResp *api.Response

	MockAddShipmentItemResp *api.ResponseID
	MockShippingItemList    *[]sdbi.ShipmentItem

	MockShipmentBoxList *[]sdbi.ShipmentBox

	MockAddInsuranceResp    *api.ResponseID
	MockUpdateInsuranceResp *api.Response
	MockInsurance           *sdbi.Insurance
	MockInsuranceList       *[]sdbi.Insurance
	MockDeleteInsuranceResp *api.Response

	MockAddPaymentGatewayResp    *api.ResponseID
	MockUpdatePaymentGatewayResp *api.Response
	MockPaymentGateway           *sdbi.PaymentGateway
	MockPaymentGatewayList       *[]sdbi.PaymentGateway
	MockDeletePaymentGatewayResp *api.Response

	MockAddPluginResp    *api.ResponseID
	MockUpdatePluginResp *api.Response
	MockPlugin           *sdbi.Plugins
	MockPluginList       *[]sdbi.Plugins
	MockDeletePluginResp *api.Response

	MockAddStorePluginResp    *api.ResponseID
	MockUpdateStorePluginResp *api.Response
	MockStorePlugin           *sdbi.StorePlugins
	MockStorePluginList       *[]sdbi.StorePlugins
	MockDeleteStorePluginResp *api.Response

	MockAddShippingCarrierResp    *api.ResponseID
	MockUpdateShippingCarrierResp *api.Response
	MockShippingCarrier           *sdbi.ShippingCarrier
	MockShippingCarrierList       *[]sdbi.ShippingCarrier
	MockDeleteShippingCarrierResp *api.Response

	MockAddShippingMethodResp    *api.ResponseID
	MockUpdateShippingMethodResp *api.Response
	MockShippingMethod           *sdbi.ShippingMethod
	MockShippingMethodList       *[]sdbi.ShippingMethod
	MockDeleteShippingMethodResp *api.Response

	MockAddRegionResp    *api.ResponseID
	MockUpdateRegionResp *api.Response
	MockRegion           *sdbi.Region
	MockRegionList       *[]sdbi.Region
	MockDeleteRegionResp *api.Response

	MockAddSubRegionResp    *api.ResponseID
	MockUpdateSubRegionResp *api.Response
	MockSubRegion           *sdbi.SubRegion
	MockSubRegionList       *[]sdbi.SubRegion
	MockDeleteSubRegionResp *api.Response

	MockAddExcludedSubRegionResp    *api.ResponseID
	MockUpdateExcludedSubRegionResp *api.Response
	MockExcludedSubRegion           *sdbi.ExcludedSubRegion
	MockExcludedSubRegionList       *[]sdbi.ExcludedSubRegion
	MockDeleteExcludedSubRegionResp *api.Response

	MockAddIncludedSubRegionResp    *api.ResponseID
	MockUpdateIncludedSubRegionResp *api.Response
	MockIncludedSubRegion           *sdbi.IncludedSubRegion
	MockIncludedSubRegionList       *[]sdbi.IncludedSubRegion
	MockDeleteIncludedSubRegionResp *api.Response

	MockAddZoneZipResp    *api.ResponseID
	MockUpdateZoneZipResp *api.Response
	MockZoneZip           *sdbi.ZoneZip
	MockZoneZipList       *[]sdbi.ZoneZip
	MockDeleteZoneZipResp *api.Response

	MockAddTaxRateResp    *api.ResponseID
	MockUpdateTaxRateResp *api.Response
	MockTaxRate           *sdbi.TaxRate
	MockTaxRateList       *[]sdbi.TaxRate
	MockDeleteTaxRateResp *api.Response
}

//GetNew GetNew
func (a *MockAPI) GetNew() api.API {
	return a
}

//SetLogLever SetLogLever
func (a *MockAPI) SetLogLever(level int) {

}

//SetStore SetStore
func (a *MockAPI) SetStore(storeName string, localDomain string) {

}

//SetRestURL SetRestURL
func (a *MockAPI) SetRestURL(url string) {

}

//SetAPIKey SetAPIKey
func (a *MockAPI) SetAPIKey(key string) {

}

//OverrideProxy OverrideProxy
func (a *MockAPI) OverrideProxy(proxy px.Proxy) {

}

//SetStoreID SetStoreID
func (a *MockAPI) SetStoreID(sid int64) {
}

//AddAddress AddAddress
func (a *MockAPI) AddAddress(ad *sdbi.Address, headers *api.Headers) *api.ResponseID {
	return a.MockAddAddressRes
}

//UpdateAddress UpdateAddress
func (a *MockAPI) UpdateAddress(ad *sdbi.Address, headers *api.Headers) *api.Response {
	return a.MockUpdateAddressRes
}

//GetAddress GetAddress
func (a *MockAPI) GetAddress(id int64, cid int64, headers *api.Headers) *sdbi.Address {
	return a.MockAddress
}

//GetAddressList GetAddressList
func (a *MockAPI) GetAddressList(cid int64, headers *api.Headers) *[]sdbi.Address {
	var rtn *[]sdbi.Address
	if !a.mockAddressList1Used {
		rtn = a.MockAddressList1
		a.mockAddressList1Used = true
	} else {
		rtn = a.MockAddressList2
	}
	return rtn
}

//DeleteAddress DeleteAddress
func (a *MockAPI) DeleteAddress(id int64, cid int64, headers *api.Headers) *api.Response {
	return a.MockDeleteAddressRes
}

//cart

//AddCart AddCart
func (a *MockAPI) AddCart(c *sdbi.Cart, headers *api.Headers) *api.ResponseID {
	return a.MockAddCartResp
}

//UpdateCart UpdateCart
func (a *MockAPI) UpdateCart(c *sdbi.Cart, headers *api.Headers) *api.Response {
	return nil
}

//GetCart GetCart
func (a *MockAPI) GetCart(cid int64, headers *api.Headers) *sdbi.Cart {
	return a.MockCart
}

//DeleteCart DeleteCart
func (a *MockAPI) DeleteCart(id int64, cid int64, headers *api.Headers) *api.Response {
	return nil
}

//cartItem

//AddCartItem AddCartItem
func (a *MockAPI) AddCartItem(ci *sdbi.CartItem, cid int64, headers *api.Headers) *api.ResponseID {
	return a.MockCartItemAddResp
}

//UpdateCartItem UpdateCartItem
func (a *MockAPI) UpdateCartItem(ci *sdbi.CartItem, cid int64, headers *api.Headers) *api.Response {
	return a.MockCartItemUpdateResp
}

//GetCartItem GetCartItem
func (a *MockAPI) GetCartItem(cid int64, prodID int64, headers *api.Headers) *sdbi.CartItem {
	return nil
}

//GetCartItemList GetCartItemList
func (a *MockAPI) GetCartItemList(cartID int64, cid int64, headers *api.Headers) *[]sdbi.CartItem {
	return a.MockCartItemList
}

//DeleteCartItem DeleteCartItem
func (a *MockAPI) DeleteCartItem(id int64, prodID int64, cartID int64, headers *api.Headers) *api.Response {
	return nil
}

//category

//AddCategory AddCategory
func (a *MockAPI) AddCategory(c *sdbi.Category, headers *api.Headers) *api.ResponseID {
	var rtn *api.ResponseID
	if !a.mockAddCategory1User {
		rtn = a.MockAddCategoryResp1
		a.mockAddCategory1User = true
	} else if !a.mockAddCategory2User {
		rtn = a.MockAddCategoryResp2
		a.mockAddCategory2User = true
	} else if !a.mockAddCategory3User {
		rtn = a.MockAddCategoryResp3
		a.mockAddCategory3User = true
	} else if !a.mockAddCategory4User {
		rtn = a.MockAddCategoryResp4
		a.mockAddCategory4User = true
	}
	return rtn
}

//UpdateCategory UpdateCategory
func (a *MockAPI) UpdateCategory(c *sdbi.Category, headers *api.Headers) *api.Response {
	return a.MockUpdateCategoryResp
}

//GetCategory GetCategory
func (a *MockAPI) GetCategory(id int64, headers *api.Headers) *sdbi.Category {
	return a.MockCategory
}

//GetHierarchicalCategoryList GetHierarchicalCategoryList
func (a *MockAPI) GetHierarchicalCategoryList(headers *api.Headers) *[]sdbi.Category {
	return a.MockCategoryList
}

//GetCategoryList GetCategoryList
func (a *MockAPI) GetCategoryList(headers *api.Headers) *[]sdbi.Category {
	return a.MockCategoryList
}

//GetSubCategoryList GetSubCategoryList
func (a *MockAPI) GetSubCategoryList(catID int64, headers *api.Headers) *[]sdbi.Category {
	return nil
}

//DeleteCategory DeleteCategory
func (a *MockAPI) DeleteCategory(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteCategoryResp
}

//AddCustomer AddCustomer
func (a *MockAPI) AddCustomer(c *sdbi.Customer, headers *api.Headers) *api.ResponseID {
	return a.MockAddCustomerResp
}

//UpdateCustomer UpdateCustomer
func (a *MockAPI) UpdateCustomer(c *sdbi.Customer, headers *api.Headers) *api.Response {
	return a.MockUpdateCustomerResp
}

//GetCustomer GetCustomer
func (a *MockAPI) GetCustomer(email string, headers *api.Headers) *sdbi.Customer {
	return a.MockCustomer
}

//GetCustomerID GetCustomerID
func (a *MockAPI) GetCustomerID(id int64, headers *api.Headers) *sdbi.Customer {
	return a.MockCustomer
}

//GetCustomerList GetCustomerList
func (a *MockAPI) GetCustomerList(start int64, end int64, headers *api.Headers) *[]sdbi.Customer {
	return a.MockCustomerList
}

//DeleteCustomer DeleteCustomer
func (a *MockAPI) DeleteCustomer(id int64, headers *api.Headers) *api.Response {
	return nil
}

//dataStoreWriteLock

//AddDataStoreWriteLock AddDataStoreWriteLock
func (a *MockAPI) AddDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *api.Headers) *api.Response {
	return nil
}

//UpdateDataStoreWriteLock UpdateDataStoreWriteLock
func (a *MockAPI) UpdateDataStoreWriteLock(w *sdbi.DataStoreWriteLock, headers *api.Headers) *api.Response {
	return nil
}

//GetDataStoreWriteLock GetDataStoreWriteLock
func (a *MockAPI) GetDataStoreWriteLock(dataStore string, headers *api.Headers) *sdbi.DataStoreWriteLock {
	return nil
}

//dataStore

//AddLocalDatastore AddLocalDatastore
func (a *MockAPI) AddLocalDatastore(d *sdbi.LocalDataStore, headers *api.Headers) *api.Response {
	return nil
}

//UpdateLocalDatastore UpdateLocalDatastore
func (a *MockAPI) UpdateLocalDatastore(d *sdbi.LocalDataStore, headers *api.Headers) *api.Response {
	return nil
}

//GetLocalDatastore GetLocalDatastore
func (a *MockAPI) GetLocalDatastore(dataStoreName string, headers *api.Headers) *sdbi.LocalDataStore {
	return nil
}

//distrubutor

//AddDistributor AddDistributor
func (a *MockAPI) AddDistributor(d *sdbi.Distributor, headers *api.Headers) *api.ResponseID {
	return a.MockAddDistributorResp
}

//UpdateDistributor UpdateDistributor
func (a *MockAPI) UpdateDistributor(d *sdbi.Distributor, headers *api.Headers) *api.Response {
	return a.MockUpdateDistributorResp
}

//GetDistributor GetDistributor
func (a *MockAPI) GetDistributor(id int64, headers *api.Headers) *sdbi.Distributor {
	return a.MockDistributor
}

//GetDistributorList GetDistributorList
func (a *MockAPI) GetDistributorList(headers *api.Headers) *[]sdbi.Distributor {
	return a.MockDistributorList
}

//DeleteDistributor DeleteDistributor
func (a *MockAPI) DeleteDistributor(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteDistributorResp
}

//excluded sub region

//AddExcludedSubRegion AddExcludedSubRegion
func (a *MockAPI) AddExcludedSubRegion(e *sdbi.ExcludedSubRegion, headers *api.Headers) *api.ResponseID {
	return a.MockAddExcludedSubRegionResp
}

//GetExcludedSubRegionList GetExcludedSubRegionList
func (a *MockAPI) GetExcludedSubRegionList(regionID int64, headers *api.Headers) *[]sdbi.ExcludedSubRegion {
	return a.MockExcludedSubRegionList
}

//DeleteExcludedSubRegion DeleteExcludedSubRegion
func (a *MockAPI) DeleteExcludedSubRegion(id int64, regionID int64, headers *api.Headers) *api.Response {
	return a.MockDeleteExcludedSubRegionResp
}

//included sub region

//AddIncludedSubRegion AddIncludedSubRegion
func (a *MockAPI) AddIncludedSubRegion(e *sdbi.IncludedSubRegion, headers *api.Headers) *api.ResponseID {
	return a.MockAddIncludedSubRegionResp
}

//GetIncludedSubRegionList GetIncludedSubRegionList
func (a *MockAPI) GetIncludedSubRegionList(regionID int64, headers *api.Headers) *[]sdbi.IncludedSubRegion {
	return a.MockIncludedSubRegionList
}

//DeleteIncludedSubRegion DeleteIncludedSubRegion
func (a *MockAPI) DeleteIncludedSubRegion(id int64, regionID int64, headers *api.Headers) *api.Response {
	return a.MockDeleteIncludedSubRegionResp
}

//instances

//AddInstance AddInstance
func (a *MockAPI) AddInstance(i *sdbi.Instances, headers *api.Headers) *api.Response {
	return nil
}

//UpdateInstance UpdateInstance
func (a *MockAPI) UpdateInstance(i *sdbi.Instances, headers *api.Headers) *api.Response {
	return nil
}

//GetInstance GetInstance
func (a *MockAPI) GetInstance(name string, dataStoreName string, headers *api.Headers) *sdbi.Instances {
	return nil
}

//GetInstanceList GetInstanceList
func (a *MockAPI) GetInstanceList(dataStoreName string, headers *api.Headers) *[]sdbi.Instances {
	return nil
}

// insurance

//AddInsurance AddInsurance
func (a *MockAPI) AddInsurance(i *sdbi.Insurance, headers *api.Headers) *api.ResponseID {
	return a.MockAddInsuranceResp
}

//UpdateInsurance UpdateInsurance
func (a *MockAPI) UpdateInsurance(i *sdbi.Insurance, headers *api.Headers) *api.Response {
	return a.MockUpdateInsuranceResp
}

//GetInsurance GetInsurance
func (a *MockAPI) GetInsurance(id int64, headers *api.Headers) *sdbi.Insurance {
	return a.MockInsurance
}

//GetInsuranceList GetInsuranceList
func (a *MockAPI) GetInsuranceList(headers *api.Headers) *[]sdbi.Insurance {
	return a.MockInsuranceList
}

//DeleteInsurance DeleteInsurance
func (a *MockAPI) DeleteInsurance(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteInsuranceResp
}

//tax rate

//AddTaxRate AddTaxRate
func (a *MockAPI) AddTaxRate(t *sdbi.TaxRate, headers *api.Headers) *api.ResponseID {
	return a.MockAddTaxRateResp
}

//UpdateTaxRate UpdateTaxRate
func (a *MockAPI) UpdateTaxRate(t *sdbi.TaxRate, headers *api.Headers) *api.Response {
	return a.MockUpdateTaxRateResp
}

//GetTaxRate GetTaxRate
func (a *MockAPI) GetTaxRate(country string, state string, headers *api.Headers) *[]sdbi.TaxRate {
	return a.MockTaxRateList
}

//GetTaxRateList GetTaxRateList
func (a *MockAPI) GetTaxRateList(headers *api.Headers) *[]sdbi.TaxRate {
	return a.MockTaxRateList
}

//DeleteTaxRate DeleteTaxRate
func (a *MockAPI) DeleteTaxRate(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteTaxRateResp
}

//order

//AddOrder AddOrder
func (a *MockAPI) AddOrder(o *sdbi.Order, headers *api.Headers) *api.ResponseID {
	return a.MockAddOrderResp
}

//UpdateOrder UpdateOrder
func (a *MockAPI) UpdateOrder(o *sdbi.Order, headers *api.Headers) *api.Response {
	return a.MockUpdateOrderResp
}

//GetOrder GetOrder
func (a *MockAPI) GetOrder(id int64, headers *api.Headers) *sdbi.Order {
	return a.MockOrder
}

//GetOrderList GetOrderList
func (a *MockAPI) GetOrderList(cid int64, headers *api.Headers) *[]sdbi.Order {
	return a.MockOrderList
}

//GetStoreOrderList GetStoreOrderList
func (a *MockAPI) GetStoreOrderList(headers *api.Headers) *[]sdbi.Order {
	return a.MockOrderList
}

//GetStoreOrderListByStatus GetStoreOrderListByStatus
func (a *MockAPI) GetStoreOrderListByStatus(status string, headers *api.Headers) *[]sdbi.Order {
	return a.MockOrderList
}

//DeleteOrder DeleteOrder
func (a *MockAPI) DeleteOrder(id int64, headers *api.Headers) *api.Response {
	return nil
}

//order comments

//AddOrderComments AddOrderComments
func (a *MockAPI) AddOrderComments(c *sdbi.OrderComment, headers *api.Headers) *api.ResponseID {
	return a.MockAddCommentResp
}

//GetOrderCommentList GetOrderCommentList
func (a *MockAPI) GetOrderCommentList(orderID int64, headers *api.Headers) *[]sdbi.OrderComment {
	return a.MockCommentList
}

//order items

//AddOrderItem AddOrderItem
func (a *MockAPI) AddOrderItem(i *sdbi.OrderItem, headers *api.Headers) *api.ResponseID {
	return a.MockAddOrderItemResp
}

//UpdateOrderItem UpdateOrderItem
func (a *MockAPI) UpdateOrderItem(i *sdbi.OrderItem, headers *api.Headers) *api.Response {
	return nil
}

//GetOrderItem GetOrderItem
func (a *MockAPI) GetOrderItem(id int64, headers *api.Headers) *sdbi.OrderItem {
	return nil
}

//GetOrderItemList GetOrderItemList
func (a *MockAPI) GetOrderItemList(orderID int64, headers *api.Headers) *[]sdbi.OrderItem {
	return a.MockOrderItemList
}

//DeleteOrderItem DeleteOrderItem
func (a *MockAPI) DeleteOrderItem(id int64, headers *api.Headers) *api.Response {
	return nil
}

//order transaction

//AddOrderTransaction AddOrderTransaction
func (a *MockAPI) AddOrderTransaction(t *sdbi.OrderTransaction, headers *api.Headers) *api.ResponseID {
	return nil
}

//GetOrderTransactionList GetOrderTransactionList
func (a *MockAPI) GetOrderTransactionList(orderID int64, headers *api.Headers) *[]sdbi.OrderTransaction {
	return nil
}

//payment gateway

//AddPaymentGateway AddPaymentGateway
func (a *MockAPI) AddPaymentGateway(pgw *sdbi.PaymentGateway, headers *api.Headers) *api.ResponseID {
	return a.MockAddPaymentGatewayResp
}

//UpdatePaymentGateway UpdatePaymentGateway
func (a *MockAPI) UpdatePaymentGateway(pgw *sdbi.PaymentGateway, headers *api.Headers) *api.Response {
	return a.MockUpdatePaymentGatewayResp
}

//GetPaymentGateway GetPaymentGateway
func (a *MockAPI) GetPaymentGateway(id int64, headers *api.Headers) *sdbi.PaymentGateway {
	return a.MockPaymentGateway
}

//GetPaymentGateways GetPaymentGateways
func (a *MockAPI) GetPaymentGateways(headers *api.Headers) *[]sdbi.PaymentGateway {
	return a.MockPaymentGatewayList
}

//DeletePaymentGateway DeletePaymentGateway
func (a *MockAPI) DeletePaymentGateway(id int64, headers *api.Headers) *api.Response {
	return a.MockDeletePaymentGatewayResp
}

//plugins

//AddPlugin AddPlugin
func (a *MockAPI) AddPlugin(p *sdbi.Plugins, headers *api.Headers) *api.ResponseID {
	return a.MockAddPluginResp
}

//UpdatePlugin UpdatePlugin
func (a *MockAPI) UpdatePlugin(p *sdbi.Plugins, headers *api.Headers) *api.Response {
	return a.MockUpdatePluginResp
}

//GetPlugin GetPlugin
func (a *MockAPI) GetPlugin(id int64, headers *api.Headers) *sdbi.Plugins {
	return a.MockPlugin
}

//GetPluginList GetPluginList
func (a *MockAPI) GetPluginList(start int64, end int64, headers *api.Headers) *[]sdbi.Plugins {
	return a.MockPluginList
}

//DeletePlugin DeletePlugin
func (a *MockAPI) DeletePlugin(id int64, headers *api.Headers) *api.Response {
	return a.MockDeletePluginResp
}

//products

//AddProduct AddProduct
func (a *MockAPI) AddProduct(p *sdbi.Product, headers *api.Headers) *api.ResponseID {
	return a.MockAddProductResp
}

//UpdateProduct UpdateProduct
func (a *MockAPI) UpdateProduct(p *sdbi.Product, headers *api.Headers) *api.Response {
	return a.MockUpdateProductResp
}

//GetProductByID GetProductByID
func (a *MockAPI) GetProductByID(id int64, headers *api.Headers) *sdbi.Product {
	return a.MockProduct
}

//GetProductBySku GetProductBySku
func (a *MockAPI) GetProductBySku(sku string, did int64, headers *api.Headers) *sdbi.Product {
	return a.MockProduct
}

//GetProductsByPromoted GetProductsByPromoted
func (a *MockAPI) GetProductsByPromoted(start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return a.MockProductList
}

//GetProductsByName GetProductsByName
func (a *MockAPI) GetProductsByName(name string, start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return a.MockProductList
}

//GetProductsByCaterory GetProductsByCaterory
func (a *MockAPI) GetProductsByCaterory(catID int64, start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return a.MockProductList
}

//GetProductList GetProductList
func (a *MockAPI) GetProductList(start int64, end int64, headers *api.Headers) *[]sdbi.Product {
	return a.MockProductList
}

//DeleteProduct DeleteProduct
func (a *MockAPI) DeleteProduct(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteProductResp
}

//product category

//AddProductCategory AddProductCategory
func (a *MockAPI) AddProductCategory(pc *sdbi.ProductCategory, headers *api.Headers) *api.Response {
	return a.MockAddProductCategoryResp
}

//GetProductCategoryList GetProductCategoryList
func (a *MockAPI) GetProductCategoryList(productID int64, headers *api.Headers) []int64 {
	return a.MockProductCategoryIDList
}

//DeleteProductCategory DeleteProductCategory
func (a *MockAPI) DeleteProductCategory(pc *sdbi.ProductCategory, headers *api.Headers) *api.Response {
	return nil
}

//region

//AddRegion AddRegion
func (a *MockAPI) AddRegion(r *sdbi.Region, headers *api.Headers) *api.ResponseID {
	return a.MockAddRegionResp
}

//UpdateRegion UpdateRegion
func (a *MockAPI) UpdateRegion(r *sdbi.Region, headers *api.Headers) *api.Response {
	return a.MockUpdateRegionResp
}

//GetRegion GetRegion
func (a *MockAPI) GetRegion(id int64, headers *api.Headers) *sdbi.Region {
	return a.MockRegion
}

//GetRegionList GetRegionList
func (a *MockAPI) GetRegionList(headers *api.Headers) *[]sdbi.Region {
	return a.MockRegionList
}

//DeleteRegion DeleteRegion
func (a *MockAPI) DeleteRegion(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteRegionResp
}

//shipment

//AddShipment AddShipment
func (a *MockAPI) AddShipment(s *sdbi.Shipment, headers *api.Headers) *api.ResponseID {
	return a.MockAddShipmentResp
}

//UpdateShipment UpdateShipment
func (a *MockAPI) UpdateShipment(s *sdbi.Shipment, headers *api.Headers) *api.Response {
	return a.MockUpdateShipmentResp
}

//GetShipment GetShipment
func (a *MockAPI) GetShipment(id int64, headers *api.Headers) *sdbi.Shipment {
	return a.MockShipment
}

//GetShipmentList GetShipmentList
func (a *MockAPI) GetShipmentList(orderID int64, headers *api.Headers) *[]sdbi.Shipment {
	return a.MockShipmentList
}

//DeleteShipment DeleteShipment
func (a *MockAPI) DeleteShipment(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteShipmentResp
}

//shipment box

//AddShipmentBox AddShipmentBox
func (a *MockAPI) AddShipmentBox(sb *sdbi.ShipmentBox, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateShipmentBox UpdateShipmentBox
func (a *MockAPI) UpdateShipmentBox(sb *sdbi.ShipmentBox, headers *api.Headers) *api.Response {
	return nil
}

//GetShipmentBox GetShipmentBox
func (a *MockAPI) GetShipmentBox(id int64, headers *api.Headers) *sdbi.ShipmentBox {
	return nil
}

//GetShipmentBoxList GetShipmentBoxList
func (a *MockAPI) GetShipmentBoxList(shipmentID int64, headers *api.Headers) *[]sdbi.ShipmentBox {
	return a.MockShipmentBoxList
}

//DeleteShipmentBox DeleteShipmentBox
func (a *MockAPI) DeleteShipmentBox(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment item

//AddShipmentItem AddShipmentItem
func (a *MockAPI) AddShipmentItem(si *sdbi.ShipmentItem, headers *api.Headers) *api.ResponseID {
	return a.MockAddShipmentItemResp
}

//UpdateShipmentItem UpdateShipmentItem
func (a *MockAPI) UpdateShipmentItem(si *sdbi.ShipmentItem, headers *api.Headers) *api.Response {
	return nil
}

//GetShipmentItem GetShipmentItem
func (a *MockAPI) GetShipmentItem(id int64, headers *api.Headers) *sdbi.ShipmentItem {
	return nil
}

//GetShipmentItemList GetShipmentItemList
func (a *MockAPI) GetShipmentItemList(shipmentID int64, headers *api.Headers) *[]sdbi.ShipmentItem {
	return a.MockShippingItemList
}

//GetShipmentItemListByBox GetShipmentItemListByBox
func (a *MockAPI) GetShipmentItemListByBox(boxNumber int64, shipmentID int64, headers *api.Headers) *[]sdbi.ShipmentItem {
	return nil
}

//DeleteShipmentItem DeleteShipmentItem
func (a *MockAPI) DeleteShipmentItem(id int64, headers *api.Headers) *api.Response {
	return nil
}

//shipment carrier

//AddShippingCarrier AddShippingCarrier
func (a *MockAPI) AddShippingCarrier(c *sdbi.ShippingCarrier, headers *api.Headers) *api.ResponseID {
	return a.MockAddShippingCarrierResp
}

//UpdateShippingCarrier UpdateShippingCarrier
func (a *MockAPI) UpdateShippingCarrier(c *sdbi.ShippingCarrier, headers *api.Headers) *api.Response {
	return a.MockUpdateShippingCarrierResp
}

//GetShippingCarrier GetShippingCarrier
func (a *MockAPI) GetShippingCarrier(id int64, headers *api.Headers) *sdbi.ShippingCarrier {
	return a.MockShippingCarrier
}

//GetShippingCarrierList GetShippingCarrierList
func (a *MockAPI) GetShippingCarrierList(headers *api.Headers) *[]sdbi.ShippingCarrier {
	return a.MockShippingCarrierList
}

//DeleteShippingCarrier DeleteShippingCarrier
func (a *MockAPI) DeleteShippingCarrier(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteShippingCarrierResp
}

//shipment method

//AddShippingMethod AddShippingMethod
func (a *MockAPI) AddShippingMethod(s *sdbi.ShippingMethod, headers *api.Headers) *api.ResponseID {
	return a.MockAddShippingMethodResp
}

//UpdateShippingMethod UpdateShippingMethod
func (a *MockAPI) UpdateShippingMethod(s *sdbi.ShippingMethod, headers *api.Headers) *api.Response {
	return a.MockUpdateShippingMethodResp
}

//GetShippingMethod GetShippingMethod
func (a *MockAPI) GetShippingMethod(id int64, headers *api.Headers) *sdbi.ShippingMethod {
	return a.MockShippingMethod
}

//GetShippingMethodList GetShippingMethodList
func (a *MockAPI) GetShippingMethodList(headers *api.Headers) *[]sdbi.ShippingMethod {
	return a.MockShippingMethodList
}

//DeleteShippingMethod DeleteShippingMethod
func (a *MockAPI) DeleteShippingMethod(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteShippingMethodResp
}

//store

//AddStore AddStore
func (a *MockAPI) AddStore(s *sdbi.Store, headers *api.Headers) *api.ResponseID {
	return nil
}

//UpdateStore UpdateStore
func (a *MockAPI) UpdateStore(s *sdbi.Store, headers *api.Headers) *api.Response {
	return nil
}

//GetStore GetStore
func (a *MockAPI) GetStore(sname string, localDomain string, headers *api.Headers) *sdbi.Store {
	return nil
}

//DeleteStore DeleteStore
func (a *MockAPI) DeleteStore(sname string, localDomain string, headers *api.Headers) *api.Response {
	return nil
}

//store plugin

//AddStorePlugin AddStorePlugin
func (a *MockAPI) AddStorePlugin(sp *sdbi.StorePlugins, headers *api.Headers) *api.ResponseID {
	return a.MockAddStorePluginResp
}

//UpdateStorePlugin UpdateStorePlugin
func (a *MockAPI) UpdateStorePlugin(sp *sdbi.StorePlugins, headers *api.Headers) *api.Response {
	return a.MockUpdateStorePluginResp
}

//GetStorePlugin GetStorePlugin
func (a *MockAPI) GetStorePlugin(id int64, headers *api.Headers) *sdbi.StorePlugins {
	return a.MockStorePlugin
}

//GetStorePluginList GetStorePluginList
func (a *MockAPI) GetStorePluginList(headers *api.Headers) *[]sdbi.StorePlugins {
	return a.MockStorePluginList
}

//DeleteStorePlugin DeleteStorePlugin
func (a *MockAPI) DeleteStorePlugin(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteStorePluginResp
}

//sub region

//AddSubRegion AddSubRegion
func (a *MockAPI) AddSubRegion(s *sdbi.SubRegion, headers *api.Headers) *api.ResponseID {
	return a.MockAddSubRegionResp
}

//UpdateSubRegion UpdateSubRegion
func (a *MockAPI) UpdateSubRegion(s *sdbi.SubRegion, headers *api.Headers) *api.Response {
	return a.MockUpdateSubRegionResp
}

//GetSubRegion GetSubRegion
func (a *MockAPI) GetSubRegion(id int64, headers *api.Headers) *sdbi.SubRegion {
	return a.MockSubRegion
}

//GetSubRegionList GetSubRegionList
func (a *MockAPI) GetSubRegionList(regionID int64, headers *api.Headers) *[]sdbi.SubRegion {
	return a.MockSubRegionList
}

//DeleteSubRegion DeleteSubRegion
func (a *MockAPI) DeleteSubRegion(id int64, headers *api.Headers) *api.Response {
	return a.MockDeleteSubRegionResp
}

//user

//AddCustomerUser AddCustomerUser
func (a *MockAPI) AddCustomerUser(u *api.User, headers *api.Headers) *api.Response {
	return a.MockAddCustomerUserRes
}

//UpdateUser UpdateUser
func (a *MockAPI) UpdateUser(u *api.User, headers *api.Headers) *api.Response {
	return a.MockUpdateUserResp
}

//GetUser GetUser
func (a *MockAPI) GetUser(u *api.User, headers *api.Headers) *api.UserResponse {
	return a.MockUser
}

//GetAdminUsers GetAdminUsers
func (a *MockAPI) GetAdminUsers(headers *api.Headers) *[]api.UserResponse {
	return nil
}

//GetCustomerUsers GetCustomerUsers
func (a *MockAPI) GetCustomerUsers(headers *api.Headers) *[]api.UserResponse {
	return nil
}

//zip code zone

//AddZoneZip AddZoneZip
func (a *MockAPI) AddZoneZip(z *sdbi.ZoneZip, headers *api.Headers) *api.ResponseID {
	return a.MockAddZoneZipResp
}

//GetZoneZipListByExclusion GetZoneZipListByExclusion
func (a *MockAPI) GetZoneZipListByExclusion(exID int64, headers *api.Headers) *[]sdbi.ZoneZip {
	return a.MockZoneZipList
}

//GetZoneZipListByInclusion GetZoneZipListByInclusion
func (a *MockAPI) GetZoneZipListByInclusion(incID int64, headers *api.Headers) *[]sdbi.ZoneZip {
	return a.MockZoneZipList
}

//DeleteZoneZip DeleteZoneZip
func (a *MockAPI) DeleteZoneZip(id int64, incID int64, exID int64, headers *api.Headers) *api.Response {
	return a.MockDeleteZoneZipResp
}
