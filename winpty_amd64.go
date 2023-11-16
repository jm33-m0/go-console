//go:build amd64
// +build amd64

package console

import (
	"embed"
)


//go:embed winpty/amd64/*
var winpty_deps embed.FS

