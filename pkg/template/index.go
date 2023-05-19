package template

import (
  "html/template"
)

var IndexTemplate = template.Must(template.New("index.html").Parse(`
<!DOCTYPE html>
<html><head>
<title>{{.title}}</title>
</style>
</head><body>
<table>
  <tr>
    <th>Request Info</th>
  </tr>
  <tr>
    <td>URL</td>
    <td>{{ .request.URL }}</td>
  </tr>
  <tr>
    <td>Method</td>
    <td>{{ .request.Method }}</td>
  </tr>
  <tr>
    <td>Remote Addr</td>
    <td>{{ .request.RemoteAddr }}</td>
  </tr>
  <tr>
    <td>Host</td>
    <td>{{ .request.Host }}</td>
  </tr>
  <tr>
    <td>Protocol</td>
    <td>{{ .request.Proto }}</td>
  </tr>
  <tr>
    <td>Websocket</td>
    <td>{{ .request.Websocket }}</td>
  </tr>
  <tr>
    <th>Headers</th>
  </tr>
  {{- range $key, $val := .request.Headers }}
  <tr>
    <td>{{ $key }}</td>
    <td>{{ index $val 0 }}</td>
  </tr>
  {{- range $i, $h := $val }}
  {{- if gt $i 0 }}
  <tr>
    <td></td>
    <td>{{ index $val $i }}</td>
  </tr>
  {{- end }}
  {{- end }}
  {{- end }}
  <tr>
    <th>Query Params</th>
  </tr>
  {{- range $key, $val := .request.Query }}
  <tr>
    <td>{{ $key }}</td>
    <td>{{ index $val 0 }}</td>
  </tr>
  {{- range $i, $h := $val }}
  {{- if gt $i 0 }}
  <tr>
    <td></td>
    <td>{{ index $val $i }}</td>
  </tr>
  {{- end }}
  {{- end }}
  {{- end }}
  {{- if .request.TLS }}
  <tr>
    <th>TLS</th>
  </tr>
  <tr>
    <td>{{ .request.TLS }}</td>
  </tr>
  {{- end }}
  <tr>
    <th>Body</th>
  </tr>
  <tr>
    <td>{{ .request.Body }}</td>
  </tr>
  {{- if .request.Error }}
  <tr>
    <td>{{ .request.Error }}</td>
  </tr>
  {{- end }}
</table>
</body></html>
`))

