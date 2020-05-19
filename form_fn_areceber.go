package goixc

import (
	"context"
	"encoding/json"
)

func (c *Client) GetFnAreceber(ctx context.Context, query interface{}) (*FnAreceberResponse, error) {
	resp := &FnAreceberResponse{}
	err := c.List("fn_areceber", query).RunJSON(ctx, resp)
	return resp, err
}

func (c *Client) GetFnAreceberAbertos(ctx context.Context, idContrato int64) (*FnAreceberResponse, error) {
	grid, err := json.Marshal([]map[string]string{
		{"TB": "fn_areceber.liberado", "OP": "=", "P": "S"},
		{"TB": "fn_areceber.status", "OP": "!=", "P": "C"},
		{"TB": "fn_areceber.status", "OP": "!=", "P": "R"},
	})
	if err != nil {
		return nil, err
	}
	return c.GetFnAreceber(ctx, map[string]interface{}{
		"qtype":      "fn_areceber.id_contrato",
		"oper":       "=",
		"query":      idContrato,
		"rp":         100,
		"sortname":   "fn_areceber.data_vencimento",
		"sortorder":  "asc",
		"grid_param": string(grid),
	})
}
