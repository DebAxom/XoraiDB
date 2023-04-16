package main

import (
	"fmt"
	"os"
	"path"
)

var DBdir string

func main() {
	cwd, _ := os.Getwd()
	DBdir = path.Join(cwd, "databases")
	fmt.Println("XoraiDB v0.01")

	// createObj("school", "leema", "student", `{name: "Leema Chutia",class: 12,branch: "Arts"}`)
	// createObj("school", "nirmali", "student", `{name: "Nirmali Xonowal",class: 12,branch: "Sci"}`)
	// createObj("school", "anupama", "teacher", `{name: "Anupama Hazarika",dept: "IP"}`)
	// deleteObj("school", "leema")
	// createBubble("school", "marksheet", "leema", "2022", `{history : 95,Assamese: 99,English: 72}`)
	// createBubble("school", "marksheet", "nirmali", "2022", `{history : 60,Assamese: 99,English: 85}`)
	// createBubble("school", "marksheet", "leema", "2023", `{history : 90,Assamese: 100,English: 90}`)
	// createBubble("school", "marksheet", "nirmali", "2023", `{history : 70,Assamese: 99,English: 95}`)
	// deleteBubble("school", "marksheet", "leema", "2023")
}
