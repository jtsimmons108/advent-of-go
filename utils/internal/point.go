package internal

import (
	"math"

	"simmons.com/advent-of-go/mathutils"
)

type Delta struct {
	Dx int64
	Dy int64
}

type Point struct {
	X int64
	Y int64
}

var (
	Origin = Point{
		X: 0,
		Y: 0,
	}
)

func NewPoint(x int64, y int64) Point {
	return Point{
		X: x,
		Y: y,
	}
}
func (p Point) NextPoint(d Delta) Point {
	return Point{
		X: p.X + d.Dx,
		Y: p.Y + d.Dy,
	}
}

func (p Point) ManhattanDistance(o Point) int64 {
	return mathutils.Abs64(p.X-o.X) + mathutils.Abs64(p.Y-o.Y)
}

func (p Point) Distance(o Point) float64 {
	return math.Sqrt(math.Pow(float64(p.X-o.X), 2) + math.Pow(float64(p.Y-o.Y), 2))
}
