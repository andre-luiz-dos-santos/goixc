package goixc

import (
	"context"
	"os"
	"strconv"
	"testing"
)

func TestClient_LiberarTemporariamente(t *testing.T) {
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
		err = ixc.LiberarTemporariamente(context.Background(), id)
		if err != nil {
			t.Errorf("Failed to LiberarTemporariamente(%v): %v", id, err)
		}
	})
	t.Run("contrato=IXC_CONTRATO_ID", func(t *testing.T) {
		idStr := os.Getenv("IXC_CONTRATO_ID")
		if idStr == "" {
			t.Skipf("IXC_CONTRATO_ID is not set")
		}
		id, _ := strconv.Atoi(idStr)
		err = ixc.LiberarTemporariamente(context.Background(), int64(id))
		if err != nil {
			t.Errorf("Failed to LiberarTemporariamente(%v): %v", id, err)
		}
	})
}
