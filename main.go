package main

/*
 Six910-ui is a shopping cart and E-commerce system.
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

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	lg "github.com/Ulbora/Level_Logger"
	carsrv "github.com/Ulbora/Six910-ui/carouselsrv"
	csrv "github.com/Ulbora/Six910-ui/contentsrv"
	cntrysrv "github.com/Ulbora/Six910-ui/countrysrv"
	csssrv "github.com/Ulbora/Six910-ui/csssrv"
	hand "github.com/Ulbora/Six910-ui/handlers"
	isrv "github.com/Ulbora/Six910-ui/imgsrv"
	m "github.com/Ulbora/Six910-ui/managers"
	musrv "github.com/Ulbora/Six910-ui/menusrv"
	stsrv "github.com/Ulbora/Six910-ui/statesrv"
	tmpsrv "github.com/Ulbora/Six910-ui/templatesrv"
	api "github.com/Ulbora/Six910API-Go"
	ml "github.com/Ulbora/go-mail-sender"
	ds "github.com/Ulbora/json-datastore"
	"github.com/gorilla/mux"
)

func main() {
	// just the start of Six910-ui Server
	// This is the storefront for Six910-ui.
	var apiURL string
	var storeName string
	var localDomain string
	var apiKey string

	var mailHost string
	var mailUser string
	var mailPassword string
	var mailPort string
	var mailSenderAddress string
	var mailSubject string
	var schemeDefault string

	var mailSubjectOrderReceived string
	var mailBodyOrderReceived string

	var mailSubjectOrderProcessing string
	var mailBodyOrderProcessing string

	var mailSubjectOrderShipped string
	var mailBodyOrderShipped string

	var companyName string
	var six910CartSite string

	if os.Getenv("API_URL") != "" {
		apiURL = os.Getenv("API_URL")
	} else {
		apiURL = "http://localhost:3002"
	}

	if os.Getenv("STORE_NAME") != "" {
		storeName = os.Getenv("STORE_NAME")
	} else {
		storeName = "defaultLocalStore"
	}

	if os.Getenv("LOCAL_DOMAIN") != "" {
		localDomain = os.Getenv("LOCAL_DOMAIN")
	} else {
		localDomain = "defaultLocalStore.mydomain.com"
	}

	if os.Getenv("API_KEY") != "" {
		apiKey = os.Getenv("API_KEY")
	} else {
		apiKey = "GDG651GFD66FD16151sss651f651ff65555ddfhjklyy5"
	}

	if os.Getenv("EMAIL_HOST") != "" {
		mailHost = os.Getenv("EMAIL_HOST")
	}

	if os.Getenv("EMAIL_USER") != "" {
		mailUser = os.Getenv("EMAIL_USER")
	}

	if os.Getenv("EMAIL_PASSWORD") != "" {
		mailPassword = os.Getenv("EMAIL_PASSWORD")
	}

	if os.Getenv("EMAIL_PORT") != "" {
		mailPort = os.Getenv("EMAIL_PORT")
	}

	if os.Getenv("MAIL_SENDER_ADDRESS") != "" {
		mailSenderAddress = os.Getenv("MAIL_SENDER_ADDRESS")
	}

	if os.Getenv("MAIL_SUBJECT") != "" {
		mailSubject = os.Getenv("MAIL_SUBJECT")
	} else {
		mailSubject = "Six910 Shopping Cart Message"
	}

	//------------order email messages----------------------------------------------------

	if os.Getenv("MAIL_SUBJECT_ORDER_RECEIVED") != "" {
		mailSubjectOrderReceived = os.Getenv("MAIL_SUBJECT_ORDER_RECEIVED")
	} else {
		mailSubjectOrderReceived = "Six910 Shopping Cart New Order Received"
	}

	if os.Getenv("MAIL_BODY_ORDER_RECEIVED") != "" {
		mailBodyOrderReceived = os.Getenv("MAIL_BODY_ORDER_RECEIVED")
	} else {
		mailBodyOrderReceived = "New Order # %s received from %s"
	}

	if os.Getenv("MAIL_SUBJECT_ORDER_PROCESSING") != "" {
		mailSubjectOrderProcessing = os.Getenv("MAIL_SUBJECT_ORDER_PROCESSING")
	} else {
		mailSubjectOrderProcessing = "%s order # %s is processing"
	}

	if os.Getenv("MAIL_BODY_ORDER_PROCESSING") != "" {
		mailBodyOrderProcessing = os.Getenv("MAIL_BODY_ORDER_PROCESSING")
	} else {
		mailBodyOrderProcessing = "%s new order # %s received and processing. " +
			"You will receive an email when the order ships."
	}

	if os.Getenv("MAIL_SUBJECT_ORDER_SHIPPED") != "" {
		mailSubjectOrderShipped = os.Getenv("MAIL_SUBJECT_ORDER_SHIPPED")
	} else {
		mailSubjectOrderShipped = "%s order # %s has shipped"
	}

	if os.Getenv("MAIL_BODY_ORDER_SHIPPED") != "" {
		mailBodyOrderShipped = os.Getenv("MAIL_BODY_ORDER_SHIPPED")
	} else {
		mailBodyOrderShipped = "%s order # %s has shipped. " +
			"Click the link in this email to see the tracking information."
	}

	//------------order email messages-----------------------------------------------

	if os.Getenv("COMPANY_NAME") != "" {
		companyName = os.Getenv("COMPANY_NAME")
	} else {
		companyName = "Six910 Shopping Cart"
	}

	if os.Getenv("SIX910_CART_SITE") != "" {
		six910CartSite = os.Getenv("SIX910_CART_SITE")
	}
	if os.Getenv("SCHEME_DEFAULT") != "" {
		schemeDefault = os.Getenv("SCHEME_DEFAULT")
	} else {
		schemeDefault = "http://"
	}

	var sapi api.Six910API
	sapi.SetAPIKey(apiKey)
	sapi.SetRestURL(apiURL)
	sapi.SetStore(storeName, localDomain)

	var ms ml.SecureSender
	ms.MailHost = mailHost
	ms.User = mailUser
	ms.Password = mailPassword
	ms.Port = mailPort

	var sh hand.Six910Handler
	sh.SchemeDefault = schemeDefault
	sh.Six910SiteURL = six910CartSite
	sh.CompanyName = companyName
	sh.LocalDomain = localDomain
	sh.StoreName = storeName
	sh.API = sapi.GetNew()
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.MailSender = &ms
	sh.MailSenderAddress = mailSenderAddress
	sh.MailSubject = mailSubject

	sh.MailSubjectOrderReceived = mailSubjectOrderReceived
	sh.MailBodyOrderReceived = mailBodyOrderReceived
	sh.MailSubjectOrderProcessing = mailSubjectOrderProcessing
	sh.MailBodyOrderProcessing = mailBodyOrderProcessing
	sh.MailSubjectOrderShipped = mailSubjectOrderShipped
	sh.MailBodyOrderShipped = mailBodyOrderShipped

	sh.ImagePath = "./static/images"
	sh.ThumbnailPath = "./static/thumbnail"

	var ts tmpsrv.Six910TemplateService
	ts.TemplateStorePath = "./data/templateStore"
	ts.Log = &l
	var tds ds.DataStore
	tds.Path = "./data/templateStore"
	ts.TemplateStore = tds.GetNew()
	sh.TemplateService = ts.GetNew()
	sh.ActiveTemplateLocation = "./static/templates"

	var sms musrv.Six910MenuService
	sms.Log = &l
	var mds ds.DataStore
	mds.Path = "./data/menuStore"
	sms.MenuStore = mds.GetNew()
	sh.MenuService = sms.GetNew()

	var ccs csrv.CmsService
	ccs.Log = &l

	var cds ds.DataStore
	cds.Path = "./data/contentStore"
	ccs.Store = cds.GetNew()
	ccs.ContentStorePath = "./data/contentStore"

	sh.ContentService = ccs.GetNew()

	var iss isrv.Six910ImageService
	iss.ImagePath = "./static/img"
	iss.ImageFullPath = "./static/img"
	iss.Log = &l

	sh.ImageService = iss.GetNew()

	var css csssrv.Six910CSSService
	css.CSSStorePath = "./data/cssStore"
	css.Log = &l
	var csds ds.DataStore
	csds.Path = "./data/cssStore"
	css.CSSStore = csds.GetNew()
	sh.CSSService = css.GetNew()

	var cars carsrv.Six910CarouselService
	cars.StorePath = "./data/carouselStore"
	cars.Log = &l
	var cards ds.DataStore
	cards.Path = "./data/carouselStore"
	cars.Store = cards.GetNew()
	sh.CarouselService = cars.GetNew()

	var st stsrv.Six910StateService
	var sds ds.DataStore
	sds.Path = "./data/stateStore"
	st.Log = &l
	st.StateStore = sds.GetNew()

	sh.StateService = st.GetNew()

	var cntry cntrysrv.Six910CountryService
	var cntds ds.DataStore
	cntds.Path = "./data/countryStore"
	//ds.Delete("books1")
	cntry.Log = &l
	cntry.CountryStore = cntds.GetNew()

	sh.CountryService = cntry.GetNew()

	sh.AdminTemplates = template.Must(template.ParseFiles("./static/admin/index.html", "./static/admin/head.html",
		"./static/admin/login.html", "./static/admin/navbar.html", "./static/admin/productList.html",
		"./static/admin/subnavs/productNavbar.html", "./static/admin/pagination.html", "./static/admin/productSkuSearch.html",
		"./static/admin/productNameSearch.html", "./static/admin/editProduct.html", "./static/admin/addProduct.html",
		"./static/admin/productCatSearch.html",
		"./static/admin/distributorList.html", "./static/admin/editDistributor.html", "./static/admin/categoryList.html",
		"./static/admin/editCategory.html", "./static/admin/shippingCarrierList.html", "./static/admin/editShippingCarrier.html",
		"./static/admin/regionList.html", "./static/admin/editRegion.html",
		"./static/admin/shippingMethodList.html", "./static/admin/editShippingMethod.html",
		"./static/admin/insuranceList.html", "./static/admin/editInsurance.html",
		"./static/admin/taxRateList.html", "./static/admin/editTaxRate.html",
		"./static/admin/pluginList.html", "./static/admin/editPlugin.html",
		"./static/admin/storePluginList.html", "./static/admin/addStorePluginFromList.html",
		"./static/admin/pluginToAdd.html", "./static/admin/editStorePlugin.html",
		"./static/admin/paymentGatewayList.html", "./static/admin/editPaymentGateway.html",
		"./static/admin/orderList.html", "./static/admin/editOrder.html",
		"./static/admin/customerList.html", "./static/admin/customerEmailSearch.html", "./static/admin/editCustomer.html",
		"./static/admin/productUpload.html", "./static/admin/editStore.html",
		"./static/admin/imageFilesUpload.html", "./static/admin/thumbnailFilesUpload.html",
		"./static/admin/addContent.html", "./static/admin/contentList.html",
		"./static/admin/updateContent.html", "./static/admin/imageList.html",
		"./static/admin/imageUpload.html", "./static/admin/menuList.html",
		"./static/admin/editMenu.html", "./static/admin/addMenu.html",
		"./static/admin/editPageCss.html", "./static/admin/editCarousel.html",
		// "./static/admin/footer.html", "./static/admin/navbar.html", "./static/admin/contentNavbar.html",
	// "./static/admin/addContent.html", "./static/admin/images.html", "./static/admin/templates.html",
	// "./static/admin/updateContent.html", "./static/admin/mailServer.html", "./static/admin/templateUpload.html",
	// "./static/admin/imageUpload.html", "./static/admin/login.html", "./static/admin/backups.html",
	// "./static/admin/backupUpload.html",
	))

	var man m.Six910Manager
	man.API = &sapi
	man.Log = &l
	sh.Manager = man.GetNew()

	sh.LoadTemplate()

	h := sh.GetNew()

	fmt.Println("Six910 (six nine ten) UI is running on port 8080!")
	router := mux.NewRouter()

	//site pages
	router.HandleFunc("/", h.Index).Methods("GET")
	router.HandleFunc("/customerLoginPage", h.CustomerLoginPage).Methods("GET")
	router.HandleFunc("/customerLogin", h.CustomerLogin).Methods("POST")
	router.HandleFunc("/register", h.CreateCustomerAccountPage).Methods("GET")
	router.HandleFunc("/customerLogout", h.CustomerLogout).Methods("GET")

	router.HandleFunc("/viewProduct/{id}", h.ViewProduct).Methods("GET")
	router.HandleFunc("/productByCategoryList/{catId}/{catName}/{start}/{end}", h.ViewProductByCatList).Methods("GET")
	router.HandleFunc("/productByCategoryAndManufacturerList/{catId}/{catName}/{manf}/{start}/{end}", h.ViewProductByCatAndManufacturerList).Methods("GET")

	router.HandleFunc("/searchProductsByName", h.SearchProductList).Methods("POST")
	router.HandleFunc("/searchProductsByName/{search}/{start}/{end}", h.SearchProductList).Methods("GET")
	router.HandleFunc("/searchProductsByManufacturerAndName/{manf}/{search}/{start}/{end}", h.SearchProductByManufacturerList).Methods("GET")

	router.HandleFunc("/addProductToCart", h.AddProductToCart).Methods("GET")
	router.HandleFunc("/addToCart/{prodId}", h.AddProductToCart).Methods("GET")

	router.HandleFunc("/updateCart", h.UpdateProductToCart).Methods("GET")

	router.HandleFunc("/shoppingCartView", h.ViewCart).Methods("GET")

	router.HandleFunc("/startCheckout", h.CheckOutView).Methods("GET")

	router.HandleFunc("/completeOrder/{transactionCode}", h.CheckOutComplateOrder).Methods("GET")

	router.HandleFunc("/checkoutContinue", h.CheckOutContinue).Methods("POST")

	router.HandleFunc("/createCustomerAccount", h.CreateCustomerAccount).Methods("POST")

	router.HandleFunc("/viewCustomerAccount", h.UpdateCustomerAccountPage).Methods("GET")

	router.HandleFunc("/updateCustomerAccount", h.UpdateCustomerAccount).Methods("POST")

	//admin pages
	router.HandleFunc("/admin", h.StoreAdminIndex).Methods("GET")
	router.HandleFunc("/admin/login", h.StoreAdminLogin).Methods("GET")
	router.HandleFunc("/admin/loginNonOAuth", h.StoreAdminLoginNonOAuthUser).Methods("POST")

	router.HandleFunc("/admin/getStore", h.StoreAdminEditStorePage).Methods("GET")
	router.HandleFunc("/admin/updateStore", h.StoreAdminEditStore).Methods("POST")

	router.HandleFunc("/admin/productListBySku", h.StoreAdminSearchProductBySkuPage).Methods("GET")
	router.HandleFunc("/admin/productListBySku", h.StoreAdminSearchProductBySkuPage).Methods("POST")

	router.HandleFunc("/admin/productListByName", h.StoreAdminSearchProductByNamePage).Methods("GET")
	router.HandleFunc("/admin/productListByName", h.StoreAdminSearchProductByNamePage).Methods("POST")

	router.HandleFunc("/admin/productListByCat", h.StoreAdminSearchProductByCategoryPage).Methods("GET")
	router.HandleFunc("/admin/productListByCat", h.StoreAdminSearchProductByCategoryPage).Methods("POST")

	router.HandleFunc("/admin/productList/{start}/{end}", h.StoreAdminViewProductList).Methods("GET")
	router.HandleFunc("/admin/addProductPage", h.StoreAdminAddProductPage).Methods("GET")
	router.HandleFunc("/admin/addProduct", h.StoreAdminAddProduct).Methods("POST")
	router.HandleFunc("/admin/getProduct/{id}", h.StoreAdminEditProductPage).Methods("GET")
	router.HandleFunc("/admin/updateProduct", h.StoreAdminEditProduct).Methods("POST")
	router.HandleFunc("/admin/deleteProduct/{id}", h.StoreAdminDeleteProduct).Methods("GET")

	router.HandleFunc("/admin/productsUploadPage", h.StoreAdminUploadProductFilePage).Methods("GET")
	router.HandleFunc("/admin/productsUpload", h.StoreAdminUploadProductFile).Methods("POST")

	router.HandleFunc("/admin/distributorList", h.StoreAdminViewDistributorList).Methods("GET")
	router.HandleFunc("/admin/addDistributor", h.StoreAdminAddDistributor).Methods("POST")
	router.HandleFunc("/admin/getDistributor/{id}", h.StoreAdminEditDistributorPage).Methods("GET")
	router.HandleFunc("/admin/updateDistributor", h.StoreAdminEditDistributor).Methods("POST")
	router.HandleFunc("/admin/deleteDistributor/{id}", h.StoreAdminDeleteDistributor).Methods("GET")
	router.HandleFunc("/admin/categoryList", h.StoreAdminViewCategoryList).Methods("GET")
	router.HandleFunc("/admin/addCategory", h.StoreAdminAddCategory).Methods("POST")
	router.HandleFunc("/admin/getCategory/{id}", h.StoreAdminEditCategoryPage).Methods("GET")
	router.HandleFunc("/admin/updateCategory", h.StoreAdminEditCategory).Methods("POST")
	router.HandleFunc("/admin/deleteCategory/{id}", h.StoreAdminDeleteCategory).Methods("GET")

	router.HandleFunc("/admin/shippingCarrierList", h.StoreAdminViewCarrierList).Methods("GET")
	router.HandleFunc("/admin/addShippingCarrier", h.StoreAdminAddCarrier).Methods("POST")
	router.HandleFunc("/admin/getShippingCarrier/{id}", h.StoreAdminEditCarrierPage).Methods("GET")
	router.HandleFunc("/admin/updateShippingCarrier", h.StoreAdminEditCarrier).Methods("POST")
	router.HandleFunc("/admin/deleteShippingCarrier/{id}", h.StoreAdminDeleteCarrier).Methods("GET")

	router.HandleFunc("/admin/shippingRegionList", h.StoreAdminViewRegionList).Methods("GET")
	router.HandleFunc("/admin/addShippingRegion", h.StoreAdminAddRegion).Methods("POST")
	router.HandleFunc("/admin/getShippingRegion/{id}", h.StoreAdminEditRegionPage).Methods("GET")
	router.HandleFunc("/admin/updateShippingRegion", h.StoreAdminEditRegion).Methods("POST")
	router.HandleFunc("/admin/deleteShippingRegion/{id}", h.StoreAdminDeleteRegion).Methods("GET")

	router.HandleFunc("/admin/shippingMethodList", h.StoreAdminViewShippingMethodList).Methods("GET")
	router.HandleFunc("/admin/addShippingMethod", h.StoreAdminAddShippingMethod).Methods("POST")
	router.HandleFunc("/admin/getShippingMethod/{id}", h.StoreAdminEditShippingMethodPage).Methods("GET")
	router.HandleFunc("/admin/updateShippingMethod", h.StoreAdminEditShippingMethod).Methods("POST")
	router.HandleFunc("/admin/deleteShippingMethod/{id}", h.StoreAdminDeleteShippingMethod).Methods("GET")

	router.HandleFunc("/admin/insuranceList", h.StoreAdminViewInsuranceList).Methods("GET")
	router.HandleFunc("/admin/addInsurance", h.StoreAdminAddInsurance).Methods("POST")
	router.HandleFunc("/admin/getInsurance/{id}", h.StoreAdminEditInsurancePage).Methods("GET")
	router.HandleFunc("/admin/updateInsurance", h.StoreAdminEditInsurance).Methods("POST")
	router.HandleFunc("/admin/deleteInsurance/{id}", h.StoreAdminDeleteInsurance).Methods("GET")

	router.HandleFunc("/admin/taxRateList", h.StoreAdminViewTaxRateList).Methods("GET")
	router.HandleFunc("/admin/addTaxRate", h.StoreAdminAddTaxRate).Methods("POST")
	router.HandleFunc("/admin/getTaxRate/{id}/{country}/{state}", h.StoreAdminEditTaxRatePage).Methods("GET")
	router.HandleFunc("/admin/updateTaxRate", h.StoreAdminEditTaxRate).Methods("POST")
	router.HandleFunc("/admin/deleteTaxRate/{id}", h.StoreAdminDeleteTaxRate).Methods("GET")

	router.HandleFunc("/admin/pluginList/{start}/{end}", h.StoreAdminViewPluginList).Methods("GET")
	router.HandleFunc("/admin/addPlugin", h.StoreAdminAddPlugin).Methods("POST")
	router.HandleFunc("/admin/getPlugin/{id}", h.StoreAdminEditPluginPage).Methods("GET")
	router.HandleFunc("/admin/updatePlugin", h.StoreAdminEditPlugin).Methods("POST")
	router.HandleFunc("/admin/deletePlugin/{id}", h.StoreAdminDeletePlugin).Methods("GET")

	router.HandleFunc("/admin/storePluginList", h.StoreAdminViewStorePluginList).Methods("GET")
	router.HandleFunc("/admin/addStorePluginFromList/{start}/{end}", h.StoreAdminAddStorePluginFromListPage).Methods("GET")
	router.HandleFunc("/admin/getPluginToAdd/{id}", h.StoreAdminGetStorePluginToAddPage).Methods("GET")
	router.HandleFunc("/admin/addStorePlugin/{id}", h.StoreAdminAddStorePlugin).Methods("GET")
	router.HandleFunc("/admin/getStorePlugin/{id}", h.StoreAdminEditStorePluginPage).Methods("GET")
	router.HandleFunc("/admin/updateStorePlugin", h.StoreAdminEditStorePlugin).Methods("POST")
	router.HandleFunc("/admin/deleteStorePlugin/{id}", h.StoreAdminDeleteStorePlugin).Methods("GET")

	router.HandleFunc("/admin/paymentGatewayList", h.StoreAdminViewPaymentGatewayList).Methods("GET")
	router.HandleFunc("/admin/addPaymentGateway", h.StoreAdminAddPaymentGateway).Methods("POST")
	router.HandleFunc("/admin/getPaymentGateway/{id}", h.StoreAdminEditPaymentGatewayPage).Methods("GET")
	router.HandleFunc("/admin/updatePaymentGateway", h.StoreAdminEditPaymentGateway).Methods("POST")
	router.HandleFunc("/admin/deletePaymentGateway/{id}", h.StoreAdminDeletePaymentGateway).Methods("GET")

	router.HandleFunc("/admin/orderList/{status}", h.StoreAdminViewOrderList).Methods("GET")
	router.HandleFunc("/admin/getOrder/{id}", h.StoreAdminEditOrderPage).Methods("GET")
	router.HandleFunc("/admin/updateOrder", h.StoreAdminEditOrder).Methods("POST")
	router.HandleFunc("/admin/addNewOrderComment", h.StoreAdminEditOrder).Methods("POST")

	router.HandleFunc("/admin/customerList/{start}/{end}", h.StoreAdminViewCustomerList).Methods("GET")
	router.HandleFunc("/admin/getCustomer/{id}", h.StoreAdminEditCustomerPage).Methods("GET")
	router.HandleFunc("/admin/updateCustomer", h.StoreAdminEditCustomer).Methods("POST")
	router.HandleFunc("/admin/customerByEmail", h.StoreAdminSearchCustomerByEmailPage).Methods("GET")
	router.HandleFunc("/admin/customerByEmail", h.StoreAdminSearchCustomerByEmailPage).Methods("POST")

	router.HandleFunc("/admin/imagesUploadPage", h.StoreAdminUploadImageFilesPage).Methods("GET")
	router.HandleFunc("/admin/imagesUpload", h.StoreAdminUploadImageFiles).Methods("POST")

	router.HandleFunc("/admin/thumbnailsUploadPage", h.StoreAdminUploadThumbnailFilesPage).Methods("GET")
	router.HandleFunc("/admin/thumbnailsUpload", h.StoreAdminUploadThumbnailFiles).Methods("POST")

	router.HandleFunc("/admin/addContentPage", h.StoreAdminAddContentPage).Methods("GET")
	router.HandleFunc("/admin/addContent", h.StoreAdminAddContent).Methods("POST")
	router.HandleFunc("/admin/contentList", h.StoreAdminContentList).Methods("GET")
	router.HandleFunc("/admin/getContent/{name}", h.StoreAdminGetContent).Methods("GET")
	router.HandleFunc("/admin/updateContent", h.StoreAdminUpdateContent).Methods("POST")
	router.HandleFunc("/admin/deleteContent/{name}", h.StoreAdminDeleteContent).Methods("GET")

	router.HandleFunc("/admin/addImagePage", h.StoreAdminAddImagePage).Methods("GET")
	router.HandleFunc("/admin/addImage", h.StoreAdminUploadImage).Methods("POST")
	router.HandleFunc("/admin/imageList", h.StoreAdminImageList).Methods("GET")
	router.HandleFunc("/admin/deleteImage/{name}", h.StoreAdminDeleteImage).Methods("GET")

	router.HandleFunc("/admin/addMenuPage", h.StoreAdminAddMenuPage).Methods("GET")
	router.HandleFunc("/admin/addMenu", h.StoreAdminAddMenu).Methods("POST")
	router.HandleFunc("/admin/menuList", h.StoreAdminMenuList).Methods("GET")
	router.HandleFunc("/admin/getMenu/{name}", h.StoreAdminGetMenu).Methods("GET")
	router.HandleFunc("/admin/updateMenu", h.StoreAdminUpdateMenu).Methods("POST")
	// router.HandleFunc("/admin/deleteContent/{name}", h.StoreAdminDeleteContent).Methods("GET")

	router.HandleFunc("/admin/getPageCss/{name}", h.StoreAdminGetPageCSS).Methods("GET")
	router.HandleFunc("/admin/updatePageCss", h.StoreAdminUpdatePageCSS).Methods("POST")

	router.HandleFunc("/admin/getCarousel/{name}", h.StoreAdminGetCarousel).Methods("GET")
	router.HandleFunc("/admin/updateCarousel", h.StoreAdminUpdateCarousel).Methods("POST")

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", router)
}

// go mod init github.com/Ulbora/Six910-ui
