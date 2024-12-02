package internal

import (
	"simmons.com/advent-of-go/mathutils"
)

func LevelsAreSafe(levels []int) bool {
	isDecreasing := levels[0] > levels[1]

	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i+1] - levels[i]
		abs := mathutils.Abs(diff)
		switch {
		case isDecreasing && diff > 0:
			return false
		case !isDecreasing && diff < 0:
			return false
		case abs < 1 || abs > 3:
			return false
		default:
			// no-op
		}
	}

	return true
}

func LevelsAreSafeWithDampener(levels []int) bool {
	if LevelsAreSafe(levels) {
		return true
	}

	for i := range levels {
		newLevels := append([]int{}, levels[:i]...)
		newLevels = append(newLevels, levels[i+1:]...)
		if LevelsAreSafe(newLevels) {
			return true
		}
	}

	return false
}
