package service

import (
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/ilaryonov/fiasdomain/address"
	addressEntity "github.com/ilaryonov/fiasdomain/address/entity"
	"github.com/ilaryonov/fiasdomain/directory/service"
	"github.com/ilaryonov/fiashelper"
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

func (his *HouseImportService) Import(filePath string, wg *sync.WaitGroup, isFull bool) {
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
			collection = insertCollection(his.houseRepo, collection, node, isFull)
			count++
		case <-done:
			break Loop
		}
	}
	if len(collection) > 0 {
		collection = insertCollection(his.houseRepo, collection, nil, isFull)
	}
	finish := time.Now()
	fmt.Println("Количество добавленных записей в адреса:", count)
	fmt.Println("Время выполнения адресов:", finish.Sub(start))
}
