package goixc

import (
	"context"
	"fmt"
)

// http://wiki.ixcsoft.com.br/index.php/Desbloqueio_de_Confian%C3%A7a_-_API
func (c *Client) DesbloqueioConfianca(ctx context.Context, idContrato int64) error {
	req := map[string]int64{"id": idContrato}
	var resp DesbloqueioConfiancaResponse
	err := c.Get("desbloqueio_confianca", req).RunJSON(ctx, &resp)
	if err != nil {
		return err
	}
	if resp.Tipo == "sucesso" {
		return nil
	}
	switch resp.Codigo {
	case "106": // {Tipo:erro Codigo:106 Mensagem:ID de contrato inexistente na base de dados.}
		return &ContratoError{idContrato, "não existe", ErrNotFound}
	case "107": // {Tipo:erro Codigo:107 Mensagem:O contrato já encontra-se ativo, não é necessário utilizar o desbloqueio de confiança.}
		return ErrAlreadyDone
	case "109": // {Tipo:erro Codigo:109 Mensagem:Não é possível utilizar o desbloqueio de confiança, o contrato está inativo.}
		return &ContratoError{idContrato, "não pode ser desbloqueado", ErrForbidden}
	case "113": // {Tipo:erro Codigo:113 Mensagem:O desbloqueio de confiança não está disponível para o contrato 1234, este recurso foi usado no dia  e não foi realizado o pagamento até o dia 01/01/2022, este recurso será habilitado novamente quando o título que vence após o dia 29/01/0000 for pago.}
		return &ContratoError{idContrato, "não pode ser desbloqueado", ErrForbidden}
	}
	return &IXCFormError{"desbloqueio_confianca", req, fmt.Sprintf("%+v", resp)}
}
