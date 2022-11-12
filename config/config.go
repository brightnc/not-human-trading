package config

import (
	"strings"

	"github.com/brightnc/not-human-trading/pkg/logger"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

// ",squash" will ignore mapstructure for that field

type Config struct {
	App appConfig
}

type appConfig struct {
	HTTPPort string `envconfig:"APP_HTTP_PORT" default: 8443`
	Env      string `envconfig:"APP_ENV" default: staging`
}

var config Config

// Init is application config initialization ...
func Init() {
	err := godotenv.Load()
	if err != nil {
		envFileNotFound := strings.Contains(err.Error(), "no such file or directory")
		if !envFileNotFound {
			logger.Infof("read config error: %v", err)
		} else {
			logger.Info("use environment from OS")
		}
	}
	err = envconfig.Process("", &config)
	if err != nil {
		logger.Fatalf("parse config error: %v", err)
	}

}

// GetConfig is use for export private variable which is config ...
func GetConfig() *Config {
	return &config
}
