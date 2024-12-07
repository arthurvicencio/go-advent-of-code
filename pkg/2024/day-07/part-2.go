package twenty_twenty_four_day_07

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	type Calibration struct {
		Target int
		Args   []int
	}

	calibrations := make([]Calibration, 0)
	for _, line := range strings.Split(input2, "\n") {

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

func (p Part2) findTarget(target int, args []int) bool {
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
			State{index + 1, p.concatInt(current.Value, args[index])},
		)
	}

	return false
}

func (p Part2) concatInt(a, b int) int {
	// Find the number of digits in b
	digits := int(math.Log10(float64(b))) + 1
	// Multiply a by 10^digits and add b
	return a*int(math.Pow(10, float64(digits))) + b
}
