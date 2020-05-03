package cli

import (
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/service"
	"log"
)

type Handler struct {
	versionService service.VersionService
}

func NewHandler(a service.VersionService) *Handler {
	return &Handler{
		versionService: a,
	}
}

func (h *Handler) GetVersionInfo() {
	//version := entity.Version{Name: "Version", Version: 615}
	//err := h.versionService.UpdateVersion(&version)
	v := h.versionService.GetLastVersionInfo()
	log.Println(v)
}
