package utils2019

import (
	"fmt"
	"strconv"

	"simmons.com/advent-of-go/utils"
)

func GetVisitedPoints(instructions ...string) (map[utils.Point]int, error) {
	visited := map[utils.Point]int{}
	current := utils.Origin
	step := 0

	for _, ins := range instructions {

		var d utils.Delta

		switch ins[0] {
		case 'R':
			d = utils.Delta{Dx: 1, Dy: 0}
		case 'U':
			d = utils.Delta{Dx: 0, Dy: 1}
		case 'L':
			d = utils.Delta{Dx: -1, Dy: 0}
		case 'D':
			d = utils.Delta{Dx: 0, Dy: -1}
		default:
			return nil, fmt.Errorf(`unrecognized direction`)
		}

		distance, err := strconv.Atoi(ins[1:])
		utils.CheckError(err, `Unable to get distance`)

		for distance > 0 {
			current = current.NextPoint(d)
			step++
			if _, found := visited[current]; !found {
				visited[current] = step
			}
			distance--
		}
	}

	return visited, nil
}
