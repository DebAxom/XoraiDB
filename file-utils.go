package main

import (
	"log"
	"os"
)

func FileExistsOrCreate(filePath string) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		file, _ := os.Create(filePath)
		file.Close()
	}
}

func FolderExistsOrCreate(folderpath string) {
	if _, err := os.Stat(folderpath); os.IsNotExist(err) {
		e := os.MkdirAll(folderpath, os.ModePerm)
		if e != nil {
			log.Fatal(e)
		}
	}
}

func AppendToFile(filePath, data string) {
	labelFile, _ := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	labelFile.WriteString(data)
	labelFile.Close()
}

func CreateAndWrite(filePath, data string) {
	File, _ := os.Create(filePath)
	File.Write([]byte(data))
	File.Close()
}
