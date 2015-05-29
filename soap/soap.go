package soap

import (
	"encoding/xml"
)

type Envelope struct {
	XMLName xml.Name    `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
//	XSI     string      `xml:"xmlns xsi,attr,omitempty"`
//	XSD     string      `xml:"xmlns xsd,attr,omitempty"`
	Soap    string      `xml:"xmlns soap,attr"`
	Body    Body `xml:"Body"`
}

func NewEnvelope() *Envelope {
	se := &Envelope{}
//	se.XSI = "http://www.w3.org/2001/XMLSchema-instance"
//	se.XSD = "http://www.w3.org/2001/XMLSchema"
	se.Soap = "http://schemas.xmlsoap.org/soap/envelope/"

	return se
}

type Body struct {
	XMLName xml.Name    `xml:"Body"`
	Content interface{} `xml:",any"`
}

type Fault struct {
	XMLName     xml.Name `xml:"Fault"`
	FaultCode   string   `xml:"faultcode"`
	FaultString string   `xml:"faultstring"`
	Detail      string   `xml:"detail"`
}
