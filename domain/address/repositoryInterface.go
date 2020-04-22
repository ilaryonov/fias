package address

import "gitlab.com/ilaryonov/fiascli-clean/models"

type AddressRepositoryInterface interface {
	GetByFormalname(term string) (*models.AddrObject, error)
	GetCityByFormalname(term string) (*models.AddrObject, error)
	BatchInsert() bool
}