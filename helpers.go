package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func readFiles(path string) {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	files, err := f.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, v := range files {
		if v.IsDir() {
			readFiles(fmt.Sprintf("%s/%s", path, v.Name()))
		} else {
			str, _ := filepath.Abs(fmt.Sprintf("%s/%s", path, v.Name()))
			filesFull = append(filesFull, str)
		}
	}
}
