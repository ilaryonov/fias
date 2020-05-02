package mysql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"log"
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

func (a *AddressRepository) BatchInsert(collection []interface{}) error {
	first := collection[0]
	switch first.(type) {
	case entity.AddrObject:
		log.Println(collection[0])
		break
	case *entity.HouseObject:
		log.Println(collection[0])
	default:
		break
	}

	return nil
}

func (a *AddressRepository) BatchInsertAddress(collection []interface{}) error {

	return nil
}

func (a *AddressRepository) BatchInsertHouse(collection []interface{}) error {

	return nil
}
