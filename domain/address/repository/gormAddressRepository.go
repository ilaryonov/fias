package repository

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
)

type gormAddressRepo struct {
	DB *gorm.DB
}

func NewGormAddressRepository(db *gorm.DB) address.RepositoryInterface {
	return &gormAddressRepo{DB: db}
}
