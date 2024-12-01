package utils

import "simmons.com/advent-of-go/utils/internal"

var (
	// Errors
	CheckError = internal.CheckError

	// Input
	DayInput                  = internal.DayInput
	ConvertInputToStringSlice = internal.ConvertInputToStringSlice
	ConvertInputToIntSlice    = internal.ConvertInputToIntSlice
	ConvertInputToBigIntSlice = internal.ConvertInputToBigIntSlice
	ExtractInts               = internal.ExtractInts

	Origin = Point{X: 0, Y: 0}
)

type (
	Day = internal.Day

	Point = internal.Point
	Delta = internal.Delta
)
