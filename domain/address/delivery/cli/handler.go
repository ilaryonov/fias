package cli

import (
	"github.com/sirupsen/logrus"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/service"
	fiasApi "gitlab.com/ilaryonov/fiascli-clean/domain/fiasApi/service"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
)

type Handler struct {
	importService *service.ImportService
	logger        logrus.Logger
}

func NewHandler(a *service.ImportService, logger logrus.Logger) *Handler {
	return &Handler{
		importService: a,
		logger:        logger,
	}
}

func (h *Handler) CheckUpdates(fiasApi *fiasApi.FiasApiService, versionService *version.VersionService) {
	v := versionService.GetLastVersionInfo()

	if v.Version > 0 {
		h.importService.CheckUpdates(fiasApi, v.Version)
	} else {
		h.importService.StartFullImport(fiasApi)
	}
}
