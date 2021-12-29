package goixc

import (
	"context"
	"errors"
	"os"
	"strconv"
	"testing"
)

func TestClient_GetBoleto(t *testing.T) {
	url := os.Getenv("IXC_URL")
	token := os.Getenv("IXC_TOKEN")
	if url == "" || token == "" {
		t.Skipf("IXC_URL or IXC_TOKEN is not set")
	}
	ixc, err := NewClient(url, token)
	if err != nil {
		t.Fatalf("Failed to create a new IXC client: %v", err)
	}
	t.Run("boleto=-1", func(t *testing.T) {
		id := int64(-1)
		_, err = ixc.GetBoleto(context.Background(), id)
		if !errors.Is(err, ErrInvalid) {
			t.Errorf("GetBoleto(%v) = %v; want ErrInvalid", id, err)
		}
	})
	t.Run("boleto=IXC_BOLETO_ID", func(t *testing.T) {
		idStr := os.Getenv("IXC_BOLETO_ID")
		if idStr == "" {
			t.Skipf("IXC_BOLETO_ID is not set")
		}
		id, _ := strconv.Atoi(idStr)
		pdf, err := ixc.GetBoleto(context.Background(), int64(id))
		if err != nil {
			t.Errorf("Failed to GetBoleto(%v): %v", id, err)
		}
		t.Logf("Returned %v bytes", len(pdf))
	})
}
