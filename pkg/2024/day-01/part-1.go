package twenty_twenty_four_day_01

import (
	_ "embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {
	left := make([]int, 0)
	right := make([]int, 0)

	for _, line := range strings.Split(input1, "\n") {
		var a, b int

		fmt.Sscanf(line, "%d   %d", &a, &b)

		left = append(left, a)
		right = append(right, b)
	}

	sort.Ints(left)
	sort.Ints(right)

	var answer int

	for i, n := range left {
		answer += p.abs(n - right[i])
	}

	fmt.Println(answer)
}

func (p Part1) abs(m int) int {
	if m < 0 {
		return -m
	}
	return m
}
