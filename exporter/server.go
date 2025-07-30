package exporter

import (
	"fmt"
	"net/http"

	"github.com/thebigmatchplayer/zilliqa-exporter/config"
	"github.com/thebigmatchplayer/zilliqa-exporter/utils"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"
)

func StartHTTPServer(cfg *config.Config) {
	http.Handle("/metrics", promhttp.Handler())
	utils.Log.Info("Starting metrics server", zap.Int("port", cfg.ExporterStruct.Port))
	if err := http.ListenAndServe(fmt.Sprintf(":%d", cfg.ExporterStruct.Port), nil); err != nil {
		utils.Log.Fatal("HTTP server failed", zap.Error(err))
	}
}
