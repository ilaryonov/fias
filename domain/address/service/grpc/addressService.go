package grpc

import (
	"github.com/sirupsen/logrus"
	"github.com/ilaryonov/fiasdomain/address"
	"github.com/ilaryonov/fiasdomain/address/entity"
)

type AddressService struct {
	addressRepo address.AddressRepositoryInterface
	logger      logrus.Logger
}

func NewAddressService(addressRepo address.AddressRepositoryInterface, logger logrus.Logger) *AddressService {
	return &AddressService{
		addressRepo: addressRepo,
		logger:      logger,
	}
}

func (a *AddressService) GetByGuid(guid string) (*entity.AddrObject, error) {
	address := a.addressRepo.GetByGuid(guid)
	return &address, nil
}

func (a *AddressService) GetCities() ([]entity.AddrObject, interface{}) {
	cities := a.addressRepo.GetCities()
	return cities, nil
}

func (a *AddressService) GetCitiesByTerm(term string, count int64) ([]entity.AddrObject, error) {
	cities := a.addressRepo.GetCitiesByTerm(term, count)
	return cities, nil
}

