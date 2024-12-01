package internal

type OpCode int

const (
	ADD    = OpCode(1)
	MUL    = OpCode(2)
	INPUT  = OpCode(3)
	OUTPUT = OpCode(4)
	JIT    = OpCode(5)
	JIF    = OpCode(6)
	LT     = OpCode(7)
	EQ     = OpCode(8)
	OFFSET = OpCode(9)
	HALT   = OpCode(99)
)
