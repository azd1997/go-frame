package config

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
)

type LoggerConfig struct {
	LogPath string `json:"log_path"`
	LogLevel string `json:"log_level"`
} 

type Config struct {
	LoggerConfig LoggerConfig `json:"logger_config"`
}

var conf Config

func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Wrap(err, "Read config file failed")
	}
	err = json.Unmarshal(configFile, &conf)
	if err != nil {
		return errors.Wrap(err, "Unmarshal config file failed")
	}
	return nil
}

func GetConfig() Config {
	return conf
}