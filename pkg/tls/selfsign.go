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
  CertKey       *rsa.PrivateKey
  ClientCert    *x509.Certificate
  ClientCertKey *rsa.PrivateKey
  ClientBytes   []byte
  CAPrivateKey  *rsa.PrivateKey
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

  clientcert, err := NewCertificate(cfg.ClientCertOptions())
  if err != nil {
    return nil, err
  }

  cakey, err := NewRSAKey(cfg.KeyOptions())
  if err != nil {
    return nil, err
  }

  certkey, err := NewRSAKey(cfg.KeyOptions())
  if err != nil {
    return nil, err
  }

  clientcertkey, err := NewRSAKey(cfg.KeyOptions())
  if err != nil {
    return nil, err
  }

  selfsignedca, err := SignCertificate(ca, ca, &cakey.PublicKey, cakey)
  if err != nil {
    return nil, err
  }

  signed, err := SignCertificate(ca, cert, &certkey.PublicKey, cakey)
  if err != nil {
    return nil, err
  }

  clientsigned, err := SignCertificate(ca, clientcert, &clientcertkey.PublicKey, cakey)
  if err != nil {
    return nil, err
  }

  return &SelfSign{
    Config: cfg,
    CACert: ca,
    CACertBytes: selfsignedca,
    Cert: cert,
    CertBytes: signed,
    CertKey: certkey,
    ClientCert: clientcert,
    ClientBytes: clientsigned,
    ClientCertKey: clientcertkey,
    CAPrivateKey: cakey,
    TLSConfig: nil,
  }, nil
}

func (ss *SelfSign) EncodeCertToPem() *bytes.Buffer {
  return EncodeToPem(ss.CertBytes)
}

func (ss *SelfSign) EncodeClientCertToPem() *bytes.Buffer {
  return EncodeToPem(ss.ClientBytes)
}

func (ss *SelfSign) EncodeCACertToPem() *bytes.Buffer {
  return EncodeToPem(ss.CACertBytes)
}

func (ss *SelfSign) EncodeCAPrivateKeyToPem() *bytes.Buffer {
  return EncodePrivateKeyToPem(ss.CAPrivateKey)
}

func (ss *SelfSign) EncodeCertPrivateKeyToPem() *bytes.Buffer {
  return EncodePrivateKeyToPem(ss.CertKey)
}

func (ss *SelfSign) EncodeClientCertPrivateKeyToPem() *bytes.Buffer {
  return EncodePrivateKeyToPem(ss.ClientCertKey)
}

func (ss *SelfSign) CertChain() string {
  return ss.EncodeClientCertToPem().String() + ss.EncodeCACertToPem().String()
}

func (ss *SelfSign) CertAndKey() string {
  return ss.EncodeClientCertToPem().String() + ss.EncodeClientCertPrivateKeyToPem().String()
}

func (ss *SelfSign) ClientCertAndKey() string {
  return ss.EncodeClientCertToPem().String() + ss.EncodeClientCertPrivateKeyToPem().String()
}

func (ss *SelfSign) String() string {
  s := ""
  s += ss.EncodeCACertToPem().String() + "\n"
  s += ss.EncodeCertToPem().String() + "\n"
  s += ss.EncodeClientCertToPem().String() + "\n"
  return s
}
