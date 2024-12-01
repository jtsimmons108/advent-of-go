package days

import (
	"fmt"

	"gonum.org/v1/gonum/stat/combin"
	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/internal"
)

type day7 struct {
	program      []int64
	permutations [][]int
}

func Day7() utils.Day {
	d := day7{}
	input := utils.DayInput(2019, 7)
	d.program = utils.ConvertInputToBigIntSlice(input, ",")
	d.permutations = combin.Permutations(5, 5)

	return d
}

func (d day7) SolvePart1() string {
	var max int64 = 0

	for _, phaseSettings := range d.permutations {
		val := d.RunThrusters(phaseSettings)
		if val > max {
			max = val
		}
	}
	return fmt.Sprintf("%d", max)
}

func (d day7) SolvePart2() string {
	var max int64 = 0

	for _, phaseSettings := range d.permutations {
		val := d.RunContinualThrusters(phaseSettings)
		if val > max {
			max = val
		}
	}
	return fmt.Sprintf("%d", max)
}

func (d day7) RunThrusters(phaseSettings []int) int64 {
	var (
		input int64
	)

	input = 0

	for _, phase := range phaseSettings {
		icp := internal.NewIntCodeProgram(d.program, int64(phase), input)
		if status := icp.Run(); status != internal.StatusComplete {
			panic(`Unable to complete program`)
		}
		input = icp.GetOutput()
	}

	return input
}

func (d day7) RunContinualThrusters(phaseSettings []int) int64 {

	inputs := []int64{0}
	amplifiers := []*internal.IntCodeProgram{}

	step := 0

	for i := range 5 {
		amplifiers = append(amplifiers, internal.NewIntCodeProgram(d.program, int64(phaseSettings[i]+5)).WithIOPauses())
	}

	finished := false

	for !finished {
		amp := amplifiers[step%5]

		status := amp.Run()
		switch status {
		case internal.StatusInput:
			if len(inputs) == 0 {
				panic(`Trying to input with no inputs`)
			}
			amp.AddInput(inputs[0])
			inputs = inputs[1:]
		case internal.StatusOutput:
			if len(inputs) != 0 {
				panic(`Stacking inputs up, not expected`)
			}
			inputs = append(inputs, amp.GetOutput())
			step++

		case internal.StatusComplete:
			if step%5 == 4 {
				finished = true
			}
			step++

		case internal.StatusFault:
			panic(`Received fault during run`)
		default:
			panic(`Unknown status received`)
		}
	}

	if len(inputs) != 1 {
		panic(`Only expected one output at finish`)
	}

	return inputs[0]
}
