package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/ka1i/wisper/internal"
	"github.com/ka1i/wisper/pkg/utils"
)

var home = utils.GetHome()

func init() {
	runtime.LockOSThread()
}

func main() {
	internal.InitApp()

	addr := internal.Serve(home)
	fw := internal.Watch(home)

	localServe := fmt.Sprintf("http://localhost:%d", addr.Port)
	log.Printf("started server on %s\n", localServe)

	app := internal.Wisper(localServe, fw)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
