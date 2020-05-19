package service

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/ilaryonov/fiasdomain/address"
	addressEntity "github.com/ilaryonov/fiasdomain/address/entity"
	"github.com/ilaryonov/fiasdomain/directory/entity"
	"github.com/ilaryonov/fiasdomain/directory/service"
	fiasApi "github.com/ilaryonov/fiasdomain/fiasApi/service"
	"regexp"
	"sync"
)

type ImportService struct {
	addressImportService *AddressImportService
	houseImportService   *HouseImportService
	logger               logrus.Logger
	directoryService     *service.DirectoryService
	isFull bool `default:"false"`
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
	is.isFull = true
	fileResult := api.GetLastDownloadFileInfo()
	if len(fileResult.FiasCompleteXmlUrl) > 0 {
		xmlFiles := is.directoryService.DownloadAndExtractFile(fileResult.FiasCompleteXmlUrl)
		is.ParseFiles(xmlFiles)
	}
}

func (is *ImportService) ParseFiles(files *[]entity.File) {
	var wg sync.WaitGroup
	for _, file := range *files {
		if r, err := regexp.MatchString(addressEntity.GetAddressXmlFile(), file.Path); err == nil && r {
			wg.Add(1)
			go is.addressImportService.Import(file.Path, &wg, is.isFull)
		}
		if r, err := regexp.MatchString(addressEntity.GetHouseXmlFile(), file.Path); err == nil && r {
			wg.Add(1)
			go is.houseImportService.Import(file.Path, &wg, is.isFull)
		}
	}
	wg.Wait()

}

func insertCollection(repo address.InsertUpdateInterface, collection []interface{}, node interface{}, isFull bool) []interface{} {
	if collection == nil {
		collection = append(collection, node)
		return collection
	}
	if node == nil {
		repo.InsertUpdateCollection(collection, isFull)
		return collection[:0]
	}
	if len(collection) < viper.GetInt("import.collectionCount") {
		collection = append(collection, node)
		return collection
	} else {
		collection = append(collection, node)
		repo.InsertUpdateCollection(collection, isFull)
		return collection[:0]
	}
}
