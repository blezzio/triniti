package templates

import "embed"

//go:embed templates/*.gohtml
//go:embed static
var FS embed.FS
