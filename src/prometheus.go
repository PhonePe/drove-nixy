package main

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const ns = "nixy"

var (
	countFailedReloads = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "reloads_failed",
			Help:      "Total number of failed Nginx reloads",
		},
	)
	countSuccessfulReloads = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "reloads_successful",
			Help:      "Total number of successful Nginx reloads",
		},
	)
	histogramReloadDuration = prometheus.NewHistogram(
		prometheus.HistogramOpts{
			Namespace: ns,
			Name:      "reload_duration",
			Help:      "Nginx reload duration",
			Buckets:   prometheus.ExponentialBuckets(0.05, 2, 10),
		},
	)
	countInvalidSubdomainLabelWarnings = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "invalid_subdomain_label_warnings",
			Help:      "Total number of warnings about invalid subdomain label",
		},
	)
	countDuplicateSubdomainLabelWarnings = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "duplicate_subdomain_label_warnings",
			Help:      "Total number of warnings about duplicate subdomain label",
		},
	)
	countEndpointCheckFails = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "endpoint_check_fails",
			Help:      "Total number of endpoint check failure errors",
		},
	)
	countEndpointDownErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "endpoint_down_errors",
			Help:      "Total number of endpoint down errors",
		},
	)
	countAllEndpointsDownErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "all_endpoints_down_errors",
			Help:      "Total number of all endpoints down errors",
		},
	)
	countDroveStreamErrors = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "drove_stream_errors",
			Help:      "Total number of Drove stream errors",
		},
	)
	countDroveStreamNoDataWarnings = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "drove_stream_no_data_warnings",
			Help:      "Total number of warnings about no data in Drove stream",
		},
	)
	countDroveEventsReceived = prometheus.NewCounter(
		prometheus.CounterOpts{
			Namespace: ns,
			Name:      "drove_events_received",
			Help:      "Total number of received Drove events",
		},
	)
)

func setupPrometheusMetrics() {
	prometheus.MustRegister(countFailedReloads)
	prometheus.MustRegister(countSuccessfulReloads)
	prometheus.MustRegister(histogramReloadDuration)
	prometheus.MustRegister(countInvalidSubdomainLabelWarnings)
	prometheus.MustRegister(countDuplicateSubdomainLabelWarnings)
	prometheus.MustRegister(countEndpointCheckFails)
	prometheus.MustRegister(countEndpointDownErrors)
	prometheus.MustRegister(countAllEndpointsDownErrors)
	prometheus.MustRegister(countDroveStreamErrors)
	prometheus.MustRegister(countDroveStreamNoDataWarnings)
	prometheus.MustRegister(countDroveEventsReceived)
}

func observeReloadTimeMetric(e time.Duration) {
	histogramReloadDuration.Observe(float64(e) / float64(time.Second))
}
