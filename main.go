package main

import (
  "github.com/namsral/flag"
  "taemon1337/http-test-server/pkg/server"
)

func main() {
  srv := server.NewServer()
  cfg := srv.Config

  flag.BoolVar(&cfg.UseTLS, "tls", cfg.UseTLS, "Start the HTTPS/TLS server")
  flag.BoolVar(&cfg.UseHTTP, "http", cfg.UseHTTP, "Start the HTTP server")
  flag.String(flag.DefaultConfigFlagname, "", "path to config file")
  flag.StringVar(&cfg.HttpAddr, "http-addr", cfg.HttpAddr, "socket address to listen for HTTP on")
  flag.StringVar(&cfg.HttpsAddr, "tls-addr", cfg.HttpsAddr, "socket address to listen for HTTPS/TLS on")
  flag.StringVar(&cfg.CommonName, "servername", cfg.CommonName, "The common name and servername for TLS")
  flag.StringVar(&cfg.SkipVerify, "skipverify", cfg.SkipVerify, "Skip TLS verification")
  flag.StringVar(&cfg.ClientAuth, "clientauth", cfg.ClientAuth, "TLS verification of client; One of 'none', 'request', 'verify', 'require', 'mutual' (mutual is most secure)")
  flag.StringVar(&cfg.RSAKeySize, "keysize", cfg.RSAKeySize, "The number of bits for the RSA key to use for TLS")
  flag.StringVar(&cfg.MinTLS, "min-tls-version", cfg.MinTLS, "The minimum TLS version to accept, 1.0-1.3")
  flag.StringVar(&cfg.MaxTLS, "max-tls-version", cfg.MaxTLS, "The maximum TLS version to accept, 1.0-1.3")
  flag.Parse()

  if cfg.UseHTTP || cfg.UseTLS {
    srv.Run()
  } else {
    flag.PrintDefaults()
    srv.Logger.Printf("No '--tls' or '--http' service enabled, nothing to do.")
  }
}
