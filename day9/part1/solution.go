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

func parseInput(filePath string) (string, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return "", fmt.Errorf("Problem reading file: %w", err)
	}
	inputStr := strings.TrimSpace(string(input))
	return inputStr, nil
}

func nTriangle(n int) int {
	return n * (n + 1) / 2
}

func fileValue(index, size, location int) int {
	id := index / 2
	return id * ((location-1)*size + nTriangle(size))
}

func firstPass(input *string) ([][2]int, int) {
	memory := make([][2]int, len(*input))
	partialAns := 0
	place := 0
	for i, char := range *input {
		toInt := int(char - '0')
		memory[i] = [2]int{toInt, place}
		if i%2 == 0 {
			test := fileValue(i, toInt, place)
			partialAns += test
		}
		place += toInt
	}
	return memory, partialAns
}

func secondPass(memory *[][2]int) int {
	leftPointer := 1
	rightPointer := len(*memory) - 1
	partialAns := 0
	if rightPointer%2 == 0 {
		rightPointer -= 1
	}
	for leftPointer <= rightPointer {
		id := rightPointer / 2
		size := min((*memory)[leftPointer][0], (*memory)[rightPointer][0])
	}
	return partialAns
}

func part1() {
	setCwdToSourceFile()
	input, err := parseInput("../input")
	if err != nil {
		fmt.Print("Error parsing input: ", err)
		return
	}
	fmt.Print(firstPass(&input))
}

func main() {
	part1()
}
