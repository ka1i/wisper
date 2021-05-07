package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/ka1i/wisper/internal"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	internal.InitApp()

	addr := internal.Serve()
	fw := internal.Watch()

	localServe := fmt.Sprintf("http://localhost:%d", addr.Port)
	log.Printf("started server on %s\n", localServe)

	app := internal.Wisper(localServe, fw)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
