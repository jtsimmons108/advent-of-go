package days

import (
	"fmt"

	"simmons.com/advent-of-go/mathutils"
	"simmons.com/advent-of-go/utils"
)

type planet struct {
	xPos int64
	yPos int64
	zPos int64
	xVel int64
	yVel int64
	zVel int64
}

type day12 struct {
	debug        bool
	planets      []*planet
	initialState []planet
}

func Day12() utils.Day {
	d := day12{
		debug: false,
	}

	input := utils.DayInput(2019, 12)
	lines := utils.ConvertInputToStringSlice(input, "\n")

	for _, line := range lines {
		ints := utils.ExtractInts(line, true)
		d.planets = append(d.planets, &planet{
			xPos: int64(ints[0]),
			yPos: int64(ints[1]),
			zPos: int64(ints[2]),
		})
		d.initialState = append(d.initialState, planet{
			xPos: int64(ints[0]),
			yPos: int64(ints[1]),
			zPos: int64(ints[2]),
		})
	}

	return d
}

func (d day12) SolvePart1() string {

	for range 100 {
		d.step()
	}

	total := int64(0)

	for _, p := range d.planets {
		total += p.potentialEnergy() * p.kineticEnergy()
	}

	return fmt.Sprintf("%d", total)
}

func (d day12) SolvePart2() string {
	for i := range d.planets {
		init := d.initialState[i]
		d.planets[i] = &planet{
			xPos: init.xPos,
			yPos: init.yPos,
			zPos: init.zPos,
		}
	}

	pointsMatch := make([]bool, 4)
	velocitiesMatch := make([]bool, 4)
	steps := make([]int, 4)
	step := 0
	for !allTrue(append(pointsMatch, velocitiesMatch...)) {
		d.step()
		step++
		for i := range d.planets {
			p := d.planets[i]
			initial := d.initialState[i]

			if p.xVel == 0 && p.yVel == 0 && p.zVel == 0 {
				fmt.Printf("Velocities zero out for Planet %d at step %04d \n", i+1, step)
				velocitiesMatch[i] = true
			}

			if p.xPos == initial.xPos && p.yPos == initial.yPos && p.zPos == initial.zPos {
				fmt.Printf("Points line up for Planet %d at step %04d \n", i+1, step)
				pointsMatch[i] = true
			}

			if *d.planets[i] == d.initialState[i] {
				steps[i] = step
				fmt.Printf("Found repeated step for Planet %d at step: %d\n", i+1, step)
			}
		}

	}

	return fmt.Sprintf("%v", steps)

}

func (d day12) step() {
	for i := range d.planets {
		for j := range d.planets {
			if i != j {
				d.planets[i].gravitate(d.planets[j])
			}
		}

	}
	for _, p := range d.planets {
		p.move()
		if d.debug {
			fmt.Printf("%s\n", p.String())
		}
	}
}

func (p *planet) gravitate(o *planet) {
	if p.xPos > o.xPos {
		p.xVel--
	} else if p.xPos < o.xPos {
		p.xVel++
	}

	if p.yPos > o.yPos {
		p.yVel--
	} else if p.yPos < o.yPos {
		p.yVel++
	}

	if p.zPos > o.zPos {
		p.zVel--
	} else if p.zPos < o.zPos {
		p.zVel++
	}
}

func (p *planet) move() {
	p.xPos += p.xVel
	p.yPos += p.yVel
	p.zPos += p.zVel
}

func (p *planet) String() string {
	return fmt.Sprintf("Pos: <%2d, %2d, %2d>, Vel: <%2d, %2d, %2d>", p.xPos, p.yPos, p.zPos, p.xVel, p.yVel, p.zVel)
}

func (p *planet) potentialEnergy() int64 {
	return mathutils.Abs64(p.xPos) + mathutils.Abs64(p.yPos) + mathutils.Abs64(p.zPos)
}

func (p *planet) kineticEnergy() int64 {
	return mathutils.Abs64(p.xVel) + mathutils.Abs64(p.yVel) + mathutils.Abs64(p.zVel)
}

func allTrue(bools []bool) bool {
	for _, b := range bools {
		if !b {
			return false
		}
	}
	return true
}
