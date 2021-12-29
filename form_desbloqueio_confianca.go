package goixc

import (
	"context"
	"fmt"
)

// http://wiki.ixcsoft.com.br/index.php/Desbloqueio_de_Confian%C3%A7a_-_API
func (c *Client) DesbloqueioConfianca(ctx context.Context, idContrato int64) error {
	req := map[string]int64{"id": idContrato}
	resp := &DesbloqueioConfiancaResponse{}
	err := c.Get("desbloqueio_confianca", req).RunJSON(ctx, resp)
	if err != nil {
		return err
	}
	if resp.Tipo == "sucesso" {
		return nil
	}
	switch resp.Codigo {
	case "106":
		return &ErrContrato{idContrato, "não existe", ErrNotFound}
	case "107":
		return nil // Já está desbloqueado
	case "109":
		return &ErrContrato{idContrato, "não pode ser desbloqueado", ErrForbidden}
	}
	return &ErrIXCForm{"desbloqueio_confianca", req, fmt.Sprintf("%+v", resp)}
}
