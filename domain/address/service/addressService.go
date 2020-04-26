package service

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
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

func (a *AddressService) CheckUpdates() {
	a.logger.Warn("run checkUpdates")
}
