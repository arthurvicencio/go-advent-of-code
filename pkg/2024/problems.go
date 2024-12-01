package twenty_twenty_four

import (
	twenty_twenty_four_day_01 "github.com/arthurvicencio/go-advent-of-code/pkg/2024/day-01"
	"github.com/arthurvicencio/go-advent-of-code/pkg/aocutils"
)

func Problems() map[string]aocutils.Problem {
	return map[string]aocutils.Problem{
		"2024-01-1": twenty_twenty_four_day_01.Part1{},
		"2024-01-2": twenty_twenty_four_day_01.Part2{},
	}
}
