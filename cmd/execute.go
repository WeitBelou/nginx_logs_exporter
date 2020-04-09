package cmd

import (
	"fmt"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/rs/zerolog"
	"github.com/spf13/cobra"
)

var (
	// Variables for cmdline flags
	debug     bool
	port      int
	inputFile string

	log = zerolog.New(os.Stderr).Level(zerolog.InfoLevel)
)

func init() {
	// Parse flags and read config
	rootCmd.PersistentFlags().StringVarP(
		&inputFile,
		"input-file", "i",
		"",
		"Input file",
	)
	_ = rootCmd.MarkPersistentFlagRequired("input-file")

	rootCmd.PersistentFlags().BoolVarP(
		&debug,
		"debug", "d",
		false,
		"Enable debug logging",
	)
	rootCmd.PersistentFlags().IntVarP(
		&port,
		"port", "p",
		8989,
		"Port where http server will be listening",
	)
}

var rootCmd = cobra.Command{
	Use: "nginx_logs_exporter",
	Run: runCommand,
	PersistentPreRun: func(*cobra.Command, []string) {
		// Enable debug logging
		if debug {
			log = log.Level(zerolog.DebugLevel)
		}
	},
}

func runCommand(*cobra.Command, []string) {
	log.Info().Msg("Starting nginx_logs_exporter")

	// Create registry for prometheus metrics
	registry := prometheus.NewPedanticRegistry()

	// Create and register metrics collector
	collector := newNginxCollector()
	err := registry.Register(collector)
	if err != nil {
		log.Error().Err(err).Msg("Failed to register metrics collector")
		os.Exit(1)
	}

	// Create and start metrics recorder
	recorder := newNginxMetricsRecorder(collector, inputFile)
	recorder.record()

	// Create server with prometheus handler
	address := fmt.Sprintf("0.0.0.0:%d", port)
	server := newServer(address, newPromHandler(registry))

	// Run server
	log.Info().Str("address", address).Msg("Server started")
	err = server.ListenAndServe()
	if err != nil {
		log.Error().Err(err).Msg("Server stopped")
		os.Exit(1)
	}
}

// Execute executes root command
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
