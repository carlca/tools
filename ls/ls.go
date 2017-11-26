package main

import (
	"fmt"
	"path/filepath"
	"os"
	"time"
)

type File struct {
	Name string
	Size int64
	IsDir bool
	ModTime time.Time
}

func main() {

	rootpath := "C://Go"
	if len(os.Args) > 1 {
		rootpath = os.Args[1]
	}
	files := make([]File, 10)
	err := filepath.Walk(rootpath, func(path string, info os.FileInfo, err error) error {
		if err != filepath.SkipDir {
			files = append(files, parseFileInfo(info))
		}
		return nil
	})
	if err != nil {
		fmt.Print("walk error [%v]\n", err)
	}
	for _, file := range files {
		fmt.Printf("%-60s", file.Name)
		fmt.Printf("%10v  ", file.Size)
		fmt.Print(file.ModTime.Format(time.RFC3339))
		if file.IsDir {
			fmt.Println("  [DIR]")
		} else {
			fmt.Println("")
		}
	}
}

func parseFileInfo(info os.FileInfo) File {
	file := File {
		info.Name(), info.Size(), info.IsDir(), info.ModTime(),
	}
	return file
}