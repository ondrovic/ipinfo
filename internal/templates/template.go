package templates

import (
	"embed"
)

//go:embed html/template.html
var HtmlTemplateFS embed.FS
