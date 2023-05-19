package templates

var CertTemplate = `
<!DOCTYPE html>
<html><head>
<title>{{.title}}</title>
</style>
</head><body>
<table>
  <tr>
    <th>TLS Certificates</th>
  </tr>
  <tr>
    <td>Server Certificate</td>
    <td><a href='/download/ca.crt'>ca.crt</a></td>
  </tr>
  <tr>
    <td>CA Certificate</td>
    <td><a href='/download/server.crt'>server.crt</a></td>
  </tr>
  <tr>
    <td>Client Certificate</td>
    <td><a href='/download/client.crt'>client.crt</a></td>
  </tr>
  <tr>
    <td>Client Certificate Key</td>
    <td><a href='/download/client.key'>client.key</a></td>
  </tr>
</table>
</body></html>
`
