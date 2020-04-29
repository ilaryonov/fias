package helper

import (
	"encoding/xml"
	"fmt"
	"github.com/spf13/viper"
	"os"
	"path/filepath"
)

type DecodeNode func()

func ParseFile(fileName string, c chan xml.StartElement) {
	path, err := filepath.Abs(viper.GetString("directory.filePath") + fileName)
	xmlFile, err := os.Open(path)
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
			c <- se
		}
	}
}