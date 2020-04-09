package cmd

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// newPromHandler creates new http handler for exporting metrics in Prometheus format
func newPromHandler(registry *prometheus.Registry) http.Handler {
	return promhttp.InstrumentMetricHandler(
		registry,
		promhttp.HandlerFor(registry, promhttp.HandlerOpts{
			Registry:          registry,
			EnableOpenMetrics: true,
		}),
	)
}

// newServer creates http server and registers prometheus handler
func newServer(address string, promHandler http.Handler) http.Server {
	// Create new router
	mux := http.NewServeMux()

	// Register prometheus endpoint
	mux.Handle("/metrics", promHandler)

	// Create http server with mux as root handler
	return http.Server{
		Addr:    address,
		Handler: mux,
	}
}
