package main

import (
	"fmt"
	"os"
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
}
