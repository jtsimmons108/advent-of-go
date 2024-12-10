package days

import (
	"strconv"

	"simmons.com/advent-of-go/utils"
)

type day8 struct {
	nodes map[rune][]pos
	rows  int
	cols  int
}

func Day8() utils.Day {
	d := day8{}

	d.nodes = map[rune][]pos{}

	lines := utils.ConvertInputToStringSlice(utils.DayInput(2024, 8), "\n")
	d.rows = len(lines)
	d.cols = len(lines[0])

	for r, line := range lines {
		for c, char := range line {
			if char != '.' {
				d.nodes[char] = append(d.nodes[char], pos{r, c})
			}
		}
	}
	return d
}

func (d day8) SolvePart1() string {

	antinodes := map[rune][]pos{}

	for char, posList := range d.nodes {
		for i, p1 := range posList {
			for j, p2 := range posList {
				if i != j {
					dr := p1.r - p2.r
					dc := p1.c - p2.c

					antiNodePos := pos{p1.r + dr, p1.c + dc}
					if antiNodePos.r >= 0 && antiNodePos.r < d.rows && antiNodePos.c >= 0 && antiNodePos.c < d.cols {
						antinodes[char] = append(antinodes[char], antiNodePos)
					}
				}
			}
		}
	}

	uniqueAntiNodes := map[pos]struct{}{}

	for _, posList := range antinodes {
		for _, pos := range posList {
			uniqueAntiNodes[pos] = struct{}{}
		}
	}
	return strconv.Itoa(len(uniqueAntiNodes))
}

func (d day8) SolvePart2() string {
	antinodes := map[rune][]pos{}

	for char, posList := range d.nodes {
		for i, p1 := range posList {
			for j, p2 := range posList {
				if i != j {
					dr := p1.r - p2.r
					dc := p1.c - p2.c

					antiNodePos := pos{p1.r, p1.c}
					for antiNodePos.r >= 0 && antiNodePos.r < d.rows && antiNodePos.c >= 0 && antiNodePos.c < d.cols {
						antinodes[char] = append(antinodes[char], antiNodePos)
						antiNodePos = pos{antiNodePos.r + dr, antiNodePos.c + dc}
					}
				}
			}
		}
	}

	uniqueAntiNodes := map[pos]struct{}{}

	for _, posList := range antinodes {
		for _, pos := range posList {
			uniqueAntiNodes[pos] = struct{}{}
		}
	}
	return strconv.Itoa(len(uniqueAntiNodes))
}
