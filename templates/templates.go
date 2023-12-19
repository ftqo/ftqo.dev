package templates

import "embed"

//go:embed */**.html
var T embed.FS
