package config

import (
	"fmt"

	"github.com/dropbox/godropbox/errors"
	"github.com/spf13/viper"
)

type IConfiguration interface {
	DatabaseConfig() Database
	ServerConfig() Server
}

type configuration struct {
	Database Database
	Server   Server
}

type Database struct {
	Username  string
	Password  string
	Hostname  string
	Port      string
	Name      string
	Parameter string
}

type Server struct {
	Address string
}

func (c *configuration) DatabaseConfig() Database {
	return c.Database
}

func (c *configuration) ServerConfig() Server {
	return c.Server
}

func NewConfig(configPath string) (IConfiguration, error) {
	var c configuration

	// TODO: use configPath
	viper.SetConfigFile("env/local.toml")

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("error reading")
		return nil, errors.Wrap(err, "")
	}

	err = viper.Unmarshal(&c)
	if err != nil {
		fmt.Println("error unmarshalig")
		return nil, errors.Wrap(err, "")
	}

	return &c, nil
}
