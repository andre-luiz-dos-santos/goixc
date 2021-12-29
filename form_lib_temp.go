package goixc

import (
	"context"
)

func (c *Client) LiberarTemporariamente(ctx context.Context, idContrato int64) error {
	req := map[string]int64{"id": idContrato}
	var resp Response
	err := c.Get("cliente_contrato_btn_lib_temp_24722", req).RunJSON(ctx, &resp)
	if err != nil {
		return err
	}
	if resp.Type == "success" {
		return nil
	}
	return &IXCFormError{"cliente_contrato_btn_lib_temp_24722", req, resp.Message}
}
