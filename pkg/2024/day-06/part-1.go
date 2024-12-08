package twenty_twenty_four_day_06

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input1 string

type Point_P1 struct{ X, Y int }

type Part1 struct{}

func (p Part1) Solve() {

	var start Point_P1

	grid := make([][]string, 0)

	for y, line := range strings.Split(input1, "\n") {
		row := make([]string, 0)

		for x, char := range strings.Split(line, "") {
			row = append(row, char)
			if char == "^" {
				start = Point_P1{x, y}
			}
		}

		grid = append(grid, row)
	}

	track := make(map[Point_P1]bool)

	var x, y = start.X, start.Y
	var dir int
	for {
		if !p.isValidCell(grid, x, y) {
			break
		}

		track[Point_P1{x, y}] = true

		directions := []Point_P1{
			{0, -1},
			{1, 0},
			{0, 1},
			{-1, 0},
		}
		for i, d := range directions {
			if i != dir {
				continue
			}
			var dx, dy = x + d.X, y + d.Y
			if p.isValidCell(grid, dx, dy) && grid[dy][dx] == "#" {
				dir = (dir + 1) % 4
			} else {
				x, y = dx, dy
			}
		}
	}

	fmt.Println(len(track))
}

func (p Part1) isValidCell(grid [][]string, x, y int) bool {
	return x >= 0 && x <= len(grid[0])-1 && y >= 0 && y <= len(grid)-1
}
