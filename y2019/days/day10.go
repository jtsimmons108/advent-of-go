package days

import (
	"fmt"
	"math"
	"sort"

	"simmons.com/advent-of-go/utils"
)

type day10 struct {
	points         []utils.Point
	targetPointMap map[string]utils.Point
	angles         map[utils.Point]map[utils.Point]float64
}

func Day10() utils.Day {
	d := day10{
		targetPointMap: map[string]utils.Point{},
	}
	input := utils.DayInput(2019, 10)
	lines := utils.ConvertInputToStringSlice(input, "\n")

	for y, row := range lines {
		for x, c := range row {
			if c == '#' {
				d.points = append(d.points, utils.Point{X: int64(x), Y: int64(y)})
			}
		}
	}

	d.angles = map[utils.Point]map[utils.Point]float64{}
	for _, p := range d.points {
		d.angles[p] = map[utils.Point]float64{}
	}
	for _, point := range d.points {
		for _, other := range d.points {
			if point != other {
				dx := other.X - point.X
				dy := point.Y - other.Y // Y is negated

				angle := math.Atan(float64(dy) / float64(dx))
				if dx < 0 {
					angle -= math.Pi
				}

				d.angles[point][other] = angle
			}

		}
	}

	return d
}

func (d day10) SolvePart1() string {

	max := 0
	for p, others := range d.angles {
		allDeltas := map[float64]struct{}{}
		for _, angle := range others {
			if _, found := allDeltas[angle]; !found {
				allDeltas[angle] = struct{}{}
			}
		}

		if len(allDeltas) > max {
			max = len(allDeltas)
			d.targetPointMap[`ans`] = p
		}
	}
	return fmt.Sprintf("%d", max)
}

func (d day10) SolvePart2() string {

	type beacon struct {
		point utils.Point
		angle float64
	}

	monitor := d.targetPointMap[`ans`]

	others, found := d.angles[monitor]
	if !found {
		panic(`Can't find answer from part 1`)
	}

	beacons := []beacon{}

	for p, angle := range others {
		beacons = append(beacons, beacon{
			point: p,
			angle: angle,
		})
	}

	sort.Slice(beacons, func(i, j int) bool {
		b1 := beacons[i]
		b2 := beacons[j]

		if b1.angle == b2.angle {
			return b1.point.Distance(monitor) < b2.point.Distance(monitor)
		}

		return b1.angle > b2.angle
	})

	destroyed := []beacon{}
	targetLength := len(beacons)

	for len(destroyed) != targetLength {
		angles := map[float64]struct{}{}
		saved := []beacon{}

		for _, b := range beacons {
			_, found := angles[b.angle]

			if found {
				saved = append(saved, b)
			} else {
				angles[b.angle] = struct{}{}
				destroyed = append(destroyed, b)
			}
		}
		beacons = saved
	}

	target := destroyed[199]
	return fmt.Sprintf("%d", target.point.X*100+target.point.Y)
}
