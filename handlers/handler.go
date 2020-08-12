package handlers

import "net/http"

// import (
// 	api "github.com/Ulbora/Six910API-Go"
// 	sdbi "github.com/Ulbora/six910-database-interface"
// )

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
	//routes
	authCodeRedirectURI = "/tokenHandler"
	adminIndex          = "/admin/index"
	adminLoginFailedURL = "/admin/login?error=Login Failed"
	adminChangePassword = "/admin/changePassword"

	//pages
	adminloginPage          = "login.html"
	adminChangePwPage       = "changePassword.html"
	adminIndexPage          = "index.html"
	productFileUploadPage   = "productUpload.html"
	productUploadResultPage = "productUploadResults.html"

	authCodeState = "ghh66555h"
	storeAdmin    = "StoreAdmin"
	customerRole  = "customer"
)

//LoginError LoginError
type LoginError struct {
	Error string
}

//ClientCreds ClientCreds
type ClientCreds struct {
	AuthCodeClient string
	AuthCodeSecret string
	AuthCodeState  string
	// SchemeDefault  string // = "http://"
}

//Handler Handler
type Handler interface {
	//--- admin methods----------------------------------------------------------

	StoreAdminLogin(w http.ResponseWriter, r *http.Request)
	StoreAdminLoginNonOAuthUser(w http.ResponseWriter, r *http.Request)
	StoreAdminHandleToken(w http.ResponseWriter, r *http.Request)
	StoreAdminLogout(w http.ResponseWriter, r *http.Request)
	StoreAdminChangePassword(w http.ResponseWriter, r *http.Request)
	StoreAdminChangeUserPassword(w http.ResponseWriter, r *http.Request)

	StoreAdminIndex(w http.ResponseWriter, r *http.Request)

	//products
	StoreAdminUploadProductFilePage(w http.ResponseWriter, r *http.Request)
	StoreAdminUploadProductFile(w http.ResponseWriter, r *http.Request)

	StoreAdminAddProductPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddProduct(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditProductPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditProduct(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewProductList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteProduct(w http.ResponseWriter, r *http.Request)

	// //orders
	// StoreAdminEditOrderPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditOrder(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewOrderList(w http.ResponseWriter, r *http.Request)

	// //shipments
	// StoreAdminAddShipmentPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddShipment(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditShipmentPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditShipment(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewShipmentList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteShipment(w http.ResponseWriter, r *http.Request)

	// //customers
	// StoreAdminEditCustomerPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCustomer(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCustomerUserPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCustomerUser(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewCustomerList(w http.ResponseWriter, r *http.Request)

	// //categories
	// StoreAdminAddCategoryPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddCategory(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCategoryPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCategory(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewCategoryList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteCategory(w http.ResponseWriter, r *http.Request)

	// //distributors
	// StoreAdminAddDistributorPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddDistributor(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditDistributorPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditDistributor(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewDistributorList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteDistributor(w http.ResponseWriter, r *http.Request)

	// //insurance
	// StoreAdminAddInsurancePage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddInsurance(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditInsurancePage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditInsurance(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewInsuranceList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteInsurance(w http.ResponseWriter, r *http.Request)

	// //payment gateway
	// StoreAdminAddPaymentGatewayPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddPaymentGateway(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditPaymentGatewayPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditPaymentGateway(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewPaymentGatewayList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeletePaymentGateway(w http.ResponseWriter, r *http.Request)

	// //plugins
	// StoreAdminAddPluginPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddPlugin(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditPluginPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditPlugin(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewPluginList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeletePlugin(w http.ResponseWriter, r *http.Request)

	// //shipment carriers
	// StoreAdminAddCarrierPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddCarrier(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCarrierPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditCarrier(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewCarrierList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteCarrier(w http.ResponseWriter, r *http.Request)

	// //shipping methods
	// StoreAdminAddShippingMethodPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddShippingMethod(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditShippingMethodPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditShippingMethod(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewShippingMethodList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteShippingMethod(w http.ResponseWriter, r *http.Request)

	// //regions
	// StoreAdminAddRegionPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminAddRegion(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditRegionPage(w http.ResponseWriter, r *http.Request)
	// StoreAdminEditRegion(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewRegionList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteRegion(w http.ResponseWriter, r *http.Request)

	// //abandoned carts
	// StoreAdminViewCart(w http.ResponseWriter, r *http.Request)
	// StoreAdminViewCartList(w http.ResponseWriter, r *http.Request)
	// StoreAdminDeleteCart(w http.ResponseWriter, r *http.Request)

	// //---customer methods-------------------------------------------------------------

	// Index(w http.ResponseWriter, r *http.Request)

	// //products
	// ViewProductList(w http.ResponseWriter, r *http.Request)
	// SearchProductList(w http.ResponseWriter, r *http.Request)
	// ViewProduct(w http.ResponseWriter, r *http.Request)

	// //cart
	// AddProductToCart(w http.ResponseWriter, r *http.Request)
	// UpdateProductToCart(w http.ResponseWriter, r *http.Request)
	// CheckOut(w http.ResponseWriter, r *http.Request)

	// //customer
	// CreateCustomerAccount(w http.ResponseWriter, r *http.Request)
	// UpdateCustomerAccount(w http.ResponseWriter, r *http.Request)

	// CustomerLogin(w http.ResponseWriter, r *http.Request)
	// CustomerLogout(w http.ResponseWriter, r *http.Request)
	// CustomerChangePassword(w http.ResponseWriter, r *http.Request)

	// //orders
	// ViewCustomerOrder(w http.ResponseWriter, r *http.Request)
	// ViewCustomerOrderList(w http.ResponseWriter, r *http.Request)
}
