package harness

import (
	_ "embed"
	"strings"
)

//go:embed VERSION
var rawVersion string

// Version is the current version of the harness module.
var Version = strings.TrimSpace(rawVersion)
