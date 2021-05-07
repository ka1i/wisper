package app

import (
	"fmt"
	"log"
	"os"

	"github.com/ka1i/wisper/internal"
	"github.com/ka1i/wisper/pkg/usage"
	"github.com/ka1i/wisper/pkg/version"
)

func service() {
	internal.InitApp()

	addr := internal.Serve()
	fw := internal.Watch()

	localServe := fmt.Sprintf("http://localhost:%d", addr.Port)
	log.Printf("started server on %s\n", localServe)

	app := internal.Wisper(localServe, fw)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func start(argc int, argv []string) {
	switch argv[0] {
	case "-h", "--help", "help":
		usage.Usage()
	case "-v", "--version", "version":
		version.Version()
	default:
		log.Println("please check usage")
	}
}

// Wisper run ...
func Wisper() int {
	if len(os.Args) > 1 {
		var argc = len(os.Args)
		var argv = os.Args[1:]
		start(argc, argv)
	} else {
		service()
	}
	return 0
}
