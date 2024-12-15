package twenty_twenty_four_day_15

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

	rawInput := strings.Split(input2, "\n\n")

	pos := Point{}

	raw := make([][]string, 0)
	for _, line := range strings.Split(rawInput[0], "\n") {
		row := make([]string, 0)
		for _, char := range strings.Split(line, "") {
			if char == "#" {
				row = append(row, "#", "#")
			} else if char == "O" {
				row = append(row, "[", "]")

			} else if char == "." {
				row = append(row, ".", ".")

			} else if char == "@" {
				row = append(row, "@", ".")
			}
		}
		raw = append(raw, row)
	}

	grid := make(map[Point]string)
	for y := range raw {
		for x := range raw[y] {
			grid[Point{x, y}] = raw[y][x]
			if raw[y][x] == "@" {
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

	type State struct {
		Left, Right Point
	}

	for _, ms := range movements {
		for _, move := range ms {
			if move == "^" {
				next := Point{pos.X, pos.Y - 1}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "[" || grid[next] == "]" {
					queue := make([]State, 0)

					objs := make([]Point, 0)

					if grid[next] == "[" {
						queue = append(queue, State{next, Point{next.X + 1, next.Y}})
						// objs = append(objs, next, Point{next.X + 1, next.Y})
					}
					if grid[next] == "]" {
						queue = append(queue, State{Point{next.X - 1, next.Y}, next})
						// objs = append(objs, Point{next.X - 1, next.Y}, next)
					}

					allGood := true
					seen := make(map[State]bool)

					for len(queue) > 0 {

						current := queue[0]
						queue = queue[1:]

						if seen[current] {
							continue
						}
						seen[current] = true

						left, right := current.Left, current.Right

						if grid[left] == "[" && grid[right] == "]" {
							objs = append(objs, left, right)
						}

						nextMiddle := State{Point{left.X, left.Y - 1}, Point{right.X, right.Y - 1}}

						if grid[nextMiddle.Left] == "[" && grid[nextMiddle.Right] == "]" {
							queue = append(queue, nextMiddle)
							continue
						} else if grid[nextMiddle.Left] == "#" && grid[nextMiddle.Right] == "#" {
							allGood = false
							break
						}

						nextleft := State{Point{left.X - 1, left.Y - 1}, Point{left.X, left.Y - 1}}
						if grid[nextleft.Left] == "[" && grid[nextleft.Right] == "]" {
							queue = append(queue, nextleft)
						} else if grid[nextleft.Left] == "#" && grid[nextleft.Right] == "#" {
							allGood = false
							break
						}

						nextRight := State{Point{right.X, left.Y - 1}, Point{right.X + 1, right.Y - 1}}
						if grid[nextRight.Left] == "[" && grid[nextRight.Right] == "]" {
							queue = append(queue, nextRight)
						} else if grid[nextRight.Left] == "#" && grid[nextRight.Right] == "#" {
							allGood = false
							break
						}
					}

					if allGood {
						for i := len(objs) - 1; i >= 0; i-- {
							grid[Point{objs[i].X, objs[i].Y - 1}] = grid[objs[i]]
							grid[objs[i]] = "."
						}
						pos = next
					}
				}
			} else if move == ">" {
				next := Point{pos.X + 1, pos.Y}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "[" || grid[next] == "]" {
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
						if grid[objPos] == "[" || grid[objPos] == "]" {
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
				} else if grid[next] == "[" || grid[next] == "]" {

					queue := make([]State, 0)

					if grid[next] == "[" {
						queue = append(queue, State{next, Point{next.X + 1, next.Y}})
					}
					if grid[next] == "]" {
						queue = append(queue, State{Point{next.X - 1, next.Y}, next})
					}

					objs := make([]Point, 0)
					allGood := true
					seen := make(map[State]bool)

					for len(queue) > 0 {

						current := queue[0]
						queue = queue[1:]

						if seen[current] {
							continue
						}
						seen[current] = true

						left, right := current.Left, current.Right

						if grid[left] == "[" && grid[right] == "]" {
							objs = append(objs, left, right)
						}

						nextMiddle := State{Point{left.X, left.Y + 1}, Point{right.X, right.Y + 1}}
						if grid[nextMiddle.Left] == "[" && grid[nextMiddle.Right] == "]" {
							queue = append(queue, nextMiddle)
							continue
						} else if grid[nextMiddle.Left] == "#" && grid[nextMiddle.Right] == "#" {
							allGood = false
							break
						}

						nextleft := State{Point{left.X - 1, left.Y + 1}, Point{left.X, left.Y + 1}}
						if grid[nextleft.Left] == "[" && grid[nextleft.Right] == "]" {
							queue = append(queue, nextleft)
						} else if grid[nextleft.Left] == "#" && grid[nextleft.Right] == "#" {
							allGood = false
							break
						}

						nextRight := State{Point{right.X, left.Y + 1}, Point{right.X + 1, right.Y + 1}}
						if grid[nextRight.Left] == "[" && grid[nextRight.Right] == "]" {
							queue = append(queue, nextRight)
						} else if grid[nextRight.Left] == "#" && grid[nextRight.Right] == "#" {
							allGood = false
							break
						}
					}

					if allGood {
						for i := len(objs) - 1; i >= 0; i-- {
							grid[Point{objs[i].X, objs[i].Y + 1}] = grid[objs[i]]
							grid[objs[i]] = "."
						}
						pos = next
					}
				}
			} else if move == "<" {
				next := Point{pos.X - 1, pos.Y}
				if grid[next] == "." {
					pos = next
				} else if grid[next] == "#" {
					continue
				} else if grid[next] == "[" || grid[next] == "]" {
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
						if grid[objPos] == "[" || grid[objPos] == "]" {
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
		if v == "[" {
			ans += 100*p.Y + p.X
		}
	}

	fmt.Println(ans)
}
