package main

import (
	"fmt"

	"github.com/ka1i/matrix/pkg/version"
)

func main() {
	fmt.Printf("MATRIX:%v\n", version.Version.ToString())
	version.Version.Print()
}
