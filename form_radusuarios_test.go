package goixc

import (
	"context"
	"os"
	"testing"
)

func TestClient_GetRadusuarioByLogin(t *testing.T) {
	url := os.Getenv("IXC_URL")
	token := os.Getenv("IXC_TOKEN")
	if url == "" || token == "" {
		t.Skipf("IXC_URL or IXC_TOKEN is not set")
	}
	ixc, err := NewClient(url, token)
	if err != nil {
		t.Fatalf("Failed to create a new IXC client: %v", err)
	}
	for _, login := range []string{"", "random-string"} {
		t.Run("login="+login, func(t *testing.T) {
			resp, err := ixc.GetRadusuarioByLogin(context.Background(), login)
			if err == nil {
				t.Errorf("GetRadusuarioByLogin(%v) = %#v; want nil", login, resp)
			}
		})
	}
	t.Run("login=IXC_LOGIN", func(t *testing.T) {
		login := os.Getenv("IXC_LOGIN")
		if login == "" {
			t.Skipf("IXC_LOGIN is not set")
		}
		resp, err := ixc.GetRadusuarioByLogin(context.Background(), login)
		if err != nil {
			t.Errorf("GetRadusuarioByLogin(%v): %v", login, err)
		}
		t.Logf("%#v", resp)
	})
}
