package main

var tmpl_intro = `package {{.PackageName}}
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
	{{exportableSymbol .Name}} {{(decodeType .Restriction.Base).Local}} {{TagDelimiter}}"xml:{{.Name}}"{{TagDelimiter}}
}
{{end}}
{{range .ComplexTypes}}
type {{exportableSymbol .Name}} struct {
{{range .Sequence}}{{exportableSymbol .Name}} {{(decodeElementType .).Local}} {{TagDelimiter}}"xml:{{.Name}}"{{TagDelimiter}}
{{end}}}
{{end}}
{{end}}`

var tmpl_elements = `{{range .Wsdl.Types.Schemas}}{{range .Elements}}
type {{exportableSymbol .Name}} struct{
	{{range .ComplexTypes.Sequence}}{{exportableSymbol .Name}} {{(decodeType .Type).Local}}  {{TagDelimiter}}xml:"{{.Name}}"{{TagDelimiter}}
	{{end}}
}
{{end}}{{end}}`

var tmpl_messages = `{{range .Wsdl.Messages}}type Ms{{exportableSymbol .Name}} struct{
	{{range .Part}}{{exportableSymbol .Element}} {{TagDelimiter}}xml:"{{getName .Element}}"{{TagDelimiter}}{{end}}
}
{{end}}`

var tmpl_operations = `
type {{.Wsdl.PortType.Name}} interface {
	{{range .Wsdl.PortType.Operations}}
	{{.Name}}(inMessage Ms{{exportableSymbol .Input.Message}})(outMessage *Ms{{exportableSymbol .Output.Message}}{{if .Fault.Message}}, faultMessage *{{.Fault.Message}}{{end}})
	{{end}}
}
`

var (
	tmplService = `
package {{.PackageName}}

import (
	ws "github.com/fedorpatlin/wsdl-go/webservice"
	"encoding/xml"
)

type Entry struct {
	Key string {{TagDelimiter}}xml:"key"{{TagDelimiter}}
	Value string {{TagDelimiter}}xml:"value"{{TagDelimiter}}
}
	
type {{.ServiceName}} struct {
	Url string
}

func New{{.ServiceName}}() *{{.ServiceName}}{
	s := {{.ServiceName}}{}
	s.Url = "{{.ServiceUrl}}"
	
	return &s
}
{{with $s := .}}
//complexTypes
{{range .Types}}
type {{.XMLName.Local}} struct {
	XMLNamespace string {{TagDelimiter}}xml:"xmlns,attr"{{TagDelimiter}}
	{{range .Fields}}{{.Name}} {{if StringHasValue .Type}}{{.Type}}{{end}} {{if StringHasValue .XMLName.Local}}{{TagDelimiter}}xml:"{{.XMLName.Space}} {{.XMLName.Local}}"{{TagDelimiter}}{{end}}
	{{end}}
}
{{end}}

//messages
{{range .Messages}}
type {{.Name}} struct {
	XMLName xml.Name        {{TagDelimiter}}xml:"{{if StringHasValue .XMLName.Space}} {{.XMLName.Space}} {{end}} {{.XMLName.Local}}"{{TagDelimiter}}
	{{if .Input}}Action  string          {{TagDelimiter}}xml:"-"{{TagDelimiter}}{{end}}
	{{range .Params}}
	{{.ParamName}} {{.ParamType}} {{TagDelimiter}}xml:"{{.XMLParamName}}"{{TagDelimiter}}
	{{end}}
}

{{if .Input}}func (si {{.Name}}) GetAction() string {
	return si.Action
}{{end}}
{{end}}
{{range .Methods}}
func (s *{{$s.ServiceName}}) {{.Name}}({{if .HasParams}}p {{.InputType}}{{end}}) (r *{{.OutputType}}, err error) {
	{{if .HasParams}}si := {{.MessageIn}}{}
	si.Action = "{{.Action}}"
	si.{{.ParamInName}} = p{{end}}

	sr, err := webservice.CallService({{if .HasParams}}si{{else}}nil{{end}}, s.Url)
	if err != nil {
		return nil, err
	}

	var so {{.MessageOut}}
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.{{.ParamOutName}}, nil
}
{{end}}{{end}}
`
)
