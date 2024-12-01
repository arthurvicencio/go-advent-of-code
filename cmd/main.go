package main

import (
	"fmt"
	"os"

	twenty_twenty_four "github.com/arthurvicencio/go-advent-of-code/pkg/2024"
	"github.com/arthurvicencio/go-advent-of-code/pkg/aocutils"
)

func main() {

	problems := make(map[string]aocutils.Problem)
	mergeMap(problems, twenty_twenty_four.Problems())

	if _, exists := problems[os.Args[1]]; !exists {
		fmt.Printf("no solution for \"%s\"\n", os.Args[1])
		os.Exit(1)
		return
	}

	problems[os.Args[1]].Solve()
}

func mergeMap[A, B comparable](a, b map[A]B) {
	for i, v := range b {
		a[i] = v
	}
}
