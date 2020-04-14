package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/config"
	"gitlab.com/ilaryonov/fiascli-clean/server/cli"
)

func main() {
	logger := logrus.Logger{}
	if err := config.Init(); err != nil {
		logger.Fatalf("%s", err.Error())
	}
	app := cli.NewApp(logger)

	if err := app.Run(); err != nil {
		app.Logger.Fatalf("%s", err.Error())
	}
}
