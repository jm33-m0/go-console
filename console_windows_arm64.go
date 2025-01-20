//go:build arm64
// +build arm64

package console

import (
	"errors"
)

func newNative(cols int, rows int) (Console, error) {
	return nil, errors.New("winpty for arm64 not implemented")
}
