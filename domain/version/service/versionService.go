package service

import (
	"context"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionService struct {
	versionRepo version.VersionRepositoryInterface
}

func NewVersionService(versionRepo version.VersionRepositoryInterface) *VersionService {
	return &VersionService{
		versionRepo: versionRepo,
	}
}


func (a *VersionService) GetVersionInfo() (*entity.Option, error) {
	version, error := a.versionRepo.GetVersion(context.Background())
	return version, error
}
