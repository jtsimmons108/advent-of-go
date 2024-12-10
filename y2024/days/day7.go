package days

import (
	"fmt"
	"strconv"
	"strings"

	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2024/internal"
)

type equation struct {
	target int64
	nums   []int64
}

type day7 struct {
	equations []equation
}

func Day7() utils.Day {
	d := day7{}
	lines := utils.ConvertInputToStringSlice(utils.DayInput(2024, 7), "\n")

	for _, line := range lines {
		split := strings.Split(line, ":")
		target, err := strconv.ParseInt(split[0], 10, 64)
		utils.CheckError(err, "Error parsing target")
		ints := utils.ExtractInts(split[1], false)

		var nums []int64
		for _, i := range ints {
			nums = append(nums, int64(i))
		}

		d.equations = append(d.equations, equation{target, nums})
	}

	return d
}

func (d day7) SolvePart1() string {

	total := int64(0)

	for _, eq := range d.equations {
		if internal.IsCalibrated(eq.target, eq.nums) {
			total += eq.target
		}

	}
	return fmt.Sprintf("%d", total)
}

func (d day7) SolvePart2() string {

	total := int64(0)

	for _, eq := range d.equations {
		if internal.IsCalibratedWithConcatenation(eq.target, eq.nums) {
			total += eq.target
		}

	}
	return fmt.Sprintf("%d", total)
}
