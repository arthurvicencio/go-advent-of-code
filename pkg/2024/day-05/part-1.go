package twenty_twenty_four_day_05

import (
	_ "embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {
	rules := make(map[int]map[int]bool)
	updates := make([][]int, 0)

	rawInput := strings.Split(input1, "\n\n")

	for _, line := range strings.Split(rawInput[0], "\n") {
		var a, b int
		fmt.Sscanf(line, "%d|%d", &a, &b)
		if _, ok := rules[b]; !ok {
			rules[b] = make(map[int]bool)
		}
		rules[b][a] = true
	}

	for _, line := range strings.Split(rawInput[1], "\n") {
		update := make([]int, 0)
		for _, char := range strings.Split(line, ",") {
			n, _ := strconv.Atoi(char)
			update = append(update, n)
		}
		updates = append(updates, update)
	}

	var answer int
	for _, update := range updates {
		sorted := p.sort(rules, update)
		if p.isSorted(sorted, update) {
			answer += sorted[len(update)/2]
		}
	}
	fmt.Println(answer)
}

func (p Part1) sort(rules map[int]map[int]bool, slice []int) []int {
	numPriorities := make(map[int]int)
	for _, num := range slice {
		numPriorities[num] = 0
	}

	for _, num := range slice {
		for numAfter, numsBefore := range rules {
			if _, ok := numsBefore[num]; ok {
				numPriorities[numAfter]++
			}
		}
	}

	sorted := make([]int, len(slice))
	copy(sorted, slice)

	sort.Slice(sorted, func(i, j int) bool {
		return numPriorities[sorted[i]] < numPriorities[sorted[j]]
	})

	return sorted
}

func (p Part1) isSorted(a, b []int) bool {
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
