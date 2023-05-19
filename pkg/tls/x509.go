package tls

import (
  "net"
  "time"
  "math/big"
  "bytes"
  "strconv"
  "crypto/x509"
  "crypto/x509/pkix"
  "crypto/rsa"
  "crypto/rand"
  "encoding/pem"
)

func NewCertificate(opts map[string]string) (*x509.Certificate, error) {
  isca := false
  years := 10
  keyusage := x509.KeyUsageDigitalSignature | x509.KeyUsageCertSign
  extkeyusage := []x509.ExtKeyUsage{}
  ips := []net.IP{net.IPv4(127, 0, 0, 1), net.IPv6loopback}
  dnsnames := []string{"localhost"}

  subj := pkix.Name{
    CommonName:    "localhost",
    Organization:  []string{""},
    Country:       []string{"US"},
    Province:      []string{""},
    Locality:      []string{""},
    StreetAddress: []string{""},
    PostalCode:    []string{""},
  }

  if v, ok := opts["cn"]; ok { subj.CommonName = v }
  if v, ok := opts["org"]; ok { subj.Organization = []string{v} }
  if v, ok := opts["country"]; ok { subj.Country = []string{v} }
  if v, ok := opts["province"]; ok { subj.Province = []string{v} }
  if v, ok := opts["locality"]; ok { subj.Locality = []string{v} }
  if v, ok := opts["street"]; ok { subj.StreetAddress = []string{v} }
  if v, ok := opts["postal"]; ok { subj.PostalCode = []string{v} }
  if v, ok := opts["isCA"]; ok { isca = v == "true" }
  if _, ok := opts["client"]; ok { extkeyusage = append(extkeyusage, x509.ExtKeyUsageClientAuth) }
  if _, ok := opts["server"]; ok { extkeyusage = append(extkeyusage, x509.ExtKeyUsageServerAuth) }
  if v, ok := opts["ips"]; ok { ips = append(ips, net.ParseIP(v)) }
  if v, ok := opts["domains"]; ok { dnsnames = append(dnsnames, v) }
  if v, ok := opts["years"]; ok {
    i, err := strconv.Atoi(v)
    if err != nil {
      return nil, err
    }
    years = i
  }

  cert := &x509.Certificate{
    SerialNumber:           big.NewInt(2019),
    Subject:                subj,
    NotBefore:              time.Now(),
    NotAfter:               time.Now().AddDate(years, 0, 0),
    IsCA:                   isca,
    IPAddresses:            ips,
    DNSNames:               dnsnames,
    ExtKeyUsage:            extkeyusage,
    KeyUsage:               keyusage,
    BasicConstraintsValid:  true,
  }

  return cert, nil
}

func SignCertificate(cert, cacert *x509.Certificate, pubkey *rsa.PublicKey, privkey *rsa.PrivateKey) ([]byte, error) {
  signed, err := x509.CreateCertificate(rand.Reader, cert, cacert, pubkey, privkey)
  if err != nil {
    return nil, err
  }

  return signed, nil
}

func EncodeToPem(cert []byte) *bytes.Buffer {
  encodedcert := new(bytes.Buffer)
  pem.Encode(encodedcert, &pem.Block{
    Type: "CERTIFICATE",
    Bytes: cert,
  })

  return encodedcert
}

func EncodePrivateKeyToPem(pri *rsa.PrivateKey) *bytes.Buffer {
  encodedcert := new(bytes.Buffer)
  pem.Encode(encodedcert, &pem.Block{
    Type: "RSA PRIVATE KEY",
    Bytes: x509.MarshalPKCS1PrivateKey(pri),
  })

  return encodedcert
}

func NewRSAKey(opts map[string]string) (*rsa.PrivateKey, error) {
  keysize := 4096

  if v, ok := opts["keysize"]; ok {
    i, err := strconv.Atoi(v)
    if err != nil {
      return nil, err
    }
    keysize = i
  }

  privkey, err := rsa.GenerateKey(rand.Reader, keysize)
  if err != nil {
    return nil, err
  }

  return privkey, nil
}


