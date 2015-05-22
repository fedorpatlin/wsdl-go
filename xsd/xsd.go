package xsd

import (
	"encoding/xml"
)

type Schema struct {
	XMLName            xml.Name      `xml:"http://www.w3.org/2001/XMLSchema schema"`
	TNS                string        `xml:"xmlns tns",attr`
	XS                 string        `xml:"xmlns xs,attr"`
	TargetNamespace    string        `xml:"targetNamespace,attr"`
	ElementFormDefault string        `xml:"elementFormDefault,attr"`
	Version            string        `xml:"version,attr"`
	Elements           []Element     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []ComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleTypes        []SimpleType  `xml:http://www.w3.org/2001/XMLSchema simpleType`
}

type Element struct {
	XMLName      xml.Name     `xml:"http://www.w3.org/2001/XMLSchema element"`
	Type         string       `xml:"type,attr"`
	Nillable     string       `xml:"nillable,attr"`
	MinOccurs    string       `xml:"minOccurs,attr"`
	MaxOccurs    string       `xml:"maxOccurs,attr"`
	Form         string       `xml:"form,attr"`
	Name         string       `xml:"name,attr"`
	ComplexTypes *ComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type ComplexType struct {
	XMLName  xml.Name        `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Name     string          `xml:"name,attr"`
	Abstract bool            `xml:"abstract,attr"`
	Sequence []Element       `xml:"sequence>element"`
	Content  *ComplexContent `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
}

type ComplexContent struct {
	XMLName   xml.Name  `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
	Extension Extension `xml:"http://www.w3.org/2001/XMLSchema extension"`
}

type Extension struct {
	XMLName  xml.Name  `xml:"http://www.w3.org/2001/XMLSchema extension"`
	Base     string    `xml:"base,attr"`
	Sequence []Element `xml:"sequence>element"`
}

type Import struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2001/XMLSchema import"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Namespace      string   `xml:"namespace,attr"`
}

type SimpleType struct {
	Name     string       `xml:"name,attr"`
	Sequence *Restriction `xml:"restriction"`
}

type Sequence struct {
	Elements []*Element `xml:"http://www.w3.org/2001/XMLSchema element"`
}

type Restriction struct {
	Base         string        `xml:"base,attr"`
	Pattern      *Pattern      `xml:"http://www.w3.org/2001/XMLSchema pattern"`
	Enumeration  []Enumeration `xml:"enumeration"`
	MinInclusive *MinInclusive `xml:"http://www.w3.org/2001/XMLSchema minInclusive"`
	MaxInclusive *MaxInclusive `xml:"http://www.w3.org/2001/XMLSchema maxInclusive"`
}

type Enumeration struct {
	Value string `xml:"value,attr"`
}

type Pattern struct {
	Value string `xml:"value,attr"`
}

type MinInclusive struct {
	Value string `xml:"value,attr"`
}

type MaxInclusive struct {
	Value string `xml:"value,attr"`
}
