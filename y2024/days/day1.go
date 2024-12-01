package days

import (
	"sort"
	"strconv"

	"simmons.com/advent-of-go/mathutils"
	"simmons.com/advent-of-go/utils"
)

type day1 struct {
	leftColumn  []int
	rightColumn []int
}

func Day1() utils.Day {
	d := day1{}
	input := utils.ConvertInputToStringSlice(utils.DayInput(2024, 1), "\n")

	for _, line := range input {
		nums := utils.ExtractInts(line, false)
		d.leftColumn = append(d.leftColumn, nums[0])
		d.rightColumn = append(d.rightColumn, nums[1])
	}

	sort.Ints(d.leftColumn)
	sort.Ints(d.rightColumn)

	return d
}

func (d day1) SolvePart1() string {
	total := 0

	for i := range d.leftColumn {
		total += mathutils.Abs(d.leftColumn[i] - d.rightColumn[i])
	}

	return strconv.Itoa(total)
}

func (d day1) SolvePart2() string {
	counts := map[int]int{}

	for _, num := range d.rightColumn {
		counts[num]++
	}

	total := 0

	for _, num := range d.leftColumn {
		total += num * counts[num]
	}

	return strconv.Itoa(total)
}
