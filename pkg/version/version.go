package version

import (
	"fmt"
	"strings"
)

type version struct {
	ver string
	tag string
	upt string
	env string
}

func (v *version) ToString() string {
	return strings.Join([]string{v.ver, v.tag}, "@")
}

func (v *version) Print() {
	fmt.Printf("Version: %s \n", string(v.ver))
	fmt.Printf("Git Tags: %s \n", string(v.tag))
	fmt.Printf("Update Time: %s \n", string(v.upt))
	fmt.Printf("Compiler Environment: %s \n", string(v.env))
}

var tagStr string
var uptStr string
var envStr string
var Version = getVersion()

func getVersion() *version {
	return &version{
		ver: "v0.0.1",
		tag: tagStr,
		upt: uptStr,
		env: envStr,
	}
}
