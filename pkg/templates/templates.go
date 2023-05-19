package templates

import (
  "github.com/gin-gonic/contrib/renders/multitemplate"
)

func GetTemplates() multitemplate.Render {
  templates := multitemplate.New()

  templates.AddFromString("index.html", IndexTemplate)
  templates.AddFromString("certs.html", CertTemplate)

  return templates
}
