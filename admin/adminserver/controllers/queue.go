// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package controllers

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/zeebo/errs"

	"ultimatedivision/internal/logger"
	"ultimatedivision/queue"
)

var (
	// ErrQueue is an internal error type for queue controller.
	ErrQueue = errs.Class("queue controller error")
)

// QueueTemplates holds all queue related templates.
type QueueTemplates struct {
	List *template.Template
	Get  *template.Template
}

// Queue is a mvc controller that handles all queue related views.
type Queue struct {
	log       logger.Logger
	queue     *queue.Service
	templates QueueTemplates
}

// NewQueue is a constructor for queue controller.
func NewQueue(log logger.Logger, queue *queue.Service, templates QueueTemplates) *Queue {
	queueController := &Queue{
		log:       log,
		queue:     queue,
		templates: templates,
	}

	return queueController
}

// List is an endpoint that will provide a web page with clients.
func (controller *Queue) List(w http.ResponseWriter, r *http.Request) {
	clients := controller.queue.List()
	err := controller.templates.List.Execute(w, clients)
	if err != nil {
		controller.log.Error("can not execute list clients template", ErrQueue.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// Get is an endpoint that will provide a web page with client by id.
func (controller *Queue) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := uuid.Parse(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	client, err := controller.queue.Get(id)
	if err != nil {
		controller.log.Error("could not get client by id", ErrQueue.Wrap(err))
		switch {
		case queue.ErrNoClient.Has(err):
			http.Error(w, err.Error(), http.StatusNotFound)
		default:
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	err = controller.templates.Get.Execute(w, client)
	if err != nil {
		controller.log.Error("can not execute get client template", ErrQueue.Wrap(err))
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
