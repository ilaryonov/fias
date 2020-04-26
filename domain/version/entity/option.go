package entity

import "github.com/jinzhu/gorm"

type Version struct {
	gorm.Model
	Version int
}

func (o *Version) TableName() string {
	return "version"
}