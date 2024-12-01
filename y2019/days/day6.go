package days

import (
	"fmt"
	"strings"

	"simmons.com/advent-of-go/utils"
)

type day6 struct {
	orbits map[string]string
}

func Day6() utils.Day {
	d := day6{
		orbits: map[string]string{},
	}

	input := utils.DayInput(2019, 6)
	lines := utils.ConvertInputToStringSlice(input, "\n")

	for _, line := range lines {
		split := strings.Split(line, ")")
		d.orbits[split[1]] = split[0]

	}
	return d
}

func (d day6) SolvePart1() string {
	total := 0

	for planet := range d.orbits {
		total += d.OrbitCount(planet)
	}

	return fmt.Sprintf("%d", total)
}

func (d day6) SolvePart2() string {
	min := len(d.orbits)

	for planet := range d.orbits {
		if planet == "YOU" || planet == "SAN" {
			continue
		}

		if d.Orbits("YOU", planet) && d.Orbits("SAN", planet) {
			hops := d.Steps("YOU", planet) + d.Steps("SAN", planet) - 2
			if hops < min {
				min = hops
			}
		}
	}

	return fmt.Sprintf("%d", min)
}

func (d day6) OrbitCount(outer string) int {
	inner, found := d.orbits[outer]
	if !found {
		return 0
	}
	return 1 + d.OrbitCount(inner)
}

func (d day6) Orbits(planet1 string, planet2 string) bool {
	inner, found := d.orbits[planet1]
	if !found {
		return false
	}
	if inner == planet2 {
		return true
	}
	return d.Orbits(inner, planet2)
}

func (d day6) Steps(planet1 string, planet2 string) int {
	if planet1 == planet2 {
		return 0
	}

	if !d.Orbits(planet1, planet2) {
		panic(`Trying to get steps to planet it does not orbit`)
	}

	inner := d.orbits[planet1]
	return 1 + d.Steps(inner, planet2)
}
