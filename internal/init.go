package internal

import (
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/ka1i/wisper/pkg/assets"
	"github.com/ka1i/wisper/pkg/utils"
)

var (
	assetsRoot = "web"
)

func InitApp() {
	if !utils.Check(utils.GetHome()) {
		log.Printf("unpkg embed assets")
		err := assetsUnpkg(assetsRoot, utils.GetHome())
		if err != nil {
			utils.Fatal(err)
		}
	}
}

func assetsUnpkg(assetsName string, template string) error {
	storage, err := assets.Storage.ReadDir(assetsName)
	if err != nil {
		return err
	}
	for _, file := range storage {
		storagePath := path.Join(assetsName, file.Name())

		localPath := path.Join(template, storagePath[len(assetsRoot)+1:])
		err = os.MkdirAll(filepath.Dir((localPath)), os.ModePerm)
		if err != nil {
			return err
		}
		if file.IsDir() {
			err := assetsUnpkg(path.Join(assetsName, file.Name()), template)
			if err != nil {
				return err
			}
		} else if file.Name()[0] == 46 {
			continue
		} else {
			fmt.Printf("unpkg: %s\n", localPath)
			in, err := assets.Storage.Open(storagePath)
			if err != nil {
				return err
			}
			defer in.Close()
			out, err := os.Create(localPath)
			if err != nil {
				return err
			}
			defer out.Close()
			io.Copy(out, in)
		}
	}
	return nil
}
