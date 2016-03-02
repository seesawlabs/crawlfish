package config

import (
	"github.com/dropbox/godropbox/errors"
	"github.com/spf13/viper"
)

type IConfiguration interface {
	ServerConfig() Server
	FirebaseConfig() Firebase
	All() *configuration
}

type configuration struct {
	Server   Server
	Firebase Firebase
}

type Server struct {
	Address string
}

type Firebase struct {
	Url  string
	Auth string
}

func (c *configuration) ServerConfig() Server {
	return c.Server
}

func (c *configuration) FirebaseConfig() Firebase {
	return c.Firebase
}

func (c *configuration) All() *configuration {
	return c
}

func NewConfig(configPath string) (IConfiguration, error) {
	var c configuration

	viper.SetConfigFile(configPath)

	err := viper.ReadInConfig()
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		return nil, errors.Wrap(err, "")
	}

	return &c, nil
}
