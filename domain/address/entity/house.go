package entity

import (
	"encoding/xml"
	"errors"
	"github.com/jinzhu/gorm"
	"time"
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

func (o *HouseObject) UnmarshalXml(decoder *xml.Decoder, se *xml.StartElement) (XmlToStructInterface, error) {
	layoutISO := "2006-01-02"
	if se.Name.Local == "House" {
		err := decoder.DecodeElement(o, se)
		o.ID = 0
		if err != nil {
			return nil, err
		}
		t, _ := time.Parse(layoutISO, o.EndDate)
		if t.Unix() < time.Now().Unix() {
			return nil, errors.New("old date for house")
		}
		return o, nil
	}
	return nil, errors.New("not house entity")
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

/*func (a *HouseObject) Import(f os.FileInfo, wg *sync.WaitGroup, db *gorm.DB) {
	defer wg.Done()

	fmt.Println(a.TableName(), f.Name())

	start := time.Now()
	path, err := filepath.Abs("/media/ilarionov/hard-disk/fias/" + f.Name())
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer xmlFile.Close()
	total := 0
	decoder := xml.NewDecoder(xmlFile)

	var element string
	var collection []interface{}

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			element = se.Name.Local
			if element == "House" {

				decoder.DecodeElement(&a, &se)

				a.ID = 0

				if len(collection) < 2500 {
					collection = append(collection, *a)
					total++
				} else {
					err := BatchInsert(db, collection)
					if err != nil {
						fmt.Println("error", err.Error())
					}
					collection = collection[:0]
				}
			}
		}
	}
	if len(collection) > 0 {
		err := BatchInsert(db, collection)
		if err != nil {
			fmt.Println("error", err.Error())
		}
	}

	finish := time.Now()
	fmt.Println("Количество добавленных записей в адреса:", total)
	fmt.Println("Время выполнения домов:", finish.Sub(start))
	fmt.Println(a.TableName(), f.Name())
}*/
