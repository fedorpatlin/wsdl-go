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

	"code.google.com/p/wsdl-go/wsdl"
	"code.google.com/p/wsdl-go/xsd"
)

var wsdlFile = flag.String("w", "", "WSDL file with full path")
var xsdFile = flag.String("x", "", "XSD file with full path")
var packageName = flag.String("p", "", "Package name")
var outFile = flag.String("o", "", "Output file")

type Method struct {
	Name         string
	Action       string
	HasParams    bool
	InputType    string
	OutputType   string
	MessageIn    string
	MessageOut   string
	ParamInName  string
	ParamOutName string
}

type Message struct {
	Name    string
	XMLName string
	Action  string
	Params  []MessageParamIn

	ParamName    string
	XMLParamName string
	ParamType    string
	Input        bool
}

type MessageParamIn struct {
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

type TemplateData struct {
	PackageName string
	ServiceName string
	ServiceUrl  string
	Messages    []Message
	Methods     []Method
	Types       []StructType
}

var data TemplateData

// wsdl -w="C:\Temp\wsdl\CartaoEndpointService.wsdl" -x="C:\Temp\wsdl\CartaoEndpointService_schema1.xsd" -p="main" -o="C:\Temp\service.go"
// wsdl -w="C:\Temp\wsdl\authendpointservice.wsdl" -x="C:\Temp\wsdl\AuthEndpointService_schema1.xsd" -p="login" -o="C:\Temp\auth_service.go"
func main() {
	flag.Parse()

	if *wsdlFile == "" || *packageName == "" || *outFile == "" {
		flag.Usage()
		return
	}

	var d wsdl.Definitions
	unmarshal(*wsdlFile, &d)

	var s xsd.Schema

	// se foi informado qual o arquivo de schema
	// o schema pode ser passado em um arquivo separado ou
	// poder estar dentro do proprio wsdl
	if *xsdFile != "" {
		unmarshal(*xsdFile, &s)
	} else {
		// TODO: na verdade podemos ter mais de um schema
		s = d.Types.Schemas[0]
	}

	buf, f := createOut(*outFile)
	defer buf.Flush()
	defer f.Close()

	// create de service file
	create(&d, &s, buf, f)
}

func unmarshal(n string, i interface{}) {
	f, err := os.Open(n)
	if err != nil {
		exit(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		exit(err)
	}

	err = xml.Unmarshal(b, i)
	if err != nil {
		exit(err)
	}
}

func createOut(n string) (*bufio.Writer, *os.File) {
	// remove o arquivo de saida
	err := os.Remove(n)
	// verificar se houve, se houve erro e o erro não for do tipo "não existe"
	// ignora o erro se retorna o erro em questão
	if err != nil && !os.IsNotExist(err) {
		exit(err)
	}

	// cria o arquivo de saída
	f, err := os.Create(n)
	if err != nil {
		exit(err)
	}

	return bufio.NewWriter(f), f
}

// exit sai da aplicação exibindo o erro se existir
func exit(err error) {
	if err != nil {
		fmt.Println(err.Error())
	}
	os.Exit(1)
}

// create cria o arquivo com o serviço a ser consumido
func create(d *wsdl.Definitions, s *xsd.Schema, b *bufio.Writer, file *os.File) {
	funcMap := template.FuncMap{
		"StringHasValue": StringHasValue,
		"TagDelimiter":   TagDelimiter,
	}

	// create the template
	tmpl, err := template.New("").Funcs(funcMap).Parse(tmplService)
	if err != nil {
		exit(err)
	}

	data.PackageName = *packageName
	data.ServiceName = d.Service.Name
	data.ServiceUrl = d.Service.Port.Address.Location
	data.Methods = make([]Method, 0)
	data.Messages = make([]Message, 0)

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
			t = StructType{
				Name:   exportableSymbol(s.ComplexTypes[i].Name),
				Fields: make([]Field, 0),
			}

			if s.ComplexTypes[i].Content == nil {
				for ii := 0; ii < len(s.ComplexTypes[i].Sequence); ii++ {
					fi := Field{
						Name:    exportableSymbol(s.ComplexTypes[i].Sequence[ii].Name),
						Type:    decodeType(s.ComplexTypes[i].Sequence[ii]),
						XMLName: s.ComplexTypes[i].Sequence[ii].Name,
					}
					t.Fields = append(t.Fields, fi)
				}
			} else {
				_, b := extractNamespace(s.ComplexTypes[i].Content.Extension.Base)
				t.Fields = append(t.Fields, Field{Name: exportableSymbol(b)})
				for ii := 0; ii < len(s.ComplexTypes[i].Content.Extension.Sequence); ii++ {
					fi := Field{
						Name:    exportableSymbol(s.ComplexTypes[i].Content.Extension.Sequence[ii].Name),
						Type:    decodeType(s.ComplexTypes[i].Content.Extension.Sequence[ii]),
						XMLName: s.ComplexTypes[i].Content.Extension.Sequence[ii].Name,
					}
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
		e := findElement(s, d.PortType.Operations[i].Input.Message, d.Messages)
		fmt.Printf("find elements message: %s\n", d.PortType.Operations[i].Input.Message)
		var c *xsd.ComplexType

		if e.ComplexTypes == nil {
			c = findComplexType(s, e.Name)
		} else {
			c = e.ComplexTypes
		}

		message := Message{}
		if c.Name != "" {
			message.Name = exportableSymbol(c.Name)
			message.XMLName = c.Name
		} else {
			message.Name = exportableSymbol(e.Name)
			message.XMLName = e.Name
		}

		if len(c.Sequence) > 0 {
			
			for _, v := range c.Sequence {
				messageParam := MessageParamIn{}
				messageParam.ParamName = exportableSymbol(v.Name)
				messageParam.XMLParamName = v.Name
				_, t := extractNamespace(v.Type)
				messageParam.ParamType = exportableSymbol(t)
				messageParam.Input = true

				message.Params = append(message.Params, messageParam)
			}
			data.Messages = append(data.Messages, message)
			_, st := extractNamespace(c.Sequence[0].Type)
			m.InputType = exportableSymbol(st)
			m.MessageIn = message.Name
			m.ParamInName = message.ParamName
			m.HasParams = true
		} else {
			m.HasParams = false
		}

		// find output parameter type
		e = findElement(s, d.PortType.Operations[i].Output.Message, d.Messages)

		if e.ComplexTypes == nil {
			c = findComplexType(s, e.Name)
		} else {
			c = e.ComplexTypes
		}


		message = Message{}

		//message.Name = exportableSymbol(c.Name)
		//message.XMLName = c.Name

		if c.Name != "" {
			message.Name = exportableSymbol(c.Name)
			message.XMLName = c.Name
		} else {
			message.Name = exportableSymbol(e.Name)
			message.XMLName = e.Name
		}

		for _, v := range c.Sequence {
			messageParam := MessageParamIn{}
			messageParam.ParamName = exportableSymbol(v.Name)
			messageParam.XMLParamName = v.Name
			_, t := extractNamespace(v.Type)
			messageParam.ParamType = exportableSymbol(t)
			messageParam.Input = false

			message.Params = append(message.Params, messageParam)
		}

		/*message.ParamName = exportableSymbol(c.Sequence[0].Name)
		message.XMLParamName = c.Sequence[0].Name
		message.ParamType = exportableSymbol(c.Sequence[0].Type[si+1:])
		message.Input = false*/

		data.Messages = append(data.Messages, message)
		_, st := extractNamespace(c.Sequence[0].Type)
		m.OutputType = exportableSymbol(st)
		m.MessageOut = message.Name
		m.ParamOutName = exportableSymbol(c.Sequence[0].Name)

		data.Methods = append(data.Methods, m)
	}

	// executa o template para geração do arquivo com
	// o serviço que será consumido
	err = tmpl.Execute(file, data)
	if err != nil {
		exit(err)
	}
}

func findElement(s *xsd.Schema, t string, msg []wsdl.Message) *xsd.Element {
	_ ,t = extractNamespace(t)

	t = strings.Replace(t, "SoapIn", "", -1)
	t = strings.Replace(t, "SoapOut", "Response", -1)

	for i := 0; i < len(s.Elements); i++ {
		if s.Elements[i].Type == t || s.Elements[i].Name == t {
			return &s.Elements[i]
		}
	}
	for _, m := range msg {
		if m.Name == t {
			ps := strings.Split(m.Part.Element, ":")

			return findElement(s, ps[len(ps)-1], msg)
		}
	}
	return nil
}

func findComplexType(s *xsd.Schema, n string) *xsd.ComplexType {
	_, n = extractNamespace(n)
	n = strings.Replace(n, "SoapIn", "", -1)
	n = strings.Replace(n, "SoapOut", "Response", -1)
	for i := 0; i < len(s.ComplexTypes); i++ {
		if s.ComplexTypes[i].Name == n {
			return &s.ComplexTypes[i]
		}
	}

	return nil
}

func exportableSymbol(s string) string {
	if strings.Contains(s,":"){
		panic(s)
	}
	fmt.Printf("DEBUG: exportableSymbol %s\n", s)
	return strings.ToUpper(s[0:1]) + s[1:]
}

func extractNamespace(elt string) (ns, e string) {
	if !strings.Contains(elt, ":") {
		ns = ""
		e = elt
		return
	}
	s := strings.Split(elt, ":")
	ns = s[0]
	e = s[len(s)-1]
	return
}

func decodeType(e xsd.Element) string {
	t := e.Type
	ns, t := extractNamespace(t)
	// TODO(dops): tratar
	if ns == "" && e.Name == "entry" {
		return "[]Entry"
	}
	if ns == "xs" || ns == "xsd" {
		switch t {
		case "string":
			return "string"
		case "boolean":
			return "bool"
		case "decimal", "double":
			return "float64"
		default:
			return "nil"
		}
	} else if ns == "s" {
		switch t {
		case "string":
			return "string"
		case "boolean":
			return "bool"
		case "decimal", "double":
			return "float64"
		default:
			return "nil"
		}
		//	} else if t[0:3] == "tns" {
	} else {

		ty := exportableSymbol(t)
		if e.MaxOccurs == "unbounded" {
			ty = "[]" + ty
		}
		fmt.Printf("%s:%s\n", ns, ty)
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
