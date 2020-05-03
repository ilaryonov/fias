package service

import (
	"github.com/sirupsen/logrus"
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionService struct {
	versionRepo version.VersionRepositoryInterface
	logger      logrus.Logger
}

func NewVersionService(versionRepo version.VersionRepositoryInterface, logger logrus.Logger) *VersionService {
	return &VersionService{
		versionRepo: versionRepo,
		logger:      logger,
	}
}

func (a *VersionService) GetLastVersionInfo() *entity.Version {
	version := a.versionRepo.GetVersion()
	return version
}

func (a *VersionService) UpdateVersion(version *entity.Version) error {
	error := a.versionRepo.UpdateVersion(version)
	return error
}
