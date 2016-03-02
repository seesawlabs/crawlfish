package main

import (
	"flag"

	"github.com/Sirupsen/logrus"
	"github.com/seesawlabs/crawlfish/server"
	"github.com/seesawlabs/crawlfish/shared/config"
	"github.com/seesawlabs/crawlfish/shared/database"
)

type Application struct {
	Config   config.IConfiguration
	Database database.IDatabase
	Server   *server.Server
}

func (a *Application) InitConfiguration(configPath string) {
	config, err := config.NewConfig(configPath)
	if err != nil {
		panic(err)
	}

	a.Config = config
}

func (a *Application) InitDatabase() {
	dbx, err := database.NewDatabase(a.Config.DatabaseConfig())
	if err != nil {
		panic(err)
	}

	a.Database = dbx
}

func (a *Application) InitServer() {
	s := server.NewServer(a.Config.ServerConfig())
	s.Init(a.Database)
}

func main() {
	// get env name from flags
	configPath := flag.String("configpath", "env/local.yml", "environment config path")

	app := Application{}

	logrus.Info("Init configuration... \n")
	app.InitConfiguration(*configPath)

	logrus.Info("Init database... \n")
	app.InitDatabase()

	logrus.Info("Start server... \n")
	app.InitServer()
}
