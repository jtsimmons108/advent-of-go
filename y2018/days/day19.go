package days

import (
	"fmt"
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2018/internal"
)

type day19 struct {
	ip         int
	ipRegister int
	ops        []string
	values     [][]int
}

func Day19() utils.Day {
	d := day19{}
	input := utils.ConvertInputToStringSlice(utils.DayInput(2018, 19), "\n")
	d.ipRegister = utils.ExtractInts(input[0], false)[0]

	for _, line := range input[1:] {
		d.ops = append(d.ops, strings.Split(line, " ")[0])
		d.values = append(d.values, utils.ExtractInts(line, false))
	}

	return d
}

func (d day19) SolvePart1() string {

	device := internal.NewDevice(6)
	executions := 0

	for d.ip < len(d.ops) {
		device.Registers[d.ipRegister] = d.ip
		device.Operate(internal.Operation(d.ops[d.ip]), d.values[d.ip][0], d.values[d.ip][1], d.values[d.ip][2])
		d.ip = device.Registers[d.ipRegister] + 1
		executions++
	}
	return strconv.Itoa(device.Registers[0])
}

func (d day19) SolvePart2() string {

	device := internal.NewDevice(6)
	device.Registers[0] = 1
	executions := 0

	for d.ip < len(d.ops) {
		device.Registers[d.ipRegister] = d.ip
		device.Operate(internal.Operation(d.ops[d.ip]), d.values[d.ip][0], d.values[d.ip][1], d.values[d.ip][2])
		d.ip = device.Registers[d.ipRegister] + 1
		executions++
	}
	fmt.Printf("Executed %d operations", executions)
	return strconv.Itoa(device.Registers[0])

}
