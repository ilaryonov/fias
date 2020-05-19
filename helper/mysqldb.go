package helper

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"github.com/ilaryonov/fiasdomain/address/entity"
	entity2 "github.com/ilaryonov/fiasdomain/version/entity"
)

func InitMysqlGormDb() *gorm.DB {
	db, err := gorm.Open("mysql", viper.GetString("db.dsn"))
	if err != nil {
		panic("db connection refused")
	}
	//defer db.Close()

	db.LogMode(viper.GetBool("db.LogMode"))
	if viper.GetBool("db.debug") {
		db.Debug()
	}
	db.Set("gorm:table_options", "charset=utf8")
	db.AutoMigrate(&entity.AddrObject{})
	db.AutoMigrate(&entity.HouseObject{})
	db.AutoMigrate(&entity2.Version{})
	return db
}
