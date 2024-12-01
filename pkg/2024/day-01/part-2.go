package twenty_twenty_four_day_01

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	left := make([]int, 0)
	right := make(map[int]int)

	for _, line := range strings.Split(input2, "\n") {
		var a, b int

		fmt.Sscanf(line, "%d   %d", &a, &b)

		left = append(left, a)
		right[b]++
	}

	var answer int

	for _, n := range left {
		answer += n * right[n]
	}

	fmt.Println(answer)
}
