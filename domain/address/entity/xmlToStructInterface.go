package entity

import "encoding/xml"

type XmlToStructInterface interface {
	UnmarshalXml(d *xml.Decoder, start *xml.StartElement) (XmlToStructInterface, error)
}
