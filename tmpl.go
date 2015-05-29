package main

const tmpl_intro = `{{define "Tag"}}{{TagDelimiter}}xml:"{{.}}"{{TagDelimiter}}{{end}}
package {{.PackageName}}
{{define "MsgPrefix"}}Msg{{.}}{{end}}
{{define "NewMsg"}}New{{template "MsgPrefix" .}}{{end}}

import "encoding/xml"

type {{.ServiceName}}Struct struct {
	Url string
}

func New{{.ServiceName}}(url string)(svc {{.ServiceName}}Struct){
	svc.Url=url
	return
}
`
const tmpl_complex_type = `{{range .Wsdl.Types.Schemas}}
{{range .SimpleTypes}}
type {{exportableSymbol .Name}} struct{
	{{exportableSymbol .Name}} {{(decodeType .Restriction.Base).Local}} {{template "Tag" .Name}}
}
{{end}}
{{range .ComplexTypes}}
type {{exportableSymbol .Name}} struct {
{{range .Sequence}}{{exportableSymbol .Name}} {{(decodeElementType .).Local}} {{template "Tag" .Name}}
{{end}}}
{{end}}
{{end}}`

const tmpl_elements = `{{range .Wsdl.Types.Schemas}}{{range .Elements}}
type {{exportableSymbol .Name}} struct{
	{{range .ComplexTypes.Sequence}}{{exportableSymbol .Name}} {{(decodeType .Type).Local}}  {{template "Tag" .Name}}
	{{end}}
}
{{end}}{{end}}`



const tmpl_operations = `{{$tns:=.Wsdl.TargetNamespace}}
// input message name is Name of Operation
// output message name is Name of a Message

type {{.Wsdl.PortType.Name}} interface {
	{{range .Wsdl.PortType.Operations}}
	{{.Name}}(inMessage *Msg{{exportableSymbol .Input.Message}})(outMessage *Msg{{exportableSymbol .Output.Message}}{{if .Fault.Message}}, faultMessage *{{.Fault.Message}}{{end}})
	{{end}}
}
{{range .Wsdl.PortType.Operations}}{{$op:=.}}{{$ms:=getMessagesFromOperation $op}}{{range $ms.input}}

type Msg{{.Name}} struct {
	XMLName xml.Name
	{{$n:=.Name}}{{range .Part}}{{exportableSymbol .Element}} {{template "Tag" $n}}{{end}}
}

func {{template "NewMsg" .Name}}() *Msg{{.Name}} {
	var msg {{template "MsgPrefix" .Name}}
	msg.XMLName.Local = "{{$op.Name}}"
	msg.XMLName.Space = "{{$tns}}"
	return &msg 
}
{{end}}

{{range $ms.output}}
type Msg{{.Name}} struct {
	XMLName xml.Name
	{{$n:=.Name}}{{range .Part}}{{exportableSymbol .Element}} {{template "Tag" $n}}{{end}}
}

func {{template "NewMsg" .Name}}() *Msg{{.Name}} {
	var msg {{template "MsgPrefix" .Name}}
	msg.XMLName.Local = "{{.Name}}"
	msg.XMLName.Space = "{{$tns}}"
	return &msg 
}
{{end}}
{{end}}
`
const tmpl_soap = `
type Envelope struct {
	XMLName xml.Name    {{template "Tag" "http://schemas.xmlsoap.org/soap/envelope/ Envelope"}}
	Soap    string      {{template "Tag" "xmlns soap,attr"}}
	Body    Body 		{{template "Tag" "Body"}}
}

func NewEnvelope() *Envelope {
	se := &Envelope{}
	se.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	return se
}

type Body struct {
	XMLName xml.Name    {{template "Tag" "Body"}}
	Content interface{} {{template "Tag" ",any"}}
}

type Fault struct {
	XMLName     xml.Name {{template "Tag" "Fault"}}
	FaultCode   string   {{template "Tag" "faultcode"}}
	FaultString string   {{template "Tag" "faultstring"}}
	Detail      string   {{template "Tag" "detail"}}
}
`