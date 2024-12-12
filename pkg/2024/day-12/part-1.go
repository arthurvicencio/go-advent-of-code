package twenty_twenty_four_day_12

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

	grid := make([][]string, 0)
	for _, line := range strings.Split(input1, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	seen := make(map[Point]bool)

	var price int

	for y := range grid {
		for x := range grid[y] {

			if seen[Point{x, y}] {
				continue
			}

			queue := []Point{{x, y}}

			var area, perimeter int

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				if seen[current] {
					continue
				}
				seen[current] = true

				var intersections int

				dirs := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
				for _, d := range dirs {

					dx, dy := current.X+d.X, current.Y+d.Y
					isInBoundsOfGrid := dx >= 0 && dx <= len(grid[y])-1 && dy >= 0 && dy <= len(grid)-1
					if !isInBoundsOfGrid {
						continue
					}

					if grid[y][x] == grid[dy][dx] {
						intersections++
						queue = append(queue, Point{dx, dy})
					}
				}

				area++
				perimeter += 4 - intersections
			}

			price += area * perimeter
		}
	}

	fmt.Println(price)
}
