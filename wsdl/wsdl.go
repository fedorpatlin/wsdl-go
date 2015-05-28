package wsdl

import (
	"encoding/xml"

	"github.com/fedorpatlin/wsdl-go/xsd"
)

type Definitions struct {
	XMLName         xml.Name  `xml:"definitions"`
	TargetNamespace string    `xml:"targetNamespace,attr"`
	Name            string    `xml:"name,attr"`
	Types           Type      `xml:"types"`
	Messages        []Message `xml:"message"`
	PortType        PortType  `xml:"portType"`
	Binding         []Binding `xml:"binding"`
	Service         Service   `xml:"service"`
}

type Type struct {
	Schemas []xsd.Schema `xml:"schema"`
}

type Message struct {
	Name string `xml:"name,attr"`
	Part []Part `xml:"part"`
}

type Part struct {
	Name    string `xml:"name,attr"`
	Element string `xml:"element,attr"`
}

type PortType struct {
	Name       string              `xml:"name,attr"`
	Operations []PortTypeOperation `xml:"operation"`
}

type PortTypeOperation struct {
	Name   string                   `xml:"name,attr"`
	Input  PortTypeOperationMessage `xml:"input"`
	Output PortTypeOperationMessage `xml:"output"`
	Fault  PortTypeOperationMessage `xml:"fault"`
}

type PortTypeOperationMessage struct {
	Name    string `xml:"name,attr,omitempty`
	Message string `xml:"message,attr"`
}

type Binding struct {
	Name        string             `xml:"name,attr"`
	Type        string             `xml:"type,attr"`
	SoapBinding SoapBinding        `xml:"binding"`
	Operations  []BindingOperation `xml:"operation"`
}

type SoapBinding struct {
	XMLName   xml.Name `xml:"binding"`
	Transport string   `xml:"transport,attr"`
	Style     string   `xml:"style,attr"`
}

type BindingOperation struct {
	Name          string        `xml:"name,attr"`
	SoapOperation SoapOperation `xml:"operation"`
	Input         SoapBodyIO    `xml:"input"`
	Output        SoapBodyIO    `xml:"output"`
	Fault         SoapBody      `xml:"fault>fault"`
}

type SoapOperation struct {
	SoapAction string `xml:"soapAction,attr"`
}

type SoapBodyIO struct {
	SoapBody SoapBody `xml:"body"`
}

type SoapBody struct {
	Name string `xml:"name,attr,omitempty"`
	Use  string `xml:"use,attr"`
}

type Service struct {
	Name string      `xml:"name,attr"`
	Port ServicePort `xml:"port"`
}

type ServicePort struct {
	XMLName xml.Name       `xml:"port"`
	Name    string         `xml:"name,attr"`
	Binding string         `xml:"binding,attr"`
	Address ServiceAddress `xml:"address"`
}

type ServiceAddress struct {
	XMLName  xml.Name `xml:"address"`
	Location string   `xml:"location,attr"`
}
