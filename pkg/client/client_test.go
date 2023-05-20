package client

import (
  "testing"
  "taemon1337/http-test-server/pkg/config"
  "taemon1337/http-test-server/pkg/server"
  "taemon1337/http-test-server/pkg/tls"
)

func TestClient(t *testing.T) {
  cfg := config.NewConfig()
  cfg.HttpAddr = "127.0.0.1:8081"
  cfg.HttpsAddr = "127.0.0.1:8444"

  ss, err := tls.NewSelfSign(cfg)
  if err != nil {
    t.Errorf("%s", err)
  }

  tlsclientconfig, err := ss.TLSClientConfig()
  if err != nil {
    t.Errorf("%s", err)
  }

  tlsserverconfig, err := ss.TLSConfig()
  if err != nil {
    t.Errorf("%s", err)
  }

  cli := NewClient(tlsclientconfig)

  srv := server.NewServer()
  srv.SelfSign = ss
  srv.TLSConfig = tlsserverconfig
  srv.LoadTLS()

  t.Run("start the server", func (t *testing.T) {
    err := srv.Run()
    if err != nil {
      t.Errorf("%s", err)
    }
  })

  t.Run("connect to test server", func (t *testing.T) {
    _, err := cli.HTTPClient.Get("http://127.0.0.1:8081/")
    if err != nil {
      t.Errorf("%s", err)
      return
    }
  })
}
