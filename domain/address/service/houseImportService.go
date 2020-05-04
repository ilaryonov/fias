package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	addressEntity "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"gitlab.com/ilaryonov/fiascli-clean/domain/directory/service"
	"gitlab.com/ilaryonov/fiascli-clean/helper"
	"sync"
	"time"
)

type HouseImportService struct {
	houseRepo        address.HouseRepositoryInterface
	logger           logrus.Logger
	directoryService *service.DirectoryService
}

func NewHouseImportService(houseRepo address.HouseRepositoryInterface, logger logrus.Logger, directoryService *service.DirectoryService) *HouseImportService {
	return &HouseImportService{
		houseRepo:        houseRepo,
		logger:           logger,
		directoryService: directoryService,
	}
}

func (his *HouseImportService) Import(filePath string, wg *sync.WaitGroup) {
	defer wg.Done()
	start := time.Now()
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
	var collection []interface{}
	count := 0

Loop:
	for {
		select {
		case node := <-houseChannel:
			collection = insertCollection(his.houseRepo, collection, node)
			count++
		case <-done:
			break Loop
		}
	}
	if len(collection) > 0 {
		collection = insertCollection(his.houseRepo, collection, nil)
	}
	finish := time.Now()
	fmt.Println("Количество добавленных записей в адреса:", count)
	fmt.Println("Время выполнения адресов:", finish.Sub(start))
}
