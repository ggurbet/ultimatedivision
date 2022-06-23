// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/BoostyLabs/evmsignature"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/store"
)

var (
	// ErrStore is an internal error type for store controller.
	ErrStore = errs.Class("store controller error")
)

// StoreTemplates holds all store related templates.
type StoreTemplates struct {
	List   *template.Template
	Update *template.Template
}

// Store is a mvc controller that handles all admins related views.
type Store struct {
	log logger.Logger

	store *store.Service

	templates StoreTemplates
}

// NewStore is a constructor for store controller.
func NewStore(log logger.Logger, store *store.Service, templates StoreTemplates) *Store {
	storeController := &Store{
		log:       log,
		store:     store,
		templates: templates,
	}

	return storeController
}

// SettingResponse entity describes values required for returning to admin panel.
type SettingResponse struct {
	ID          int     `json:"id"`
	CardsAmount int     `json:"cardsAmount"`
	IsRenewal   bool    `json:"isRenewal"`
	HourRenewal int     `json:"dateRenewal"`
	Price       float64 `json:"price"`
}

// List is an endpoint that will provide a web page with all settings.
func (controller *Store) List(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	settings, err := controller.store.List(ctx)
	if err != nil {
		controller.log.Error("could not get store list", ErrStore.Wrap(err))
		http.Error(w, "could not get store list", http.StatusInternalServerError)
		return
	}

	var settingResponses []SettingResponse
	for _, setting := range settings {
		settingResponse := SettingResponse{
			ID:          setting.ID,
			CardsAmount: setting.CardsAmount,
			IsRenewal:   setting.IsRenewal,
			HourRenewal: setting.HourRenewal,
			Price:       evmsignature.WeiBigToEthereumFloat(&setting.Price),
		}

		settingResponses = append(settingResponses, settingResponse)
	}

	err = controller.templates.List.Execute(w, settingResponses)
	if err != nil {
		controller.log.Error("can not execute list store template", ErrStore.Wrap(err))
		http.Error(w, "can not execute list store template", http.StatusInternalServerError)
		return
	}
}

// Update is an endpoint that will update setting of store.
func (controller *Store) Update(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	params := mux.Vars(r)

	settingID, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "could not parse setting id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		setting, err := controller.store.Get(ctx, settingID)
		if err != nil {
			controller.log.Error("could not get setting", ErrStore.Wrap(err))
			if store.ErrNoSetting.Has(err) {
				http.Error(w, "no setting with such id", http.StatusNotFound)
				return
			}
			http.Error(w, "could not get setting", http.StatusInternalServerError)
			return
		}

		price, err := strconv.ParseFloat(setting.Price.String(), 64)
		if err != nil {
			controller.log.Error("could not convert price", ErrStore.Wrap(err))
			http.Error(w, "could not convert price", http.StatusInternalServerError)
			return
		}

		settingResponse := SettingResponse{
			ID:          setting.ID,
			CardsAmount: setting.CardsAmount,
			IsRenewal:   setting.IsRenewal,
			HourRenewal: setting.HourRenewal,
			Price:       price / float64(evmsignature.WeiInEthereum),
		}

		if err = controller.templates.Update.Execute(w, settingResponse); err != nil {
			controller.log.Error("could not execute update store template", ErrStore.Wrap(err))
			http.Error(w, "could not execute update store template", http.StatusInternalServerError)
			return
		}
	case http.MethodPost:
		if err = r.ParseForm(); err != nil {
			http.Error(w, "could not get store form", http.StatusBadRequest)
			return
		}

		cardsAmount, err := strconv.Atoi(r.FormValue("cardsAmount"))
		if err != nil {
			http.Error(w, "invalid cardsAmount", http.StatusBadRequest)
			return
		}

		isRenewal, err := strconv.ParseBool(r.FormValue("isRenewal"))
		if err != nil {
			http.Error(w, "invalid isRenewal", http.StatusBadRequest)
			return
		}

		hourRenewal, err := strconv.Atoi(r.FormValue("hourRenewal"))
		if err != nil {
			http.Error(w, "invalid hourRenewal", http.StatusBadRequest)
			return
		}
		if hourRenewal < store.HourOfDayMin || hourRenewal > store.HourOfDayMax {
			http.Error(w, fmt.Sprintf("hourRenewal should be in the range of %d to %d", store.HourOfDayMin, store.HourOfDayMax), http.StatusBadRequest)
			return
		}

		price, err := strconv.ParseFloat(r.FormValue("price"), 64)
		if err != nil {
			http.Error(w, "invalid price", http.StatusBadRequest)
			return
		}

		bigIntPrice, err := evmsignature.EthereumFloatToWeiBig(price)
		if err != nil {
			http.Error(w, "invalid price", http.StatusBadRequest)
			return
		}

		setting := store.Setting{
			ID:          settingID,
			CardsAmount: cardsAmount,
			IsRenewal:   isRenewal,
			HourRenewal: hourRenewal,
			Price:       *bigIntPrice,
		}

		if err = controller.store.Update(ctx, setting); err != nil {
			controller.log.Error("could not update store status", ErrStore.Wrap(err))
			http.Error(w, "could not update store status", http.StatusInternalServerError)
			return
		}
		Redirect(w, r, "/store", http.MethodGet)
	}
}
