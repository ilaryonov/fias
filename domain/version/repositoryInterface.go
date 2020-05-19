package addressGetVersion

import (
	"github.com/ilaryonov/fias/domain/version/entity"
)

type VersionRepositoryInterface interface {
	GetVersion() *entity.Version
	UpdateVersion(version *entity.Version) error
	CreateVersion(version *entity.Version) error
}
