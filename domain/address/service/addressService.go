package service

import (
	"gitlab.com/ilaryonov/fiascli-clean/domain/address"
)

type AddressService struct {
	addressRepo address.AddressRepositoryInterface
}

func NewAddressService(addressRepo address.AddressRepositoryInterface) *AddressService {
	return &AddressService{
		addressRepo: addressRepo,
	}
}
