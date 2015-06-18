package main

const tmpl_intro = `{{define "Tag"}}{{TagDelimiter}}xml:"{{.}}"{{TagDelimiter}}{{end}}
// Autogenerated by github.com/fedorpatlin/wsdl-go

package {{.PackageName}}
{{define "MsgPrefix"}}Msg{{exportableSymbol .}}{{end}}
{{define "NewMsg"}}New{{template "MsgPrefix" .}}{{end}}

import "encoding/xml"
import "strings"
import "io/ioutil"
import "net/http"

{{$GlobSvcName := exportableSymbol .ServiceName}}
type {{$GlobSvcName}} struct {
	Url string
}

func New{{$GlobSvcName}}(url string)(svc {{$GlobSvcName}}){
	svc.Url=url
	return
}
`
const tmpl_complex_type = `{{range .Wsdl.Types.Schemas}}
// simple types
{{range .SimpleTypes}}
type {{exportableSymbol .Name}} struct{
	{{exportableSymbol .Name}} {{(decodeType .Restriction.Base).Local}} {{template "Tag" .Name}}
}
{{end}}
// complex types
{{range .ComplexTypes}}
type {{$t := exportableSymbol .Name}} {{exportableSymbol .Name}} struct {
{{range .Sequence}}	{{exportableSymbol .Name}} {{(decodeElementType .).Local}} {{template "Tag" .Name}}
{{end}}}
{{end}}
// anonymous types
{{range .Elements}}
type {{exportableSymbol .Name}} struct{
	{{range .ComplexTypes.Sequence}}{{exportableSymbol .Name}} {{(decodeType .Type).Local}}  {{template "Tag" .Name}}
	{{end}}
}
{{end}}

{{end}}`

const tmpl_elements = ""
const tmpl_operations = `{{$tns:=.Wsdl.TargetNamespace}}
{{$t := exportableSymbol .Wsdl.PortType.Name}}
{{range .Wsdl.PortType.Operations}}

// {{.Documentation}}
func (s *{{$GlobSvcName}}) {{.Name}}(inMs *{{template "MsgPrefix" .Input.Message}}, outMs *{{template "MsgPrefix" .Output.Message}}) error {
	var envSnd = NewSoapEnvelope()
	var envRcv = NewSoapEnvelope()
	envSnd.Body.Content = inMs
	envRcv.Body.Content = outMs
	if err := s.SendRequest(envSnd, envRcv); err != nil {
		return err
	}
	outMs = envRcv.Body.Content.(*{{template "MsgPrefix" .Output.Message}})
	return nil
}
{{end}}

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
type SoapEnvelope struct {
	XMLName xml.Name    {{template "Tag" "http://schemas.xmlsoap.org/soap/envelope/ Envelope"}}
	Soap    string      {{template "Tag" "xmlns soap,attr"}}
	Body    Body 		{{template "Tag" "Body"}}
}

func NewSoapEnvelope() *SoapEnvelope {
	se := &SoapEnvelope{}
	se.Soap = "http://schemas.xmlsoap.org/soap/envelope/"
	return se
}

type Body struct {
	XMLName xml.Name    {{template "Tag" "Body"}}
	Content interface{} {{template "Tag" ",any"}}
}

type SoapFault struct {
	XMLName     xml.Name {{template "Tag" "Fault"}}
	FaultCode   string   {{template "Tag" "faultcode"}}
	FaultString string   {{template "Tag" "faultstring"}}
	Detail      string   {{template "Tag" "detail"}}
}

func (s *{{$GlobSvcName}})SendRequest(send, receive *SoapEnvelope) error {
	sendMarshalled, err := xml.MarshalIndent(&send, "", "  ")
	if err != nil {
		return nil
	}
	var soapreader = strings.NewReader(string(sendMarshalled))
	resp, err := http.Post(s.Url, "text/xml; charset=utf-8", soapreader)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	buffer, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	xml.Unmarshal(buffer, receive)
	return nil
}
`
