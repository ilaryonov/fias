package grpc

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	grpcHandler "gitlab.com/ilaryonov/fiascli-clean/domain/address/delivery/grpc"
	addressMysql "gitlab.com/ilaryonov/fiascli-clean/domain/address/repository/mysql"
	grpcAddress "gitlab.com/ilaryonov/fiascli-clean/domain/address/service/grpc"
	grpcService "gitlab.com/ilaryonov/fiascli-clean/domain/address/service/grpc"
	versionMysql "gitlab.com/ilaryonov/fiascli-clean/domain/version/repository/mysql"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
	"gitlab.com/ilaryonov/fiascli-clean/helper"
)

type App struct {
	Server         *grpcHandler.Handler
	Logger         logrus.Logger
	DB             *gorm.DB
	AddressService *grpcAddress.AddressService
	//HouseService     *address.HouseImportService
	VersionService *version.VersionService
}

func NewApp(logger logrus.Logger) *App {
	defer func() {
		if r := recover(); r != nil {
			logger.Fatal(r)
		}
	}()
	db := helper.InitMysqlGormDb()
	addressRepo := addressMysql.NewMysqlAddressRepository(db)
	//houseRepo := addressMysql.NewMysqlHouseRepository(db)
	versionRepo := versionMysql.NewMysqlVersionRepository(db)
	addressService := grpcService.NewAddressService(addressRepo, logger)

	handler := grpcHandler.NewHandler(addressService)

	return &App{
		Server:         handler,
		Logger:         logger,
		DB:             db,
		AddressService: addressService,
		//HouseService:     houseImportService,
		VersionService: version.NewVersionService(versionRepo, logger),
	}
}

func (a *App) Run() error {
	err := a.Server.Serve()
	if err != nil {
		a.Logger.Fatal(err)
	}
	return nil
}
