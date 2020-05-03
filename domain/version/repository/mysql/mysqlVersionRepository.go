package mysql

import (
	"errors"
	"github.com/jinzhu/gorm"
	versionRepo "gitlab.com/ilaryonov/fiascli-clean/domain/version"
	"gitlab.com/ilaryonov/fiascli-clean/domain/version/entity"
)

type VersionRepository struct {
	DB *gorm.DB
}

func (v *VersionRepository) GetVersion() *entity.Version {
	version := entity.Version{}
	//v.DB.Create(&entity.Version{Version: 610})
	v.DB.Last(&version)
	if version.Version <= 0 {
		return &entity.Version{}
	} else {
		return &version
	}
}

func (v *VersionRepository) UpdateVersion(version *entity.Version) error {
	var ErrNotFound = errors.New("not found")
	v.DB.Save(version)
	return ErrNotFound
}

func (v *VersionRepository) CreateVersion(version *entity.Version) error {
	panic("implement me")
}

func NewMysqlVersionRepository(db *gorm.DB) versionRepo.VersionRepositoryInterface {
	return &VersionRepository{DB: db}
}
