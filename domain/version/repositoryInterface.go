package addressGetVersion

import (
	"context"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionRepositoryInterface interface {
	GetVersion(ctx context.Context) (*entity.Version, error)
	UpdateVersion(ctx context.Context, version *entity.Version) error
	CreateVersion(ctx context.Context, version *entity.Version) error
}
