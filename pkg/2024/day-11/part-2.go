package twenty_twenty_four_day_11

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

type State_P2 struct {
	Blink int
	Num   int
}

var cache_p2 = make(map[State_P2]int)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	nums := make([]int, 0)
	for _, num := range strings.Split(input2, " ") {
		n, _ := strconv.Atoi(num)
		nums = append(nums, n)
	}

	var stones int
	for _, n := range nums {
		stones += p.blink(0, n, 75)
	}

	fmt.Println(stones)
}

func (p Part2) blink(iter int, num int, length int) int {
	state := State_P2{iter, num}
	if _, ok := cache_p2[state]; ok {
		return cache_p2[state]
	}

	if iter >= length {
		return 1
	}

	if num == 0 {
		cache_p2[state] = p.blink(iter+1, 1, length)
	} else {
		digits := strconv.Itoa(num)
		if len(digits)%2 == 0 {
			mid := len(digits) / 2
			left, _ := strconv.Atoi(digits[:mid])
			right, _ := strconv.Atoi(digits[mid:])

			cache_p2[state] += p.blink(iter+1, left, length)
			cache_p2[state] += p.blink(iter+1, right, length)
		} else {
			cache_p2[state] = p.blink(iter+1, num*2024, length)
		}
	}

	return cache_p2[state]
}
