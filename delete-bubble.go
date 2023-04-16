package main

import (
	"os"
	"path"
)

func deleteBubble(dbname, bubblename, objname, id string) (bool, string) {

	if !ObjectExists(dbname, objname) {
		return false, "Object " + "@" + objname + " doesn't exists !"
	}
	bubblePath := path.Join(DBdir, dbname, "bubble", bubblename, "@"+objname, id)
	RemoveBubbleFromObj(path.Join(DBdir, dbname, "obj", "@"+objname), bubblename, id)
	os.Remove(bubblePath)
	return true, ""
}
