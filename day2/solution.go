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

func determineSafety(report []int) (bool, int) {
	type Direction int
	const (
		Increasing Direction = iota
		Decreasing
	)
	diff := report[0] - report[1]
	for i := 0; i < len(report)-1; i++ {
		report_diff := report[i] - report[i+1]
		if report_diff == 0 || report_diff > 3 || report_diff < -3 {
			return false, i + 1
		}
		if (diff > 0) != (report_diff > 0) {
			return false, i + 1
		}
	}
	return true, 0
}

func removeIndex(slice []int, index int) []int {
	// copySlice := make([]int, len(slice))
	// copy(copySlice, slice)
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	data, err := parseInput("/Users/samtang/Programming/advent-of-code-20204/day2/input.txt")
	if err != nil {
		return
	}
	answer := 0
	for _, report := range data {
		initiallySafe, problemIndex := determineSafety(report)
		if initiallySafe {
			answer += 1
			continue
		}
		updatedReport := removeIndex(report, problemIndex)
		safe, _ := determineSafety(updatedReport)
		updatedReport2 := removeIndex(report, problemIndex-1)
		safe2, _ := determineSafety(updatedReport2)
		safe3 := false
		if problemIndex-2 >= 0 {
			updatedReport3 := removeIndex(report, problemIndex-2)
			safe3, _ = determineSafety(updatedReport3)
		}
		if safe || safe2 || safe3 {
			answer += 1
			continue
		}
	}
	fmt.Print(answer)
}
