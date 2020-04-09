package cmd

import (
	"strconv"

	"github.com/prometheus/client_golang/prometheus"
)

// nginxLogEvent describes nginx access log event
type nginxLogEvent struct {
	HTTPHost      string  `json:"http_host"`
	URI           string  `json:"uri"`
	Status        int     `json:"status"`
	RequestTime   float64 `json:"request_time"`
	RequestMethod string  `json:"request_method"`
}

// convertToLabels converts log event to prometheus Labels
func (e nginxLogEvent) convertToLabels() prometheus.Labels {
	return prometheus.Labels{
		"host":   e.HTTPHost,
		"uri":    e.URI,
		"status": strconv.Itoa(e.Status),
		"method": e.RequestMethod,
	}
}
