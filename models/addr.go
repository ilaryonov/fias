package models

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

type AddrObjects struct {
	Object []AddrObject
}

func (a AddrObject) GetXmlFile() string {
	return "AS_ADDROBJ_"
}

func (a AddrObject) TableName() string {
	return "fias_address"
}

/*func (a *AddrObject) Import(f os.FileInfo, wg *sync.WaitGroup, db *gorm.DB) {
	defer wg.Done()

	start := time.Now()
	path, err := filepath.Abs("/media/ilarionov/hard-disk/fias/" + f.Name())
	xmlFile, err := os.Open(path)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	total := 0

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
			if element == "Object" {

				decoder.DecodeElement(&a, &se)
				a.ID = 0
				//db.Create(&a)
				if a.Actstatus == "0" {
					continue
				}

				//fmt.Println(object.Formalname)
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
	fmt.Println("Время выполнения адресов:", finish.Sub(start))
	fmt.Println(a.TableName(), f.Name())
}*/
