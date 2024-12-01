package days

import (
	"fmt"
	"strconv"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/internal"
)

type day2 struct {
	program []int64
}

func Day2() utils.Day {

	d := day2{}
	input := utils.DayInput(2019, 2)
	d.program = utils.ConvertInputToBigIntSlice(input, ",")

	return d
}

func (d day2) SolvePart1() string {

	d.program[1] = 12
	d.program[2] = 2
	icp := internal.NewIntCodeProgram(d.program)

	if status := icp.Run(); status != internal.StatusComplete {
		return `unable to calculate day 2`
	}

	return fmt.Sprintf("%d", icp.Program[0])
}

func (d day2) SolvePart2() string {
	noun := 0
	target := int64(19690720)

	for noun < 100 {
		verb := 0
		for verb < 100 {
			d.program[1] = int64(noun)
			d.program[2] = int64(verb)

			icp := internal.NewIntCodeProgram(d.program)

			if status := icp.Run(); status != internal.StatusComplete {
				return `unable to calculate day 2 part 2`
			}
			if icp.Program[0] == target {
				return strconv.Itoa(noun*100 + verb)
			}
			verb += 1
		}
		noun += 1
	}
	panic(`No Answer Found`)
}
