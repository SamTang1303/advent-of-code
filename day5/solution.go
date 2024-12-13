package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
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

type Set map[int]struct{}

type Update struct {
	pages           []int
	comparisonPairs *map[int]Set
	wasIncorrect    *bool
}

func comparePages(page1, page2 int, comparisonPairs *map[int]Set) bool {
	if _, exists := (*comparisonPairs)[page1][page2]; exists {
		return true
	}
	if _, exists := (*comparisonPairs)[page2][page1]; exists {
		return false
	}
	return true
}

func (a Update) Len() int { return len(a.pages) }
func (a Update) Swap(i, j int) {
	*a.wasIncorrect = true
	a.pages[i], a.pages[j] = a.pages[j], a.pages[i]
}
func (a Update) Less(i, j int) bool { return comparePages(a.pages[i], a.pages[j], a.comparisonPairs) } // Change to '>' for descending

func main() {
	setCwdToSourceFile()
	rules, updates, err := parseInput("input")
	if err != nil {
		return
	}
	answer := 0
	comparisonPairs := make(map[int]Set)
	for _, rule := range rules {
		illegalFollowsSet, exists := comparisonPairs[rule[0]]
		if !exists {
			comparisonPairs[rule[0]] = Set{rule[1]: struct{}{}}
			continue
		}
		illegalFollowsSet[rule[1]] = struct{}{}
	}
	for i := range updates {
		boolRef := false
		update := Update{
			pages:           updates[i],
			comparisonPairs: &comparisonPairs,
			wasIncorrect:    &boolRef,
		}
		sort.Sort(update)
		if !*update.wasIncorrect {
			continue
		}
		middleIndex := len(update.pages) / 2
		answer += update.pages[middleIndex]
	}
	fmt.Print(answer)
}
