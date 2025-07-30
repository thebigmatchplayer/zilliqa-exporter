# Zilliqa Metrics Exporter

A Prometheus exporter for monitoring a Zilliqa RPC node. It periodically scrapes key RPC metrics and exposes them via an HTTP endpoint in a format compatible with Prometheus.

## Features

- Export Zilliqa node metrics:
  - Block height
  - Syncing status
  - Peer count
  - Listening status
- Expose metrics on `/metrics` endpoint
- Configurable via `config.toml`
ß- Developed and maintained by [Luganodes](https://www.luganodes.com/)

## Getting Started

### Installation

```bash
git clone https://github.com/thebigmatchplayer/zilliqa-exporter.git
cd zilliqa-exporter
go build -o zilliqa-exporter ./cmd
```

### Configuration

Edit the `config.toml` file:

```toml
[exporter]
rpc_endpoint = "https://api.zilliqa.com/"  # Default: official Zilliqa RPC
scrape_interval = 15                       # Default: 15 seconds
port = 6969                                # Default: 6969
```

### Usage

```bash
zilliqa-exporter -config /path/to/config.toml
```

### Prometheus Scrape Config

Add this to your Prometheus `prometheus.yml`:

```yaml
scrape_configs:
  - job_name: "zilliqa"
    static_configs:
      - targets: ["localhost:6969"]
```

## Metrics

| Metric Name            | Description                                      |
|------------------------|--------------------------------------------------|
| `zilliqa_block_height` | Current Zilliqa block height                     |
| `zilliqa_syncing`      | Syncing status (`1 = true`, `0 = false`)         |
| `zilliqa_peer_count`   | Number of connected peers                        |
| `zilliqa_listening`    | Whether the node is listening (`1 = true`, `0`)  |

## License

MIT License

## Maintainers

This project is maintained by [Luganodes](https://www.luganodes.com/).
