package main

import (
	"encoding/xml"
)

type Definitions struct {
	XMLName             xml.Name `xml:"defintions"`
	AttrTargetNamespace string   `xml:"targetNamespace,attr"`
	AttrName            string   `xml:"name,attr"`
	Services            []Service
}

type Types struct {
	XMLName xml.Name `xml:"types"`
	Schemas []Schemas
}

type TypeSchema struct {
	XMLName xml.Name `xml:"xsd:schema"`
	AttrXsd string   `xml:"xmlns:xsd,attr"`
	Imports TypeImport
}

type TypeImport struct {
	XMLName            xml.Name `xml:"xsd:import"`
	AttrSchemaLocation string   `xml:"schemaLocation"`
	AttrNamespace      string   `xml:"namespace"`
}

type Service struct {
	XMLName  xml.Name `xml:"service"`
	AttrName string   `xml:"name,attr"`
}

type ServicePort struct {
	XMLName     xml.Name `xml:"port"`
	AttrName    string   `xml:"name,attr"`
	AttrBinding string   `xml:"binding,attr"`
	Address     ServiceAddress
}

type ServiceAddress struct {
	XMLName      xml.Name `xml:"soap:address"`
	AttrLocation string   `xml:"location,attr"`
}

func main() {
}
