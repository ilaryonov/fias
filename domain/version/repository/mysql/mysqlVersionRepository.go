package mysql

import (
	"context"
	"github.com/jinzhu/gorm"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionRepository struct {
	DB *gorm.DB
}

func (v *VersionRepository) GetVersion(ctx context.Context) (*entity.Option, error) {
	panic("implement me")
}

func (v *VersionRepository) UpdateVersion(ctx context.Context, version *entity.Option) error {
	panic("implement me")
}

func (v *VersionRepository) CreateVersion(ctx context.Context, version *entity.Option) error {
	panic("implement me")
}

func NewMysqlVersionRepository(db *gorm.DB) address.VersionRepositoryInterface {
	return &VersionRepository{DB: db}
}