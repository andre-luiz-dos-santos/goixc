package goixc

import (
	"context"
	"strings"
)

func (c *Client) GetClientes(ctx context.Context, query interface{}) (*ClientesResponse, error) {
	resp := &ClientesResponse{}
	err := c.List("cliente", query).RunJSON(ctx, resp)
	return resp, err
}

func (c *Client) GetClienteByCPFCNPJ(ctx context.Context, cpf string) (*Cliente, error) {
	cpf = strings.TrimSpace(cpf)
	for _, s := range cpfVariations(cpf) {
		if s == "" {
			continue
		}
		resp, err := c.GetClientes(ctx, map[string]string{
			"qtype": "cliente.cnpj_cpf",
			"oper":  "=",
			"query": s,
		})
		if err != nil {
			return nil, err
		}
		if resp.Total >= 1 || len(resp.Registros) >= 1 {
			return resp.Registros[0], nil
		}
	}
	return nil, &NotFoundError{"cliente", cpf}
}

func cpfVariations(cpfCNPJ string) []string {
	ss := make([]string, 0, 3)
	s := notNumbersRE.ReplaceAllString(cpfCNPJ, "")
	ss = append(ss, s)
	if len(s) == 11 {
		ss = append(ss, s[0:3]+"."+s[3:6]+"."+s[6:9]+"-"+s[9:11])
	}
	if len(s) == 14 {
		ss = append(ss, s[0:2]+"."+s[2:5]+"."+s[5:8]+"/"+s[8:12]+"-"+s[12:14])
	}
	return ss
}
