package fileutils

import (
	"fmt"
	"os"
)

func ReadTextFile(filename string) (string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Printf("Failed to read file: %s", err)
		return "", err
	}
	return string(content), nil
}
