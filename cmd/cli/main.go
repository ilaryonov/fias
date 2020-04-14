package main

import (
	"fiascli-clean/config"
	"github.com/sirupsen/logrus"
)

func main() {
	logger := logrus.Logger{}
	if err := config.Init(); err != nil {
		logger.Fatalf("%s", err.Error())
	}

}
