# HTTP Test Server

A simply useful HTTP and TLS test server designed to run in any environment and always be reachable and useful for debugging/verification.

### Example Usage

```bash
  ./http-test-server --tls --servername $HOSTNAME
```

### Help Usage

```bash
Usage of ./http-test-server:
  -addr=":8080": socket address to listen for HTTP on
  -clientauth="none": TLS verification of client; One of 'none', 'request', 'verify', 'require', 'mutual' (mutual is most secure)
  -config="": path to config file
  -http=false: Start the HTTP server
  -keysize="4096": The number of bits for the RSA key to use for TLS
  -saddr=":8443": socket address to listen for HTTPS/TLS on
  -servername="localhost": The common name and servername for TLS
  -skipverify="false": Skip TLS verification
  -tls=false: Start the HTTPS/TLS server
```

