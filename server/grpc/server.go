package grpc

import (
	"context"
	"flag"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	grpcHandler "github.com/ilaryonov/fias/domain/address/delivery/grpc"
	grpcHandlerAddress "github.com/ilaryonov/fias/domain/address/delivery/grpc/address"
	addressMysql "github.com/ilaryonov/fias/domain/address/repository/mysql"
	grpcAddress "github.com/ilaryonov/fias/domain/address/service/grpc"
	grpcService "github.com/ilaryonov/fias/domain/address/service/grpc"
	versionMysql "github.com/ilaryonov/fias/domain/version/repository/mysql"
	version "github.com/ilaryonov/fias/domain/version/service"
	"github.com/ilaryonov/fias/helper"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
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
	wg := sync.WaitGroup{}
	wg.Add(2)
	go a.grpcService(&wg)
	go proxyService(&wg)
	wg.Wait()
	return nil
}

func(a *App) grpcService(wg *sync.WaitGroup) {
	defer wg.Done()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		oscall := <-c
		log.Printf("system call:%+v", oscall)
		a.Server.Server.GracefulStop()
	}()
	if err := a.Server.Serve(); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}



func proxyService(wg *sync.WaitGroup) {
	defer wg.Done()
	var grpcServerEndpoint = flag.String("grpc-server-endpoint",  "localhost:50051", "gRPC server endpoint")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// Register gRPC server endpoint
	// Note: Make sure the gRPC server is running properly and accessible
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := grpcHandlerAddress.RegisterAddressHandlerHandlerFromEndpoint(ctx, mux,  *grpcServerEndpoint, opts)
	if err != nil {
		log.Fatal("error reg endpoint", err)
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	log.Println("Start Http server on port: 8081")
	log.Fatal(http.ListenAndServe(":8081", mux))
}
