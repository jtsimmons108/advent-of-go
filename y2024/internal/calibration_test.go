package internal_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"simmons.com/advent-of-go/y2024/internal"
)

func TestCalibration(t *testing.T) {
	// test cases

	t.Run(`Part 1: EquationIsCalibrated`, func(t *testing.T) {
		testCases := []struct {
			target       int64
			nums         []int64
			IsCalibrated bool
		}{
			{190, []int64{10, 19}, true},
			{3267, []int64{81, 40, 27}, true},
			{83, []int64{17, 5}, false},
			{156, []int64{15, 6}, false},
			{7290, []int64{6, 8, 6, 15}, false},
			{161011, []int64{16, 10, 13}, false},
			{192, []int64{17, 8, 14}, false},
			{21037, []int64{9, 7, 18, 13}, false},
			{292, []int64{11, 6, 16, 20}, true},
		}

		// run tests
		for _, tc := range testCases {
			assert.True(t, internal.IsCalibrated(tc.target, tc.nums) == tc.IsCalibrated)
		}
	})

	t.Run(`Part 2: EquationIsCalibratedWithConcatenation`, func(t *testing.T) {
		testCases := []struct {
			target       int64
			nums         []int64
			IsCalibrated bool
		}{
			{190, []int64{10, 19}, true},
			{3267, []int64{81, 40, 27}, true},
			{83, []int64{17, 5}, false},
			{156, []int64{15, 6}, true},
			{7290, []int64{6, 8, 6, 15}, true},
			{161011, []int64{16, 10, 13}, false},
			{192, []int64{17, 8, 14}, true},
			{21037, []int64{9, 7, 18, 13}, false},
			{292, []int64{11, 6, 16, 20}, true},
		}

		// run tests
		for _, tc := range testCases {
			assert.True(t, internal.IsCalibratedWithConcatenation(tc.target, tc.nums) == tc.IsCalibrated)
		}
	})

}
