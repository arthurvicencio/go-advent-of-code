package twenty_twenty_four_day_08

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	grid := make([][]string, 0)

	antenas := make([]Point, 0)
	antinodeLocations := make(map[Point]bool)

	for y, line := range strings.Split(input2, "\n") {
		row := make([]string, 0)
		for x, cell := range strings.Split(line, "") {
			row = append(row, cell)
			if cell != "." {
				antenas = append(antenas, Point{x, y})
				antinodeLocations[Point{x, y}] = true
			}
		}

		grid = append(grid, row)
	}

	for i := 0; i < len(antenas)-1; i++ {
		for j := i + 1; j < len(antenas); j++ {

			antenaA, antenaB := antenas[i], antenas[j]
			if grid[antenaA.Y][antenaA.X] != grid[antenaB.Y][antenaB.X] {
				continue
			}

			angle := p.angle(antenaA, antenaB)
			dist := p.distance(antenaA, antenaB)

			p.addPointToMap(
				p.getAllNextPointsInBounds(antenaB, angle, dist, grid),
				antinodeLocations,
			)

			reverseAngle := math.Mod(angle+180, 360)
			p.addPointToMap(
				p.getAllNextPointsInBounds(antenaA, reverseAngle, dist, grid),
				antinodeLocations,
			)
		}
	}

	fmt.Println(len(antinodeLocations))
}

func (p Part2) inBoundsOfgrid(x, y int, grid [][]string) bool {
	return x >= 0 && x <= len(grid[0])-1 && y >= 0 && y <= len(grid)-1
}

func (p Part2) getAllNextPointsInBounds(pnt Point, angle, distance float64, grid [][]string) []Point {
	points := make([]Point, 0)
	nextPoint := pnt
	for {
		nextPoint = p.getNextPoint(nextPoint, angle, distance)
		if !p.inBoundsOfgrid(nextPoint.X, nextPoint.Y, grid) {
			break
		}
		points = append(points, nextPoint)
	}
	return points
}

func (p Part2) addPointToMap(pnts []Point, m map[Point]bool) {
	for _, p := range pnts {
		m[p] = true
	}
}

func (p Part2) angle(p1, p2 Point) float64 {
	// Calculate the differences
	dx := float64(p2.X - p1.X)
	dy := float64(p2.Y - p1.Y)

	// Calculate the angle in radians using atan2
	angle := math.Atan2(dy, dx)

	// Convert the angle to degrees if needed
	angleDegrees := angle * (180 / math.Pi)

	// Ensure the angle is positive (0 to 360 degrees)
	if angleDegrees < 0 {
		angleDegrees += 360
	}

	return angleDegrees
}

func (p Part2) getNextPoint(pnt Point, angle, distance float64) Point {
	radians := angle * (math.Pi / 180) // Convert angle to radians
	newX := float64(pnt.X) + distance*math.Cos(radians)
	newY := float64(pnt.Y) + distance*math.Sin(radians)
	return Point{int(math.Round(newX)), int(math.Round(newY))}
}

func (p Part2) distance(p1, p2 Point) float64 {
	// Apply the distance formula
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
