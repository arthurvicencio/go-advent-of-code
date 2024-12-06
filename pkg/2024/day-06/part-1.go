package twenty_twenty_four_day_06

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input1 string

type Point_P1 struct{ Y, X int }

type Part1 struct{}

func (p Part1) Solve() {

	var start Point_P1

	grid := make([][]string, 0)

	for y, line := range strings.Split(input1, "\n") {
		row := make([]string, 0)

		for x, char := range strings.Split(line, "") {
			row = append(row, char)
			if char == "^" {
				start = Point_P1{y, x}
			}
		}

		grid = append(grid, row)
	}

	track := make(map[Point_P1]bool)

	var y, x = start.Y, start.X
	var dir int
	for {
		if !p.isValidCell(grid, y, x) {
			break
		}

		track[Point_P1{y, x}] = true

		directions := []Point_P1{
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

	fmt.Println(len(track))
}

func (p Part1) isValidCell(grid [][]string, y, x int) bool {
	return y >= 0 && y <= len(grid)-1 && x >= 0 && x <= len(grid[0])-1
}
