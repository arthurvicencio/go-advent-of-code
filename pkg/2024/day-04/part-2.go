package twenty_twenty_four_day_04

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {

	type Point struct{ X, Y int }

	const TARGET = "MAS"

	grid := make([][]string, 0)

	for _, line := range strings.Split(input2, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	var answer int

	foundLocations := make([][]Point, 0)

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			directions := []Point{
				{-1, -1}, // up, left
				{1, -1},  // up, right
				{1, 1},   // down, right
				{-1, 1},  // down, left
			}

			for _, direction := range directions {

				str := grid[y][x]

				var dy, dx = y, x

				locations := []Point{{dy, dx}}

				for i := 0; i < 2; i++ {
					dy = dy + direction.Y
					if dy < 0 || dy > len(grid)-1 {
						continue
					}

					dx = dx + direction.X
					if dx < 0 || dx > len(grid[0])-1 {
						continue
					}

					str += grid[dy][dx]
					locations = append(locations, Point{dy, dx})
				}

				if str == TARGET {
					foundLocations = append(foundLocations, locations)
				}
			}
		}
	}

	for i := 0; i < len(foundLocations)-1; i++ {
		for j := i + 1; j < len(foundLocations); j++ {
			if foundLocations[i][1] == foundLocations[j][1] {
				answer++
			}
		}
	}

	fmt.Println(answer)
}
