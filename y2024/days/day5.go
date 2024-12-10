package days

import (
	"fmt"
	"strconv"

	"simmons.com/advent-of-go/utils"
)

type day5 struct {
	orderingRules map[int][]int
	inverseRules  map[int][]int
	printList     [][]int
}

func Day5() utils.Day {
	d := day5{}

	d.orderingRules = make(map[int][]int)
	d.inverseRules = make(map[int][]int)

	splitInput := utils.ConvertInputToStringSlice(utils.DayInput(2024, 5), "\n\n")
	rules := utils.ConvertInputToStringSlice(splitInput[0], "\n")
	prints := utils.ConvertInputToStringSlice(splitInput[1], "\n")

	for _, rule := range rules {
		nums := utils.ExtractInts(rule, false)
		d.orderingRules[nums[0]] = append(d.orderingRules[nums[0]], nums[1])
		d.inverseRules[nums[1]] = append(d.inverseRules[nums[1]], nums[0])
	}

	for _, print := range prints {
		d.printList = append(d.printList, utils.ConvertInputToIntSlice(print, ","))
	}

	return d
}

func (d day5) SolvePart1() string {
	total := 0
	for _, print := range d.printList {

		if d.isValidRule(print) {
			middle := print[len(print)/2]
			total += middle
		}
	}
	return strconv.Itoa(total)
}

func (d day5) SolvePart2() string {
	total := 0
	for _, print := range d.printList {

		if d.isValidRule(print) {
			total += print[len(print)/2]
		} else {
			fmt.Printf("Need to reorder: %v\n", print)
			d.reorder(print)
			// fixed := d.reorder(print)
			// total += fixed[len(fixed)/2]

		}
	}
	return strconv.Itoa(total)
}

func (d day5) isValidRule(print []int) bool {
	// Build index map
	indices := make(map[int]int)
	for i, num := range print {
		indices[num] = i
	}

	// Check ordering rules
	for num, nextNums := range d.orderingRules {
		for _, nextNum := range nextNums {
			first, firstFound := indices[num]
			second, secondFound := indices[nextNum]
			if firstFound && secondFound && first > second {
				return false
			}
		}
	}

	return true
}

func (d day5) reorder(print []int) []int {
	// Build index map
	indices := make(map[int]int)
	for i, num := range print {
		indices[num] = i
	}

	return nil

}
