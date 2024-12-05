package twenty_twenty_four_day_03

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {
	var answer int

	reg := regexp.MustCompile(`mul\(\d+,\d+\)`)

	for _, line := range strings.Split(input1, "\n") {

		muls := reg.FindAllString(line, -1)
		for _, mul := range muls {
			var a, b int
			fmt.Sscanf(mul, "mul(%d,%d)", &a, &b)
			answer += a * b
		}
	}

	fmt.Println(answer)
}
