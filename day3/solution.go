package main

import (
	"fmt"
	"os"
)

func parseInput(filePath string) (string, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Problem reading file")
	}
	return string(input), nil
}

func main() {
	// lol, did this one in vim
	data, err := parseInput("/Users/samtang/Programming/advent-of-code-20204/day3/input.txt")
	if err != nil {
		fmt.Print(err)
		return
	}
	fmt.Print(data)
}
