package days

import (
	"fmt"
	"strconv"

	"simmons.com/advent-of-go/utils"
)

type day6 struct {
	rows     int
	cols     int
	startPos pos
	blocks   map[pos]struct{}
}

func Day6() utils.Day {
	d := day6{}
	d.blocks = make(map[pos]struct{})

	input := utils.ConvertInputToStringSlice(utils.DayInput(2024, 6), "\n")
	d.rows = len(input)
	d.cols = len(input[0])

	for r, row := range input {
		for c, char := range row {
			switch char {
			case '#':
				d.blocks[pos{r, c}] = struct{}{}
			case '^':
				d.startPos = pos{r, c}
			default:
				// no-op
			}
		}
	}
	return d
}

func (d day6) SolvePart1() string {
	currentPos := d.startPos
	direction := complex(0, 1)
	visited := map[pos]struct{}{currentPos: {}}

	for currentPos.r >= 0 && currentPos.r < d.rows && currentPos.c >= 0 && currentPos.c < d.cols {
		nextPos := pos{currentPos.r + int(imag(direction)*-1), currentPos.c + int(real(direction))}

		_, found := d.blocks[nextPos]
		switch found {
		case true:
			direction *= complex(0, -1)
			fmt.Printf("Hit block at %v, turning to %v\n", currentPos, direction)
		case false:
			visited[currentPos] = struct{}{}
			fmt.Printf("Moving from %v to %v\n", currentPos, nextPos)
			currentPos = nextPos
		}

	}
	return strconv.Itoa(len(visited))
}

func (d day6) SolvePart2() string {

	return ``
}
