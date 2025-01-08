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

func nthPower(base, exponent int) int {
	if exponent == 0 {
		return 1
	}
	if exponent == 1 {
		return base
	}
	sqrtAnswer := nthPower(base, exponent/2)
	if exponent%2 == 0 {
		return sqrtAnswer * sqrtAnswer
	}
	return sqrtAnswer * sqrtAnswer * base
}

func nthBit(num, n int) bool {
	return (num>>n)&1 == 1
}

func permutations[T comparable](length int, symbols []T) [][]T {
	if length == 1 {
		baseCase := make([][]T, len(symbols))
		for i, symbol := range symbols {
			baseCase[i] = []T{symbol}
		}
		return baseCase
	}
	possibilities := make([][]T, nthPower(len(symbols), length))
	oneLessPossibilities := permutations(length-1, symbols)
	index := 0
	for _, symbol := range symbols {
		for _, shorter := range oneLessPossibilities {
			shorterCopy := make([]T, len(shorter))
			copy(shorterCopy, shorter)
			possibilities[index] = append(shorterCopy, symbol)
			index += 1
		}

	}
	return possibilities
}

type Operation int

const (
	Addition Operation = iota
	Multiplication
	Concat
)

var operations []Operation = []Operation{
	Addition,
	Multiplication,
	Concat,
}

func concat(num1, num2 int) int {
	num1Str := strconv.Itoa(num1)
	num2Str := strconv.Itoa(num2)
	returnVal, _ := strconv.Atoi(num1Str + num2Str)
	return returnVal
}

func operation(num1, num2 int, operation Operation) int {
	switch operation {
	case Addition:
		return num1 + num2
	case Multiplication:
		return num1 * num2
	case Concat:
		return concat(num1, num2)
	default:
		return -1
	}
}

func evaluate(equation []int, operations []Operation) int {
	answer := equation[0]
	for i := 1; i < len(equation); i++ {
		answer = operation(answer, equation[i], operations[i-1])
	}
	return answer
}

func part2() {
	setCwdToSourceFile()
	values, equations, err := parseInput("input")
	if err != nil {
		fmt.Print("Error parsing input: ", err)
		return
	}
	answer := 0
valuesLoop:
	for i := range values {
		allOperations := permutations(len(equations[i])-1, operations)
		for _, operation := range allOperations {
			evaluation := evaluate(equations[i], operation)
			if evaluation == values[i] {
				answer += values[i]
				continue valuesLoop
			}
		}
	}
	fmt.Print(answer, "\n")

	// Part one answer: 1582598718861
	// Part one answer: 165278151522644
}

func main() {
	part2()
}
