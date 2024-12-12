package twenty_twenty_four_day_12

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

	grid := make([][]string, 0)
	for _, line := range strings.Split(input2, "\n") {
		grid = append(grid, strings.Split(line, ""))
	}

	var price int

	seen := make(map[Point]bool)

	for y := range grid {
		for x := range grid[y] {

			if seen[Point{x, y}] {
				continue
			}

			queue := []Point{{x, y}}

			region := make(map[Point]bool)

			for len(queue) > 0 {
				current := queue[0]
				queue = queue[1:]

				if seen[current] {
					continue
				}
				seen[current] = true
				region[current] = true

				dirs := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
				for _, d := range dirs {

					dx, dy := current.X+d.X, current.Y+d.Y
					isInBoundsOfGrid := dx >= 0 && dx <= len(grid[y])-1 && dy >= 0 && dy <= len(grid)-1
					if !isInBoundsOfGrid {
						continue
					}

					if grid[y][x] == grid[dy][dx] {
						queue = append(queue, Point{dx, dy})
					}
				}
			}

			var side int
			for loc := range region {

				up := Point{loc.X, loc.Y - 1}
				right := Point{loc.X + 1, loc.Y}
				down := Point{loc.X, loc.Y + 1}
				left := Point{loc.X - 1, loc.Y}

				upLeft := Point{loc.X - 1, loc.Y - 1}
				upRight := Point{loc.X + 1, loc.Y - 1}
				downLeft := Point{loc.X - 1, loc.Y + 1}
				downRight := Point{loc.X + 1, loc.Y + 1}

				if !region[up] && !region[left] {
					side++
				}
				if !region[up] && !region[right] {
					side++
				}
				if !region[down] && !region[left] {
					side++
				}
				if !region[down] && !region[right] {
					side++
				}
				if region[up] && region[left] && !region[upLeft] {
					side++
				}
				if region[up] && region[right] && !region[upRight] {
					side++
				}
				if region[down] && region[left] && !region[downLeft] {
					side++
				}
				if region[down] && region[right] && !region[downRight] {
					side++
				}
			}

			area := len(region)
			price += area * side
		}
	}

	fmt.Println(price)
}
