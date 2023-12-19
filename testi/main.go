package main

import (
	"os"
)

func main() {
	appendFile := false
	data := "Jotain \n muuta \n ja vielä \n jotain"
	if appendFile {
		os.WriteFile("testfile.txt", []byte(data), os.ModeAppend)
	} else {
		var before []byte
		before, _ = os.ReadFile("testfile.txt")
		before = append(before, []byte("\n")...)
		before = append(before, before...)
		os.WriteFile("testfile.txt", before, 0666)

	}
}
