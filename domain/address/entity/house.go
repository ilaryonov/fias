package entity

import (
	"github.com/jinzhu/gorm"
)

type HouseObject struct {
	gorm.Model
	ParentGuid string `xml:"AOGUID,attr"`
	Houseguid  string `xml:"HOUSEGUID,attr"`
	Housenum   string `xml:"HOUSENUM,attr"`
	Buildnum   string `xml:"BUILDNUM,attr"`
	Structnum  string `xml:"STRUCTNUM,attr"`
	Postalcode string `xml:"POSTALCODE,attr"`
	EndDate string `xml:"ENDDATE,attr"`
}

type HouseObjects struct {
	Object []HouseObjects
}

func GetHouseXmlFile() string {
	return "AS_HOUSE_"
}

func (o HouseObject) TableName() string {
	return "fias_house"
}
