package mysql

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	"gitlab.com/ilaryonov/fiascli-clean/models"
)

type AddressRepository struct {
	DB *gorm.DB
}

func (a AddressRepository) BatchInsert() bool {

}

func NewMysqlAddressRepository(db *gorm.DB) address.AddressRepositoryInterface {
	return &AddressRepository{DB: db}
}


func (a AddressRepository) GetByFormalname(term string) (*models.AddrObject, error) {
	panic("implement me")
}

func (a AddressRepository) GetCityByFormalname(term string) (*models.AddrObject, error) {
	panic("implement me")
}