package webservice

import (
	"encoding/xml"
	"testing"
)

type ObterUsuariosSoapIn struct {
	XMLName xml.Name        `xml:"http://webservice.auth.app.bsbr.altec.com/ obterUsuarios"`
	Action  string          `xml:"-"`
	Param   UsuarioParamDTO `xml:"param"`
}

func (si ObterUsuariosSoapIn) GetAction() string {
	return si.Action
}

type ObterUsuariosSoapOut struct {
	XMLName xml.Name          `xml:"http://webservice.auth.app.bsbr.altec.com/ obterUsuariosResponse"`
	Return  UsuarioResultList `xml:"return"`
}

func NewObterUsuariosSoapIn() ObterUsuariosSoapIn {
	si := ObterUsuariosSoapIn{}
	si.Action = ""

	return si
}

type UsuarioParamDTO struct {
	XMLNamespace string `xml:"xmlns,attr"`
	Agencia      string `xml:"agencia"`
	Conta        string `xml:"conta"`
	Banco        string `xml:"banco"`
	Imei         string `xml:"IMEI"`
	Linha        string `xml:"linha"`
}

type UsuarioResultList struct {
	Agencia       string    `xml:"agencia"`
	Conta         string    `xml:"conta"`
	Banco         string    `xml:"banco"`
	CodErro       string    `xml:"codErro"`
	MsgErro       string    `xml:"msgErro"`
	TipoConta     string    `xml:"tipoConta"`
	ListaUsuarios []Usuario `xml:"listaUsuarios"`
}

type Usuario struct {
	CodigoUsuario      string `xml:"codigoUsuario"`
	CodigoUsuarioConta string `xml:"codigoUsuarioConta"`
	FlagCadastro       string `xml:"flagCadastro"`
	HsEkscPbe          string `xml:"hsEkscPbe"`
	Nome               string `xml:"nome"`
	Penumper           string `xml:"penumper"`
	TipoUsuario        string `xml:"tipoUsuario"`
	SKa                string `xml:"sKa"`
	SRand              string `xml:"sRand"`
	SSingJs            string `xml:"sSignJs"`
}

// http://localhost/Santander/Webservices/AuthService/AuthEndpointService.asmx

type AuthEndpointService struct {
	Url string
}

func NewAuthEndpointService() *AuthEndpointService {
	s := &AuthEndpointService{}
	s.Url = "http://localhost/Santander/Webservices/AuthService/AuthEndpointService.asmx"

	return s
}

// função de chamada de exemplo
func (s AuthEndpointService) ObterUsuarios(param UsuarioParamDTO) (r *UsuarioResultList, err error) {
	si := NewObterUsuariosSoapIn()
	si.Param = param

	// chama o serviço apontando para determina url
	sr, err := callService(si, s.Url)
	if err != nil {
		return nil, err
	}

	// monta a estrutura de retorno
	var so ObterUsuariosSoapOut
	err = xml.Unmarshal([]byte(sr.Body.Content), &so)
	if err != nil {
		return nil, err
	}

	return &so.Return, nil
}

func TestAuthEndpointService(t *testing.T) {
	s := NewAuthEndpointService()

	p := UsuarioParamDTO{}
	p.Banco = "0033"
	p.Agencia = "1111"
	p.Conta = "000111111111"
	p.Imei = "01020304"
	p.Linha = "983157131"

	l, err := s.ObterUsuarios(p)
	if err != nil {
		t.Fatal(err)
	}

	if c := len(l.ListaUsuarios); c != 2 {
		t.Fatalf("esperado 2 usuarios recebido %d usuarios", c)
	}

}
