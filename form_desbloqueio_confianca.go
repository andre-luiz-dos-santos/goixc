package goixc

import (
	"context"
	"regexp"
)

var (
	// Textos do site
	// http://wiki.ixcsoft.com.br/index.php/Desbloqueio_de_Confian%C3%A7a_-_API
	// Uma das atualizações do IXC mudou a resposta de objeto para lista.
	// Segundo o suporte do IXC, a resposta voltará a ser objeto com código indicando o erro.
	dcNil          = regexp.MustCompile(`\b(com sucesso|j(a|\\u00e1) encontra-se ativo)\b`)
	dcErrNotFound  = regexp.MustCompile(`\b(contrato inexistente)\b`)
	dcErrForbidden = regexp.MustCompile(`\b(n(a|\\u00e3)o pode ser desbloqueado|desabilitado para|n(a|\\u00e3)o est(a|\\u00e1) ativo)\b`)
)

func (c *Client) DesbloqueioConfianca(ctx context.Context, idContrato int64) error {
	req := map[string]int64{"id": idContrato}
	b, err := c.Get("desbloqueio_confianca", req).Run(ctx)
	if err != nil {
		return err
	}
	switch {
	case dcNil.Match(b):
		return nil
	case dcErrNotFound.Match(b):
		return &ErrContrato{idContrato, "não existe", ErrNotFound}
	case dcErrForbidden.Match(b):
		return &ErrContrato{idContrato, "não pode ser desbloqueado", ErrForbidden}
	}
	return &ErrIXCForm{"desbloqueio_confianca", req, string(b)}
}
