package tls

import (
  "crypto/tls"
  "crypto/x509"
)

func (ss *SelfSign) LoadTLSConfig() error {
  var clientauth tls.ClientAuthType = tls.NoClientCert
  var mintls uint16 = tls.VersionTLS10
  var maxtls uint16 = tls.VersionTLS13
  certPem := ss.EncodeCertToPem()
  certPrivKeyPem := ss.EncodeCertPrivateKeyToPem()

  serverCert, err := tls.X509KeyPair(certPem.Bytes(), certPrivKeyPem.Bytes())
  if err != nil {
    return err
  }

  mintls = ParseTlsVersion(ss.Config.MinTLS, mintls)
  maxtls = ParseTlsVersion(ss.Config.MaxTLS, maxtls)

  switch ss.Config.ClientAuth {
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

  ss.TLSConfig = &tls.Config{
    Certificates:         []tls.Certificate{serverCert},
    ServerName:           ss.Config.CommonName,
    ClientAuth:           clientauth,
    InsecureSkipVerify:   ss.Config.SkipVerify == "true",
    MinVersion:           mintls,
    MaxVersion:           maxtls,
//    CipherSuites:
  }

  if ss.TLSConfig.ClientAuth == tls.VerifyClientCertIfGiven || ss.TLSConfig.ClientAuth == tls.RequireAndVerifyClientCert {
    certpool := x509.NewCertPool()
    certpool.AddCert(ss.CACert)
    ss.TLSConfig.ClientCAs = certpool
  }

  return nil
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
