package main

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ilaryonov/fiascli-clean/config"
	"gitlab.com/ilaryonov/fiascli-clean/server/cli"
	"strconv"
)

func main() {
	logger := logrus.Logger{}
	if err := config.Init(); err != nil {
		logger.Fatalf("%s", err.Error())
	}
	app := cli.NewApp(logger)

	db, err := gorm.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {

	}
	logMode, err := strconv.ParseBool(viper.GetString("db.logMode"))
	if err != nil {
		logMode = false
	}
	db.LogMode(logMode)

	//addressRepo := addressRepo.NewGormAddressRepository(db)

	if err := app.Run(); err != nil {
		app.Logger.Fatalf("%s", err.Error())
	}
}
