package address

import "gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"

type AddressRepositoryInterface interface {
	GetByFormalname(term string) (*entity.AddrObject, error)
	GetCityByFormalname(term string) (*entity.AddrObject, error)
	InsertUpdateCollection(collection []interface{}, isFull bool)
	GetByGuid(guid string) entity.AddrObject
	GetCities() []entity.AddrObject
	GetCitiesByTerm(term string, count int64) []entity.AddrObject
}

type HouseRepositoryInterface interface {
	GetByAddressGuid(term string) (*entity.AddrObject, error)
	InsertUpdateCollection(collection []interface{}, isFull bool)
}

type InsertUpdateInterface interface {
	InsertUpdateCollection(collection []interface{}, isFull bool)
}
