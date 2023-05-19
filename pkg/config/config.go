package config

type Config struct {
  HttpAddr          string
  HttpsAddr         string
  UseTLS            bool
  SkipVerify        string
  CommonName        string
  ClientAuth        string
  CACommonName      string
  CAOrganization    string
  CACountry         string
  CAProvince        string
  CALocality        string
  CAStreetAddress   string
  CAPostalCode      string
  CAYearsValid      string
  CertOrganization  string
  CertCountry       string
  CertProvince      string
  CertLocality      string
  CertStreetAddress string
  CertPostalCode    string
  CertIPAddresses   string
  CertDomains       string
  CertYearsValid    string
  RSAKeySize        string
}

func NewConfig() *Config {
  return &Config{
    HttpAddr:         ":8080",
    HttpsAddr:        ":8443",
    UseTLS:           false,
    SkipVerify:       "false",
    CommonName:       "localhost",
    ClientAuth:       "none",
    CACommonName:     "root.localhost",
    CAOrganization:   "",
    CACountry:        "US",
    CAProvince:       "",
    CALocality:       "",
    CAStreetAddress:  "",
    CAPostalCode:     "",
    CAYearsValid:     "10",
    CertOrganization: "",
    CertCountry:      "US",
    CertProvince:     "",
    CertLocality:     "",
    CertStreetAddress:  "",
    CertPostalCode:   "",
    CertIPAddresses:  "",
    CertDomains:      "",
    CertYearsValid:   "1",
    RSAKeySize:       "4096",
  }
}

func (c *Config) CAOptions() map[string]string {
  return map[string]string{
    "cn":  c.CACommonName,
    "org": c.CAOrganization,
    "country": c.CACountry,
    "province": c.CAProvince,
    "locality": c.CALocality,
    "street": c.CAStreetAddress,
    "postal": c.CAPostalCode,
    "isCA": "true",
    "client": "true",
    "server": "true",
    "years": c.CAYearsValid,
  }
}

func (c *Config) CertOptions() map[string]string {
  return map[string]string{
    "cn":  c.CommonName,
    "org": c.CertOrganization,
    "country": c.CertCountry,
    "province": c.CertProvince,
    "locality": c.CertLocality,
    "street": c.CertStreetAddress,
    "postal": c.CertPostalCode,
    "isCA": "true",
    "client": "true",
    "server": "true",
    "ips": c.CertIPAddresses,
    "domains":  c.CertDomains,
    "years": c.CertYearsValid,
  }
}

func (c *Config) KeyOptions() map[string]string {
  return map[string]string{
    "keysize": c.RSAKeySize,
  }
}
