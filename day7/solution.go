package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func setCwdToSourceFile() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		fmt.Println("Error getting caller")
		return
	}
	dir := filepath.Dir(filename)

	// Change the current working directory
	err := os.Chdir(dir)
	if err != nil {
		fmt.Println("Error changing directory:", err)
		return
	}
}

func parseInput(filePath string) ([]string, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return []string{""}, fmt.Errorf("Problem reading file")
	}
	inputStr := strings.TrimSpace(string(input))
	return strings.Split(inputStr, "\n"), nil
}

func main() {
	setCwdToSourceFile()
	input, err := parseInput("input")
	if err != nil {
		fmt.Print("Error parsing input: ", err)
		return
	}
	fmt.Print(input[len(input)-1])
}
