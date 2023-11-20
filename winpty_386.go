//go:build windows && 386
// +build windows,386

package console

import (
	"embed"
)

//go:embed winpty/386/*
var winpty_deps embed.FS
