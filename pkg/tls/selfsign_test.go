package tls

import (
  "testing"
  "taemon1337/http-test-server/pkg/config"
)

func TestSelfSign(t *testing.T) {
  cfg := config.NewConfig()

  ss, err := NewSelfSign(cfg)
  if err != nil {
    t.Fatalf("%s", err)
  }

  t.Logf("%s", ss)
}
