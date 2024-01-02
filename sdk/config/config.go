package config

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	GRpcServerConf GRpcServerConf `toml:"grpc_server"`
}

type GRpcServerConf struct {
	Enable   bool   `toml:"enable"`
	Port     int    `toml:"port"`
	Protocol string `toml:"protocol"`
}

var Cfg Config

func InitConfig(filePath string) error {
	if _, err := toml.DecodeFile(filePath, &Cfg); err != nil {
		return err
	}
	return nil
}
