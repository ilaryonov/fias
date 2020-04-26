package cli

import (
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/service"
)

type Handler struct {
	addressService service.AddressService
}

func NewHandler(a service.AddressService) *Handler {
	return &Handler{
		addressService: a,
	}
}

func (h *Handler) CheckUpdates() {
	h.addressService.CheckUpdates()
}