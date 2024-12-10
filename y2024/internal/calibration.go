package internal

import (
	"fmt"
	"strconv"

	"simmons.com/advent-of-go/utils"
)

func IsCalibrated(target int64, nums []int64) bool {
	queue := []int64{nums[0]}

	for _, num := range nums[1:] {
		newQueue := []int64{}
		for _, q := range queue {
			newQueue = append(newQueue, q+num, q*num)
		}
		queue = newQueue
	}

	for _, q := range queue {
		if q == target {
			return true
		}
	}

	return false
}

func IsCalibratedWithConcatenation(target int64, nums []int64) bool {
	queue := []int64{nums[0]}

	for _, num := range nums[1:] {
		newQueue := []int64{}
		for _, q := range queue {
			newQueue = append(newQueue, q+num, q*num)
			concat, err := strconv.ParseInt(fmt.Sprintf("%d%d", q, num), 10, 64)
			utils.CheckError(err, "Error parsing concatenated number")
			newQueue = append(newQueue, concat)
		}
		queue = newQueue
	}

	for _, q := range queue {
		if q == target {
			return true
		}
	}

	return false
}
