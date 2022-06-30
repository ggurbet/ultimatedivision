// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

package metrics

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// Counter exposes functionality to count metric.
type Counter interface {
	// Inc used to increment Counter.
	Inc()
}

// Metric is an implementation of metrics using prometheus.
type Metric struct {
	handler  http.Handler
	newUsers Counter
	logins   Counter
	logouts  Counter
	purchase Counter
}

// GetHandler return http.Handler to access metrics.
func (metric *Metric) GetHandler() http.Handler {
	return metric.handler
}

// NewUsersInc increment Counter newUsers.
func (metric *Metric) NewUsersInc() {
	metric.newUsers.Inc()
}

// LoginsInc increment Counter logins.
func (metric *Metric) LoginsInc() {
	metric.logins.Inc()
}

// LogoutsInc increment Counter logouts.
func (metric *Metric) LogoutsInc() {
	metric.logouts.Inc()
}

// PurchaseInc increment Counter purchase.
func (metric *Metric) PurchaseInc() {
	metric.purchase.Inc()
}

// NewMetric is a constructor for a Metric.
func NewMetric() *Metric {
	newUsers := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_registrations",
		Help: "The total number of successful registrations.",
	})

	logins := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_logins",
		Help: "The total number of successful logins.",
	})

	logouts := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_logouts",
		Help: "The total number of successful logouts.",
	})

	purchase := prometheus.NewCounter(prometheus.CounterOpts{
		Name: "number_purchase",
		Help: "The total number of successful purchase.",
	})

	// Create a custom registry.
	registry := prometheus.NewRegistry()

	// Register using our custom registry gauge.
	registry.MustRegister(newUsers)
	registry.MustRegister(logins)
	registry.MustRegister(logouts)
	registry.MustRegister(purchase)

	// Register system metrics.
	registry.MustRegister(collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}))
	registry.MustRegister(collectors.NewGoCollector())

	return &Metric{
		// Expose metrics.
		handler:  promhttp.HandlerFor(registry, promhttp.HandlerOpts{Registry: registry}),
		newUsers: newUsers,
		logins:   logins,
		logouts:  logouts,
		purchase: purchase,
	}
}
