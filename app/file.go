package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func findFile(filename string) (string, error) {
	directory := os.Args[2]
	filepath := filepath.Join(directory, filename)
	fmt.Println("searching for file: " + filepath)
	file_contents, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return string(file_contents[:]), nil
}

func saveFile(file_path string, file_contents string) {
	directory := os.Args[2]
	file_path = filepath.Join(directory, file_path)
	os.WriteFile(file_path, []byte(file_contents), 0666)
}
