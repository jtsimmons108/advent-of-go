package internal

type ParameterMode int

const (
	PositionMode  = ParameterMode(0)
	ImmediateMode = ParameterMode(1)
	RelativeMode  = ParameterMode(2)
)
