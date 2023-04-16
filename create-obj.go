package main

import (
	"path"
	"strings"
)

func createObj(dbname, objname, label, data string) (bool, string) {

	objDir := path.Join(DBdir, dbname, "obj")

	if ObjectExists(dbname, objname) {
		return false, "An object with the identifier " + "@" + objname + " already exists !"
	}

	if label == "" {
		label = "?"
		objpath := path.Join(DBdir, dbname, "obj", "@"+objname)
		data = strings.ReplaceAll(data, ",", "\n")
		metadata := "!?;\n"
		CreateAndWrite(objpath, metadata+data)
		AppendToFile(path.Join(objDir, "label", "null"), "@"+objname+"-")
	}

	objpath := path.Join(DBdir, dbname, "obj", "@"+objname)
	metadata := "!" + label + ";\n"
	CreateAndWrite(objpath, metadata+data)

	labelFilePath := path.Join(objDir, "label", label)

	// Check if the label path exists, or else create it.
	FileExistsOrCreate(labelFilePath)

	// Add the object identifier to the label file
	AppendToFile(labelFilePath, "@"+objname+"-")

	return true, ""
}
