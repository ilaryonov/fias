package service

import (
	version "gitlab.com/ilaryonov/fiascli-clean/domain/version"
)

type VersionService struct {
	versionRepo version.VersionRepositoryInterface
}

func NewVersionService(versionRepo version.VersionRepositoryInterface) *VersionService {
	return &VersionService{
		versionRepo: versionRepo,
	}
}
