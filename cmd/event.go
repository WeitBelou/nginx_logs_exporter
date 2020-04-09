package cmd

import (
	"strconv"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

// nginxLogEvent describes nginx access log event
type nginxLogEvent struct {
	Status      int
	RequestTime time.Duration
	URI         string
}

// convertToLabels converts log event to prometheus Labels
func (e nginxLogEvent) convertToLabels() prometheus.Labels {
	return prometheus.Labels{
		"uri":    e.URI,
		"status": strconv.Itoa(e.Status),
	}
}
