package twenty_twenty_four_day_10

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {

	type Point struct{ X, Y int }

	const (
		LOWEST_POINT  = 0
		HIGHEST_POINT = 9
	)

	startingPositions := make([]Point, 0)

	grid := make([][]int, 0)
	for y, line := range strings.Split(input2, "\n") {

		row := make([]int, 0)
		for x, char := range strings.Split(line, "") {
			n, _ := strconv.Atoi(char)

			if n == LOWEST_POINT {
				startingPositions = append(startingPositions, Point{x, y})
			}

			row = append(row, n)
		}

		grid = append(grid, row)
	}

	queue := startingPositions

	var score int

	for len(queue) > 0 {

		current := queue[0]
		queue = queue[1:]

		x, y := current.X, current.Y

		if grid[y][x] == HIGHEST_POINT {
			score++
			continue
		}

		directions := []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}
		for _, d := range directions {

			dx, dy := x+d.X, y+d.Y

			isPointValid := dx >= 0 && dx <= len(grid[0])-1 && dy >= 0 && dy <= len(grid)-1
			if !isPointValid {
				continue
			}

			if grid[dy][dx] == grid[y][x]+1 {
				queue = append(queue, Point{dx, dy})
			}
		}
	}

	fmt.Println(score)
}
