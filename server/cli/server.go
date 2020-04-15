package cli

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"github.com/jinzhu/gorm"
	"os"
)

type App struct {
	server *cli.App
	Logger logrus.Logger
	DB *gorm.DB
}

func NewApp(logger logrus.Logger) App {
	app := &App{}
	app.server = initCli()
	app.DB = initDb()
	app.Logger = logger

	return *app
}

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {

	}
	db.LogMode(true)
	return db
}

func initCli() *cli.App {
	app := cli.App{}
	app.Name = "fiascli"
	app.Usage = "cli fias program"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name:  "version",
			Usage: "fias version",
			Action: func(c *cli.Context) {
				versionHandler()
			},
		},
		/*{
			Name:  "checkdelta",
			Usage: "check deltas from fias.nalog.ru",
			Action: func(c *cli.Context) {
				controllers.CheckUpdates()
			},
		},*/
	}
	return &app
}

func versionHandler() {
	fmt.Printf("%s", "hello")
}

func (a *App) Run() error {
	err := a.server.Run(os.Args)
	return err
}
