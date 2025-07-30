package exporter

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type rpcRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
	ID      int    `json:"id"`
}

type rpcResponse struct {
	Result json.RawMessage `json:"result"`
}

func StartScraper(metrics *Metrics, rpcEndpoint string, interval int) {
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		if res, err := callRPC("eth_blockNumber", rpcEndpoint); err == nil {
			if val, err := parseHexToInt(res); err == nil {
				metrics.BlockHeight.Set(float64(val))
			}
		}
		if res, err := callRPC("eth_syncing", rpcEndpoint); err == nil {
			if val, err := parseBoolToInt(res); err == nil {
				metrics.Syncing.Set(float64(val))
			}
		}
		if res, err := callRPC("net_peerCount", rpcEndpoint); err == nil {
			if val, err := parseHexToInt(res); err == nil {
				metrics.PeerCount.Set(float64(val))
			}

		}
		if res, err := callRPC("net_listening", rpcEndpoint); err == nil {
			if val, err := parseBoolToInt(res); err == nil {
				metrics.Listening.Set(float64(val))
			}
		}
	}
}

func callRPC(method string, rpcEndpoint string) (json.RawMessage, error) {
	reqBody := &rpcRequest{
		JSONRPC: "2.0",
		Method:  method,
		Params:  make([]any, 0),
		ID:      rand.Intn(100) + 1,
	}
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(reqBody); err != nil {
		return nil, err
	}

	//send http post with buffer as body
	resp, err := http.Post(rpcEndpoint, "application/json", &buf)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var rpcResp rpcResponse
	if err := json.NewDecoder(resp.Body).Decode(&rpcResp); err != nil {
		return nil, err
	}

	return rpcResp.Result, nil
}

func parseHexToInt(raw json.RawMessage) (int64, error) {
	hexStr := strings.Trim(string(raw), `"`)
	if !strings.HasPrefix(hexStr, "0x") {
		return 0, fmt.Errorf("invalid hex string: %s", hexStr)
	}
	return strconv.ParseInt(strings.TrimPrefix(hexStr, "0x"), 16, 64)
}

func parseBoolToInt(raw json.RawMessage) (int64, error) {
	boolStr := strings.Trim(string(raw), `"`)
	switch boolStr {
	case "false":
		return 0, nil
	case "true":
		return 1, nil
	default:
		return -1, fmt.Errorf("invalid bool string: %s", boolStr)
	}
}
