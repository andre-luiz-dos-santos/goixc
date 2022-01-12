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

func (c *Client) GetFnAreceberAbertos(ctx context.Context, key string, value interface{}) ([]*FnAreceber, error) {
	grid, err := json.Marshal([]map[string]string{
		{"TB": "fn_areceber.liberado", "OP": "=", "P": "S"},
		{"TB": "fn_areceber.status", "OP": "!=", "P": "C"},
		{"TB": "fn_areceber.status", "OP": "!=", "P": "R"},
	})
	if err != nil {
		return nil, err
	}
	resp, err := c.GetFnAreceber(ctx, map[string]interface{}{
		"qtype":      "fn_areceber." + key,
		"oper":       "=",
		"query":      value,
		"rp":         100,
		"sortname":   "fn_areceber.data_vencimento",
		"sortorder":  "asc",
		"grid_param": string(grid),
	})
	return resp.Registros, err
}
