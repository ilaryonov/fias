package cli

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"gitlab.com/ilaryonov/fiascli-clean/commands"
	addressMysql "gitlab.com/ilaryonov/fiascli-clean/domain/address/repository/mysql"
	address "gitlab.com/ilaryonov/fiascli-clean/domain/address/service"
	versionMysql "gitlab.com/ilaryonov/fiascli-clean/domain/version/repository/mysql"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
	"os"
)

type App struct {
	server         *cli.App
	Logger         logrus.Logger
	DB             *gorm.DB
	addressService *address.AddressService
	versionService *version.VersionService
}

func NewApp(logger logrus.Logger) *App {
	server := initCli()
	db := initDb()
	addressRepo := addressMysql.NewMysqlAddressRepository(db)
	versionRepo := versionMysql.NewMysqlVersionRepository(db)
	return &App{
		server:         server,
		Logger:         logger,
		DB:             db,
		addressService: address.NewAddressService(addressRepo),
		versionService: version.NewVersionService(versionRepo),
	}
}

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {

	}

	db.LogMode(viper.GetBool("db.LogMode"))
	if viper.GetBool("db.debug") {
		db.Debug()
	}
	return db
}

func initCli() *cli.App {
	app := cli.App{}
	app.Name = "fiascli"
	app.Usage = "cli fias program"
	app.Version = "0.1.0"
	app.Commands = []cli.Command{
		{
			Name:  "version",
			Usage: "fias version",
			Action: func(c *cli.Context) {
				commands.GetVersion()
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
	fmt.Printf("%s", viper.GetString("db.dsn"))
}

func (a *App) Run() error {
	err := a.server.Run(os.Args)
	return err
}
