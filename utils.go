package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

// CheckLock checks whether a lock file with a given name exists in a directory
// name of file musst be {name}.lock
func CheckLock(name string, weekFolder string) bool {
	fileName := fmt.Sprintf("%s.lock", name)
	filePath := filepath.Join(weekFolder, fileName)
	_, err := os.Stat(filePath)
	return err == nil
}

// GetImageFiles returns all images files located in the given folder
// image files end with jpg,jpeg,gif or png
func GetImageFiles(folder string) ([]os.FileInfo, error) {
	imgFiles, err := ioutil.ReadDir(folder)
	if err != nil {
		return nil, err
	}
	images := []os.FileInfo{}
	for _, img := range imgFiles {
		if !img.IsDir() {
			switch filepath.Ext(img.Name())[1:] {
			case
				"jpg",
				"jpeg",
				"gif",
				"png":
				images = append(images, img)
			}
		}
	}
	return images, nil
}