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

	if len(os.Args) < 2 {
		fmt.Println("missing name argument")
		os.Exit(1)
		return
	}

	name := os.Args[1]
	if _, exists := problems[name]; !exists {
		fmt.Printf("no solution for \"%s\"\n", name)
		os.Exit(1)
		return
	}

	problems[name].Solve()
}

func mergeMap[A comparable, B any](a, b map[A]B) {
	for i, v := range b {
		a[i] = v
	}
}
