package twenty_twenty_four_day_03

import (
	_ "embed"
	"fmt"
	"regexp"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	var answer int

	reg := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)

	enabled := true

	for _, line := range strings.Split(input2, "\n") {

		exps := reg.FindAllString(line, -1)
		for _, exp := range exps {

			if strings.HasPrefix(exp, "mul") {

				if !enabled {
					continue
				}

				var a, b int
				fmt.Sscanf(exp, "mul(%d,%d)", &a, &b)
				answer += a * b
			} else if exp == "do()" {
				enabled = true
			} else if exp == "don't()" {
				enabled = false
			}
		}
	}

	fmt.Println(answer)
}
