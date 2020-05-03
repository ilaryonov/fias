package mysql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"log"
)

type HouseRepository struct {
	DB *gorm.DB
}

func (a *HouseRepository) GetByAddressGuid(term string) (*entity.AddrObject, error) {
	panic("implement me")
}

func NewMysqlHouseRepository(db *gorm.DB) address.HouseRepositoryInterface {
	return &HouseRepository{DB: db}
}

func (a *HouseRepository) InsertUpdateCollection(collection []interface{}) error {
	var aoguid []string
	var forInsert []interface{}

	for _, item := range collection {
		aoguid = append(aoguid, item.(entity.AddrObject).Aoguid)
	}
	foundedAddresses := a.CheckByGuids(aoguid)

	for _, item := range collection {
		if len(foundedAddresses[item.(entity.AddrObject).Aoguid].Aoguid) > 0 {
			addr := item.(entity.AddrObject)
			addr.ID = foundedAddresses[item.(entity.AddrObject).Aoguid].ID
			a.DB.Save(&addr)
		} else {
			forInsert = append(forInsert, item.(entity.AddrObject))
		}
	}

	first := collection[0]
	var tableName string
	switch first.(type) {
	case entity.AddrObject:
		log.Println(collection[0])
		tableName = entity.AddrObject{}.TableName()
		break
	case entity.HouseObject:
		tableName = entity.HouseObject{}.TableName()
	default:
		break
	}
	// If there is no data, nothing to do.
	if len(collection) == 0 {
		return nil
	}
	var err error
	if len(forInsert) > 0 {
		batchInsert(a.DB, forInsert, tableName)
	}

	return err
}

func (a *HouseRepository) CheckByGuids(guids []string) map[string]entity.AddrObject {
	var addresses []entity.AddrObject
	result := make(map[string]entity.AddrObject)
	a.DB.Select([]string{"id, aoguid"}).Where("aoguid IN (?)", guids).Find(&addresses)
	for _, item := range addresses {
		result[item.Aoguid] = item
	}
	return result
}

func (a *HouseRepository) Update(item interface{}) {

}
