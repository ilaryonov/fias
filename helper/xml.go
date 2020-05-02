package helper

import (
	"encoding/xml"
	"fmt"
	"os"
)

type ParseElement func(decoder *xml.Decoder, element *xml.StartElement) interface{}

func ParseFile(fileName string, c chan interface{}, done chan bool, ParseElement ParseElement) {
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
			data := ParseElement(decoder, &se)
			if err == nil {
				c <- data
			}
		}
	}
	done <- true
}
