package main

var tmpl_intro = `
{{define "Tag"}}{{TagDelimiter}}xml:"{{.}}"{{TagDelimiter}}{{end}}
package {{.PackageName}}

import "encoding/xml"

type {{.ServiceName}}Struct struct {
	Url string
}

func New{{.ServiceName}}(url string)(svc {{.ServiceName}}Struct){
	svc.Url=url
	return
}
`
var tmpl_complex_type = `{{range .Wsdl.Types.Schemas}}
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

var tmpl_elements = `{{range .Wsdl.Types.Schemas}}{{range .Elements}}
type {{exportableSymbol .Name}} struct{
	{{range .ComplexTypes.Sequence}}{{exportableSymbol .Name}} {{(decodeType .Type).Local}}  {{template "Tag" .Name}}
	{{end}}
}
{{end}}{{end}}`

var tmpl_operations = `
type {{.Wsdl.PortType.Name}} interface {
	{{range .Wsdl.PortType.Operations}}
	{{.Name}}(inMessage *Msg{{exportableSymbol .Input.Message}})(outMessage *Msg{{exportableSymbol .Output.Message}}{{if .Fault.Message}}, faultMessage *{{.Fault.Message}}{{end}})
	{{end}}
}
{{range .Wsdl.PortType.Operations}}{{$op:=.}}{{$ms:=getMessagesFromOperation $op}}{{range $ms.input}}
// input message name is Name of Operation
type Msg{{.Name}} struct {
	XMLName xml.Name
	{{$n:=.Name}}{{range .Part}}{{exportableSymbol .Element}} {{template "Tag" $n}}{{end}}
}

func NewMsg{{.Name}}() *Msg{{.Name}} {
	var msg Msg{{.Name}}
	msg.XMLName.Local = "{{$op.Name}}"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}
{{end}}
// output message name is Name of a message
{{range $ms.output}}
type Msg{{.Name}} struct {
	XMLName xml.Name
	{{$n:=.Name}}{{range .Part}}{{exportableSymbol .Element}} {{template "Tag" $n}}{{end}}
}

func NewMsg{{.Name}}() *Msg{{.Name}} {
	var msg Msg{{.Name}}
	msg.XMLName.Local = "{{.Name}}"
	msg.XMLName.Space = "urn:SAPCCMS"
	return &msg 
}
{{end}}
{{end}}
`
