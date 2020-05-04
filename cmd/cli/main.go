package main

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/config"
	addressDelivery "gitlab.com/ilaryonov/fiascli-clean/domain/address/delivery/cli"
	versionDelivery "gitlab.com/ilaryonov/fiascli-clean/domain/version/delivery/cli"
	"gitlab.com/ilaryonov/fiascli-clean/server/cli"
	"os"
)

func main() {
	/*numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	runtime.GOMAXPROCS(numcpu)*/
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
