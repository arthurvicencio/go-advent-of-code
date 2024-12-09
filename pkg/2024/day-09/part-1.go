package twenty_twenty_four_day_09

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input1 string

type Part1 struct{}

func (p Part1) Solve() {
	if len(input1)%2 != 0 {
		input1 += "0"
	}

	rawInput := make([]int, 0)
	for _, num := range strings.Split(input1, "") {
		n, _ := strconv.Atoi(num)
		rawInput = append(rawInput, n)
	}

	disk := make([]int, 0)
	fileBlocks := make([]int, 0)
	freeSpaceLocations := make([]int, 0)

	for i := 0; i < len(rawInput)-1; i += 2 {

		blkCount := rawInput[i]
		for j := 0; j < blkCount; j++ {
			disk = append(disk, i/2)
			fileBlocks = append(fileBlocks, len(disk)-1)
		}

		freeCount := rawInput[i+1]
		for j := 0; j < freeCount; j++ {
			disk = append(disk, 0)
			freeSpaceLocations = append(freeSpaceLocations, len(disk)-1)
		}
	}

	for i := len(fileBlocks) - 1; fileBlocks[i] > freeSpaceLocations[0]; i-- {
		freeSpaceIndex, fileBlockIndex := freeSpaceLocations[0], fileBlocks[i]

		disk[freeSpaceIndex], disk[fileBlockIndex] = disk[fileBlockIndex], disk[freeSpaceIndex]

		freeSpaceLocations = freeSpaceLocations[1:]
	}

	var checksum int
	for i := 0; i < len(disk); i++ {
		checksum += i * disk[i]
	}

	fmt.Println(checksum)
}
