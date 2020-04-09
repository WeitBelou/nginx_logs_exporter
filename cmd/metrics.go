package cmd

import (
	"github.com/prometheus/client_golang/prometheus"
)

// Common prefix for all nginx metrics
const nginxMetricsNamespace = "nginx"

// List of labels for metrics
var nginxMetricsLabelNames = []string{"host", "uri", "status", "method"}

// nginxMetrics specifies set of metrics collectors for nginx metrics
type nginxMetrics struct {
	httpRequestTotal           *prometheus.CounterVec
	httpRequestDurationSeconds *prometheus.HistogramVec
}

// newNginxCollector create struct with several nginx metrics
func newNginxCollector() nginxMetrics {
	return nginxMetrics{
		httpRequestDurationSeconds: prometheus.NewHistogramVec(
			prometheus.HistogramOpts{
				Namespace: nginxMetricsNamespace,
				Name:      "http_request_duration_seconds",
			},
			nginxMetricsLabelNames,
		),
		httpRequestTotal: prometheus.NewCounterVec(
			prometheus.CounterOpts{
				Namespace: nginxMetricsNamespace,
				Name:      "http_request_total",
			},
			nginxMetricsLabelNames,
		),
	}
}

// Describe implements Collector.
// Runs Describe on all metrics.
func (m nginxMetrics) Describe(ch chan<- *prometheus.Desc) {
	m.httpRequestTotal.Describe(ch)
	m.httpRequestDurationSeconds.Describe(ch)
}

// Collect implements Collector.
// Runs Collect on all metrics.
func (m nginxMetrics) Collect(ch chan<- prometheus.Metric) {
	m.httpRequestTotal.Collect(ch)
	m.httpRequestDurationSeconds.Collect(ch)
}

// update updates metrics from nginx log event
func (m nginxMetrics) update(e nginxLogEvent) {
	// Labels
	labels := e.convertToLabels()

	// Increment number of requests
	m.httpRequestTotal.
		With(labels).
		Inc()

	// Save request duration
	m.httpRequestDurationSeconds.
		With(labels).
		Observe(e.RequestTime)
}
