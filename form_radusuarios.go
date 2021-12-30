package goixc

import (
	"context"
	"strings"
)

func (c *Client) GetRadusuarios(ctx context.Context, query interface{}) (*RadusuariosResponse, error) {
	resp := &RadusuariosResponse{}
	err := c.List("radusuarios", query).RunJSON(ctx, resp)
	return resp, err
}

func (c *Client) GetRadusuarioByLogin(ctx context.Context, login string) (*Radusuario, error) {
	login = strings.TrimSpace(login)
	if login == "" {
		return nil, &NotFoundError{"login", login}
	}
	resp, err := c.GetRadusuarios(ctx, map[string]string{
		"qtype": "radusuarios.login",
		"oper":  "=",
		"query": login,
	})
	if err != nil {
		return nil, err
	}
	if resp.Total < 1 || len(resp.Registros) < 1 {
		return nil, &NotFoundError{"login", login}
	}
	return resp.Registros[0], nil
}
