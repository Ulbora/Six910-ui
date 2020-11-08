package handlers

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
	adminIndex          = "/admin"
	adminLogin          = "/admin/login"
	adminLoginFailedURL = "/admin/login?error=Login Failed"
	adminChangePassword = "/admin/changePassword"

	//routes product upload
	adminProdUploadView = "/admin/productUploadView"
	adminProdUpload     = "/admin/productUpload"

	//routes product
	adminAddProdView         = "/admin/addProdView"
	adminAddProdViewFail     = "/admin/addProdView?error=Add Failed"
	adminAddProduct          = "/admin/addProduct"
	adminEditProdView        = "/admin/editProductView"
	adminEditProdViewFail    = "/admin/editProductView?error=Update Failed"
	adminEditProd            = "/admin/editProduct"
	adminDeleteProd          = "/admin/deleteProduct"
	adminProductListView     = "/admin/productListView"
	adminProductList         = "/admin/productList/0/100"
	adminProductListError    = "/admin/productList/0/100?error=Update Failed"
	adminProductListViewFail = "/admin/productListView?error=Add Failed"

	//routes shipment
	adminAddShipmentView      = "/admin/addShipmentView"
	adminAddShipmentViewFail  = "/admin/addShipmentView?error=Add Failed"
	adminAddShipment          = "/admin/addShipment"
	adminEditShipmentView     = "/admin/editShipmentView"
	adminEditShipmentViewFail = "/admin/editShipmentView?error=Update Failed"
	adminEditShipment         = "/admin/editShipment"
	adminDeleteShipment       = "/admin/deleteShipment"
	adminShipmentListView     = "/admin/shipmentListView"
	adminShipmentListViewFail = "/admin/shipmentListView?error=Add Failed"

	//routes order
	adminEditOrderView     = "/admin/editOrderView"
	adminEditOrderViewFail = "/admin/editOrderView?error=Update Failed"
	adminEditOrder         = "/admin/editOrder"
	adminOrderListView     = "/admin/orderList/New"
	adminOrderListViewFail = "/admin/orderList?error=Update Failed"

	//routes customer
	adminEditCustomerView         = "/admin/editCustomerView"
	adminEditCustomerViewFail     = "/admin/editCustomerView?error=Update Failed"
	adminEditCustomer             = "/admin/editCustomer"
	adminCustomerListView         = "/admin/customerList"
	adminCustomerListViewFail     = "/admin/customerList?error=Update Failed"
	adminEditCustomerUserView     = "/admin/editCustomerUserView"
	adminEditCustomerUserViewFail = "/admin/editCustomerUserView?error=Update Failed"
	adminEditCustomerUser         = "/admin/editCustomerUser"

	//routes category
	adminAddCategoryView      = "/admin/addCategoryView"
	adminAddCategoryViewFail  = "/admin/addCategoryView?error=Add Failed"
	adminAddCategory          = "/admin/addCategory"
	adminEditCategoryView     = "/admin/editCategoryView"
	adminEditCategoryViewFail = "/admin/editCategoryView?error=Update Failed"
	adminEditCategory         = "/admin/editCategory"
	adminDeleteCategory       = "/admin/deleteCategory"
	adminCategoryListView     = "/admin/categoryList"
	adminCategoryListViewFail = "/admin/categoryList?error=Add Failed"

	//routes Distributor
	adminAddDistributorView      = "/admin/addDistributorView"
	adminAddDistributorViewFail  = "/admin/addDistributorView?error=Add Failed"
	adminAddDistributor          = "/admin/addDistributor"
	adminEditDistributorView     = "/admin/editDistributorView"
	adminEditDistributorViewFail = "/admin/editDistributorView?error=Update Failed"
	adminEditDistributor         = "/admin/editDistributor"
	adminDeleteDistributor       = "/admin/deleteDistributor"
	adminDistributorListView     = "/admin/distributorList"
	adminDistributorListViewFail = "/admin/distributorList?error=Add Failed"

	//routes Insurance
	adminAddInsuranceView      = "/admin/addInsuranceView"
	adminAddInsuranceViewFail  = "/admin/addInsuranceView?error=Add Failed"
	adminAddInsurance          = "/admin/addInsurance"
	adminEditInsuranceView     = "/admin/editInsuranceView"
	adminEditInsuranceViewFail = "/admin/editInsuranceView?error=Update Failed"
	adminEditInsurance         = "/admin/editInsurance"
	adminDeleteInsurance       = "/admin/deleteInsurance"
	adminInsuranceList         = "/admin/insuranceList"
	adminInsuranceListFail     = "/admin/insuranceList?error=Add Failed"

	//routes TaxRate
	adminAddTaxRateView      = "/admin/addTaxRate"
	adminAddTaxRateViewFail  = "/admin/addTaxRate?error=Add Failed"
	adminAddTaxRate          = "/admin/addTaxRate"
	adminEditTaxRateView     = "/admin/editTaxRate"
	adminEditTaxRateViewFail = "/admin/editTaxRate?error=Update Failed"
	adminEditTaxRate         = "/admin/editTaxRate"
	adminDeleteTaxRate       = "/admin/deleteTaxRate"
	adminTaxRateListView     = "/admin/taxRateList"
	adminTaxRateListViewFail = "/admin/taxRateList?error=Add Failed"

	//routes Payment Gateway
	adminAddPaymentGatewayView      = "/admin/addPaymentGatewayView"
	adminAddPaymentGatewayViewFail  = "/admin/addPaymentGatewayView?error=Add Failed"
	adminAddPaymentGateway          = "/admin/addPaymentGateway"
	adminEditPaymentGatewayView     = "/admin/editPaymentGatewayView"
	adminEditPaymentGatewayViewFail = "/admin/editPaymentGatewayView?error=Update Failed"
	adminEditPaymentGateway         = "/admin/editPaymentGateway"
	adminDeletePaymentGateway       = "/admin/deletePaymentGateway"
	adminPaymentGatewayListView     = "/admin/paymentGatewayList"
	adminPaymentGatewayListViewFail = "/admin/paymentGatewayList?error=Add Failed"

	//routes Payment Gateway
	adminAddPluginView      = "/admin/addPluginView"
	adminAddPluginViewFail  = "/admin/addPluginView?error=Add Failed"
	adminAddPlugin          = "/admin/addPlugin"
	adminEditPluginView     = "/admin/editPluginView"
	adminEditPluginViewFail = "/admin/editPluginView?error=Update Failed"
	adminEditPlugin         = "/admin/editPlugin"
	adminDeletePlugin       = "/admin/deletePlugin"
	adminPluginListView     = "/admin/pluginList"
	adminPluginListViewFail = "/admin/pluginList?error=Add Failed"

	//routes Payment Gateway
	adminAddStorePluginView      = "/admin/addStorePluginView"
	adminAddStorePluginViewFail  = "/admin/addStorePluginView?error=Add Failed"
	adminAddStorePlugin          = "/admin/addStorePlugin"
	adminEditStorePluginView     = "/admin/editStorePluginView"
	adminEditStorePluginViewFail = "/admin/editStorePluginView?error=Update Failed"
	adminEditStorePlugin         = "/admin/editStorePlugin"
	adminDeleteStorePlugin       = "/admin/deleteStorePlugin"
	adminStorePluginListView     = "/admin/storePluginList"
	adminStorePluginListViewFail = "/admin/storePluginList?error=Add Failed"

	//routes shipping carrier
	adminAddShippingCarrierView      = "/admin/addShippingCarrierView"
	adminAddShippingCarrierViewFail  = "/admin/addShippingCarrierView?error=Add Failed"
	adminAddShippingCarrier          = "/admin/addShippingCarrier"
	adminEditShippingCarrierView     = "/admin/editShippingCarrierView"
	adminEditShippingCarrierViewFail = "/admin/editShippingCarrierView?error=Update Failed"
	adminEditShippingCarrier         = "/admin/editShippingCarrier"
	adminDeleteShippingCarrier       = "/admin/deleteShippingCarrier"
	adminShippingCarrierList         = "/admin/shippingCarrierList"
	adminShippingCarrierListFail     = "/admin/shippingCarrierList?error=Add Failed"

	//routes shipping method
	adminAddShippingMethodView      = "/admin/addShippingMethodView"
	adminAddShippingMethodViewFail  = "/admin/addShippingMethodView?error=Add Failed"
	adminAddShippingMethod          = "/admin/addShippingMethod"
	adminEditShippingMethodView     = "/admin/editShippingMethodView"
	adminEditShippingMethodViewFail = "/admin/editShippingMethodView?error=Update Failed"
	adminEditShippingMethod         = "/admin/editShippingMethod"
	adminDeleteShippingMethod       = "/admin/deleteShippingMethod"
	adminShippingMethodList         = "/admin/shippingMethodList"
	adminShippingMethodListFail     = "/admin/shippingMethodList?error=Add Failed"

	//routes Region
	adminAddRegionView      = "/admin/addRegionView"
	adminAddRegionViewFail  = "/admin/addRegionView?error=Add Failed"
	adminAddRegion          = "/admin/addRegion"
	adminEditRegionView     = "/admin/editRegionView"
	adminEditRegionViewFail = "/admin/editRegionView?error=Update Failed"
	adminEditRegion         = "/admin/editRegion"
	adminDeleteRegion       = "/admin/deleteRegion"
	adminRegionListView     = "/admin/shippingRegionList"
	adminRegionListViewFail = "/admin/shippingRegionList?error=Add Failed"

	//routes Sub Region
	adminAddSubRegionView      = "/admin/addSubRegionView"
	adminAddSubRegionViewFail  = "/admin/addSubRegionView?error=Add Failed"
	adminAddSubRegion          = "/admin/addSubRegion"
	adminEditSubRegionView     = "/admin/editSubRegionView"
	adminEditSubRegionViewFail = "/admin/editSubRegionView?error=Update Failed"
	adminEditSubRegion         = "/admin/editSubRegion"
	adminDeleteSubRegion       = "/admin/deleteSubRegion"
	adminSubRegionListView     = "/admin/subRegionView"
	adminSubRegionListViewFail = "/admin/subRegionListView?error=Add Failed"

	//routes Ex Sub Region
	adminAddExSubRegionView      = "/admin/addExcludedSubRegionView"
	adminAddExSubRegionViewFail  = "/admin/addExcludedSubRegionView?error=Add Failed"
	adminAddExSubRegion          = "/admin/addExcludedSubRegion"
	adminEditExubRegionView      = "/admin/editExcludedSubRegionView"
	adminEditExSubRegionViewFail = "/admin/editExcludedSubRegionView?error=Update Failed"
	adminEditExSubRegion         = "/admin/editExcludedSubRegion"
	adminDeleteExSubRegion       = "/admin/deleteExcludedSubRegion"
	adminExSubRegionListView     = "/admin/excludedSubRegionView"
	adminExSubRegionListViewFail = "/admin/excludedSubRegionListView?error=Add Failed"

	//routes Inc Sub Region
	adminAddIncSubRegionView      = "/admin/addIncludedSubRegionView"
	adminAddIncSubRegionViewFail  = "/admin/addIncludedSubRegionView?error=Add Failed"
	adminAddIncSubRegion          = "/admin/addIncludedSubRegion"
	adminEditIncSubRegionView     = "/admin/editIncludedSubRegionView"
	adminEditIncSubRegionViewFail = "/admin/editIncludedSubRegionView?error=Update Failed"
	adminEditIncSubRegion         = "/admin/editIncludedSubRegion"
	adminDeleteIncSubRegion       = "/admin/deleteIncludedSubRegion"
	adminIncSubRegionListView     = "/admin/includedSubRegionView"
	adminIncSubRegionListViewFail = "/admin/includedSubRegionListView?error=Add Failed"

	//routes Inc Sub Region
	adminAddZipZoneView         = "/admin/addZipZoneView"
	adminAddZipZoneViewFail     = "/admin/addZipZoneView?error=Add Failed"
	adminAddZipZoneRegion       = "/admin/addZipZone"
	adminDeleteZipZoneRegion    = "/admin/deleteZipZone"
	adminIncZipZoneListView     = "/admin/includedZipZone"
	adminIncZipZoneListViewFail = "/admin/zipZoneListView?error=Add Failed"

	//content
	adminAddContent     = "/admin/addContent"
	adminAddContentFail = "/admin/addContent?error=Add Failed"
	adminGetContent     = "/admin/getContent"
	adminContentList    = "/admin/contentList"

	adminImageList     = "/admin/imageList"
	adminImageListFail = "/admin/imageList?error=Add Failed"

	//menu
	adminAddMenu     = "/admin/addMenu"
	adminAddMenuFail = "/admin/addMenu?error=Add Failed"
	adminGetMenu     = "/admin/getMenu"
	adminMenuList    = "/admin/menuList"

	//routes customer
	customerIndexView               = "/"
	customerIndexViewUserFail       = "/?error=user_update_failed"
	customerIndexViewAddressFail    = "/?error=address_update_failed"
	customerIndexViewPasswordFail   = "/?error=password_change_failed"
	customerIndexViewAddressAddFail = "/?error=address_add_failed"

	//routes shopping cart
	customerLoginView            = "/customerLoginPage"
	customerLoginViewFail        = "/customerLoginPage?error=Login Failed"
	customerShoppingCartView     = "/shoppingCartView"
	customerShoppingCartViewFail = "/shoppingCartView?error=Add Failed"

	//routes customer
	createCustomerView      = "/register"
	createCustomerViewFail  = "/register?error=Customer Email Already Used"
	createCustomerViewError = "/register?error=Customer Create Error"
	createCustomer          = "/createCustomer"
	customerInfoView        = "/customerInfo"
	customerInfoViewFail    = "/customerInfo?error=Add Failed"
	customerOrderView       = "/customerOrder"
	customerOrderListView   = "/customerOrderList"
	//adminDeleteZipZoneRegion    = "/admin/deleteZipZone"
	//adminIncZipZoneListView     = "/admin/includedZipZone"
	//adminIncZipZoneListViewFail = "/admin/zipZoneListView?error=Add Failed"

	//pages
	adminloginPage    = "login.html"
	adminChangePwPage = "changePassword.html"
	adminIndexPage    = "index.html"

	//pages product upload
	productFileUploadPage   = "productUpload.html"
	productUploadResultPage = "productUploadResults.html"

	//pages product
	adminAddProductPage        = "addProduct.html"
	adminEditProductPage       = "editProduct.html"
	adminProductSkuSearchPage  = "productSkuSearch.html"
	adminProductNameSearchPage = "productNameSearch.html"
	adminProductCatSearchPage  = "productCatSearch.html"
	adminProductListPage       = "productList.html"

	//pages product
	adminAddShipmentPage  = "addShipment.html"
	adminEditShipmentPage = "editShipment.html"
	adminShipmentListPage = "shipmentList.html"

	//pages order
	adminEditOrderPage = "editOrder.html"
	adminOrderListPage = "orderList.html"

	//pages customer
	adminEditCustomerPage        = "editCustomer.html"
	adminEditCustomerUserPage    = "editCustomerUser.html"
	adminCustomerListPage        = "customerList.html"
	adminCustomerEmailSearchPage = "customerEmailSearch.html"

	//pages Distributor
	adminAddDistributorPage  = "addDistributor.html"
	adminEditDistributorPage = "editDistributor.html"
	adminDistributorListPage = "distributorList.html"

	//pages product
	adminAddCategoryPage  = "addCategory.html"
	adminEditCategoryPage = "editCategory.html"
	adminCategoryListPage = "categoryList.html"

	//pages Insurance
	adminAddInsurancePage  = "addInsurance.html"
	adminEditInsurancePage = "editInsurance.html"
	adminInsuranceListPage = "insuranceList.html"

	//pages TaxRate
	adminAddTaxRatePage  = "addTaxRate.html"
	adminEditTaxRatePage = "editTaxRate.html"
	adminTaxRateListPage = "taxRateList.html"

	//pages Payment Gateway
	adminAddPaymentGatewayPage  = "addPaymentGateway.html"
	adminEditPaymentGatewayPage = "editPaymentGateway.html"
	adminPaymentGatewayListPage = "paymentGatewayList.html"

	//pages Plugin
	adminAddPluginPage  = "addPlugin.html"
	adminEditPluginPage = "editPlugin.html"
	adminPluginListPage = "pluginList.html"

	//pages store Plugin
	adminAddStorePluginPage         = "addStorePlugin.html"
	adminAddStorePluginFromListPage = "addStorePluginFromList.html"
	adminStorePlugintoAddPage       = "pluginToAdd.html"
	adminEditStorePluginPage        = "editStorePlugin.html"
	adminStorePluginListPage        = "storePluginList.html"

	//pages shipping carrier
	adminAddShippingCarrierPage  = "addShippingCarrier.html"
	adminEditShippingCarrierPage = "editShippingCarrier.html"
	adminShippingCarrierListPage = "shippingCarrierList.html"

	//pages shipping method
	adminAddShippingMethodPage  = "addShippingMethod.html"
	adminEditShippingMethodPage = "editShippingMethod.html"
	adminShippingMethodListPage = "shippingMethodList.html"

	//pages region
	adminAddRegionPage  = "addRegion.html"
	adminEditRegionPage = "editRegion.html"
	adminRegionListPage = "regionList.html"

	//pages sub region
	adminAddSubRegionPage  = "addSubRegion.html"
	adminEditSubRegionPage = "editSubRegion.html"
	adminSubRegionListPage = "subRegionList.html"

	//pages ex sub region
	adminAddExSubRegionPage  = "addExcludedSubRegion.html"
	adminEditExSubRegionPage = "editExcludedSubRegion.html"
	adminExSubRegionListPage = "excludedSubRegionList.html"

	//pages Inc sub region
	adminAddIncSubRegionPage  = "addIncludedSubRegion.html"
	adminEditIncSubRegionPage = "editIncludedSubRegion.html"
	adminIncSubRegionListPage = "includedSubRegionList.html"

	//pages Inc sub region
	adminAddZipZonePage  = "addZipZone.html"
	adminZipZoneListPage = "zipZoneList.html"

	adminEditStorePage = "editStore.html"

	imageFilesUploadPage     = "imageFilesUpload.html"
	thumbnailFilesUploadPage = "thumbnailFilesUpload.html"

	adminAddContentPage  = "addContent.html"
	adminUpdateContent   = "updateContent.html"
	adminContentListPage = "contentList.html"

	adminImageUploadPage = "imageUpload.html"
	adminImageListPage   = "imageList.html"

	adminAddMenuPage  = "addMenu.html"
	adminUpdateMenu   = "editMenu.html"
	adminMenuListPage = "menuList.html"

	adminEditPageCSS = "editPageCss.html"

	adminEditCarousel = "editCarousel.html"

	//customer pages
	customerLoginPage                = "login.html"
	customerIndexPage                = "index.html"
	customerProductListPage          = "productList.html"
	customerProductByCatPage         = "productsByCategory.html"
	customerProductsSearchListPage   = "productsSearch.html"
	customerProductPage              = "product.html"
	customerShoppingCartPage         = "shoppingCart.html"
	customerShoppingCartPage2        = "shoppingCart2.html"
	customerShoppingCartContinuePage = "shoppingCartContinue.html"
	checkoutReceiptPage              = "checkoutReceipt.html"
	customerCreatePage               = "createCustomer.html"
	customerUpdatePage               = "updateCustomer.html"
	customerCreateAddressPage        = "createCustomerAddress.html"
	customerInfoPage                 = "createInfo.html"
	customerOrderPage                = "customerOrder.html"
	customerOrderListPage            = "customerOrderList.html"

	authCodeState = "ghh66555h"
	storeAdmin    = "StoreAdmin"
	customerRole  = "customer"

	billingAddressType  = "Billing"
	shippingAddressType = "Shipping"

	timeFormat = "2006-01-02 15:04:05"

	//content pages
	indexContent               = "index"
	productListContent         = "productList"
	productCategoryListContent = "productCategoryList"
	productContent             = "product"
	shoppingCartContent        = "shoppingCart"
	shoppingCartContent2       = "shoppingCart2"
	shoppingCartContent3       = "shoppingCartContinue"

	orderListContent = "orderList"

	usernameForAddedNotes = "System Admin"

	mailMessageUploadComplete = "Product file upload completed with a product upload count of: "
)
