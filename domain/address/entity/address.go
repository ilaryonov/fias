package entity

import (
	"github.com/jinzhu/gorm"
)

type AddrObject struct {
	gorm.Model
	Aoguid     string `xml:"AOGUID,attr" gorm:"unique_index`
	Aolevel    string `xml:"AOLEVEL,attr" gorm:"index:city;index:street`
	Parentguid string `xml:"PARENTGUID,attr" gorm:"index:parent;index:street`
	Shortname  string `xml:"SHORTNAME,attr" gorm:"index:city`
	Formalname string `xml:"FORMALNAME,attr" gorm:"index:city;index:street`
	Offname    string `xml:"OFFNAME,attr"`
	Postalcode string `xml:"POSTALCODE,attr"`
	Actstatus  string `xml:"ACTSTATUS,attr"`
}

func GetAddressXmlFile() string {
	return "AS_ADDROBJ_"
}

func (*AddrObject) TableName() string {
	return "fias_address"
}
