package main

import (
	"fmt"
	"os"
	"strings"
)

func parseInput(filePath string) []string {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return []string{}
	}
	lines := strings.Split(string(input), "\n")
	return lines[:len(lines)-1]
}

var xmasPosition = map[byte]int{
	'X': 0,
	'M': 1,
	'A': 2,
	'S': 3,
}

var directions = [8][2]int{
	{1, 1},
	{0, 1},
	{1, 0},
	{-1, 0},
	{0, -1},
	{-1, 1},
	{1, -1},
	{-1, -1},
}

func validSubstr(direction [2]int, position [2]int, puzzle *[]string) bool {
	letter := (*puzzle)[position[1]][position[0]]
	xOutOfBounds := position[0]+direction[0] > len((*puzzle)[0])-1 || position[0]+direction[0] < 0
	yOutOfBounds := position[1]+direction[1] > len(*puzzle)-1 || position[1]+direction[1] < 0

	if xOutOfBounds || yOutOfBounds {
		return false
	}
	nextPuzzleLetter := (*puzzle)[position[1]+direction[1]][position[0]+direction[0]]
	nextXmasLetter := "XMAS"[xmasPosition[letter]+1]
	validNextLetter := nextPuzzleLetter == nextXmasLetter
	if !validNextLetter {
		return false
	}

	if nextPuzzleLetter == 'S' {
		return true
	}
	return validSubstr(direction, [2]int{position[0] + direction[0], position[1] + direction[1]}, puzzle)
}

func validXmas(position [2]int, puzzle *[]string) bool {
	validX := position[0] > 0 && position[0] < len((*puzzle)[0])-1
	validY := position[1] > 0 && position[1] < len(*puzzle)-1
	if !validX || !validY {
		return false
	}
	topLeft := (*puzzle)[position[1]-1][position[0]-1]
	topRight := (*puzzle)[position[1]-1][position[0]+1]
	bottomRight := (*puzzle)[position[1]+1][position[0]+1]
	bottomLeft := (*puzzle)[position[1]+1][position[0]-1]

	if !((bottomLeft == 'M' && topRight == 'S') || (bottomLeft == 'S' && topRight == 'M')) {
		return false
	}

	if !((topLeft == 'M' && bottomRight == 'S') || (topLeft == 'S' && bottomRight == 'M')) {
		return false
	}
	print(string(topLeft), " . ", string(topRight), "\n")
	print(". ", string((*puzzle)[position[1]][position[0]]), " .", "\n")
	print(string(bottomLeft), " . ", string(bottomRight), "\n")

	return true
}

func main() {
	puzzle := parseInput("/Users/samtang/Programming/advent-of-code-20204/day4/input.txt")
	answer := 0
	for y, line := range puzzle {
		for x, character := range line {
			if character == 'A' && validXmas([2]int{x, y}, &puzzle) {
				answer += 1
			}
		}
	}
	fmt.Print(answer)
}
