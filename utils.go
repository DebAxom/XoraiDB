package main

import (
	"io/fs"
	"log"
	"os"
	"sort"
	"strings"
)

func ListFiles(dir string) []fs.FileInfo {
	files, err := ReadDir(dir)

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func ReadDir(dirname string) ([]fs.FileInfo, error) {
	f, err := os.Open(dirname)
	if err != nil {
		return nil, err
	}
	list, err := f.Readdir(-1)
	f.Close()
	if err != nil {
		return nil, err
	}
	sort.Slice(list, func(i, j int) bool { return list[i].Name() < list[j].Name() })
	return list, nil
}

func TextExists(text, word string) bool {
	if strings.Contains(text, word) {
		return true
	}
	return false
}
