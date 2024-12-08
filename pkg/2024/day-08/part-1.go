package twenty_twenty_four_day_08

import (
	_ "embed"
	"fmt"
	"math"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

type Point_P1 struct{ X, Y int }

func (p Part1) Solve() {
	grid := make([][]string, 0)

	antenaLocations := make([]Point_P1, 0)
	antinodeLocations := make(map[Point_P1]bool)

	for y, line := range strings.Split(input1, "\n") {
		row := make([]string, 0)
		for x, cell := range strings.Split(line, "") {
			row = append(row, cell)
			if cell != "." {
				antenaLocations = append(antenaLocations, Point_P1{x, y})
			}
		}

		grid = append(grid, row)
	}

	for i := 0; i < len(antenaLocations)-1; i++ {
		for j := i + 1; j < len(antenaLocations); j++ {

			antenaA, antenaB := antenaLocations[i], antenaLocations[j]
			if grid[antenaA.Y][antenaA.X] != grid[antenaB.Y][antenaB.X] {
				continue
			}

			angle := p.angle(antenaA, antenaB)
			dist := p.distance(antenaA, antenaB)

			nextPoint := p.getNextPoint(antenaB, angle, dist)
			if p.inBoundsOfgrid(nextPoint.X, nextPoint.Y, grid) {
				antinodeLocations[nextPoint] = true
			}

			reverseAngle := math.Mod(angle+180, 360)
			nextPoint2 := p.getNextPoint(antenaA, reverseAngle, dist)
			if p.inBoundsOfgrid(nextPoint2.X, nextPoint2.Y, grid) {
				antinodeLocations[nextPoint2] = true
			}
		}
	}

	fmt.Println(len(antinodeLocations))
}

func (p Part1) inBoundsOfgrid(x, y int, grid [][]string) bool {
	return x >= 0 && x <= len(grid[0])-1 && y >= 0 && y <= len(grid)-1
}

func (p Part1) angle(p1, p2 Point_P1) float64 {
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

func (p Part1) getNextPoint(pnt Point_P1, angle, distance float64) Point_P1 {
	radians := angle * (math.Pi / 180) // Convert angle to radians
	newX := float64(pnt.X) + distance*math.Cos(radians)
	newY := float64(pnt.Y) + distance*math.Sin(radians)
	return Point_P1{int(math.Round(newX)), int(math.Round(newY))}
}

func (p Part1) distance(p1, p2 Point_P1) float64 {
	// Apply the distance formula
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
