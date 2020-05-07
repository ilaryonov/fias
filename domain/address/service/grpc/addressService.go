package grpc

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
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
