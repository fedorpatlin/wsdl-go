package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Definitions struct {
	XMLName         xml.Name  `xml:"definitions"`
	TargetNamespace string    `xml:"targetNamespace,attr"`
	Name            string    `xml:"name,attr"`
	Types           Type      `xml:"types"`
	Messages        []Message `xml:"message"`
	PortType        PortType  `xml:"portType"`
	Binding         Binding   `xml:"binding"`
	Service         Service   `xml:"service"`
}

type Type struct {
	Schemas []Schema `xml:"http://www.w3.org/2001/XMLSchema schema"`
}

type Message struct {
	Name string `xml:"name,attr"`
	Part Part   `xml:"part"`
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
	Name    string             `xml:"name,attr"`
	Type    string             `xml:"type,attr"`
	SoapBinding SoapBinding        `xml:"binding"`
	Operations  []BindingOperation `xml:"operation"`
}

type SoapBinding struct {
	XMLName       xml.Name `xml:"http://schemas.xmlsoap.org/wsdl/soap/ binding"`
	Transport string   `xml:"transport,attr"`
	Style     string   `xml:"style,attr"`
}

type BindingOperation struct {
	Name      string        `xml:"name,attr"`
	SoapOperation SoapOperation `xml:"http://schemas.xmlsoap.org/wsdl/soap/ operation"`
	Input         SoapBodyIO    `xml:"input"`
	Output        SoapBodyIO    `xml:"output"`
	Fault         SoapBody      `xml:"fault>http://schemas.xmlsoap.org/wsdl/soap/ fault"`
}

type SoapOperation struct {
	SoapAction string `xml:"soapAction,attr"`
}

