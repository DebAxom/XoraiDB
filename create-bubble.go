package main

import (
	"path"
)

func createBubble(dbname, bubblename, objname, id, data string) (bool, string) {

	if !ObjectExists(dbname, objname) {
		return false, "Object " + "@" + objname + " doesn't exists ! \n You need to create the object first before creating a bubble !"
	}

	if BubbleExists(dbname, bubblename, objname, id) {
		return false, "Another bubble of @" + objname + " with the same type and id already exists."
	}

	FolderExistsOrCreate(path.Join(DBdir, dbname, "bubble", bubblename, "@"+objname))

	bubblePath := path.Join(DBdir, dbname, "bubble", bubblename, "@"+objname, id)
	addBubbleToObj(path.Join(DBdir, dbname, "obj", "@"+objname), bubblename, id)
	CreateAndWrite(bubblePath, data)

	return true, ""
}
