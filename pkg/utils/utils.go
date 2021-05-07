package utils

import (
	"log"
	"os"
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

func Check(path string) bool {
	s, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return s.IsDir()
}
