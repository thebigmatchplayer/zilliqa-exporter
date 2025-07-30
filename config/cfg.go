package config

import (
	"os"

	"github.com/thebigmatchplayer/zilliqa-exporter/utils"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

type Config struct {
	ExporterStruct struct {
		RpcEndpoint    string `toml:"rpc_endpoint"`
		ScrapeInterval int    `toml:"scrape_interval"`
		Port           int    `toml:"port"`
	} `toml:"exporter"`
}

func LoadConfig(path string) *Config {
	file, err := os.OpenFile(path, os.O_RDONLY, 0440)
	if err != nil {
		utils.Log.Fatal("Unable to Load config", zap.String("path", path), zap.Error(err))
	}
	defer file.Close()

	var cfg Config
	decoder := toml.NewDecoder(file)
	if _, err := decoder.Decode(&cfg); err != nil {
		utils.Log.Fatal("Unable to parse config", zap.Error(err))
	}

	setDefaultConfig(&cfg.ExporterStruct.RpcEndpoint, "127.0.0.1", "rpc_endpoint")
	setDefaultConfig(&cfg.ExporterStruct.ScrapeInterval, 15, "scrape_interval")
	setDefaultConfig(&cfg.ExporterStruct.Port, 6969, "port")

	return &cfg
}

func setDefaultConfig[T comparable](field *T, defaultValue T, fieldName string) {
	var zeroValue T
	if *field == zeroValue {
		*field = defaultValue
		utils.Log.Warn("Config value missing, using default",
			zap.String("field", fieldName),
			zap.Any("default", defaultValue),
		)
	}
}
