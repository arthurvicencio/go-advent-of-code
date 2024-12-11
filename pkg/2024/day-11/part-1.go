package twenty_twenty_four_day_11

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type State struct {
	Blink int
	Num   int
}

var cache = make(map[State]int)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {
	nums := make([]int, 0)
	for _, num := range strings.Split(input1, " ") {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}

	var stones int
	for _, n := range nums {
		stones += p.blink(0, n, 25)
	}

	fmt.Println(stones)
}

func (p Part1) blink(iter int, num int, length int) int {
	state := State{iter, num}
	if _, ok := cache[state]; ok {
		return cache[state]
	}

	if iter >= length {
		return 1
	}

	if num == 0 {
		cache[state] = p.blink(iter+1, 1, length)
	} else {
		digits := strconv.Itoa(num)
		if len(digits)%2 == 0 {
			mid := len(digits) / 2
			left, _ := strconv.Atoi(digits[:mid])
			right, _ := strconv.Atoi(digits[mid:])
			cache[state] += p.blink(iter+1, left, length)
			cache[state] += p.blink(iter+1, right, length)
		} else {
			cache[state] = p.blink(iter+1, num*2024, length)
		}
	}

	return cache[state]
}
