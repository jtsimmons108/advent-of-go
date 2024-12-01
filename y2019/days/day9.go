package days

import (
	"fmt"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/internal"
)

type day9 struct {
	program []int64
}

func Day9() utils.Day {
	d := day9{}
	input := utils.DayInput(2019, 9)
	d.program = utils.ConvertInputToBigIntSlice(input, ",")

	return d
}

func (d day9) SolvePart1() string {
	icp := internal.NewIntCodeProgram(d.program, 1)
	if status := icp.Run(); status != internal.StatusComplete {
		return `Unable to run program`
	}
	return fmt.Sprintf("%d", icp.GetOutput())
}

func (d day9) SolvePart2() string {
	icp := internal.NewIntCodeProgram(d.program, 2)
	if status := icp.Run(); status != internal.StatusComplete {
		return `Unable to run program`
	}
	return fmt.Sprintf("%d", icp.GetOutput())
}
