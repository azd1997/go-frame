package config

import (
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type LoggerConfig struct {
	LogPath string `json:"log_path" yaml:"log_path"`
	LogLevel string `json:"log_level" yaml:"log_level"`
}

type DBConfig struct {
	DbHost string `json:"db_host" yaml:"db_host"`
	DbPort string `json:"db_port" yaml:"db_port"`
	DbUser string `json:"db_user" yaml:"db_user"`
	DbPassword string `json:"db_password" yaml:"db_password"`
	DbName string `json:"db_name" yaml:"db_name"`
}

type RedisConfig struct {
	Addr string `json:"addr" yaml:"addr"`
	Password string `json:"password" yaml:"password"`
	DB int `json:"db" yaml:"db"`
}

type HttpConfig struct {
	Addr string `json:"addr" yaml:"addr"`
}

type RpcConfig struct {
	Addr string `json:"addr" yaml:"addr"`
}

type NsqConfig struct {
	Topic string `json:"topic" yaml:"topic"`
	Channel string `json:"channel" yaml:"channel"`
	NsqLookupdAddr string `json:"nsqlookupd_addr" yaml:"nsqlookupd_addr"`
	NsqdAddr string `json:"nsqd_addr" yaml:"nsqd_addr"`
}

type Config struct {
	LoggerConfig LoggerConfig `json:"logger_config" yaml:"logger_config"`
	DBConfig DBConfig `json:"db_config" yaml:"db_config"`
	RedisConfig RedisConfig `json:"redis_config" yaml:"redis_config"`
	HttpConfig HttpConfig `json:"http_config" yaml:"http_config"`
	RpcConfig RpcConfig `json:"rpc_config" yaml:"rpc_config"`
	NsqConfig NsqConfig `json:"nsq_config" yaml:"nsq_config"`
}

var conf Config

func InitConfig(configPath string) error {
	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		return errors.Wrap(err, "Read config file failed")
	}
	err = yaml.Unmarshal(configFile, &conf)
	if err != nil {
		return errors.Wrap(err, "Unmarshal config file failed")
	}
	return nil
}

func Conf() Config {
	return conf
}