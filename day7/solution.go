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

func parseInput(filePath string) ([]int, [][]int, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return []int{}, [][]int{}, fmt.Errorf("Problem reading file")
	}
	inputStr := strings.TrimSpace(string(input))
	lines := strings.Split(inputStr, "\n")
	values := make([]int, len(lines))
	equations := make([][]int, len(lines))
	for i, line := range lines {
		splitLine := strings.Split(line, ": ")
		values[i], _ = strconv.Atoi(splitLine[0])
		equationsStr := strings.Split(splitLine[1], " ")
		for _, numStr := range equationsStr {
			num, _ := strconv.Atoi(numStr)
			equations[i] = append(equations[i], num)
		}
	}
	return values, equations, nil
}

func nthBit(num, n int) bool {
	return (num>>n)&1 == 1
}

func possibleOperations(n int) [][]bool {
	length := 1 << n
	possible_operations := make([][]bool, length)
	for i := 0; i < length; i++ {
		possible_operations[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			possible_operations[i][j] = nthBit(i, j)
		}
	}
	return possible_operations
}

func operation(num1, num2 int, operation bool) int {
	if operation {
		return num1 + num2
	}
	return num1 * num2
}

func evaluate(equation []int, operations []bool) int {
	answer := equation[0]
	for i := 1; i < len(equation); i++ {
		answer = operation(answer, equation[i], operations[i-1])
	}
	return answer
}

func main() {
	setCwdToSourceFile()
	values, equations, err := parseInput("input")
	if err != nil {
		fmt.Print("Error parsing input: ", err)
		return
	}
	answer := 0
valuesLoop:
	for i := range values {
		possible_operations := possibleOperations(len(equations[i]) - 1)
		for _, operation := range possible_operations {
			evaluation := evaluate(equations[i], operation)
			fmt.Print(equations[i], operation, evaluation, values[i], "\n")
			if evaluation == values[i] {
				answer += values[i]
				continue valuesLoop
			}
		}
	}
	fmt.Print(answer)
	// Part one answer: 1582598718861
}
