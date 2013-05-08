package main

var (
	tmplService = `
package {{.PackageName}}

import (
	"encoding/xml"
	"code.google.com/p/wsdl-go/wsdl"
)
	
type {{.ServiceName}} struct {
	Url string
}

func New{{.ServiceName}}() *{{.ServiceName}}{
	s := {{.ServiceName}}{}
	s.Url = "{{.ServiceUrl}}"
	
	return &s
}
{{with $s := .}}
{{range .Types}}
type {{.Name}} struct {
	XMLNamespace string {{TagDelimiter}}xml:"xmlns,attr"{{TagDelimiter}}
	{{range .Fields}}{{.Name}} {{if StringHasValue .Type}}{{.Type}}{{end}} {{if StringHasValue .XMLName}}{{TagDelimiter}}xml:"{{.XMLName}}"{{TagDelimiter}}{{end}}
	{{end}}
}
{{end}}
{{range .Messages}}
type {{.Name}} struct {
	XMLName xml.Name        {{TagDelimiter}}xml:"http://webservice.auth.app.bsbr.altec.com/ {{.XMLName}}"{{TagDelimiter}}
	{{if .Input}}Action  string          {{TagDelimiter}}xml:"-"{{TagDelimiter}}{{end}}
	{{.ParamName}} {{.ParamType}} {{TagDelimiter}}xml:"{{.XMLParamName}}"{{TagDelimiter}}
}

{{if .Input}}func (si {{.Name}}) GetAction() string {
	return si.Action
}{{end}}
{{end}}{{range .Methods}}
func (s *{{$s.ServiceName}}) {{.Name}}(p {{.InputType}}) (r *{{.OutputType}}, err error) {
	si := {{.MessageIn}}{}
	si.Action = "{{.Action}}"
	si.{{.ParamInName}} = p

	sr, err := webservice.CallService(si, s.Url)
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
