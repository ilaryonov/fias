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
	v, err := h.versionService.GetVersionInfo()
	log.Println(v, err)
}