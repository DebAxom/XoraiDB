package main

import (
	"os"
	"path"
)

func BubbleExists(dbname, objname, bubblename, id string) bool {
	filePath := path.Join(DBdir, dbname, "bubble", bubblename, objname, id)
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}
