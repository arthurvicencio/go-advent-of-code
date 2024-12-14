package twenty_twenty_four_day_14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {

	type Point struct {
		X, Y int
	}

	type Velocity struct {
		X, Y int
	}

	type Robot struct {
		Pos      Point
		Velocity Velocity
	}

	const (
		TOTAL_SECONDS = 100
		ROW_LENGTH    = 101
		COL_LENGTH    = 103
	)

	robots := make([]Robot, 0)
	robotLocator := make(map[Point]int)
	for _, line := range strings.Split(input1, "\n") {
		var point Point
		var velocity Velocity
		fmt.Sscanf(line, "p=%d,%d v=%d,%d", &point.X, &point.Y, &velocity.X, &velocity.Y)

		robots = append(robots, Robot{point, velocity})
		robotLocator[point] = 1
	}

	simulate := func() {
		for i, r := range robots {
			if robotLocator[robots[i].Pos] > 0 {
				robotLocator[robots[i].Pos]--
			}

			robots[i].Pos.X = (r.Pos.X + r.Velocity.X) % ROW_LENGTH
			if robots[i].Pos.X < 0 {
				robots[i].Pos.X = (robots[i].Pos.X + ROW_LENGTH) % ROW_LENGTH
			}

			robots[i].Pos.Y = (r.Pos.Y + r.Velocity.Y) % COL_LENGTH
			if robots[i].Pos.Y < 0 {
				robots[i].Pos.Y = (robots[i].Pos.Y + COL_LENGTH) % COL_LENGTH
			}

			robotLocator[robots[i].Pos]++
		}
	}

	for seconds := 0; seconds < TOTAL_SECONDS; seconds++ {
		simulate()
	}

	type QuadrantRange struct {
		FromY int
		ToY   int
		FromX int
		ToX   int
	}

	quadrantRanges := []QuadrantRange{
		{
			FromY: 0,
			ToY:   COL_LENGTH / 2,
			FromX: 0,
			ToX:   ROW_LENGTH / 2,
		},
		{
			FromY: 0,
			ToY:   COL_LENGTH / 2,
			FromX: (ROW_LENGTH / 2) + 1,
			ToX:   ROW_LENGTH,
		},
		{
			FromY: (COL_LENGTH / 2) + 1,
			ToY:   COL_LENGTH,
			FromX: 0,
			ToX:   ROW_LENGTH / 2,
		},
		{
			FromY: (COL_LENGTH / 2) + 1,
			ToY:   COL_LENGTH,
			FromX: (ROW_LENGTH / 2) + 1,
			ToX:   ROW_LENGTH,
		},
	}

	quadrants := make([]int, 4)
	for i, quadrantRange := range quadrantRanges {
		for y := quadrantRange.FromY; y < quadrantRange.ToY; y++ {
			for x := quadrantRange.FromX; x < quadrantRange.ToX; x++ {
				quadrants[i] += robotLocator[Point{x, y}]
			}
		}
	}

	fmt.Println(quadrants[0] * quadrants[1] * quadrants[2] * quadrants[3])
}
