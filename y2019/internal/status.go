package internal

type Status int

const (
	StatusComplete = Status(0)
	StatusContinue = Status(1)
	StatusFault    = Status(2)
	StatusInput    = Status(3)
	StatusOutput   = Status(4)
)
