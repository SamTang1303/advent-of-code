package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Point [2]int
type Vec [2]int
type Nil struct{}

var null Nil = Nil{}

func parseInput(filePath string) ([]string, error) {
	input, err := os.ReadFile(filePath)
	if err != nil {
		return []string{""}, fmt.Errorf("Problem reading file")
	}
	inputStr := strings.TrimSpace(string(input))
	lines := strings.Split(inputStr, "\n")
	return lines, nil
}

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

func (self Vec) neg() Vec {
	return Vec{
		-1 * self[0],
		-1 * self[1],
	}
}

func (self Point) move(vector Vec) Point {
	return Point{
		self[0] + vector[0],
		self[1] + vector[1],
	}
}

func diff(point1, point2 Point) Vec {
	return Vec{
		point2[0] - point1[0],
		point2[1] - point1[1],
	}
}

func inBounds(point Point, input *[]string) bool {
	if point[0] < 0 || point[0] >= len((*input)[0]) {
		return false
	}
	if point[1] < 0 || point[1] >= len((*input)) {
		return false
	}
	return true
}

func findAntinodes(input *[]string, antenna1, antenna2 Point) []Point {
	antinodes := make([]Point, 0, 5)
	diffVec := diff(antenna1, antenna2)
	antinode1 := antenna1
	antinode2 := antenna2
	for inBounds(antinode1, input) {
		antinodes = append(antinodes, antinode1)
		antinode1 = antinode1.move(diffVec.neg())
	}
	for inBounds(antinode2, input) {
		antinodes = append(antinodes, antinode2)
		antinode2 = antinode2.move(diffVec)
	}
	return antinodes
}

func evalPoint(point Point, input *[]string, antennasMap *map[byte][]Point, antinodes *map[Point]struct{}) {
	char := (*input)[point[1]][point[0]]
	if char == '.' {
		return
	}
	antennas, exits := (*antennasMap)[char]
	if !exits {
		(*antennasMap)[char] = []Point{point}
		return
	}
	for _, antenna := range antennas {
		charAntinodes := findAntinodes(input, point, antenna)
		// fmt.Println(point, antenna, charAntinodes)
		for _, antinode := range charAntinodes {
			(*antinodes)[antinode] = struct{}{}
			// fmt.Println()
		}
	}
	(*antennasMap)[char] = append(antennas, point)
}

func part1() {
	setCwdToSourceFile()
	input, err := parseInput("../input")
	if err != nil {
		fmt.Print("Error parsing input: ", err)
		return
	}
	antennas := make(map[byte][]Point)
	antinodes := make(map[Point]struct{})
	for y, line := range input {
		for x := range line {
			evalPoint(Point{x, y}, &input, &antennas, &antinodes)
		}
	}
	fmt.Println(len(antinodes))
}

func main() {
	part1()
}
