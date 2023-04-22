package main

import (
	"bufio"
	"encoding/json"
	"os"
	"path"
	"regexp"
	"strings"

	"golang.org/x/exp/slices"
)

func ObjectExists(dbname, objname string) bool {
	res := false
	root := path.Join(DBdir, dbname, "obj", "label")
	files := ListFiles(root)
	for _, file := range files {
		content, _ := os.ReadFile(path.Join(root, file.Name()))
		if TextExists(string(content), "@"+objname+"-") {
			res = true
			break
		}
	}
	return res
}

func getLabel(objPath string) string {
	r, _ := regexp.Compile("!(.*?);")
	objFile, _ := os.Open(objPath)
	scanner := bufio.NewScanner(objFile)
	scanner.Scan()
	objFile.Close()
	return r.FindStringSubmatch(scanner.Text())[1]
}

func addBubbleToObj(objPath, bubbleName, id string) {
	content, _ := os.ReadFile(objPath)
	data := strings.Replace(string(content), "\n", "-"+bubbleName+"#"+id+";\n", 1)
	os.WriteFile(objPath, []byte(data), 0644)
}

func RemoveBubbleFromObj(objPath, bubbleName, id string) {
	content, _ := os.ReadFile(objPath)
	data := strings.Replace(string(content), "-"+bubbleName+"#"+id+";", "", 1)
	os.WriteFile(objPath, []byte(data), 0644)
}

func getAllBubbles(objPath string) []string {
	bubbles := []string{}
	r, _ := regexp.Compile("-(.*?);")

	objFile, _ := os.Open(objPath)
	scanner := bufio.NewScanner(objFile)
	scanner.Scan()
	metadata := scanner.Text()
	objFile.Close()

	matches := r.FindAllStringSubmatch(metadata, -1)
	for _, v := range matches {
		bubbles = append(bubbles, v[1])
	}

	return bubbles
}

func getAllBubblesTypes(objPath string) []string {
	bubbles := []string{}
	r, _ := regexp.Compile("-(.*?);")

	objFile, _ := os.Open(objPath)
	scanner := bufio.NewScanner(objFile)
	scanner.Scan()
	metadata := scanner.Text()
	objFile.Close()

	matches := r.FindAllStringSubmatch(metadata, -1)
	for _, bubble := range matches {
		bubbleType := strings.Split(bubble[1], "#")[0]
		if slices.Contains(bubbles, bubbleType) {
			continue
		}
		bubbles = append(bubbles, bubbleType)
	}

	return bubbles
}

func getData(objPath string) map[string]any {
	data := map[string]any{}
	objFile, _ := os.ReadFile(objPath)
	content := strings.Split(string(objFile), "\n")[1]
	json.Unmarshal([]byte(content), &data)
	return data
}
