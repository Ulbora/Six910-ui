package handlers

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"

	six910api "github.com/Ulbora/Six910API-Go"
	mll "github.com/Ulbora/go-mail-sender"
	sdbi "github.com/Ulbora/six910-database-interface"
	"github.com/gorilla/mux"
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

//PgwPage PgwPage
type PgwPage struct {
	Error              string
	PaymentGatway      *sdbi.PaymentGateway
	PaymentGatwayList  *[]sdbi.PaymentGateway
	StorePluginPgwList *[]sdbi.StorePlugins
	StorePluginPgw     *sdbi.StorePlugins
}

//StoreAdminAddPaymentGatewayPage StoreAdminAddPaymentGatewayPage
func (h *Six910Handler) StoreAdminAddPaymentGatewayPage(w http.ResponseWriter, r *http.Request) {
	apgs, suc := h.getSession(r)
	h.Log.Debug("session suc in PWG add view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(apgs) {
			aiErr := r.URL.Query().Get("error")
			var apgpg PgwPage
			apgpg.Error = aiErr
			h.AdminTemplates.ExecuteTemplate(w, adminAddPaymentGatewayPage, &apgpg)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminAddPaymentGateway StoreAdminAddPaymentGateway
func (h *Six910Handler) StoreAdminAddPaymentGateway(w http.ResponseWriter, r *http.Request) {
	apgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw add", suc)
	if suc {
		if h.isStoreAdminLoggedIn(apgs) {
			apg := h.processPgw(r)
			h.Log.Debug("pgw add", *apg)
			hd := h.getHeader(apgs)
			var found bool
			pgl := h.API.GetPaymentGateways(hd)
			for _, pg := range *pgl {
				if pg.StorePluginsID == apg.StorePluginsID {
					found = true
					break
				}
			}
			var suc bool
			if !found {
				//if pgw is BtcPayServer then
				spi := h.API.GetStorePlugin(apg.StorePluginsID, hd)
				h.Log.Debug("spi", *spi)
				if spi.PluginName == btcPayServer {
					apg.Name = spi.PluginName
					bpay := h.BTCPlugin.NewPairConnect(apg.CheckoutURL)
					apg.ClientID = bpay.ClientID
					apg.ClientKey = bpay.PrivateKey
					apg.Token = bpay.Token
					apg.PostOrderURL = h.Six910SiteURL + "/completeOrder/"
					h.Log.Debug("pgw add resp", *bpay)
					//var bpic btc.PayPlugin
					//sent mail to store email
					if h.MailSenderAddress != "" {
						var adminMail mll.Mailer
						adminMail.Body = "Activate your BTCPay Server Token here: " + bpay.PairingURL
						adminMail.Subject = "Activation Link"
						str := h.API.GetStore(h.StoreName, h.LocalDomain, hd)
						adminMail.Recipients = []string{str.Email}
						adminMail.SenderAddress = h.MailSenderAddress
						h.MailSender.SendMail(&adminMail)
					}
				}
				prres := h.API.AddPaymentGateway(apg, hd)
				h.Log.Debug("pgw add resp", *prres)
				log.Println("pgw add resp", *prres)
				//done------------Need to add a new field in payment gateway table for TOKEN--------------------
				//done------------there is no place to store the token from btspay server
				//add call to btcserver plugin to create new
				//instance that contains a btcclient and
				//inject in h.BtcPayPlugin
				//plugin should have New(pg *PaymentGateway), CreateInvoice and maybe more
				suc = prres.Success
			} else {
				suc = true
			}
			if suc {
				http.Redirect(w, r, adminPaymentGatewayListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminPaymentGatewayListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditPaymentGatewayPage StoreAdminEditPaymentGatewayPage
func (h *Six910Handler) StoreAdminEditPaymentGatewayPage(w http.ResponseWriter, r *http.Request) {
	epgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw edit view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(epgs) {
			hd := h.getHeader(epgs)
			epgpErr := r.URL.Query().Get("error")
			epgvars := mux.Vars(r)
			idstr := epgvars["id"]
			pgID, _ := strconv.ParseInt(idstr, 10, 64)
			h.Log.Debug("pgw id in edit", pgID)
			var dgp PgwPage
			dgp.Error = epgpErr

			var wg sync.WaitGroup
			wg.Add(1)
			go func(id int64, header *six910api.Headers) {
				defer wg.Done()
				dgp.PaymentGatway = h.API.GetPaymentGateway(id, header)
				h.Log.Debug("dgp.PaymentGatway", *dgp.PaymentGatway)
			}(pgID, hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				var espipgwlst []sdbi.StorePlugins
				espigl := h.API.GetStorePluginList(header)
				for i := range *espigl {
					if (*espigl)[i].IsPGW {
						espipgwlst = append(espipgwlst, (*espigl)[i])
					}
				}
				dgp.StorePluginPgwList = &espipgwlst
				h.Log.Debug("espipgwlst", espipgwlst)
			}(hd)

			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminEditPaymentGatewayPage, &dgp)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminEditPaymentGateway StoreAdminEditPaymentGateway
func (h *Six910Handler) StoreAdminEditPaymentGateway(w http.ResponseWriter, r *http.Request) {
	epgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw edit", suc)
	if suc {
		if h.isStoreAdminLoggedIn(epgs) {
			epg := h.processPgw(r)
			h.Log.Debug("pgw update", *epg)
			hd := h.getHeader(epgs)
			res := h.API.UpdatePaymentGateway(epg, hd)
			h.Log.Debug("Pgw update resp", *res)
			if res.Success {
				http.Redirect(w, r, adminPaymentGatewayListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminPaymentGatewayListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminViewPaymentGatewayList StoreAdminViewPaymentGatewayList
func (h *Six910Handler) StoreAdminViewPaymentGatewayList(w http.ResponseWriter, r *http.Request) {
	gpgls, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw view", suc)
	if suc {
		if h.isStoreAdminLoggedIn(gpgls) {
			var pgwPage PgwPage
			hd := h.getHeader(gpgls)

			var wg sync.WaitGroup
			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				var spipgwlst []sdbi.StorePlugins
				spigl := h.API.GetStorePluginList(header)
				for i := range *spigl {
					if (*spigl)[i].IsPGW {
						spipgwlst = append(spipgwlst, (*spigl)[i])
					}
				}
				pgwPage.StorePluginPgwList = &spipgwlst
				h.Log.Debug("spipgwlst", spipgwlst)
			}(hd)

			wg.Add(1)
			go func(header *six910api.Headers) {
				defer wg.Done()
				pgl := h.API.GetPaymentGateways(header)
				h.Log.Debug("pgw  in list", pgl)
				pgwPage.PaymentGatwayList = pgl
			}(hd)
			wg.Wait()

			h.AdminTemplates.ExecuteTemplate(w, adminPaymentGatewayListPage, &pgwPage)
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

//StoreAdminDeletePaymentGateway StoreAdminDeletePaymentGateway
func (h *Six910Handler) StoreAdminDeletePaymentGateway(w http.ResponseWriter, r *http.Request) {
	dpgs, suc := h.getSession(r)
	h.Log.Debug("session suc in pgw list delete", suc)
	if suc {
		if h.isStoreAdminLoggedIn(dpgs) {
			hd := h.getHeader(dpgs)
			dpgvars := mux.Vars(r)
			idstrd := dpgvars["id"]
			idddpg, _ := strconv.ParseInt(idstrd, 10, 64)
			res := h.API.DeletePaymentGateway(idddpg, hd)
			h.Log.Debug("pgw delete resp", *res)
			if res.Success {
				http.Redirect(w, r, adminPaymentGatewayListView, http.StatusFound)
			} else {
				http.Redirect(w, r, adminPaymentGatewayListViewFail, http.StatusFound)
			}
		} else {
			http.Redirect(w, r, adminLogin, http.StatusFound)
		}
	}
}

func (h *Six910Handler) processPgw(r *http.Request) *sdbi.PaymentGateway {
	var p sdbi.PaymentGateway
	id := r.FormValue("id")
	p.ID, _ = strconv.ParseInt(id, 10, 64)
	p.CheckoutURL = r.FormValue("checkoutUrl")
	p.ClientID = r.FormValue("clientId")
	p.ClientKey = r.FormValue("clientKey")
	p.LogoURL = r.FormValue("logoUrl")
	p.PostOrderURL = r.FormValue("postOrderUrl")
	p.Token = r.FormValue("token")
	storePID := r.FormValue("storePluginId")
	p.StorePluginsID, _ = strconv.ParseInt(storePID, 10, 64)

	if len(p.CheckoutURL) > 3 && strings.LastIndex(p.CheckoutURL, "/") == len(p.CheckoutURL)-1 {
		p.CheckoutURL = p.CheckoutURL[0 : len(p.CheckoutURL)-1]
	}

	return &p
}
