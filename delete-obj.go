package main

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
)

func deleteObj(dbname, objname string) (bool, string) {
	objPath := path.Join(DBdir, dbname, "obj", "@"+objname)

	if !ObjectExists(dbname, objname) {
		return false, "Object " + objname + " doesn't exist !"
	}

	label := getLabel(objPath)
	if label == "?" {
		label = "null"
	}

	allBubbles := getAllBubblesTypes(objPath) // Get all the different types of bubbles that belongs to the object
	for _, bubble := range allBubbles {
		os.RemoveAll(path.Join(DBdir, dbname, "bubble", bubble, "@"+objname)) // Delete the /bubble/bubbleType/@obj directory
	}

	labelFilePath := path.Join(DBdir, dbname, "obj", "label", label)
	removeLabel(labelFilePath, objname, label) // Remove @obj- from the label file

	os.Remove(objPath)
	return true, ""
}

func removeLabel(labelFilePath, objname, label string) {
	content, _ := ioutil.ReadFile(labelFilePath)
	newContent := strings.Replace(string(content), "@"+objname+"-", "", 1)
	ioutil.WriteFile(labelFilePath, []byte(newContent), 0644)

	// Delete the label file if it's empty
	f, _ := os.Stat(labelFilePath)
	if f.Size() == 0 {
		os.Remove(labelFilePath)
	}
}
