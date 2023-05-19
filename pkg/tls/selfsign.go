package tls

import (
  "bytes"
  "crypto/rsa"
  "crypto/x509"
  "crypto/tls"
  "taemon1337/http-test-server/pkg/config"
)

type SelfSign struct {
  Config        *config.Config
  CACert        *x509.Certificate
  CACertBytes   []byte
  Cert          *x509.Certificate
  CertBytes     []byte
  PrivateKey    *rsa.PrivateKey
  TLSConfig     *tls.Config
}

func NewSelfSign(cfg *config.Config) (*SelfSign, error) {
  ca, err := NewCertificate(cfg.CAOptions())
  if err != nil {
    return nil, err
  }

  cert, err := NewCertificate(cfg.CertOptions())
  if err != nil {
    return nil, err
  }

  key, err := NewRSAKey(cfg.KeyOptions())
  if err != nil {
    return nil, err
  }

  selfsignedca, err := SignCertificate(ca, ca, key)
  if err != nil {
    return nil, err
  }

  signed, err := SignCertificate(ca, cert, key)
  if err != nil {
    return nil, err
  }

  return &SelfSign{
    Config: cfg,
    CACert: ca,
    CACertBytes: selfsignedca,
    Cert: cert,
    CertBytes: signed,
    PrivateKey: key,
    TLSConfig: nil,
  }, nil
}

func (ss *SelfSign) EncodeCertToPem() *bytes.Buffer {
  return EncodeToPem(ss.CertBytes)
}

func (ss *SelfSign) EncodeCACertToPem() *bytes.Buffer {
  return EncodeToPem(ss.CACertBytes)
}

func (ss *SelfSign) EncodePrivateKeyToPem() *bytes.Buffer {
  return EncodePrivateKeyToPem(ss.PrivateKey)
}

func (ss *SelfSign) String() string {
  s := ""
  s += ss.EncodeCACertToPem().String() + "\n"
  s += ss.EncodeCertToPem().String() + "\n"
  return s
}


