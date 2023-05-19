package tls

import (
  "crypto/tls"
  "crypto/x509"
)

func (ss *SelfSign) LoadTLSConfig() error {
  var clientauth tls.ClientAuthType = tls.NoClientCert
  certPem := ss.EncodeCertToPem()
  certPrivKeyPem := ss.EncodePrivateKeyToPem()
  caPem := ss.EncodeCACertToPem()

  serverCert, err := tls.X509KeyPair(certPem.Bytes(), certPrivKeyPem.Bytes())
  if err != nil {
    return err
  }

  certpool := x509.NewCertPool()
  certpool.AppendCertsFromPEM(caPem.Bytes())

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
    RootCAs:              nil, // use host CA set
    ClientCAs:            nil,
    Certificates:         []tls.Certificate{serverCert},
    ServerName:           ss.Config.CommonName,
    ClientAuth:           clientauth,
    InsecureSkipVerify:   ss.Config.SkipVerify == "true",
//    CipherSuites:
//    MinVersion:           
//    MaxVersion:
  }

  return nil
}

