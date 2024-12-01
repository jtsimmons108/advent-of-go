package utils2019_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"simmons.com/advent-of-go/y2019/utils2019"
)

func TestPart1(t *testing.T) {

	t.Run(`TestGetFuelFromMass`, func(t *testing.T) {
		assert.Equal(t, 2, utils2019.GetFuelFromMass(12))
		assert.Equal(t, 2, utils2019.GetFuelFromMass(14))
		assert.Equal(t, 654, utils2019.GetFuelFromMass(1969))
		assert.Equal(t, 33583, utils2019.GetFuelFromMass(100756))
	})

	t.Run(`TestGetTotalFuelFromMass`, func(t *testing.T) {
		assert.Equal(t, 2, utils2019.GetTotalFuelFromMass(14))
		assert.Equal(t, 966, utils2019.GetTotalFuelFromMass(1969))
		assert.Equal(t, 50346, utils2019.GetTotalFuelFromMass(100756))
	})
}
