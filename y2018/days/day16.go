package days

import (
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2018/internal"
)

type day16 struct {
	device          *internal.Device
	instructions    []string
	possibleMatches map[int]map[internal.Operation]struct{}
}

func Day16() utils.Day {
	d := day16{}
	d.device = internal.NewDevice(4)
	d.instructions = utils.ConvertInputToStringSlice(utils.DayInput(2018, 16), "\n")
	d.possibleMatches = map[int]map[internal.Operation]struct{}{}
	return d
}

func (d day16) SolvePart1() string {

	total := 0
	for i := 0; strings.HasPrefix(d.instructions[i], "Before"); i += 4 {
		before := utils.ExtractInts(d.instructions[i], false)
		ops := utils.ExtractInts(d.instructions[i+1], false)
		after := utils.ExtractInts(d.instructions[i+2], false)
		count := 0
		for _, op := range internal.AllOps {
			d.device.SetRegisters(append([]int{}, before...))
			d.device.Operate(op, ops[1], ops[2], ops[3])
			if d.device.RegistersEqual(after) {
				count++
				if _, found := d.possibleMatches[ops[0]]; !found {
					d.possibleMatches[ops[0]] = map[internal.Operation]struct{}{}
				}
				d.possibleMatches[ops[0]][op] = struct{}{}
			}
		}

		if count >= 3 {
			total++
		}

	}
	return strconv.Itoa(total)
}

func (d day16) SolvePart2() string {

	changed := true
	matched := map[int]internal.Operation{}
	for changed {
		changed = false
		for opNum, m := range d.possibleMatches {
			if len(m) == 1 {
				var op internal.Operation
				for o := range m {
					op = o
				}
				if _, found := matched[opNum]; found {
					continue
				}
				matched[opNum] = op
				changed = true

				for otherOpNum, otherM := range d.possibleMatches {
					if opNum != otherOpNum {
						delete(otherM, op)
					}
				}
			}
		}
	}

	i := 0
	for ; strings.HasPrefix(d.instructions[i], "Before"); i += 4 {
	}

	i += 2
	device := internal.NewDevice(4)

	for ; i < len(d.instructions); i++ {
		line := utils.ExtractInts(d.instructions[i], false)
		op, found := matched[line[0]]
		if !found {
			panic(`unable to find operation`)
		}
		device.Operate(op, line[1], line[2], line[3])
	}

	return strconv.Itoa(device.Registers[0])
}
