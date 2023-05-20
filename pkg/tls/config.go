package tls

import (
  "crypto/tls"
  "crypto/x509"
)

func NewTLSConfig(cn, inclientauth, inmintls, inmaxtls string, skipverify bool) *tls.Config {
  var clientauth tls.ClientAuthType = tls.NoClientCert
  var mintls uint16 = tls.VersionTLS10
  var maxtls uint16 = tls.VersionTLS13

  mintls = ParseTlsVersion(inmintls, mintls)
  maxtls = ParseTlsVersion(inmaxtls, maxtls)

  switch inclientauth {
    case "none":
      clientauth = tls.NoClientCert
    case "request":
      clientauth = tls.RequestClientCert
    case "require":
      clientauth = tls.RequireAnyClientCert
    case "verify":
      clientauth = tls.VerifyClientCertIfGiven
    case "mutual":
      clientauth = tls.RequireAndVerifyClientCert
    default:
      clientauth = tls.NoClientCert
  }

  return &tls.Config{
    ServerName:           cn,
    ClientAuth:           clientauth,
    InsecureSkipVerify:   skipverify,
    MinVersion:           mintls,
    MaxVersion:           maxtls,
  }
}


func (ss *SelfSign) TLSConfig() (*tls.Config, error) {
  cfg := NewTLSConfig(ss.Config.CommonName, ss.Config.ClientAuth, ss.Config.MinTLS, ss.Config.MaxTLS, ss.Config.SkipVerify)

  cert, err := tls.X509KeyPair(ss.EncodeCertToPem().Bytes(), ss.EncodeCertPrivateKeyToPem().Bytes())
  if err != nil {
    return nil, err
  }

  cfg.Certificates = []tls.Certificate{cert}

  if cfg.ClientAuth == tls.VerifyClientCertIfGiven || cfg.ClientAuth == tls.RequireAndVerifyClientCert {
    certpool := x509.NewCertPool()
    certpool.AddCert(ss.CACert)
//    certpool.AppendCertsFromPEM(ss.EncodeCACertToPem().Bytes())
    cfg.ClientCAs = certpool
  }

  return cfg, nil
}

func (ss *SelfSign) TLSClientConfig() (*tls.Config, error) {
  cfg := NewTLSConfig(ss.Config.CommonName, ss.Config.ClientAuth, ss.Config.MinTLS, ss.Config.MaxTLS, ss.Config.SkipVerify)

  cert, err := tls.X509KeyPair(ss.EncodeClientCertToPem().Bytes(), ss.EncodeClientCertPrivateKeyToPem().Bytes())
  if err != nil {
    return nil, err
  }

  cfg.Certificates = []tls.Certificate{cert}

  certpool := x509.NewCertPool()
  certpool.AddCert(ss.CACert)
  cfg.RootCAs = certpool

  return cfg, nil
}


func ParseTlsVersion(s string, defaultvalue uint16) uint16 {
  switch s {
    case "1.0":
      return tls.VersionTLS10
    case "1.1":
      return tls.VersionTLS11
    case "1.2":
      return tls.VersionTLS12
    case "1.3":
      return tls.VersionTLS13
    default:
      return defaultvalue
  }
}
