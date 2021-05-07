package version

import "fmt"

var (
	// VERSION ...
	VERSION string
	// UPTIME ...
	UPDATE string
)

// Version ...
func Version() {
	fmt.Printf("Version     :%s \n", string(VERSION))
	fmt.Printf("Last Update :%s\n", string(UPDATE))
}
