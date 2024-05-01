package assets

import (
	"embed"
)

//go:embed templates static
var EmbeddedFiles embed.FS

// As of now templates/partials/ and base.tmpl are not used, only the pages directory is used