type SoapBodyIO struct {
	SoapBody SoapBody `xml:"http://schemas.xmlsoap.org/wsdl/soap/ body"`
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
	Address ServiceAddress `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
}

type ServiceAddress struct {
	XMLName  xml.Name `xml:"http://schemas.xmlsoap.org/wsdl/soap/ address"`
	Location string   `xml:"location,attr"`
}

var wsdlFile = flag.String("w", "", "WSDL file with full path")
var xsdFile = flag.String("x", "", "XSD file with full path")
var packageName = flag.String("p", "", "Package name")
var outFile = flag.String("o", "", "Output file")

// wsdl -w="C:\Temp\wsdl\CartaoEndpointService.wsdl" -x="C:\Temp\wsdl\CartaoEndpointService_schema1.xsd" -p="main" -o="C:\Temp\service.go"
// wsdl -w="C:\Temp\wsdl\authendpointservice.wsdl" -x="C:\Temp\wsdl\AuthEndpointService_schema1.xsd" -p="login" -o="C:\Temp\auth_service.go"
func main() {
	flag.Parse()

	if *wsdlFile == "" || *xsdFile == "" || *packageName == "" || *outFile == "" {
		flag.Usage()
		return
	}

	wf, err := os.Open(*wsdlFile)
	if err != nil {
		exit(err)
	}
	defer wf.Close()

	bw, err := ioutil.ReadAll(wf)
	if err != nil {
		exit(err)
	}

	var d Definitions
	err = xml.Unmarshal(bw, &d)
	if err != nil {
		exit(err)
	}

	xf, err := os.Open(*xsdFile)
	if err != nil {
		exit(err)
	}
	defer xf.Close()

	bx, err := ioutil.ReadAll(xf)
	if err != nil {
		exit(err)
	}

	var s Schema
	err = xml.Unmarshal(bx, &s)
	if err != nil {
		exit(err)
	}

	err = os.Remove(*outFile)
	if err != nil && !os.IsNotExist(err) {
		exit(err)
	}

	f, err := os.Create(*outFile)
	if err != nil {
		exit(err)
	}
	defer f.Close()

	buf := bufio.NewWriter(f)

	// create de service file
	create(&d, &s, buf)

	buf.Flush()
}

func exit(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

func create(d *Definitions, s *Schema, b *bufio.Writer) {
	funcMap := template.FuncMap{
		// The name "title" is what the function will be called in the template text.
		"StringHasValue": StringHasValue,
		"TagDelimiter": TagDelimiter,
	}
	
	tmpl, err := template.New("").Funcs(funcMap).Parse(tmplService)
	if err != nil {
		exit(err)
	}

	type Method struct {
		Name         string
		Action       string
		InputType    string
		OutputType   string
		MessageIn    string
		MessageOut   string
		ParamInName  string
		ParamOutName string
	}

	type SoapMessage struct {
		Name         string
		XMLName      string
		Action       string
		ParamName    string
		XMLParamName string
		ParamType    string
		Input        bool
	}

	type Field struct {
		Name    string
		Type    string
		XMLName string
	}

	type StructType struct {
		Name   string
		Fields []Field
	}

	var data struct {
		PackageName string
		ServiceName string
		ServiceUrl  string
		Messages    []SoapMessage
		Methods     []Method
		Types       []StructType
	}

	data.PackageName = *packageName
	data.ServiceName = d.Service.Name
	data.ServiceUrl = d.Service.Port.Address.Location
	data.Methods = make([]Method, 0)
	data.Messages = make([]SoapMessage, 0)

	for i := 0; i < len(s.ComplexTypes); i++ {
		// check if complex type not in elements
		found := false
		for j := 0; j < len(s.Elements); j++ {
			if s.ComplexTypes[i].Name == s.Elements[j].Name {
				found = true
				break
			}
		}
		var t StructType
		if !found {
			t = StructType{}
			t.Name = exportableSymbol(s.ComplexTypes[i].Name)
			t.Fields = make([]Field, 0)
			if s.ComplexTypes[i].Content == nil {
				for ii := 0; ii < len(s.ComplexTypes[i].Sequence); ii++ {
					fi := Field{}
					fi.Name = exportableSymbol(s.ComplexTypes[i].Sequence[ii].Name)
					fi.Type = decodeType(s.ComplexTypes[i].Sequence[ii])
					fi.XMLName = s.ComplexTypes[i].Sequence[ii].Name
					t.Fields = append(t.Fields, fi)
				}
			} else {
				t.Fields = append(t.Fields, Field{Name: exportableSymbol(s.ComplexTypes[i].Content.Extension.Base[4:])})
				for ii := 0; ii < len(s.ComplexTypes[i].Content.Extension.Sequence); ii++ {
					fi := Field{}
					fi.Name = exportableSymbol(s.ComplexTypes[i].Content.Extension.Sequence[ii].Name)
					fi.Type = decodeType(s.ComplexTypes[i].Content.Extension.Sequence[ii])
					fi.XMLName = s.ComplexTypes[i].Content.Extension.Sequence[ii].Name
					t.Fields = append(t.Fields, fi)
				}
			}
			data.Types = append(data.Types, t)
		}
	}

	for i := 0; i < len(d.PortType.Operations); i++ {
		m := Method{}
		m.Name = exportableSymbol(d.PortType.Operations[i].Name)
		// TODO: get correct action in binding area
		m.Action = ""

		// find input parameter type
		e := findElement(s, d.PortType.Operations[i].Input.Message)
		c := findComplexType(s, e.Name)

		si := strings.Index(c.Sequence[0].Type, ":")

		message := SoapMessage{}
		message.Name = exportableSymbol(c.Name)
		message.XMLName = c.Name
		message.ParamName = exportableSymbol(c.Sequence[0].Name)
		message.XMLParamName = c.Sequence[0].Name
		message.ParamType = exportableSymbol(c.Sequence[0].Type[si+1:])
		message.Input = true

		data.Messages = append(data.Messages, message)

		m.InputType = exportableSymbol(c.Sequence[0].Type[si+1:])
		m.MessageIn = message.Name
		m.ParamInName = message.ParamName

		// find output parameter type
		e = findElement(s, d.PortType.Operations[i].Output.Message)
		c = findComplexType(s, e.Name)

		si = strings.Index(c.Sequence[0].Type, ":")

		message = SoapMessage{}
		message.Name = exportableSymbol(c.Name)
		message.XMLName = c.Name
		message.ParamName = exportableSymbol(c.Sequence[0].Name)
		message.XMLParamName = c.Sequence[0].Name
		message.ParamType = exportableSymbol(c.Sequence[0].Type[si+1:])
		message.Input = false

		data.Messages = append(data.Messages, message)

		m.OutputType = exportableSymbol(c.Sequence[0].Type[si+1:])
		m.MessageOut = message.Name
		m.ParamOutName = message.ParamName

		data.Methods = append(data.Methods, m)
	}

	err = tmpl.Execute(b, data)
	if err != nil {
		exit(err)
	}
}

func findElement(s *Schema, t string) *Element {
	for i := 0; i < len(s.Elements); i++ {
		if s.Elements[i].Type == t {
			return &s.Elements[i]
		}
	}

	return nil
}

func findComplexType(s *Schema, n string) *ComplexType {
	for i := 0; i < len(s.ComplexTypes); i++ {
		if s.ComplexTypes[i].Name == n {
			return &s.ComplexTypes[i]
		}
	}

	return nil
}

func exportableSymbol(s string) string {
	s = strings.ToUpper(s[0:1]) + s[1:]
	return s
}

func decodeType(e Element) string {
	t := e.Type
	if t[0:2] == "xs" {
		switch t[3:] {
		case "string":
			return "string"
		case "boolean":
			return "bool"
		case "decimal", "double":
			return "float64"
		default:
			return "nil"
		}
	} else if t[0:3] == "tns" {	
		ty := exportableSymbol(t[4:])
		if e.MaxOccurs == "unbounded" {
			ty = "[]" + ty
		}
		return ty
	}

	panic("unknown type")
}

func StringHasValue(s string) bool {
	if s != "" {
		return true
	}
	return false
}

func TagDelimiter() string {
	return "`"
}