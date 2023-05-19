package yamlable

import (
  "fmt"
  "github.com/gin-gonic/gin"
)

type YamlableRequest struct {
  URL           string                      `yaml:"url" json:"url"`
  Method        string                      `yaml:"method" json:"method"`
  Headers       map[string][]string         `yaml:"headers" json:"headers"`
  Host          string                      `yaml:"host" json:"host"`
  Proto         string                      `yaml:"proto" json:"proto"`
  RemoteAddr    string                      `yaml:"remote_addr" json:"remote_addr"`
  ContentLength string                      `yaml:"content_length" json:"content_length"`
  Websocket     bool                        `yaml:"websocket" json:"websocket"`
  Body          string                      `yaml:"body" json:"body"`
  Query         map[string][]string         `yaml:"query" json:"query"`
  TLS           *YamlableTlsConn            `yaml:"tls" json:"tls"`
  Error         string                      `yaml:"error" json:"error"`
}

func NewYamlableRequest(c *gin.Context) *YamlableRequest {
  body, err := c.GetRawData()

  return &YamlableRequest{
    URL:              fmt.Sprintf("%s", c.Request.URL),
    Method:           c.Request.Method,
    Headers:          c.Request.Header,
    Host:             c.Request.Host,
    Proto:            c.Request.Proto,
    RemoteAddr:       c.Request.RemoteAddr,
    ContentLength:    fmt.Sprintf("%d", c.Request.ContentLength),
    Websocket:        c.IsWebsocket(),
    Body:             string(body),
    Query:            c.Request.URL.Query(),
    Error:            fmt.Sprintf("%s", err),
    TLS:              NewYamlableTlsConn(c.Request),
  }
}

