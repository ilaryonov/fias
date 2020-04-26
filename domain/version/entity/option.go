package entity

import "github.com/jinzhu/gorm"

type Option struct {
	gorm.Model
	Name string
	Value int
}

func(o *Option) GetValueByName() string {
	return o.Name
}

func (o *Option) TableName() string {
	return "option"
}