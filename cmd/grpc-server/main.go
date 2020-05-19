package main

import (
	"github.com/sirupsen/logrus"
	"github.com/ilaryonov/fiasconfig"
	"github.com/ilaryonov/fiasserver/grpc"
	"os"
)

func main() {
	logger := logrus.Logger{}
	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)
	logger.Info("run grpc server")
	logger.SetLevel(logrus.DebugLevel)
	if err := config.Init(); err != nil {
		logger.Fatalf("%s", err.Error())
	}
	app := grpc.NewApp(logger)
	//defer app.DB.Close()

	//addressRepo := addressRepo.NewGormAddressRepository(db)

	if err := app.Run(); err != nil {
		app.Logger.Fatalf("%s", err.Error())
	}
}
