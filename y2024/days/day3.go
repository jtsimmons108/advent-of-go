package days

import (
	"regexp"
	"strconv"

	"simmons.com/advent-of-go/utils"
)

type day3 struct {
	input string
}

func Day3() utils.Day {
	d := day3{}

	d.input = utils.DayInput(2024, 3)

	return d
}

func (d day3) SolvePart1() string {
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	muls := re.FindAllString(d.input, -1)

	total := 0

	for _, mul := range muls {
		ints := utils.ExtractInts(mul, false)
		total += ints[0] * ints[1]
	}

	return strconv.Itoa(total)
}

func (d day3) SolvePart2() string {

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`)
	matches := re.FindAllString(d.input, -1)

	doMul := true
	total := 0

	for _, match := range matches {
		switch match {
		case `do()`:
			doMul = true
		case `don't()`:
			doMul = false
		default:
			if doMul {
				ints := utils.ExtractInts(match, false)
				total += ints[0] * ints[1]
			}

		}
	}

	return strconv.Itoa(total)
}
