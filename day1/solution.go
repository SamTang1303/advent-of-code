package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("/Users/samtang/Programming/advent-of-code-20204/day1/input.txt")
	if err != nil {
		fmt.Print("Error!!", err)
		return
	}
	input_str := string(input)
	lines := strings.Split(input_str, "\n")
	col1 := make([]int, len(lines)-1)
	col2 := make([]int, len(lines)-1)
	for index, val := range lines {
		nums := strings.Fields(val)
		if len(nums) == 0 {
			continue
		}
		num1, bad1 := strconv.Atoi(nums[0])
		num2, bad2 := strconv.Atoi(nums[1])
		if bad1 != nil || bad2 != nil {
			fmt.Print("\n-----------------------")
			fmt.Print("Bad!!!!!!!!!!!")
			return
		}
		col1[index] = num1
		col2[index] = num2
	}
	sort.Ints(col1)
	sort.Ints(col2)
	answer := 0
	for i := 0; i < len(col1); i++ {
		absDiff := col1[i] - col2[i]
		if absDiff < 0 {
			absDiff *= -1
		}
		answer += absDiff
	}
	// fmt.Print(answer)
	// Part one answer: 2285373

	// Part two:
	var col1_occurances = make(map[int]int)
	var col2_occurances = make(map[int]int)
	for i := 0; i < len(col1); i++ {
		val1 := col1[i]
		_, ok1 := col1_occurances[col1[i]]
		if ok1 {
			col1_occurances[val1] += 1
		} else {
			col1_occurances[val1] = 1
		}
		val2 := col2[i]
		_, ok2 := col2_occurances[col2[i]]
		if ok2 {
			col2_occurances[val2] += 1
		} else {
			col2_occurances[val2] = 1
		}
	}
	result := 0
	for key := range col1_occurances {
		result += key * col1_occurances[key] * col2_occurances[key]
	}
	fmt.Print(result)
	// Part two answer: 21142653
}
