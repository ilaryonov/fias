package service

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	fiasApi "gitlab.com/ilaryonov/fiascli-clean/domain/fiasApi/service"
)

type AddressImportService struct {
	addressRepo address.AddressRepositoryInterface
	logger      logrus.Logger
}

func NewAddressService(addressRepo address.AddressRepositoryInterface, logger logrus.Logger) *AddressImportService {
	return &AddressImportService{
		addressRepo: addressRepo,
		logger:      logger,
	}
}

func (a *AddressImportService) CheckUpdates(api *fiasApi.FiasApiService) {
	result := api.GetAllDownloadFileInfo()
	a.logger.Warn("run checkUpdates", result)
}

func (a *AddressImportService) StartFullImport(api *fiasApi.FiasApiService) {
	fileResult := api.GetLastDownloadFileInfo()
	if len(fileResult.FiasCompleteXmlUrl) > 0 {

	}

	/*file := helpers.FileData{FileResult.VersionId, FileResult.FiasCompleteXmlUrl}
	downloadedFileName := app.DownloadFile(&file)
	app.UnzipFiles(&file, downloadedFileName)
	app.Extract()
	v := models.Option{Name: "version", Value: FileResult.VersionId}
	app.Config.GetDb().Create(&v)*/
}
