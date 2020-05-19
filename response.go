package goixc

import "github.com/shopspring/decimal"

type Response struct {
	Message string
	Type    string
	// Algumas das API's usam palavras em português.
	Mensagem string
	Tipo     string
}

type FnAreceberResponse struct {
	Pagina    int64         `json:"page,string"`
	Total     int64         `json:"total,string"`
	Registros []*FnAreceber `json:"registros"`
}

type FnAreceber struct {
	ID             int64           `json:"id,string"`
	ClienteID      int64           `json:"id_cliente,string"`
	Liberado       string          `json:"liberado"`
	Status         string          `json:"status"`
	NossoNumero    string          `json:"nn_boleto"`
	DataEmissao    DateYMD         `json:"data_emissao"`
	DataVencimento DateYMD         `json:"data_vencimento"`
	Valor          decimal.Decimal `json:"valor"`
	Obs            string          `json:"obs"`
}

type RadusuariosResponse struct {
	Pagina    int64         `json:"page,string"`
	Total     int64         `json:"total,string"`
	Registros []*Radusuario `json:"registros"`
}

type Radusuario struct {
	ID         int64  `json:"id,string"`
	Usuario    string `json:"login"`
	Senha      string `json:"senha"`
	Ativo      string `json:"ativo"`
	ClienteID  int64  `json:"id_cliente,string"`
	ContratoID int64  `json:"id_contrato,string"`
}
