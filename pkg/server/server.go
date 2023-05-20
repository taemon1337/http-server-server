package server

import (
  "os"
  "log"
  "fmt"
  "strings"
  "net/http"
  "io/ioutil"
  gotls "crypto/tls"
  "crypto/x509"
  "github.com/gin-gonic/gin"
  "golang.org/x/sync/errgroup"
  "taemon1337/http-test-server/pkg/yamlable"
  "taemon1337/http-test-server/pkg/config"
  "taemon1337/http-test-server/pkg/templates"
  "taemon1337/http-test-server/pkg/tls"
)

type Server struct {
  Router            *gin.Engine
  Config            *config.Config
  Logger            *log.Logger
  SelfSign          *tls.SelfSign
  TLSConfig         *gotls.Config
}

func NewServer() *Server {
  r := gin.Default()
  r.HTMLRender = templates.GetTemplates()

  return &Server{
    Router:   r,
    Logger:   log.New(os.Stderr, "", 0),
    Config:   config.NewConfig(),
    SelfSign: nil,
    TLSConfig: nil,
  }
}

func (s *Server) SetRoutes() {
  s.Router.Any("/*any", func(c *gin.Context) {
    req := yamlable.NewYamlableRequest(c)

    accept := c.GetHeader("Accept")
    format := c.Query("format")

    if format != "" {
      accept = format
    }

    if req.URL == "/health" {
      c.String(http.StatusOK, "OK")
      return
    }

    if req.URL == "/certs" && s.ServingTLS() {
      c.HTML(http.StatusOK, "certs.html", gin.H{})
      return
    }

    if strings.HasPrefix(req.URL, "/download/") && s.ServingTLS() {
      switch req.URL {
        case "/download/ca.crt":
          c.String(http.StatusOK, s.SelfSign.EncodeCACertToPem().String())
        case "/download/server.crt":
          c.String(http.StatusOK, s.SelfSign.EncodeCertToPem().String())
        case "/download/client.crt":
          c.String(http.StatusOK, s.SelfSign.EncodeClientCertToPem().String())
        case "/download/client.key":
          c.String(http.StatusOK, s.SelfSign.EncodeClientCertPrivateKeyToPem().String())
      }
      return
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
}

func (s *Server) LoadTLS() error {
  if s.Config.CertFile == "" {
    s.Logger.Printf("[SERVER] Generating self-signed certs as --tls is enabled and no --cert-file specified.")
    return s.Configure()
  }

  s.TLSConfig = tls.NewTLSConfig(s.Config.CommonName, s.Config.ClientAuth, s.Config.MinTLS, s.Config.MaxTLS, s.Config.SkipVerify)

  if s.Config.CAFile != "" {
    s.Logger.Printf("[SERVER] Loading CA file from '%s'", s.Config.CAFile)
    cacert, err := ioutil.ReadFile(s.Config.CAFile)
    if err != nil {
      return fmt.Errorf("cannot load CA file from '%s' - %s", s.Config.CAFile, err)
    }

    capool := x509.NewCertPool()
    capool.AppendCertsFromPEM(cacert)
    s.TLSConfig.ClientCAs = capool
  }

  if s.Config.CertFile != "" {
    s.Logger.Printf("[SERVER] Loading server cert file from '%s' and key from '%s'", s.Config.CertFile, s.Config.KeyFile)
    var err error
    s.TLSConfig.Certificates = make([]gotls.Certificate, 1)
    s.TLSConfig.Certificates[0], err = gotls.LoadX509KeyPair(s.Config.CertFile, s.Config.KeyFile)
    if err != nil {
      return fmt.Errorf("cannot load X509 key pair from '%s' and '%s': %s", s.Config.CertFile, s.Config.KeyFile, err)
    }
  }

  return nil
}

func (s *Server) Configure() error {
  if s.SelfSign != nil {
    s.Logger.Printf("[SERVER] Server is already configured with a self-signer")
    return nil
  }

  ss, err := tls.NewSelfSign(s.Config)
  if err != nil {
    return fmt.Errorf("cannot generate certs - %s", err)
  }

  s.SelfSign = ss

  tlscfg, err := s.SelfSign.TLSConfig()
  if err != nil {
    return fmt.Errorf("cannot load server TLS config - %s", err)
  }

  s.TLSConfig = tlscfg

  return nil
}

func (s *Server) ServingTLS() bool {
  return s.Config.UseTLS && s.TLSConfig != nil
}

func (s *Server) Run() error {
  s.SetRoutes()

  g := new(errgroup.Group)

  if s.Config.UseHTTP {
    g.Go(func() error {
      s.Logger.Printf("[SERVER] Starting HTTP test server on %s...", s.Config.HttpAddr)
      return s.Router.Run(s.Config.HttpAddr)
    })
  }

  if s.Config.UseTLS {
    err := s.LoadTLS()
    if err != nil {
      return fmt.Errorf("cannot configure TLS for server - %s", err)
    }
  }

  if s.ServingTLS() {
    g.Go(func() error {
      s.Logger.Printf("[SERVER] Starting HTTPS test server on %s...", s.Config.HttpsAddr)
      srv := http.Server{
        Addr: s.Config.HttpsAddr,
        Handler: s.Router,
        TLSConfig:  s.TLSConfig,
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
