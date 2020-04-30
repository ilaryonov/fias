package entity

import (
	"encoding/xml"
	"errors"
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

func (a *AddrObject) UnmarshalXml(decoder *xml.Decoder, se *xml.StartElement) (XmlToStructInterface, error) {
	if se.Name.Local == "Object" {
		err := decoder.DecodeElement(a, se)
		a.ID = 0
		if a.Actstatus == "0" {
			return nil, errors.New("not active")
		}
		if err != nil {
			return nil, err
		}
		return a, nil
	}
	return nil, errors.New("not entity")
}

func GetAddressXmlFile() string {
	return "AS_ADDROBJ_"
}

func (a *AddrObject) TableName() string {
	return "fias_address"
}
