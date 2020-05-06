package grpc

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
	addressGrpc "gitlab.com/ilaryonov/fiascli-clean/domain/address/delivery/grpc/address"
)

type AddressService struct {
	addressRepo *address.AddressRepositoryInterface
	logger      logrus.Logger
}

func NewAddressService(addressRepo address.AddressRepositoryInterface, logger logrus.Logger) *AddressService {
	return &AddressService{
		addressRepo: &addressRepo,
		logger:      logger,
	}
}

func (a *AddressService) GetByGuid(guid *addressGrpc.GuidRequest) (*addressGrpc.Address, error) {
	add := addressGrpc.Address{}
	return &add, nil
}
