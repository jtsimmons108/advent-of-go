package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"simmons.com/advent-of-go/y2024/internal"
)

func TestLevelsAreSafe(t *testing.T) {
	// test cases

	t.Run(`Part 1: LevelsAreSafe`, func(t *testing.T) {
		testCases := []struct {
			levels []int
			isSafe bool
		}{
			{[]int{7, 6, 4, 2, 1}, true},
			{[]int{1, 2, 7, 8, 9}, false},
			{[]int{9, 7, 6, 2, 1}, false},
			{[]int{1, 3, 2, 4, 5}, false},
			{[]int{8, 6, 4, 4, 1}, false},
			{[]int{1, 3, 6, 7, 9}, true},
		}

		// run tests
		for _, tc := range testCases {
			assert.True(t, internal.LevelsAreSafe(tc.levels) == tc.isSafe)
		}
	})

	t.Run(`Part 2: LevelsAreSafeWithDampener`, func(t *testing.T) {
		testCases := []struct {
			levels []int
			isSafe bool
		}{
			{[]int{7, 6, 4, 2, 1}, true},
			{[]int{1, 2, 7, 8, 9}, false},
			{[]int{9, 7, 6, 2, 1}, false},
			{[]int{1, 3, 2, 4, 5}, true},
			{[]int{8, 6, 4, 4, 1}, true},
			{[]int{1, 3, 6, 7, 9}, true},
		}

		// run tests
		for _, tc := range testCases {
			assert.True(t, internal.LevelsAreSafeWithDampener(tc.levels) == tc.isSafe)
		}
	})

}
