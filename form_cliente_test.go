package goixc

import (
	"context"
	"os"
	"testing"
)

func TestClient_GetClienteByCPFCNPJ(t *testing.T) {
	url := os.Getenv("IXC_URL")
	token := os.Getenv("IXC_TOKEN")
	if url == "" || token == "" {
		t.Skipf("IXC_URL or IXC_TOKEN is not set")
	}
	ixc, err := NewClient(url, token)
	if err != nil {
		t.Fatalf("Failed to create a new IXC client: %v", err)
	}
	for _, cpf := range []string{"", "random-string"} {
		t.Run("cpf="+cpf, func(t *testing.T) {
			resp, err := ixc.GetClienteByCPFCNPJ(context.Background(), cpf)
			if err == nil {
				t.Errorf("GetClienteByCPFCNPJ(%v) = %#v; want nil", cpf, resp)
			}
		})
	}
	t.Run("login=IXC_CPF", func(t *testing.T) {
		cpf := os.Getenv("IXC_CPF")
		if cpf == "" {
			t.Skipf("IXC_CPF is not set")
		}
		resp, err := ixc.GetClienteByCPFCNPJ(context.Background(), cpf)
		if err != nil {
			t.Errorf("GetClienteByCPFCNPJ(%v): %v", cpf, err)
		}
		t.Logf("%#v", resp)
	})
}
