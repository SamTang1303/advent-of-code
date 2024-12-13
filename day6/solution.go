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

func parseInput(filePath string) ([]string, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return []string{""}, fmt.Errorf("Problem reading file")
	}
	inputStr := strings.TrimSpace(string(input))
	return strings.Split(inputStr, "\n"), nil
}

type Direction complex64

func (a Direction) x() int { return int(real(a)) }
func (a Direction) y() int { return int(imag(a)) }

func invalidSquare(position [2]int, room *[]string) bool {
	return position[0] > len((*room)[0])-1 || position[1] > len((*room))-1 || position[0] < 0 || position[1] < 0
}

func hash(position [2]int, direction Direction) [4]int {
	return [4]int{position[0], position[1], direction.x(), direction.y()}
}

type Set[K comparable] map[K]struct{}

var seen Set[[4]int] = make(map[[4]int]struct{})

func findLoop(position [2]int, direction Direction, room *[]string) bool {
	stateHash := hash(position, direction)
	if _, exists := seen[stateHash]; exists {
		return true
	}
	seen[stateHash] = struct{}{}
	newPos := [2]int{position[0] + direction.x(), position[1] + direction.y()}
	if invalidSquare(newPos, room) {
		return false
	}
	if (*room)[newPos[1]][newPos[0]] == '#' {
		direction = direction * complex(0, 1)
		return findLoop(position, direction, room)
	}
	return findLoop(newPos, direction, room)
}

func main() {
	setCwdToSourceFile()
	input, err := parseInput("input")
	if err != nil {
		fmt.Print("Error parsing input: ", err)
		return
	}
	startingPosition := [2]int{91, 69}
	answer := 0
	for y, line := range input {
		for x, character := range line {
			if character == '^' || character == '#' {
				continue
			}
			lineRunes := []rune(input[y])
			lineRunes[x] = '#'
			input[y] = string(lineRunes)
			if findLoop(startingPosition, Direction(complex(0, -1)), &input) {
				answer += 1
			}
			lineRunes[x] = '.'
			input[y] = string(lineRunes)
			seen = Set[[4]int]{}
		}
	}
	fmt.Print("\n\nAnswer: ", answer)
	//Part two answer: 1309
}
