package main

import (
	"fmt"
	"os"

	"example.com/files/fileutils"
)

func main() {
	rootPath, _ := os.Getwd()
	filePath := rootPath + "/data/test.txt"
	contents, err := fileutils.ReadTextFile(filePath)
	if err != nil {
		fmt.Printf("Error reading test file %v", err)
	} else {
		fmt.Println(contents)
		newContent := fmt.Sprintf("Original: \n%v\n\nDouble Original: \n%v\n%v", contents, contents, contents)
		fmt.Println(newContent)
		fileutils.WriteToFile(filePath+".output.txt", newContent)
	}
}
