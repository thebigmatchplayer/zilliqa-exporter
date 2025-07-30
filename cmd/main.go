package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thebigmatchplayer/zilliqa-exporter/config"
	"github.com/thebigmatchplayer/zilliqa-exporter/exporter"
	"github.com/thebigmatchplayer/zilliqa-exporter/utils"
)

func main() {
	configPath := flag.String("config", "./config.toml", "Path to configuration file")
	flag.Parse()

	if _, err := os.Stat(*configPath); os.IsNotExist(err) {
		fmt.Printf("Config file not found: %s\n", *configPath)
	}

	utils.InitLogger()
	utils.Log.Info("STARTED EXPORTER")

	cfg := config.LoadConfig(*configPath)

	metrics := exporter.InitMetrics()

	// individual routine to update the metrics without interrupting http handler
	go exporter.StartScraper(metrics, cfg.ExporterStruct.RpcEndpoint, cfg.ExporterStruct.ScrapeInterval)

	exporter.StartHTTPServer(cfg)
}
