package cmd

import (
	"encoding/json"

	"github.com/nxadm/tail"
)

type nginxMetricsRecorder struct {
	// Path to nginx log file
	logPath string

	// Nginx metrics to update
	metrics nginxMetrics
}

func newNginxMetricsRecorder(metrics nginxMetrics, logPath string) nginxMetricsRecorder {
	return nginxMetricsRecorder{
		logPath: logPath,
		metrics: metrics,
	}
}

func (r nginxMetricsRecorder) record() {
	go func() {
		// Open log file in tail mode
		t, err := tail.TailFile(r.logPath, tail.Config{Follow: true})
		if err != nil {
			log.Error().Err(err).Msg("Failed to open nginx log file")
			return
		}

		// Read lines
		for line := range t.Lines {
			event := nginxLogEvent{}

			// Parse log line as json
			err = json.Unmarshal([]byte(line.Text), &event)
			if err != nil {
				log.Err(err).Str("line", line.Text).Msg("Failed to parse log line to event")
				continue
			}

			// Update metrics in collector
			r.metrics.update(event)
		}
	}()
}
