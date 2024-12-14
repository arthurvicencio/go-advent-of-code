package twenty_twenty_four_day_14

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {

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
		ROW_LENGTH = 101
		COL_LENGTH = 103
	)

	robots := make([]Robot, 0)
	robotLocator := make(map[Point]int)
	for _, line := range strings.Split(input2, "\n") {
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

	display := func() {
		for y := 0; y < COL_LENGTH; y++ {
			for x := 0; x < ROW_LENGTH; x++ {
				if robotLocator[Point{x, y}] > 0 {
					fmt.Print(robotLocator[Point{x, y}])
				} else {
					fmt.Print(".")
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	for i := 0; ; i++ {
		simulate()

		allOccupiedByOneRobot := true
		for _, p := range robotLocator {
			if p > 1 {
				allOccupiedByOneRobot = false
				break
			}
		}
		if allOccupiedByOneRobot {
			display()
			fmt.Println(i + 1)
			break
		}
	}
}
