package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
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

func parseInput(filePath string) ([][2]int, [][]int, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return [][2]int{}, [][]int{}, fmt.Errorf("Problem reading file")
	}
	inputStr := string(input)
	dividedInput := strings.Split(inputStr, "\n\n")
	rulesStr, updatesStr := strings.Split(dividedInput[0], "\n"), strings.Split(dividedInput[1], "\n")
	rules := make([][2]int, len(rulesStr))
	updates := make([][]int, len(updatesStr))
	for i, line := range rulesStr {
		splitLine := strings.Split(line, "|")
		rules[i][0], _ = strconv.Atoi(splitLine[0])
		rules[i][1], _ = strconv.Atoi(splitLine[1])
	}
	for i, line := range updatesStr {
		splitLine := strings.Split(line, ",")
		for _, val := range splitLine {
			num, _ := strconv.Atoi(val)
			updates[i] = append(updates[i], num)
		}
	}
	return rules, updates, nil
}

func main() {
	setCwdToSourceFile()
	rules, updates, err := parseInput("input")
	if err != nil {
		return
	}
	fmt.Print(rules, "\n\n----------------------------\n\n", updates)
}
