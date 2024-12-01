package days

import (
	"strconv"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/utils2019"
)

type day1 struct {
	masses []int
}

func Day1() utils.Day {
	d := day1{}

	input := utils.DayInput(2019, 1)
	d.masses = utils.ConvertInputToIntSlice(input, "\n")

	return d
}

func (d day1) SolvePart1() string {

	total := 0

	for _, n := range d.masses {
		total += utils2019.GetFuelFromMass(n)
	}

	return strconv.Itoa(total)
}

func (d day1) SolvePart2() string {
	total := 0

	for _, n := range d.masses {
		total += utils2019.GetTotalFuelFromMass(n)
	}

	return strconv.Itoa(total)
}
