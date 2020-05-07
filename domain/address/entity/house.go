package entity

type HouseObject struct {
	ID         uint   `gorm:"primary_key"`
	ParentGuid string `xml:"AOGUID,attr"`
	Houseguid  string `xml:"HOUSEGUID,attr" gorm:"primary_key"`
	Housenum   string `xml:"HOUSENUM,attr" gorm:"index:number"`
	Buildnum   string `xml:"BUILDNUM,attr"`
	Structnum  string `xml:"STRUCTNUM,attr"`
	Postalcode string `xml:"POSTALCODE,attr"`
	EndDate    string `xml:"ENDDATE,attr"`
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
