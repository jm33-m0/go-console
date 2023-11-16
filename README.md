# go-console

[![Main](https://github.com/runletapp/go-console/actions/workflows/main.yml/badge.svg)](https://github.com/runletapp/go-console/actions/workflows/main.yml)
[![GoDoc](https://godoc.org/github.com/runletapp/go-console?status.svg)](https://godoc.org/github.com/runletapp/go-console)

`go-console` is a cross-platform `PTY` interface. On *nix platforms we rely on [pty](github.com/creack/pty) and on windows [go-winpty](https://github.com/iamacarpet/go-winpty) (go-console will ship with `winpty.dll` and `winpty-agent.exe` from `c:\Program Files\Microsoft VS Code\resources\app\node_modules.asar.unpacked\node-pty\build\Release`  using `go:embed`, so there's no need to include winpty binaries)

## Credits:

- creack for [pty](github.com/creack/pty)
- iamacarpet for [go-winpty](https://github.com/iamacarpet/go-winpty)
- runletapp for [go-console](https://github.com/runletapp/go-console)
- jm33-m0 for [go-console](https://github.com/jm33-m0/go-console)


## Example

```go
package main

import (
	"io"
	"log"
	"os"
	"runtime"
	"sync"

	"github.com/abakum/go-console"
)

func main() {

	proc, err := console.New(120, 60)
	if err != nil {
		panic(err)
	}
	defer proc.Close()

	var args []string

	if runtime.GOOS == "windows" {
		args = []string{"cmd.exe", "/c", "dir"}
	} else {
		args = []string{"ls", "-lah", "--color"}
	}

	if err := proc.Start(args); err != nil {
		panic(err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()

		_, err = io.Copy(os.Stdout, proc)
		if err != nil {
			log.Printf("Error: %v\n", err)
		}
	}()

	if _, err := proc.Wait(); err != nil {
		log.Printf("Wait err: %v\n", err)
	}

	wg.Wait()
}

```
