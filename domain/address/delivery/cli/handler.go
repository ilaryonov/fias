package cli

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/service"
	fiasApi "gitlab.com/ilaryonov/fiascli-clean/domain/fiasApi/service"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
)

type Handler struct {
	addressService *service.AddressImportService
	logger logrus.Logger
}

func NewHandler(a *service.AddressImportService, logger logrus.Logger) *Handler {
	return &Handler{
		addressService: a,
		logger: logger,
	}
}

func (h *Handler) CheckUpdates(fiasApi *fiasApi.FiasApiService, versionService *version.VersionService) {
	v, err := versionService.GetLastVersionInfo()
	if err != nil {
		h.logger.Error(err.Error())
	}
	if v.Version > 0 {
		h.addressService.CheckUpdates(fiasApi)
	} else {
		h.addressService.StartFullImport(fiasApi)
	}
}