package twenty_twenty_four_day_04

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {

	type Point struct{ X, Y int }

	target := "XMAS"

	grid := make([][]string, 0)

	for _, line := range strings.Split(input1, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	var answer int

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			directions := []Point{
				{0, -1},  // up
				{1, 0},   // right
				{0, 1},   // down
				{-1, 0},  // left
				{-1, -1}, // up, left
				{1, -1},  // up, right
				{1, 1},   // down, right
				{-1, 1},  // down, left
			}

			for _, direction := range directions {

				str := grid[y][x]

				var dy, dx = y, x

				for i := 0; i < 3; i++ {
					dy = dy + direction.Y
					if dy < 0 || dy > len(grid)-1 {
						continue
					}

					dx = dx + direction.X
					if dx < 0 || dx > len(grid[0])-1 {
						continue
					}

					str += grid[dy][dx]
				}

				if str == target {
					answer++
				}
			}
		}
	}

	fmt.Println(answer)
}
