package addressGetVersion

import (
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionRepositoryInterface interface {
	GetVersion() *entity.Version
	UpdateVersion(version *entity.Version) error
	CreateVersion(version *entity.Version) error
}
