package twenty_twenty_four_day_09

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var input2 string

type Part2 struct{}

func (p Part2) Solve() {
	type Marker struct {
		Start  int
		Length int
	}

	if len(input2)%2 != 0 {
		input2 += "0"
	}

	rawInput := make([]int, 0)
	for _, num := range strings.Split(input2, "") {
		n, _ := strconv.Atoi(num)
		rawInput = append(rawInput, n)
	}

	disk := make([]int, 0)

	fileMarkers := make([]Marker, 0)
	freeSpaceMarkers := make([]Marker, 0)

	for i := 0; i < len(rawInput)-1; i += 2 {

		blkCount := rawInput[i]
		fileMarkers = append(fileMarkers, Marker{
			Start:  len(disk),
			Length: blkCount,
		})

		freeCount := rawInput[i+1]
		freeSpaceMarkers = append(freeSpaceMarkers, Marker{
			Start:  len(disk) + blkCount,
			Length: freeCount,
		})

		for j := 0; j < blkCount; j++ {
			disk = append(disk, i/2)
		}

		for j := 0; j < freeCount; j++ {
			disk = append(disk, 0)
		}
	}

	for i := len(fileMarkers) - 1; i >= 0; i-- {

		lastFileMarker := fileMarkers[i]

		for j, freeSpaceMarker := range freeSpaceMarkers {

			if freeSpaceMarker.Start >= lastFileMarker.Start {
				continue
			}

			if freeSpaceMarker.Length >= lastFileMarker.Length {

				freeSpaceMarkers[j].Start += lastFileMarker.Length
				freeSpaceMarkers[j].Length -= lastFileMarker.Length

				fileStart := lastFileMarker.Start
				frssSpaceStart := freeSpaceMarker.Start

				for i := 0; i < lastFileMarker.Length; i++ {
					disk[fileStart], disk[frssSpaceStart] = disk[frssSpaceStart], disk[fileStart]
					fileStart++
					frssSpaceStart++
				}
				break
			}
		}

	}

	var checksum int
	for i := 0; i < len(disk); i++ {
		checksum += i * disk[i]
	}

	fmt.Println(checksum)
}
