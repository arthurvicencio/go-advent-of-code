package twenty_twenty_four_day_02

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	var safe int

	for _, line := range strings.Split(input2, "\n") {

		report := make([]int, 0)
		for _, char := range strings.Split(line, " ") {
			level, _ := strconv.Atoi(char)
			report = append(report, level)
		}

		for i := 0; i < len(report); i++ {
			if p.isSafe(p.intSliceRemove(report, i)) {
				safe++
				break
			}
		}
	}

	fmt.Println(safe)
}

func (p Part2) isSafe(report []int) bool {
	if !p.isIncreasing(report) && !p.isDecreasing(report) {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		isWithinRange := p.abs(report[i]-report[i+1]) >= 1 && p.abs(report[i]-report[i+1]) <= 3
		if !isWithinRange {
			return false
		}
	}

	return true
}

func (p Part2) abs(m int) int {
	if m < 0 {
		return -m
	}
	return m
}

func (p Part2) isIncreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}

func (p Part2) isDecreasing(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] <= nums[i+1] {
			return false
		}
	}
	return true
}

func (p Part2) intSliceRemove(s []int, i int) []int {
	n := make([]int, len(s))
	copy(n, s)
	return append(n[:i], n[i+1:]...)
}
