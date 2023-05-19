package server

import (
  "os"
  "log"
  "net/http"
  "github.com/gin-gonic/gin"
  "golang.org/x/sync/errgroup"
  "taemon1337/http-test-server/pkg/yamlable"
  "taemon1337/http-test-server/pkg/template"
  "taemon1337/http-test-server/pkg/config"
  "taemon1337/http-test-server/pkg/tls"
)

type Server struct {
  Router            *gin.Engine
  Config            *config.Config
  Logger            *log.Logger
  SelfSign          *tls.SelfSign
}

func NewServer() *Server {
  r := gin.Default()
  r.SetHTMLTemplate(template.IndexTemplate)

  r.Any("/*any", func(c *gin.Context) {
    req := yamlable.NewYamlableRequest(c)

    accept := c.GetHeader("Accept")
    format := c.Query("format")

    if format != "" {
      accept = format
    }

    switch accept {
      case "application/json":
      case "json":
        c.JSON(http.StatusOK, req)
      case "application/yaml":
      case "yaml":
        c.YAML(http.StatusOK, req)
      default:
        c.HTML(http.StatusOK, "index.html", gin.H{
          "request": req,
        })
    }
  })

  return &Server{
    Router:   r,
    Logger:   log.New(os.Stderr, "", 0),
    Config:   config.NewConfig(),
    SelfSign: nil,
  }
}

func (s *Server) Configure() error {
  ss, err := tls.NewSelfSign(s.Config)
  if err != nil {
    return err
  }

  s.SelfSign = ss

  return s.SelfSign.LoadTLSConfig()
}

func (s *Server) ServingTLS() bool {
  return s.SelfSign != nil && s.SelfSign.TLSConfig != nil
}

func (s *Server) Run() error {
  g := new(errgroup.Group)

  if s.Config.UseTLS {
    err := s.Configure()
    if err != nil {
      s.Logger.Printf("Could not configure TLS - %s", err)
      return err
    }
  }

  g.Go(func() error {
    s.Logger.Printf("[SERVER] Starting HTTP test server on %s...", s.Config.HttpAddr)
    return s.Router.Run(s.Config.HttpAddr)
  })

  if s.ServingTLS() {
    g.Go(func() error {
      s.Logger.Printf("[SERVER] Starting HTTPS test server on %s...", s.Config.HttpsAddr)
      srv := http.Server{
        Addr: s.Config.HttpsAddr,
        Handler: s.Router,
        TLSConfig:  s.SelfSign.TLSConfig,
      }
      return srv.ListenAndServeTLS("", "")
    })
  } else {
    s.Logger.Printf("[SERVER] Skipping TLS server as no config was present")
  }

  if err := g.Wait(); err != nil {
    s.Logger.Fatal(err)
  } else {
    s.Logger.Println("Server stopped.")
  }
  return nil
}
