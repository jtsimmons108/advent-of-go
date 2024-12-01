package days

import (
	"fmt"
	"strconv"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/internal"
)

type Tile int

const (
	Empty  = Tile(0)
	Wall   = Tile(1)
	Block  = Tile(2)
	Paddle = Tile(3)
	Ball   = Tile(4)
)

type day13 struct {
	program []int64
}

func Day13() utils.Day {
	d := day13{}

	input := utils.DayInput(2019, 13)
	d.program = utils.ConvertInputToBigIntSlice(input, ",")
	return d
}

func (d day13) SolvePart1() string {

	icp := internal.NewIntCodeProgram(d.program)
	if status := icp.Run(); status != internal.StatusComplete {
		panic(`Program did not run to completion`)
	}

	total := 0
	for i := 2; i < len(icp.Outputs); i += 3 {
		if Tile(icp.Outputs[i]) == Block {
			total++
		}
	}

	return strconv.Itoa(total)
}

func (d day13) SolvePart2() string {

	type outputInfo struct {
		xPos     int64
		val      int64
		complete bool
	}

	var (
		score       int64
		ballX       int64
		paddleX     int64
		outputCount = 0
		info        = outputInfo{}
		status      = internal.StatusContinue
	)

	d.program[0] = 2

	icp := internal.NewIntCodeProgram(d.program).WithIOPauses()
	for status != internal.StatusComplete {
		status = icp.Run()

		switch status {
		case internal.StatusInput:
			var input int64
			switch {
			case ballX < paddleX:
				input = -1
			case ballX == paddleX:
				input = 0
			case ballX > paddleX:
				input = 1
			default:
				panic(`This should never happen`)
			}
			icp.AddInput(input)

		case internal.StatusOutput:
			outputCount++
			val := icp.GetOutput()
			switch outputCount % 3 {
			case 0:
				info.val = val
				info.complete = true
			case 1:
				info.xPos = val
			}

			if info.complete {
				if info.xPos == -1 {
					score = info.val
					continue
				}

				switch Tile(info.val) {
				case Ball:
					ballX = info.xPos
				case Paddle:
					paddleX = info.xPos
				}

				info = outputInfo{}
			}

		case internal.StatusFault:
			panic(`Unexpected fault in program`)
		}
	}

	return fmt.Sprintf("%d", score)

}
