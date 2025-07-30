package exporter

import "github.com/prometheus/client_golang/prometheus"

type Metrics struct {
	BlockHeight prometheus.Gauge
	Syncing     prometheus.Gauge
	Listening   prometheus.Gauge
	PeerCount   prometheus.Gauge
}

func InitMetrics() *Metrics {
	m := &Metrics{
		BlockHeight: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "zilliqa_block_height",
			Help: "Current Zilliqa block height",
		}),
		Syncing: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "zilliqa_syncing",
			Help: "Whether the Zilliqa node is syncing (1 = true, 0 = false)",
		}),
		Listening: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "zilliqa_listening",
			Help: "Whether the Zilliqa node is listening for peers (1 = true, 0 = false)",
		}),
		PeerCount: prometheus.NewGauge(prometheus.GaugeOpts{
			Name: "zilliqa_peer_count",
			Help: "Number of peers connected to the Zilliqa node",
		}),
	}

	prometheus.MustRegister(m.BlockHeight, m.Syncing, m.Listening, m.PeerCount)
	return m
}
