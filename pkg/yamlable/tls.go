package yamlable

import (
  "fmt"
  "time"
  "net/http"
  "crypto/x509"
)

type YamlableTlsCert struct {
  Subject             string                `yaml:"subject" json:"subject"`
  Version             int                   `yaml:"version" json:"version"`
  Serial              string                `yaml:"serial" json:"serial"`
  Issuer              string                `yaml:"issuer" json:"issuer"`
  KeyUsage            string                `yaml:"key_usage" json:"key_usage"`
  IsCA                bool                  `yaml:"is_ca" json:"is_ca"`
  DNSNames            []string              `yaml:"dns_names" json:"dns_names"`
  Emails              []string              `yaml:"emails" json:"emails"`
  IPAddrs             []string              `yaml:"ips" json:"ips"`
  URIs                []string              `yaml:"uris" json:"uris"`
  NotBefore           time.Time             `yaml:"not_before" json:"not_before"`
  NotAfter            time.Time             `yaml:"not_after" json:"not_after"`
}

type YamlableTlsConn struct {
  Handshook           bool                  `yaml:"handshook" json:"handshook"`
  Cipher              string                `yaml:"cipher" json:"cipher"`
  Proto               string                `yaml:"proto" json:"proto"`
  Servername          string                `yaml:"sni" json:"sni"`
  PeerCerts           []*YamlableTlsCert    `yaml:"peer_certs" json:"peer_certs"`
  VerifiedCerts       []*YamlableTlsCert    `yaml:"verified_certs" json:"verified_certs"`
}

func NewYamlableTlsConn(r *http.Request) *YamlableTlsConn {
  if r.TLS == nil {
    return nil
  }

  verified := make([]*YamlableTlsCert, 0)
  for _, chain := range r.TLS.VerifiedChains {
    for _, crt := range chain {
      verified = append(verified, NewYamlableTlsCert(crt))
    }
  }

  return &YamlableTlsConn{
    Handshook:          r.TLS.HandshakeComplete,
    Cipher:             CipherSuiteToString(r.TLS.CipherSuite),
    Proto:              r.TLS.NegotiatedProtocol,
    Servername:         r.TLS.ServerName,
    VerifiedCerts:      verified,
  }
}

func NewYamlableTlsCert(cert *x509.Certificate) *YamlableTlsCert {
  return &YamlableTlsCert{
    Subject: cert.Subject.CommonName,
    Version: cert.Version,
    Serial:  fmt.Sprintf("%b", cert.SerialNumber),
    Issuer:  cert.Issuer.CommonName,
    KeyUsage: fmt.Sprintf("%s", cert.KeyUsage),
  }
}
