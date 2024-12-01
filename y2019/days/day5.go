package days

import (
	"fmt"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/internal"
)

type day5 struct {
	program []int64
}

func Day5() utils.Day {
	d := day5{}
	input := utils.DayInput(2019, 5)
	d.program = utils.ConvertInputToBigIntSlice(input, ",")

	return d
}

func (d day5) SolvePart1() string {
	icp := internal.NewIntCodeProgram(d.program, 1)
	if status := icp.Run(); status != internal.StatusComplete {
		return `unable to run program`
	}
	return fmt.Sprintf("%d", icp.Outputs[len(icp.Outputs)-1])
}

func (d day5) SolvePart2() string {
	icp := internal.NewIntCodeProgram(d.program, 5)
	if status := icp.Run(); status != internal.StatusComplete {
		return `unable to run program`
	}
	return fmt.Sprintf("%d", icp.Outputs[len(icp.Outputs)-1])
}
