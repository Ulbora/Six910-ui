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
	hand "github.com/Ulbora/Six910-ui/handlers"
	api "github.com/Ulbora/Six910API-Go"
	"github.com/gorilla/mux"
)

func main() {
	// just the start of Six910-ui Server
	// This is the storefront for Six910-ui.
	var apiURL string
	var storeName string
	var localDomain string
	var apiKey string

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

	var sapi api.Six910API
	sapi.SetAPIKey(apiKey)
	sapi.SetRestURL(apiURL)
	sapi.SetStore(storeName, localDomain)

	var sh hand.Six910Handler
	sh.API = sapi.GetNew()
	var l lg.Logger
	l.LogLevel = lg.AllLevel
	sh.Log = &l
	sh.AdminTemplates = template.Must(template.ParseFiles("./static/admin/index.html", "./static/admin/head.html",
		"./static/admin/login.html", "./static/admin/navbar.html", "./static/admin/productList.html",
		"./static/admin/subnavs/productNavbar.html", "./static/admin/pagination.html", "./static/admin/productSkuSearch.html",
		"./static/admin/productNameSearch.html", "./static/admin/editProduct.html", "./static/admin/addProduct.html",
		"./static/admin/distributorList.html", "./static/admin/editDistributor.html", "./static/admin/categoryList.html",
		"./static/admin/editCategory.html", "./static/admin/shippingCarrierList.html", "./static/admin/editShippingCarrier.html",
		"./static/admin/regionList.html", "./static/admin/editRegion.html",
		"./static/admin/shippingMethodList.html", "./static/admin/editShippingMethod.html",
		"./static/admin/insuranceList.html", "./static/admin/editInsurance.html",
		// "./static/admin/footer.html", "./static/admin/navbar.html", "./static/admin/contentNavbar.html",
	// "./static/admin/addContent.html", "./static/admin/images.html", "./static/admin/templates.html",
	// "./static/admin/updateContent.html", "./static/admin/mailServer.html", "./static/admin/templateUpload.html",
	// "./static/admin/imageUpload.html", "./static/admin/login.html", "./static/admin/backups.html",
	// "./static/admin/backupUpload.html",
	))

	h := sh.GetNew()

	fmt.Println("Six910 (six nine ten) UI is running on port 8080!")
	router := mux.NewRouter()

	router.HandleFunc("/admin", h.StoreAdminIndex).Methods("GET")
	router.HandleFunc("/admin/login", h.StoreAdminLogin).Methods("GET")
	router.HandleFunc("/admin/loginNonOAuth", h.StoreAdminLoginNonOAuthUser).Methods("POST")
	router.HandleFunc("/admin/productListBySku", h.StoreAdminSearchProductBySkuPage).Methods("GET")
	router.HandleFunc("/admin/productListBySku", h.StoreAdminSearchProductBySkuPage).Methods("POST")
	router.HandleFunc("/admin/productListByName", h.StoreAdminSearchProductByNamePage).Methods("GET")
	router.HandleFunc("/admin/productListByName", h.StoreAdminSearchProductByNamePage).Methods("POST")
	router.HandleFunc("/admin/productList/{start}/{end}", h.StoreAdminViewProductList).Methods("GET")
	router.HandleFunc("/admin/addProductPage", h.StoreAdminAddProductPage).Methods("GET")
	router.HandleFunc("/admin/addProduct", h.StoreAdminAddProduct).Methods("POST")
	router.HandleFunc("/admin/getProduct/{id}", h.StoreAdminEditProductPage).Methods("GET")
	router.HandleFunc("/admin/updateProduct", h.StoreAdminEditProduct).Methods("POST")
	router.HandleFunc("/admin/deleteProduct/{id}", h.StoreAdminDeleteProduct).Methods("GET")
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

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.ListenAndServe(":8080", router)
}

// go mod init github.com/Ulbora/Six910-ui
