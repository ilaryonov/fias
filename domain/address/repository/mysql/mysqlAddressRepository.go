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
func (a *AddressRepository) GetByGuid(guid string) entity.AddrObject {
	addr := entity.AddrObject{}
	a.DB.Where("aoguid = ?", guid).First(&addr)
	return addr
}

func (a AddressRepository) GetCityByFormalname(term string) (*entity.AddrObject, error) {
	panic("implement me")
}

func (a *AddressRepository) InsertUpdateCollection(collection []interface{}, isFull bool) {
	tableName := entity.AddrObject{}.TableName()
	var forInsert []interface{}
	if isFull {
		forInsert = collection
	} else {
		//TODO узкое место, тормозит выгрузка из-за проверок на наличие
		var aoguid []string

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
	}
	if len(forInsert) > 0 {
		batchInsert(a.DB, forInsert, tableName)
	}
}

func (a *AddressRepository) CheckByGuids(guids []string) map[string]entity.AddrObject {
	var addresses []entity.AddrObject
	result := make(map[string]entity.AddrObject)
	a.DB.Select([]string{"aoguid"}).Where("aoguid IN (?)", guids).Find(&addresses)
	for _, item := range addresses {
		result[item.Aoguid] = item
	}
	return result
}