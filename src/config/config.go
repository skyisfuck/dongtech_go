package config

import (
	"github.com/BurntSushi/toml"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

const (
	config_file_address = `../resources/config.toml` //go build正式环境用
	//config_file_address = `resources/config.toml` //goland本地启动用
	config_file = `config`
)

func init() {
	viper.SetDefault(config_file, config_file_address)
}

func GetConfig() (*Config, error) {
	configFilePath := viper.GetString(config_file)
	var conf Config
	if _, err := toml.DecodeFile(configFilePath, &conf); err != nil {
		logrus.WithError(err).Println("get config err")
		return nil, err
	}
	return &conf, nil
}

type Config struct {
	Base     Base
	Database Database
	Grpc     Grpc
	Web      Web
}

type Base struct {
	Author string
	Age    int
}

type Database struct {
	User     string
	Password string
	Addr     string
	Database string
	PoolSize int
	Slow     int
	Port     int
}

type Grpc struct {
	Addr string
}

type Web struct {
	Addr string
}
