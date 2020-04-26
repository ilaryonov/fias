package mysql

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	versionRepo "gitlab.com/ilaryonov/fiascli-clean/domain/version"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
	"log"
)

type VersionRepository struct {
	DB *gorm.DB
}

func (v *VersionRepository) GetVersion(ctx context.Context) (*entity.Version, error) {
	var ErrNotFound = errors.New("not found")
	version := entity.Version{}
	//v.DB.Create(&entity.Version{Version: 610})
	v.DB.Last(&version)
	if version.Version <= 0 {
		return nil, ErrNotFound
	} else {
		return &version, nil
	}
}

func (v *VersionRepository) UpdateVersion(ctx context.Context, version *entity.Version) error {
	var ErrNotFound = errors.New("not found")
	v.DB.Save(version)
	log.Println(213)
	return ErrNotFound
}

func (v *VersionRepository) CreateVersion(ctx context.Context, version *entity.Version) error {
	panic("implement me")
}

func NewMysqlVersionRepository(db *gorm.DB) versionRepo.VersionRepositoryInterface {
	return &VersionRepository{DB: db}
}