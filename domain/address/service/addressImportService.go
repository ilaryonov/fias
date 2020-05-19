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

func (a *AddressImportService) Import(filePath string, wg *sync.WaitGroup, isFull bool) {
	defer wg.Done()
	start := time.Now()
	addressChannel := make(chan interface{})
	done := make(chan bool)
	//defer close(addressChannel)
	go helper.ParseFile(filePath, addressChannel, done, func(decoder *xml.Decoder, se *xml.StartElement) (interface{}, error) {
		if se.Name.Local == "Object" {
			result := addressEntity.AddrObject{}
			err := decoder.DecodeElement(&result, se)

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
			collection = insertCollection(a.addressRepo, collection, node, isFull)
			count++
		case <-done:
			break Loop
		}
	}
	if len(collection) > 0 {
		collection = insertCollection(a.addressRepo, collection, nil, isFull)
	}
	finish := time.Now()
	fmt.Println("Количество добавленных записей в адреса:", count)
	fmt.Println("Время выполнения адресов:", finish.Sub(start))
}
