package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func parseInput(filePath string) ([][]int, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return [][]int{}, fmt.Errorf("Problem reading file")
	}
	inputStr := string(input)
	lines := strings.Split(inputStr, "\n")
	var data [][]int
	for _, val := range lines {
		numsStr := strings.Fields(val)
		nums_int := make([]int, len(numsStr))
		if len(numsStr) == 0 {
			continue
		}
		for i := range numsStr {
			nums_int[i], err = strconv.Atoi(numsStr[i])
			if err != nil {
				fmt.Print("Error converting string")
				break
			}
		}
		data = append(data, nums_int)
	}
	return data, nil
}

func main() {
	data, err := parseInput("/Users/samtang/Programming/advent-of-code-20204/day2/input.txt")
	if err != nil {
		return
	}
	fmt.Print(data)
}
