package days

import (
	"fmt"
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
)

type direction int

type pos struct {
	r int
	c int
}

type delta struct {
	dr int
	dc int
}

const (
	up        = direction(0)
	upright   = direction(1)
	right     = direction(2)
	downright = direction(3)
	down      = direction(4)
	downleft  = direction(5)
	left      = direction(6)
	upleft    = direction(7)
)

var (
	deltas = map[direction]delta{
		up:        {dr: -1, dc: 0},
		upright:   {dr: -1, dc: 1},
		right:     {dr: 0, dc: 1},
		downright: {dr: 1, dc: 1},
		down:      {dr: 1, dc: 0},
		downleft:  {dr: 1, dc: -1},
		left:      {dr: 0, dc: -1},
		upleft:    {dr: -1, dc: -1},
	}
)

type day4 struct {
	puzzle map[pos]byte
}

func Day4() utils.Day {
	d := day4{}

	input := utils.ConvertInputToStringSlice(utils.DayInput(2024, 4), "\n")
	d.puzzle = make(map[pos]byte)

	for r := range len(input) {
		for c := range len(input[0]) {
			d.puzzle[pos{r, c}] = input[r][c]
		}
	}
	return d
}

func (d day4) SolvePart1() string {
	total := 0

	for pos, char := range d.puzzle {
		if char == 'X' {
			words := d.scanForXMasFromPos(pos)
			for _, word := range words {
				if word == `XMAS` {
					total++
				}
			}
		}
	}

	return strconv.Itoa(total)
}

func (d day4) SolvePart2() string {
	total := 0

	for pos, char := range d.puzzle {
		if char == 'A' && d.isCrossedMasAtPos(pos) {
			total++
		}
	}

	return strconv.Itoa(total)
}

func (d day4) scanForXMasFromPos(p pos) []string {
	var words []string

	for _, delta := range deltas {
		currentR, currentC := p.r, p.c
		word := strings.Builder{}
		for range 4 {
			word.WriteByte(d.puzzle[pos{currentR, currentC}])
			currentR += delta.dr
			currentC += delta.dc
		}
		words = append(words, word.String())
	}

	return words
}

func (d day4) isCrossedMasAtPos(p pos) bool {
	left := fmt.Sprintf(`%c%c`, d.puzzle[p.fromDelta(deltas[upleft])], d.puzzle[p.fromDelta(deltas[downright])])
	right := fmt.Sprintf(`%c%c`, d.puzzle[p.fromDelta(deltas[downleft])], d.puzzle[p.fromDelta(deltas[upright])])

	return (left == `MS` || left == `SM`) && (right == `MS` || right == `SM`)
}

func (p pos) fromDelta(d delta) pos {
	return pos{p.r + d.dr, p.c + d.dc}
}
