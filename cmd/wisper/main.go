package main

import (
	"os"
	"runtime"

	"github.com/ka1i/wisper/cmd/app"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	os.Exit(app.Wisper())
}
