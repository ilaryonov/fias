package service

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	addressEntity "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/service"
	fiasApi "gitlab.com/ilaryonov/fiascli-clean/domain/fiasApi/service"
	"regexp"
	"sync"
)

type ImportService struct {
	addressImportService *AddressImportService
	houseImportService   *HouseImportService
	logger               logrus.Logger
	directoryService     *service.DirectoryService
}

func NewImportService(logger logrus.Logger, directoryService *service.DirectoryService, addressImportService *AddressImportService, houseImportService *HouseImportService) *ImportService {
	return &ImportService{
		addressImportService: addressImportService,
		houseImportService:   houseImportService,
		logger:               logger,
		directoryService:     directoryService,
	}
}

func (is *ImportService) CheckUpdates(api *fiasApi.FiasApiService, version int) {
	result := api.GetAllDownloadFileInfo()
	for _, file := range result {
		if file.VersionId == 613 {
			xmlFiles := is.directoryService.DownloadAndExtractFile(file.FiasDeltaXmlUrl)
			is.ParseFiles(xmlFiles)
		}
	}
}

func (is *ImportService) StartFullImport(api *fiasApi.FiasApiService) {
	fileResult := api.GetLastDownloadFileInfo()
	if len(fileResult.FiasDeltaXmlUrl) > 0 {
		xmlFiles := is.directoryService.DownloadAndExtractFile(fileResult.FiasCompleteXmlUrl)
		is.ParseFiles(xmlFiles)
	}
}

func (is *ImportService) ParseFiles(files *[]entity.File) {
	var wg sync.WaitGroup
	for _, file := range *files {
		if r, err := regexp.MatchString(addressEntity.GetAddressXmlFile(), file.Path); err == nil && r {
			wg.Add(1)
			go is.addressImportService.Import(file.Path, &wg)
		}
		if r, err := regexp.MatchString(addressEntity.GetHouseXmlFile(), file.Path); err == nil && r {
			wg.Add(1)
			go is.houseImportService.Import(file.Path, &wg)
		}
	}
	wg.Wait()
}

func insertCollection(repo address.InsertUpdateInterface, collection []interface{}, node interface{}) []interface{} {
	if collection == nil {
		collection = append(collection, node)
		return collection
	}
	if node == nil {
		repo.InsertUpdateCollection(collection)
		return collection[:0]
	}
	if len(collection) < viper.GetInt("import.collectionCount") {
		collection = append(collection, node)
		return collection
	} else {
		collection = append(collection, node)
		repo.InsertUpdateCollection(collection)
		return collection[:0]
	}
}
