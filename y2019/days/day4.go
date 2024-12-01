package days

import (
	"strconv"

	"simmons.com/advent-of-go/utils"
)

type day4 struct {
	low  int
	high int
}

func Day4() utils.Day {
	return day4{
		low:  254032,
		high: 789860,
	}
}

func (d day4) SolvePart1() string {
	current := d.low
	total := 0

	for current <= d.high {
		if IsValidPasswordPart1(current) {
			total++
		}
		current++
	}

	return strconv.Itoa(total)
}

func (d day4) SolvePart2() string {
	current := d.low
	total := 0

	for current <= d.high {
		if IsValidPasswordPart2(current) {
			total++
		}
		current++
	}

	return strconv.Itoa(total)
}

func IsValidPasswordPart1(pass int) bool {

	str := strconv.Itoa(pass)
	i := 1
	var (
		adjacent      = false
		nondecreasing = true
	)

	for i < len(str) {
		a := str[i]
		b := str[i-1]

		if a == b {
			adjacent = true
		}

		if int(a-'0') < int(b-'0') {
			nondecreasing = false
		}
		i++
	}

	return adjacent && nondecreasing
}

func IsValidPasswordPart2(pass int) bool {

	str := strconv.Itoa(pass)
	i := 1
	var (
		twoAdjacent   = false
		nondecreasing = true
		consecutive   = 1
	)

	for i < len(str) {
		a := str[i]
		b := str[i-1]

		if a == b {
			consecutive++
		} else {
			if consecutive == 2 {
				twoAdjacent = true
			}
			consecutive = 1
		}

		if int(a-'0') < int(b-'0') {
			nondecreasing = false
		}
		i++
	}

	return (consecutive == 2 || twoAdjacent) && nondecreasing
}
