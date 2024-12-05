package twenty_twenty_four

import (
	twenty_twenty_four_day_01 "github.com/arthurvicencio/go-advent-of-code/pkg/2024/day-01"
	twenty_twenty_four_day_02 "github.com/arthurvicencio/go-advent-of-code/pkg/2024/day-02"
	twenty_twenty_four_day_03 "github.com/arthurvicencio/go-advent-of-code/pkg/2024/day-03"
	twenty_twenty_four_day_04 "github.com/arthurvicencio/go-advent-of-code/pkg/2024/day-04"
	twenty_twenty_four_day_05 "github.com/arthurvicencio/go-advent-of-code/pkg/2024/day-05"
	"github.com/arthurvicencio/go-advent-of-code/pkg/aocutils"
)

func Problems() map[string]aocutils.Problem {
	return map[string]aocutils.Problem{
		"2024-d01-p1": twenty_twenty_four_day_01.Part1{},
		"2024-d01-p2": twenty_twenty_four_day_01.Part2{},

		"2024-d02-p1": twenty_twenty_four_day_02.Part1{},
		"2024-d02-p2": twenty_twenty_four_day_02.Part2{},

		"2024-d03-p1": twenty_twenty_four_day_03.Part1{},
		"2024-d03-p2": twenty_twenty_four_day_03.Part2{},

		"2024-d04-p1": twenty_twenty_four_day_04.Part1{},
		"2024-d04-p2": twenty_twenty_four_day_04.Part2{},

		"2024-d05-p1": twenty_twenty_four_day_05.Part1{},
		"2024-d05-p2": twenty_twenty_four_day_05.Part2{},
	}
}
