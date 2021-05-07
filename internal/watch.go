package internal

import (
	"time"

	"github.com/ka1i/wisper/pkg/utils"
	"github.com/progrium/watcher"
)

func Watch(dir string) *watcher.Watcher {
	fw := watcher.New()

	utils.Fatal(fw.AddRecursive(dir))

	go fw.Start(450 * time.Millisecond)

	return fw
}
