package days

import (
	"regexp"
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
)

type day3 struct {
	input   string
	pattern *regexp.Regexp
}

func Day3() utils.Day {
	d := day3{}

	d.input = utils.DayInput(2024, 3)
	d.pattern = regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

	return d
}

func (d day3) SolvePart1() string {
	muls := d.pattern.FindAllString(d.input, -1)

	total := 0

	for _, mul := range muls {
		ints := utils.ExtractInts(mul, false)
		total += ints[0] * ints[1]
	}

	return strconv.Itoa(total)
}

func (d day3) SolvePart2() string {
	muls := d.pattern.FindAllString(d.input, -1)
	parts := d.pattern.Split(d.input, -1)

	doMul := true
	total := 0

	for i := range muls {
		before := parts[i]
		doIndex := strings.Index(before, `do()`)
		dontIndex := strings.Index(before, `don't()`)

		switch {
		case doIndex > -1 && doIndex > dontIndex:
			doMul = true
		case dontIndex > -1 && dontIndex > doIndex:
			doMul = false
		}

		if doMul {
			ints := utils.ExtractInts(muls[i], false)
			total += ints[0] * ints[1]
		}

	}

	return strconv.Itoa(total)
}
