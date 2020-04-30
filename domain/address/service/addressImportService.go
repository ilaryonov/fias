package service

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	addressEntity "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/service"
	fiasApi "gitlab.com/ilaryonov/fiascli-clean/domain/fiasApi/service"
	"gitlab.com/ilaryonov/fiascli-clean/helper"
	"regexp"
	"sync"
)

type AddressImportService struct {
	addressRepo      address.AddressRepositoryInterface
	logger           logrus.Logger
	directoryService *service.DirectoryService
}

func NewAddressService(addressRepo address.AddressRepositoryInterface, logger logrus.Logger, directoryService *service.DirectoryService) *AddressImportService {
	return &AddressImportService{
		addressRepo:      addressRepo,
		logger:           logger,
		directoryService: directoryService,
	}
}

func (a *AddressImportService) CheckUpdates(api *fiasApi.FiasApiService, version int) {
	result := api.GetAllDownloadFileInfo()
	for _, file := range result {
		if file.VersionId == 613 {
			xmlFiles := a.directoryService.DownloadAndExtractFile(file.FiasDeltaXmlUrl)
			a.ParseFiles(xmlFiles)
		}
	}
}

func (a *AddressImportService) StartFullImport(api *fiasApi.FiasApiService) {
	fileResult := api.GetLastDownloadFileInfo()
	if len(fileResult.FiasDeltaXmlUrl) > 0 {
		xmlFiles := a.directoryService.DownloadAndExtractFile(fileResult.FiasDeltaXmlUrl)
		a.ParseFiles(xmlFiles)
	}
}

func (a *AddressImportService) ParseFiles(files *[]entity.File) {
	var wg sync.WaitGroup
	for _, file := range *files {
		if r, err := regexp.MatchString(addressEntity.GetAddressXmlFile(), file.Path); err == nil && r {
			wg.Add(1)
			go a.ImportAddress(file.Path, &wg)
		}
		if r, err := regexp.MatchString(addressEntity.GetHouseXmlFile(), file.Path); err == nil && r {
			wg.Add(1)
			go a.ImportHouse(file.Path, &wg)
		}
	}
	wg.Wait()
}

func (a *AddressImportService) ImportAddress(filePath string, wg *sync.WaitGroup) {
	var address addressEntity.AddrObject
	defer wg.Done()
	addressChannel := make(chan addressEntity.XmlToStructInterface)
	done := make(chan bool)
	defer close(addressChannel)
	go helper.ParseFile(filePath, addressChannel, done, &address)
	count := 0

Loop:
	for {
		select {
		case node := <-addressChannel:
			a.logger.Info(node.(*addressEntity.AddrObject).Aoguid)
			count++
		case <-done:
			break Loop
		}
	}
	close(addressChannel)
	//close(done)
	a.logger.Info("done import addresses. Count: ", count)
}

func (a *AddressImportService) ImportHouse(filePath string, wg *sync.WaitGroup) {
	var house addressEntity.HouseObject
	defer wg.Done()
	houseChannel := make(chan addressEntity.XmlToStructInterface)
	done := make(chan bool)
	defer close(houseChannel)
	go helper.ParseFile(filePath, houseChannel, done, &house)
	count := 0

Loop:
	for {
		select {
		case node := <-houseChannel:
			a.logger.Info(node.(*addressEntity.HouseObject).Houseguid)
			count++
		case <-done:
			break Loop
		}
	}
	close(houseChannel)
	//close(done)
	a.logger.Info("done import addresses. Count: ", count)
}
