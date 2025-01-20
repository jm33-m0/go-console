//go:build arm64
// +build arm64

package console

import (
	"errors"

	"github.com/iamacarpet/go-winpty"
)

type consoleWindows struct {
	initialCols int
	initialRows int

	file *winpty.WinPTY

	cwd string
	env []string
}

func newNative(cols int, rows int) (Console, error) {
	return nil, errors.New("winpty for arm64 not implemented")
}
