package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ilaryonov/fias/config"
	addressDelivery "github.com/ilaryonov/fias/domain/address/delivery/cli"
	versionDelivery "github.com/ilaryonov/fias/domain/version/delivery/cli"
	"github.com/ilaryonov/fias/server/cli"
	"os"
)

func main() {
	logger := logrus.Logger{}
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.Info("run cli server")
	logger.SetLevel(logrus.DebugLevel)
	if err := config.Init(); err != nil {
		logger.Fatalf("%s", err.Error())
	}
	app := cli.NewApp(logger)
	defer app.DB.Close()
	versionDelivery.RegisterCliEndpoints(app)
	addressDelivery.RegisterCliEndpoints(app)

	//addressRepo := addressRepo.NewGormAddressRepository(db)

	if err := app.Run(); err != nil {
		app.Logger.Fatalf("%s", err.Error())
	}
}
