package main

import (
	"bufio"
	"encoding/xml"
	"flag"
	"fmt"
	"github.com/fedorpatlin/wsdl-go/wsdl"
	"github.com/fedorpatlin/wsdl-go/xsd"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"text/template"
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
	XMLName xml.Name
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

type StructType struct {
	XMLName xml.Name
	Fields []Field
}

type Field struct {
	XMLName xml.Name
	Name    string
	Type	string
	Ns	string
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
	//data.Messages = make([]Message, 0)
	data.Types = createStructsType(d.Types.Schemas)
//	data.Messages = createStructsMessages(d)
	
//	for _, op := range d.PortType.Operations {

//		m := Method{}
//		m.Name = exportableSymbol(op.Name)
//		// TODO: get correct action in binding area
//		m.Action = op.Name
//		// find input parameter type
//		e := findElement(s, op.Input.Message, d.Messages)
//		//fmt.Printf("DEBUG find elements message: %s\n", d.PortType.Operations[i].Input.Message)
//		var c *xsd.ComplexType

//		if e.ComplexTypes == nil {
//			c = findComplexType(s, e.Name)
//		} else {
//			c = e.ComplexTypes
//		}

//		message := Message{}
//		if c.Name != "" {
//			message.Name = exportableSymbol(c.Name)
//			message.XMLName.Local = c.Name
//		} else {
//			message.Name = exportableSymbol(e.Name)
//			message.XMLName.Local = e.Name
//		}

//		if len(c.Sequence) > 0 {

//			for _, v := range c.Sequence {
//				messageParam := MessageParamIn{}
//				messageParam.ParamName = exportableSymbol(v.Name)
//				messageParam.XMLParamName = v.Name
//				_, t := extractNamespace(v.Type)
//				messageParam.ParamType = t
//				messageParam.Input = true

//				message.Params = append(message.Params, messageParam)
//			}
//			data.Messages = append(data.Messages, message)
//			_, st := extractNamespace(c.Sequence[0].Type)
//			m.InputType = st
//			m.MessageIn = message.Name
//			m.ParamInName = message.ParamName
//			fmt.Printf("DEBUG: message param %s\n",message.Name)
//			m.HasParams = true
//		} else {
//			m.HasParams = false
//		}

//		// find output parameter type
//		e = findElement(s, op.Output.Message, d.Messages)

//		if e.ComplexTypes == nil {
//			c = findComplexType(s, e.Name)
//		} else {
//			c = e.ComplexTypes
//		}

//		message = Message{}

//		//message.Name = exportableSymbol(c.Name)
//		//message.XMLName = c.Name

//		if c.Name != "" {
//			message.Name = exportableSymbol(c.Name)
//			message.XMLName.Local = c.Name
//		} else {
//			message.Name = exportableSymbol(e.Name)
//			message.XMLName.Local = e.Name
//		}

//		for _, v := range c.Sequence {
//			messageParam := MessageParamIn{}
//			messageParam.ParamName = exportableSymbol(v.Name)
//			messageParam.XMLParamName = v.Name
//			_, t := extractNamespace(v.Type)
//			messageParam.ParamType = exportableSymbol(t)
//			messageParam.Input = false

//			message.Params = append(message.Params, messageParam)
//		}

//		/*message.ParamName = exportableSymbol(c.Sequence[0].Name)
//		message.XMLParamName = c.Sequence[0].Name
//		message.ParamType = exportableSymbol(c.Sequence[0].Type[si+1:])
//		message.Input = false*/

//		data.Messages = append(data.Messages, message)
//		_, st := extractNamespace(c.Sequence[0].Type)
//		m.OutputType = exportableSymbol(st)
//		m.MessageOut = message.Name
//		m.ParamOutName = exportableSymbol(c.Sequence[0].Name)
//		data.Methods = append(data.Methods, m)
//	}

	// executa o template para geração do arquivo com
	// o serviço que será consumido
	err = tmpl.Execute(file, data)
	if err != nil {
		exit(err)
	}
}

func createStructsType(schemas []xsd.Schema) (st []StructType) {

	for _, s := range schemas {
		for _, simt := range s.SimpleTypes {
			st = append(st, structSimpleType(simt, s.TargetNamespace))
			fmt.Printf("DEBUG: simpletype added: %s\n", st)
		}
		for _, ct := range s.ComplexTypes {
			st = append(st, structComplexType(ct, s.TargetNamespace))
		}
	}

	return
}

func structSimpleType(simt xsd.SimpleType, ns string) (st StructType) {
	st.XMLName.Local = simt.Name
	st.XMLName.Space = ns
	st.Fields = append(st.Fields, Field{Name: simt.Name, Type: simt.Restriction.Base})
	return
}

func structComplexType(ct xsd.ComplexType, ns string) (st StructType) {
	st.XMLName.Local = exportableSymbol(ct.Name)
	st.XMLName.Space = ns
	fmt.Println("DEBUG: namespace "+ns)
	for _, seq := range ct.Sequence {
		var f Field
		f.Name = exportableSymbol(seq.Name)
		f.XMLName = decodeElementType(seq)
		f.Type = f.XMLName.Local
		st.Fields = append(st.Fields, f)
	}

	return
}

func structMessage(m wsdl.Message)(st StructType){
	st.XMLName.Local = m.Name
	for _,p := range m.Part {
		f := Field{}
		f.Name=p.Element
	}
	return
}



func createStructMessage(d wsdl.Definitions){
	
}

func findElement(s *xsd.Schema, t string, msg []wsdl.Message) *xsd.Element {
	_, t = extractNamespace(t)

	t = strings.Replace(t, "SoapIn", "", -1)
	t = strings.Replace(t, "SoapOut", "Response", -1)

	for i := 0; i < len(s.Elements); i++ {
		if s.Elements[i].Type == t || s.Elements[i].Name == t {
			return &s.Elements[i]
		}
	}
	for _, m := range msg {
		if m.Name == t {
			var ps []string
			for _,p := range m.Part{
			ps = strings.Split(p.Element, ":")}

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
	if strings.Contains(s, ":") {
		panic(s)
	}
	//fmt.Printf("DEBUG: exportableSymbol %s\n", s)
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

func decodeElementType(e xsd.Element) (ty xml.Name) {
	t := e.Type
	ns, t := extractNamespace(t)
	if ns == "xs" || ns == "xsd" {
		switch t {
		case "string":
			ty.Local="string"
			ty.Space=ns
			return
		case "boolean":
			ty.Local="string"
			ty.Space=ns
			return
		case "decimal", "double":
			ty.Local="string"
			ty.Space=ns
			return
		case "int":
			ty.Local="string"
			ty.Space=ns
			return
		case "unsignedInt":
			ty.Local="string"
			ty.Space=ns
			return
		default:
			ty.Local=""
			ty.Space=""
			return
		}
	} else if ns==""{
				if e.MaxOccurs == "unbounded" {
			ty.Local = "[]" + t
		} else if v, _ := strconv.ParseInt(e.MaxOccurs, 10, 64); v > 1 {
			ty.Local = "[" + e.MaxOccurs + "]" + t
		}
		return ty
	} else {

		t := exportableSymbol(t)
		if e.MaxOccurs == "unbounded" {
			ty.Local = "[]" + t
		} else if v, _ := strconv.ParseInt(e.MaxOccurs, 10, 64); v > 1 {
			ty.Local = "[" + e.MaxOccurs + "]" + t		}
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
