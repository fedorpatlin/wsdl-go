package main

import "testing"
import "fmt"
import "encoding/xml"
import "github.com/fedorpatlin/wsdl-go/soap"

type Ws struct {
	SAPCCMS
}

func (s *Ws) MtGetTidByName(inms MsgMtGetTidByNameRequest) (outms MsgMtGetTidByNameResponse) {
	return
}

func Test_request(t *testing.T) {
	var e = soap.NewEnvelope()

	var s Ws
	//	var xmlstr []byte
	var r = NewMsgMtGetTidByNameRequest()
	var items = make([]ALMTNAMEL, 1)
	items[0].CompleteName = "ERP\\Dialog_ascserp1_ERP_00\\MEDICAL_WORKPLACE_RESPONSE_TIME\\ResponseTime"
	r.SoapRequest.Item = items
	e.Body.Content = r
	soapmesage, err := xml.MarshalIndent(&e, "", "	")
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(soapmesage))
	s.MtGetTidByName(*r)
}

func Test_response(t *testing.T) {
	var response = `<?xml version="1.0" encoding="UTF-8"?><SOAP-ENV:Envelope xmlns:SOAP-ENV="http://schemas.xmlsoap.org/soap/envelope/" xmlns:SOAP-ENC="http://schemas.xmlsoap.org/soap/encoding/" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:SAPControl="urn:SAPControl" xmlns:SAPCCMS="urn:SAPCCMS" xmlns:SAPHostControl="urn:SAPHostControl" xmlns:SAPOscol="urn:SAPOscol" xmlns:SAPDSR="urn:SAPDSR"><SOAP-ENV:Body><SAPCCMS:MtGetTidByNameResponse><TidTable><item><Tid><Mtsysid>ERP     </Mtsysid><Mtmcname>Dialog_ascserp1_ERP_00                  </Mtmcname><Mtnumrange>004</Mtnumrange><Mtuid>1973600000</Mtuid><Mtclass>100</Mtclass><Mtindex>0000003517</Mtindex><Extindex>0000000696</Extindex></Tid><CompleteName>ERP\Dialog_ascserp1_ERP_00\MEDICAL_WORKPLACE_RESPONSE_TIME\ResponseTime                                                                                                                                                                                         </CompleteName><Rc>0</Rc></item></TidTable></SAPCCMS:MtGetTidByNameResponse></SOAP-ENV:Body></SOAP-ENV:Envelope>`
	var responseSoap soap.Envelope
	responseStruct := MsgMtGetTidByNameResponse{}
	responseSoap.Body.Content = &responseStruct
	xml.Unmarshal([]byte(response), &responseSoap)
	s,_ := xml.MarshalIndent(responseSoap,"", "  ")
	v := responseSoap.Body.Content.(*MsgMtGetTidByNameResponse)
	fmt.Println(v.MtGetTidByNameResponse.TidTable.Item[0].Tid)
	fmt.Println(string(s))
}
