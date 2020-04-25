package address

import (
	"context"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionRepositoryInterface interface {
	GetVersion(ctx context.Context) (*entity.Option, error)
	UpdateVersion(ctx context.Context, version *entity.Option) error
	CreateVersion(ctx context.Context, version *entity.Option) error
}
