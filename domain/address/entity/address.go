package entity

type AddrObject struct {
	ID        uint `gorm:"primary_key"`
	Aoguid     string `xml:"AOGUID,attr" gorm:"primary_key"`
	Aolevel    string `xml:"AOLEVEL,attr" gorm:"index:city;index:street"`
	Parentguid string `xml:"PARENTGUID,attr" gorm:"index:parent;index:street"`
	Shortname  string `xml:"SHORTNAME,attr" gorm:"index:city"`
	Formalname string `xml:"FORMALNAME,attr" gorm:"index:city;index:street"`
	Offname    string `xml:"OFFNAME,attr"`
	Postalcode string `xml:"POSTALCODE,attr"`
	Actstatus  string `xml:"ACTSTATUS,attr"`
}

func GetAddressXmlFile() string {
	return "AS_ADDROBJ_"
}

func (a AddrObject) TableName() string {
	return "fias_address"
}
