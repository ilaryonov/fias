package mysql

import (
	"github.com/jinzhu/gorm"
	"github.com/ilaryonov/fias/domain/address"
	"github.com/ilaryonov/fias/domain/address/entity"
)

type HouseRepository struct {
	DB *gorm.DB
}

func (hr *HouseRepository) GetByAddressGuid(term string) (*entity.AddrObject, error) {
	panic("implement me")
}

func NewMysqlHouseRepository(db *gorm.DB) address.HouseRepositoryInterface {
	return &HouseRepository{DB: db}
}

func (hr *HouseRepository) InsertUpdateCollection(collection []interface{}, isFull bool) {
	tableName := entity.HouseObject{}.TableName()
	var aoguid []string
	var forInsert []interface{}

	if isFull {
		forInsert = collection
	} else {
		for _, item := range collection {
			aoguid = append(aoguid, item.(entity.HouseObject).Houseguid)
		}
		foundedAddresses := hr.CheckByGuids(aoguid)

		for _, item := range collection {
			if len(foundedAddresses[item.(entity.HouseObject).Houseguid].Houseguid) > 0 {
				house := item.(entity.HouseObject)
				house.Houseguid = foundedAddresses[item.(entity.HouseObject).Houseguid].Houseguid
				hr.DB.Save(&house)
			} else {
				forInsert = append(forInsert, item.(entity.HouseObject))
			}
		}
	}

	if len(collection) > 0 {
		batchInsert(hr.DB, collection, tableName)
	}
}

func (hr *HouseRepository) CheckByGuids(guids []string) map[string]entity.HouseObject {
	var houses []entity.HouseObject
	result := make(map[string]entity.HouseObject)
	hr.DB.Select([]string{"houseguid"}).Where("houseguid IN (?)", guids).Find(&houses)
	for _, item := range houses {
		result[item.Houseguid] = item
	}
	return result
}

func (hr *HouseRepository) Update(item interface{}) {

}
