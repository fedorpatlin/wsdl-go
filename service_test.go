package main

import "testing"
import "fmt"
import "encoding/xml"

type Ws struct {
	SAPCCMS
}

func (s *Ws) MtGetTidByName(inms MsMtGetTidByNameRequest) (outms MsMtGetTidByNameResponse) {
	return
}

func Test_unmarshal(t *testing.T) {
	var s Ws
	var xmlstr []byte
	var r MsMtGetTidByNameRequest
	var items = make([]ALMTNAMEL, 1)
	items[0].CompleteName = "asdfasdfasdf"
	r.SoapRequest.Item = items
	xmlstr, err := xml.Marshal(&r)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(xmlstr))
	s.MtGetTidByName(r)
}
