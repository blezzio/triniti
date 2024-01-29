package assets

import "embed"

//go:embed templates/*.gohtml
//go:embed static
var FS embed.FS
