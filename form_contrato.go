package goixc

import (
	"context"
	"encoding/json"
)

func (c *Client) GetContratos(ctx context.Context, query interface{}) (*ContratosResponse, error) {
	resp := &ContratosResponse{}
	err := c.List("cliente_contrato", query).RunJSON(ctx, resp)
	return resp, err
}

func (c *Client) GetContratosByClienteID(ctx context.Context, clienteID int64) ([]*Contrato, error) {
	grid, err := json.Marshal([]map[string]string{
		{"TB": "cliente_contrato.status", "OP": "=", "P": "A"},
	})
	if err != nil {
		return nil, err
	}
	resp, err := c.GetContratos(ctx, map[string]interface{}{
		"qtype":      "cliente_contrato.id_cliente",
		"oper":       "=",
		"query":      clienteID,
		"rp":         100,
		"sortname":   "cliente_contrato.id",
		"sortorder":  "asc",
		"grid_param": string(grid),
	})
	return resp.Registros, err
}
