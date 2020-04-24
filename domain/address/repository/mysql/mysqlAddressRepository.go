package mysql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
)

type AddressRepository struct {
	DB *gorm.DB
}

func (a AddressRepository) BatchInsert() bool {

}

func NewMysqlAddressRepository(db *gorm.DB) address.AddressRepositoryInterface {
	return &AddressRepository{DB: db}
}


func (a AddressRepository) GetByFormalname(term string) (*address.AddrObject, error) {
	panic("implement me")
}

func (a AddressRepository) GetCityByFormalname(term string) (*address.AddrObject, error) {
	panic("implement me")
}