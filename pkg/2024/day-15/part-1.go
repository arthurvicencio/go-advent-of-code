package twenty_twenty_four_day_15

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

	rawInput := strings.Split(input1, "\n\n")

	pos := Point{}

	grid := make(map[Point]string)
	for y, line := range strings.Split(rawInput[0], "\n") {
		for x, char := range strings.Split(line, "") {
			grid[Point{x, y}] = char
			if grid[Point{x, y}] == "@" {
				pos = Point{x, y}
				grid[Point{x, y}] = "."
			}
		}
	}

	movements := make([][]string, 0)
	for _, line := range strings.Split(rawInput[1], "\n") {
		movements = append(movements, strings.Split(line, ""))
	}

	display := func() {
		for y := 0; y < 10; y++ {
			for x := 0; x < 20; x++ {
				if pos == (Point{x, y}) {
					fmt.Print("@")
				} else {
					fmt.Print(grid[Point{x, y}])
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}

	for _, ms := range movements {
		for _, move := range ms {
			if move == "^" {
				next := Point{pos.X, pos.Y - 1}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "O" {
					objs := make([]Point, 0)
					objPos := next
					for {
						if grid[objPos] == "" || grid[objPos] == "#" {
							break
						}

						if grid[objPos] == "." {
							for i := len(objs) - 1; i >= 0; i-- {
								grid[objs[i]] = "."
								grid[Point{objs[i].X, objs[i].Y - 1}] = "O"
							}
							pos = next
							break
						}
						if grid[objPos] == "O" {
							objs = append(objs, objPos)
						}

						objPos = Point{objPos.X, objPos.Y - 1}
					}
				}
			} else if move == ">" {
				next := Point{pos.X + 1, pos.Y}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "O" {
					objs := make([]Point, 0)
					objPos := next
					for {
						if grid[objPos] == "" || grid[objPos] == "#" {
							break
						}

						if grid[objPos] == "." {
							for i := len(objs) - 1; i >= 0; i-- {
								grid[Point{objs[i].X + 1, objs[i].Y}] = grid[objs[i]]
								grid[objs[i]] = "."
							}
							pos = next
							break
						}
						if grid[objPos] == "O" {
							objs = append(objs, objPos)
						}

						objPos = Point{objPos.X + 1, objPos.Y}
					}
				}
			} else if move == "v" {
				next := Point{pos.X, pos.Y + 1}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "O" {
					objs := make([]Point, 0)
					objs = append(objs, next)
					objPos := next
					for {
						if grid[objPos] == "" || grid[objPos] == "#" {
							break
						}

						if grid[objPos] == "." {
							for i := len(objs) - 1; i >= 0; i-- {
								grid[objs[i]] = "."
								grid[Point{objs[i].X, objs[i].Y + 1}] = "O"
							}
							pos = next
							break
						}
						if grid[objPos] == "O" {
							objs = append(objs, objPos)
						}

						objPos = Point{objPos.X, objPos.Y + 1}
					}
				}
			} else if move == "<" {
				next := Point{pos.X - 1, pos.Y}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "O" {
					objs := make([]Point, 0)
					objPos := next
					for {
						if grid[objPos] == "" || grid[objPos] == "#" {
							break
						}

						if grid[objPos] == "." {
							for i := len(objs) - 1; i >= 0; i-- {
								grid[Point{objs[i].X - 1, objs[i].Y}] = grid[objs[i]]
								grid[objs[i]] = "."
							}
							pos = next
							break
						}
						if grid[objPos] == "O" {
							objs = append(objs, objPos)
						}

						objPos = Point{objPos.X - 1, objPos.Y}
					}
				}
			}
		}
	}
	display()

	var ans int
	for p, v := range grid {
		if v == "O" {
			ans += 100*p.Y + p.X
		}
	}

	fmt.Println(ans)
}
