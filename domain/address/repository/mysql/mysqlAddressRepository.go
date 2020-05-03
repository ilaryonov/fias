package mysql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
)

type AddressRepository struct {
	DB *gorm.DB
}

func NewMysqlAddressRepository(db *gorm.DB) address.AddressRepositoryInterface {
	return &AddressRepository{DB: db}
}

func (a *AddressRepository) GetByFormalname(term string) (*entity.AddrObject, error) {
	panic("implement me")
}

func (a AddressRepository) GetCityByFormalname(term string) (*entity.AddrObject, error) {
	panic("implement me")
}

func (a *AddressRepository) InsertUpdateCollection(collection []interface{}) {
	var aoguid []string
	var forInsert []interface{}

	for _, item := range collection {
		aoguid = append(aoguid, item.(entity.AddrObject).Aoguid)
	}
	foundedAddresses := a.CheckByGuids(aoguid)

	for _, item := range collection {
		if len(foundedAddresses[item.(entity.AddrObject).Aoguid].Aoguid) > 0 {
			/*addr := item.(entity.AddrObject)
			addr.ID = foundedAddresses[item.(entity.AddrObject).Aoguid].ID
			a.DB.Save(&addr)*/
		} else {
			forInsert = append(forInsert, item.(entity.AddrObject))
		}
	}

	first := collection[0]
	var tableName string
	switch first.(type) {
	case entity.AddrObject:
		tableName = entity.AddrObject{}.TableName()
		break
	case entity.HouseObject:
		tableName = entity.HouseObject{}.TableName()
	default:
		break
	}
	if len(forInsert) > 0 {
		batchInsert(a.DB, forInsert, tableName)
	}
}

func (a *AddressRepository) BatchInsertAddress(collection []interface{}) error {

	return nil
}

func (a *AddressRepository) BatchInsertHouse(collection []interface{}) error {

	return nil
}

func (a *AddressRepository) CheckByGuids(guids []string) map[string]entity.AddrObject {
	var addresses []entity.AddrObject
	result := make(map[string]entity.AddrObject)
	a.DB.Select([]string{"id, aoguid"}).Where("aoguid IN (?)", guids).Find(&addresses)
	for _, item := range addresses {
		result[item.Aoguid] = item
	}
	return result
}

func (a *AddressRepository) Update(item interface{}) {

}
