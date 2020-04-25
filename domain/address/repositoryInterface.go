package address

import "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"

type AddressRepositoryInterface interface {
	GetByFormalname(term string) (*entity.AddrObject, error)
	GetCityByFormalname(term string) (*entity.AddrObject, error)
	BatchInsert() bool
}