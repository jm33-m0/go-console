//go:build windows && amd64
// +build windows,amd64

package console

import (
	"embed"
)

//go:embed winpty/amd64/*
var winpty_deps embed.FS
