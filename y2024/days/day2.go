package days

import (
	"strconv"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2024/internal"
)

type day2 struct {
	levels [][]int
}

func Day2() utils.Day {
	d := day2{}

	lines := utils.ConvertInputToStringSlice(utils.DayInput(2024, 2), "\n")

	for _, line := range lines {
		d.levels = append(d.levels, utils.ExtractInts(line, false))
	}

	return d
}

func (d day2) SolvePart1() string {
	total := 0

	for _, level := range d.levels {
		if internal.LevelsAreSafe(level) {
			total++
		}
	}

	return strconv.Itoa(total)
}

func (d day2) SolvePart2() string {
	total := 0

	for _, level := range d.levels {
		if internal.LevelsAreSafeWithDampener(level) {
			total++
		}
	}

	return strconv.Itoa(total)
}
