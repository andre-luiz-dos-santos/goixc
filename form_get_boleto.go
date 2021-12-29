package goixc

import (
	"bytes"
	"context"
)

func (c *Client) GetBoleto(ctx context.Context, idBoleto int64) ([]byte, error) {
	req := map[string]interface{}{
		"boletos":          idBoleto,
		"juro":             "N",
		"multa":            "N",
		"atualiza_boleto":  "N",
		"tipo_boleto":      "arquivo",
		"base64":           "N",
		"layout_impressao": "boleto_um_por_pagina",
	}
	pdf, err := c.Get("get_boleto", req).Run(ctx)
	if err != nil {
		return nil, err
	}
	if !bytes.HasPrefix(pdf, []byte("%PDF")) {
		return nil, &InvalidPDFError{pdf}
	}
	return pdf, nil
}
