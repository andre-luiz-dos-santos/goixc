package goixc

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"
)

func TestClient_DesbloqueioConfianca(t *testing.T) {
	url := os.Getenv("IXC_URL")
	token := os.Getenv("IXC_TOKEN")
	if url == "" || token == "" {
		t.Skipf("IXC_URL or IXC_TOKEN is not set")
	}
	ixc, err := NewClient(url, token)
	if err != nil {
		t.Fatalf("Failed to create a new IXC client: %v", err)
	}
	t.Run("contrato=-1", func(t *testing.T) {
		id := int64(-1)
		err := ixc.DesbloqueioConfianca(context.Background(), id)
		if !errors.Is(err, ErrNotFound) {
			t.Fatalf("DesbloqueioConfianca(%v) = %v; want ErrNotFound", id, err)
		}
	})
	t.Run("contrato=IXC_CONTRATO_ATIVO_ID", func(t *testing.T) {
		idStr := os.Getenv("IXC_CONTRATO_ATIVO_ID")
		if idStr == "" {
			t.Skipf("IXC_CONTRATO_ATIVO_ID is not set")
		}
		id, _ := strconv.Atoi(idStr)
		err := ixc.DesbloqueioConfianca(context.Background(), int64(id))
		if err != nil {
			t.Fatalf("Failed to DesbloqueioConfianca(%v); %v", id, err)
		}
	})
}
