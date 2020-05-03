package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	addressEntity "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/service"
	fiasApi "gitlab.com/ilaryonov/fiascli-clean/domain/fiasApi/service"
	"gitlab.com/ilaryonov/fiascli-clean/helper"
	"regexp"
	"sync"
	"time"
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
	defer wg.Done()
	addressChannel := make(chan interface{})
	done := make(chan bool)
	//defer close(addressChannel)
	go helper.ParseFile(filePath, addressChannel, done, func(decoder *xml.Decoder, se *xml.StartElement) (interface{}, error) {
		if se.Name.Local == "Object" {
			result := addressEntity.AddrObject{}
			err := decoder.DecodeElement(&result, se)
			result.ID = 0
			if result.Actstatus == "0" {
				return nil, errors.New("не активный адрес")
			}
			if err != nil {
				return nil, err
			}
			return result, nil
		}
		return nil, errors.New("не адрес")
	})
	count := 0
	var collection []interface{}

Loop:
	for {
		select {
		case node := <-addressChannel:
			//a.logger.Info(node.(*addressEntity.AddrObject).Aoguid)
			collection = a.insert(collection, node)
			count++
		case <-done:
			break Loop
		}
	}
	if len(collection) > 0 {
		collection = a.insert(collection, nil)
	}
	a.logger.Info("done import addresses. Count: ", count)
}

func (a *AddressImportService) ImportHouse(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	houseChannel := make(chan interface{})
	done := make(chan bool)
	defer close(houseChannel)
	go helper.ParseFile(filePath, houseChannel, done, func(decoder *xml.Decoder, se *xml.StartElement) (interface{}, error) {
		layoutISO := "2006-01-02"
		result := addressEntity.HouseObject{}
		if se.Name.Local == "House" {
			err := decoder.DecodeElement(&result, se)
			result.ID = 0
			if err != nil {
				return nil, err
			}
			t, _ := time.Parse(layoutISO, result.EndDate)
			if t.Unix() < time.Now().Unix() {
				return nil, errors.New("не активен по дате")
			}
			return result, nil
		}
		return nil, errors.New("Не дом")
	})
	count := 0

Loop:
	for {
		select {
		case node := <-houseChannel:
			a.logger.Info(node.(addressEntity.HouseObject).Houseguid)
			count++
		case <-done:
			break Loop
		}
	}
	close(houseChannel)
	a.logger.Info("done import houses. Count: ", count)
}

func (a *AddressImportService) insert(collection []interface{}, node interface{}) []interface{} {
	if node == nil {
		err := a.addressRepo.InsertUpdateCollection(collection)
		if err != nil {
			a.logger.Error(err.Error())
		}
		return collection[:0]
	}
	if len(collection) < viper.GetInt("import.collectionCount") {
		collection = append(collection, node)
		return collection
	} else {
		collection = append(collection, node)
		err := a.addressRepo.InsertUpdateCollection(collection)
		if err != nil {
			fmt.Println("error", err.Error())
		}
		return collection[:0]
	}
}
