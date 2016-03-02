package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/seesawlabs/crawlfish/server"
	"github.com/seesawlabs/crawlfish/shared/config"
)

type Application struct {
	Config config.IConfiguration
	Server *server.Server
}

func (a *Application) InitConfiguration(configPath string) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	a.Config = config
}

func (a *Application) InitServer() {
	s := server.NewServer(a.Config)
	s.Init()
}

func main() {
	configPath := flag.String("configpath", "env/dev.toml", "environment config path")

	app := Application{}

	logrus.Info("Init configuration... \n")
	app.InitConfiguration(*configPath)

	logrus.Info("Start server... \n")
	app.InitServer()
}
