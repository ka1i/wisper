package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/ka1i/wisper/internal"
	"github.com/ka1i/wisper/pkg/usage"
	"github.com/ka1i/wisper/pkg/version"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	if len(os.Args) > 1 {
		var argv = os.Args[1:]
		switch argv[0] {
		case "-h", "--help", "help":
			usage.Usage()
		case "-v", "--version", "version":
			version.Version()
		default:
			log.Println("please check usage")
		}
	}
	internal.InitApp()

	//addr := internal.Serve()
	fw := internal.Watch()

	localServe := fmt.Sprintf("http://localhost:%d", 3000)//addr.Port)
	log.Printf("started server on %s\n", localServe)

	internal.Wisper(localServe, fw)
}
