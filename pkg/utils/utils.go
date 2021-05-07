package utils

import (
	"log"
	"os/user"
	"path/filepath"
)

func GetHome() string {
	user, err := user.Current()
	if err != nil {
		Fatal(err)
	}
	return filepath.Join(user.HomeDir, ".wisper")
}

func Fatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
