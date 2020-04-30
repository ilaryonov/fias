package helper

import (
	"encoding/xml"
	"fmt"
	"gitlab.com/ilaryonov/fiascli-clean/domain/address/entity"
	"os"
)

func ParseFile(fileName string, c chan entity.XmlToStructInterface, done chan bool, str entity.XmlToStructInterface) {
	xmlFile, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file: ", err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)

	for {
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		switch se := t.(type) {
		case xml.StartElement:
			data, err := str.UnmarshalXml(decoder, &se)
			if err == nil {
				c <- data
			}
		}
	}
	done <- true
}
