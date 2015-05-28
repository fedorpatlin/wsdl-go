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

type StructType struct {
	XMLName xml.Name
	Fields  []Field
}

type Field struct {
	XMLName xml.Name
	Name    string
	Type    string
	Ns      string
}

type TemplateData struct {
	PackageName string
	ServiceName string
	ServiceUrl  string
	Wsdl        *wsdl.Definitions
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
		"StringHasValue":    StringHasValue,
		"TagDelimiter":      TagDelimiter,
		"decodeElementType": decodeElementType,
		"decodeType":        decodeType,
		"getName":           getName,
		"getNsPrefix":       getNsPrefix,
		"exportableSymbol":  exportableSymbol,
	}
	data.PackageName = *packageName
	data.ServiceName = d.Service.Name
	data.ServiceUrl = d.Service.Port.Address.Location
	data.Wsdl = d

	// create the template
	//	tmpl, err := template.New("").Funcs(funcMap).Parse(tmplService)
	//	if err != nil {
	//		exit(err)
	//	}
	tmpl1, err := template.New("").Funcs(funcMap).Parse(tmpl_intro + tmpl_complex_type + tmpl_elements + tmpl_messages + tmpl_operations)
	err = tmpl1.Execute(file, data)
	if err == nil {
		exit(err)
	}

}

func exportableSymbol(s string) string {
	if strings.Contains(s, ":") {
		_, s = extractNamespace(s)
	}
	//fmt.Printf("DEBUG: exportableSymbol %s\n", s)
	return strings.ToUpper(s[0:1]) + s[1:]
}

func getNsPrefix(elt string) (ns string) {
	ns, _ = extractNamespace(elt)
	return
}

func getName(elt string) (n string) {
	_, n = extractNamespace(elt)
	return
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
	ty = decodeType(t)
	if e.MaxOccurs == "unbounded" {
		ty.Local = "[]" + ty.Local
	} else if v, _ := strconv.ParseInt(e.MaxOccurs, 10, 64); v > 1 {
		ty.Local = "[" + e.MaxOccurs + "]" + ty.Local
	}
	return ty
	return
}

func decodeType(t string) (ty xml.Name) {
	ns, t := extractNamespace(t)
	if ns == "xs" || ns == "xsd" {
		switch t {
		case "string":
			ty.Local = "string"
			ty.Space = ns
			return
		case "boolean":
			ty.Local = "string"
			ty.Space = ns
			return
		case "decimal", "double":
			ty.Local = "string"
			ty.Space = ns
			return
		case "int":
			ty.Local = "string"
			ty.Space = ns
			return
		case "unsignedInt":
			ty.Local = "string"
			ty.Space = ns
			return
		default:
			ty.Local = "-----"
			ty.Space = "-----"
			return
		}
	} else {
		ty.Local = exportableSymbol(t)
		return
	}
	panic("unknown type")
	return
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
