package days

import (
	"fmt"
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/utils2019"
)

type day3 struct {
	visited1 map[utils.Point]int
	visited2 map[utils.Point]int
}

func Day3() utils.Day {
	d := day3{}

	var err error
	input := utils.ConvertInputToStringSlice(utils.DayInput(2019, 3), "\n")

	d.visited1, err = utils2019.GetVisitedPoints(strings.Split(input[0], ",")...)
	utils.CheckError(err, `Cannot compute first path`)

	d.visited2, err = utils2019.GetVisitedPoints(strings.Split(input[1], ",")...)
	utils.CheckError(err, `Cannot compute second path`)

	return d
}

func (d day3) SolvePart1() string {
	min := int64(1_000_000_000)

	for p := range d.visited1 {
		if _, found := d.visited2[p]; found {
			dist := p.ManhattanDistance(utils.Origin)
			if dist < min {
				min = dist
			}
		}
	}

	return fmt.Sprintf("%d", min)
}

func (d day3) SolvePart2() string {
	min := 1_000_000_000

	for p, wireOneSteps := range d.visited1 {
		if wireTwoSteps, found := d.visited2[p]; found {
			if wireOneSteps+wireTwoSteps < min {
				min = wireOneSteps + wireTwoSteps
			}
		}
	}

	return strconv.Itoa(min)
}
