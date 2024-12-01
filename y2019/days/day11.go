package days

import (
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/internal"
)

type day11 struct {
	program []int64
	deltas  map[complex128]utils.Delta
}

func Day11() utils.Day {
	d := day11{}

	input := utils.DayInput(2019, 11)
	d.program = utils.ConvertInputToBigIntSlice(input, ",")

	d.deltas = map[complex128]utils.Delta{
		1i:  {Dx: 0, Dy: 1},
		-1i: {Dx: 0, Dy: -1},
		1:   {Dx: 1, Dy: 0},
		-1:  {Dx: -1, Dy: 0},
	}

	return d
}

func (d day11) SolvePart1() string {

	colors := map[utils.Point]int{}
	current := utils.Origin
	dir := 0 + 1i

	icp := internal.NewIntCodeProgram(d.program)
	status := internal.StatusContinue
	outputs := 0

	for status != internal.StatusComplete {
		status = icp.WithIOPauses().Run()

		switch status {
		case internal.StatusInput:
			paint := 0
			if c, found := colors[current]; found {
				paint = c
			}
			icp.AddInput(int64(paint))
		case internal.StatusOutput:
			val := icp.GetOutput()
			switch outputs % 2 {
			case 0:
				colors[current] = int(val)
			case 1:
				if val == 0 {
					dir *= 1i
				} else {
					dir *= -1i
				}
				delta, found := d.deltas[dir]
				if !found {
					panic(`Trying to move in unrecognized direction`)
				}

				current = current.NextPoint(delta)
			}

			outputs++
		case internal.StatusFault:
			panic(`Unexpected fault in program`)
		default:
			// no-opp
		}

	}

	return strconv.Itoa(len(colors))

}

func (d day11) SolvePart2() string {

	current := utils.Origin
	colors := map[utils.Point]int{current: 1}
	dir := 0 + 1i

	icp := internal.NewIntCodeProgram(d.program)
	status := internal.StatusContinue
	outputs := 0

	for status != internal.StatusComplete {
		status = icp.WithIOPauses().Run()

		switch status {
		case internal.StatusInput:
			paint := 0
			if c, found := colors[current]; found {
				paint = c
			}
			icp.AddInput(int64(paint))
		case internal.StatusOutput:
			val := icp.GetOutput()
			switch outputs % 2 {
			case 0:
				colors[current] = int(val)
			case 1:
				if val == 0 {
					dir *= 1i
				} else {
					dir *= -1i
				}
				delta, found := d.deltas[dir]
				if !found {
					panic(`Trying to move in unrecognized direction`)
				}

				current = current.NextPoint(delta)
			}

			outputs++
		case internal.StatusFault:
			panic(`Unexpected fault in program`)
		default:
			// no-opp
		}

	}

	r := struct {
		minX int64
		minY int64
		maxX int64
		maxY int64
	}{}

	for p := range colors {
		if p.X > r.maxX {
			r.maxX = p.X
		} else if p.X < r.minX {
			r.minX = p.X
		}

		if p.Y > r.maxY {
			r.maxY = p.Y
		} else if p.Y < r.minY {
			r.minY = p.Y
		}
	}

	y := r.maxY

	res := strings.Builder{}
	res.WriteString("\n")
	tiles := []string{" ", "#"}
	for y >= r.minY {
		x := r.minX
		for x <= r.maxX {
			color := 0
			if c, found := colors[utils.Point{X: x, Y: y}]; found {
				color = c
			}
			res.WriteString(tiles[color])
			x++
		}
		res.WriteString("\n")
		y--
	}

	return res.String()

}
