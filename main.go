package main

import (
  "github.com/namsral/flag"
  "taemon1337/http-test-server/pkg/server"
)

func main() {
  srv := server.NewServer()
  cfg := srv.Config

  flag.String(flag.DefaultConfigFlagname, "", "path to config file")
  flag.StringVar(&cfg.HttpAddr, "addr", cfg.HttpAddr, "socket address to listen for HTTP on")
  flag.StringVar(&cfg.HttpsAddr, "saddr", cfg.HttpsAddr, "socket address to listen for HTTPS/TLS on")
  flag.BoolVar(&cfg.UseTLS, "tls", cfg.UseTLS, "Server HTTPS as well as HTTP (will")
  flag.StringVar(&cfg.CommonName, "servername", cfg.CommonName, "The common name and servername for TLS")
  flag.StringVar(&cfg.SkipVerify, "skipverify", cfg.SkipVerify, "Skip TLS verification")
  flag.StringVar(&cfg.ClientAuth, "clientauth", cfg.ClientAuth, "TLS verification of client; One of 'none', 'request', 'verify', 'require', 'mutual' (mutual is most secure)")
  flag.StringVar(&cfg.RSAKeySize, "keysize", cfg.RSAKeySize, "The number of bits for the RSA key to use for TLS")
  flag.Parse()

  srv.Run()
}
