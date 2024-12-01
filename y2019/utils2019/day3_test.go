package utils2019_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"simmons.com/advent-of-go/utils"
	"simmons.com/advent-of-go/y2019/utils2019"
)

func TestGetVistedPointsFromPath(t *testing.T) {

	t.Run(`NoPoints`, func(t *testing.T) {

		visited, err := utils2019.GetVisitedPoints()
		require.NoError(t, err, `Could not compute with no points`)
		assert.Empty(t, visited)
	})

	t.Run(`OnePoint`, func(t *testing.T) {

		visited, err := utils2019.GetVisitedPoints(`R1`)
		require.NoError(t, err, `Could not compute`)
		assert.Len(t, visited, 1)
		assert.Contains(t, visited, utils.Point{X: 1, Y: 0})

	})

	t.Run(`FourPoints`, func(t *testing.T) {

		visited, err := utils2019.GetVisitedPoints(`R1`, `U1`, `L1`, `D1`)
		require.NoError(t, err, `Could not compute`)
		assert.Len(t, visited, 4)
		assert.Contains(t, visited, utils.Point{X: 1, Y: 0})
		assert.Contains(t, visited, utils.Point{X: 1, Y: 1})
		assert.Contains(t, visited, utils.Point{X: 0, Y: 1})
		assert.Contains(t, visited, utils.Point{X: 0, Y: 0})
	})
}
