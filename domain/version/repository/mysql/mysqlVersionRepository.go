package mysql

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	versionRepo "gitlab.com/ilaryonov/fiascli-clean/domain/version"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionRepository struct {
	DB *gorm.DB
}

func (v *VersionRepository) GetVersion(ctx context.Context) (*entity.Option, error) {
	var ErrNotFound = errors.New("not found")
	version := entity.Option{Name: "version"}
	//v.DB.Create(&entity.Option{Name: "version", Value: 610})
	v.DB.Where(entity.Option{Name: "version"}).First(&version)
	if version.Value <= 0 {
		return nil, ErrNotFound
	} else {
		return &version, nil
	}
}

func (v *VersionRepository) UpdateVersion(ctx context.Context, version *entity.Option) error {
	panic("implement me")
}

func (v *VersionRepository) CreateVersion(ctx context.Context, version *entity.Option) error {
	panic("implement me")
}

func NewMysqlVersionRepository(db *gorm.DB) versionRepo.VersionRepositoryInterface {
	return &VersionRepository{DB: db}
}