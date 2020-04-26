package cli

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/urfave/cli"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	addressMysql "gitlab.com/ilaryonov/fiascli-clean/domain/address/repository/mysql"
	address "gitlab.com/ilaryonov/fiascli-clean/domain/address/service"
	entity2 "gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
	versionMysql "gitlab.com/ilaryonov/fiascli-clean/domain/version/repository/mysql"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
	"os"
)

type App struct {
	Server         *cli.App
	Logger         logrus.Logger
	DB             *gorm.DB
	addressService *address.AddressService
	VersionService *version.VersionService
}

func NewApp(logger logrus.Logger) *App {
	server := initCli()
	db := initDb()
	addressRepo := addressMysql.NewMysqlAddressRepository(db)
	versionRepo := versionMysql.NewMysqlVersionRepository(db)
	return &App{
		Server:         server,
		Logger:         logger,
		DB:             db,
		addressService: address.NewAddressService(addressRepo),
		VersionService: version.NewVersionService(versionRepo),
	}
}

func initDb() *gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {

	}
	//defer db.Close()

	db.LogMode(viper.GetBool("db.LogMode"))
	if viper.GetBool("db.debug") {
		db.Debug()
	}
	db.Set("gorm:table_options", "charset=utf8")
	db.AutoMigrate(&entity.AddrObject{})
	db.AutoMigrate(&entity.HouseObject{})
	db.AutoMigrate(&entity2.Option{})
	return db
}

func initCli() *cli.App {
	app := cli.App{}
	app.Name = "fiascli"
	app.Usage = "cli fias program"
	app.Version = "0.1.0"

	return &app
}

func (a *App) Run() error {
	err := a.Server.Run(os.Args)
	return err
}
