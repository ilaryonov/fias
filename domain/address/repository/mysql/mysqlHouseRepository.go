package mysql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
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

func (hr *HouseRepository) InsertUpdateCollection(collection []interface{}) {
	/*var aoguid []string
	var forInsert []interface{}

	for _, item := range collection {
		aoguid = append(aoguid, item.(entity.HouseObject).Houseguid)
	}
	foundedAddresses := hr.CheckByGuids(aoguid)

	for _, item := range collection {
		if len(foundedAddresses[item.(entity.HouseObject).Houseguid].Houseguid) > 0 {
			addr := item.(entity.HouseObject)
			addr.ID = foundedAddresses[item.(entity.HouseObject).Houseguid].ID
			hr.DB.Save(&addr)
		} else {
			forInsert = append(forInsert, item.(entity.HouseObject))
		}
	}*/

	tableName := entity.HouseObject{}.TableName()

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
