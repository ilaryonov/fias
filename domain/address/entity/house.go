package entity

type HouseObject struct {
	ParentGuid string `xml:"AOGUID,attr" gorm:"primary_key"`
	Houseguid  string `xml:"HOUSEGUID,attr" gorm:"index:number"`
	Housenum   string `xml:"HOUSENUM,attr" gorm:"index:number"`
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
