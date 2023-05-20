package client

import (
  "net/http"
  "crypto/tls"
)

type Client struct {
  HTTPClient      http.Client
}

func NewClient(cfg *tls.Config) *Client {
  return &Client{
    HTTPClient: http.Client{
      Transport: &http.Transport{
        TLSClientConfig: cfg,
      },
    },
  }
}
