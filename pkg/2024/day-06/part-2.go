package twenty_twenty_four_day_06

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input2 string

type Point_P2 struct{ Y, X int }

type State_P2 struct {
	Pos Point_P2
	Dir int
}

type Part2 struct{}

func (p Part2) Solve() {

	var start Point_P2

	grid := make([][]string, 0)

	for y, line := range strings.Split(input2, "\n") {
		row := make([]string, 0)

		for x, char := range strings.Split(line, "") {
			row = append(row, char)
			if char == "^" {
				start = Point_P2{y, x}
			}
		}

		grid = append(grid, row)
	}

	var loopCount int

	// go throgh every possible combination of an added obstruction
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == "#" {
				continue
			}

			grid[y][x] = "#"

			if p.isLoop(grid, start) {
				loopCount++
			}

			grid[y][x] = "."
		}
	}

	fmt.Println(loopCount)
}

func (p Part2) isLoop(grid [][]string, start Point_P2) bool {
	seen := make(map[State_P2]bool)

	var y, x = start.Y, start.X
	var dir int
	for {
		if !p.isValidCell(grid, y, x) {
			return false
		}

		state := State_P2{
			Point_P2{y, x},
			dir,
		}
		if seen[state] {
			return true
		}
		seen[state] = true

		directions := []Point_P2{
			{-1, 0},
			{0, 1},
			{1, 0},
			{0, -1},
		}
		for i, d := range directions {
			if i != dir {
				continue
			}
			var dy, dx = y + d.Y, x + d.X
			if p.isValidCell(grid, dy, dx) && grid[dy][dx] == "#" {
				dir = (dir + 1) % 4
			} else {
				y, x = dy, dx
			}
		}
	}
}

func (p Part2) isValidCell(grid [][]string, y, x int) bool {
	return y >= 0 && y <= len(grid)-1 && x >= 0 && x <= len(grid[0])-1
}
