package twenty_twenty_four_day_07

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {
	type Calibration struct {
		Target int
		Args   []int
	}

	calibrations := make([]Calibration, 0)
	for _, line := range strings.Split(input1, "\n") {

		parts := strings.Split(line, ": ")

		target, _ := strconv.Atoi(parts[0])

		args := make([]int, 0)
		for _, num := range strings.Split(parts[1], " ") {
			n, _ := strconv.Atoi(num)

			args = append(args, n)
		}

		calibrations = append(calibrations, Calibration{target, args})
	}

	var ans int
	for _, calib := range calibrations {

		if p.findTarget(calib.Target, calib.Args) {
			ans += calib.Target
		}
	}

	fmt.Println(ans)
}

func (p Part1) findTarget(target int, args []int) bool {
	type State struct {
		Index int
		Value int
	}

	stack := []State{{}}

	for len(stack) > 0 {
		current := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if current.Value == target {
			return true
		}

		if current.Index >= len(args) {
			continue
		}

		index := current.Index
		stack = append(
			stack,
			State{index + 1, current.Value + args[index]},
			State{index + 1, current.Value * args[index]},
		)
	}

	return false
}
