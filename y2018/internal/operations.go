package internal

type Operation string

var (
	AllOps = []Operation{
		addr,
		addi,
		mulr,
		muli,
		banr,
		bani,
		borr,
		bori,
		setr,
		seti,
		gtir,
		gtri,
		gtrr,
		eqir,
		eqri,
		eqrr,
	}
)

const (
	addr = Operation(`addr`)
	addi = Operation(`addi`)
	mulr = Operation(`mulr`)
	muli = Operation(`muli`)
	banr = Operation(`banr`)
	bani = Operation(`bani`)
	borr = Operation(`borr`)
	bori = Operation(`bori`)
	setr = Operation(`setr`)
	seti = Operation(`seti`)
	gtir = Operation(`gtir`)
	gtri = Operation(`gtri`)
	gtrr = Operation(`gtrr`)
	eqir = Operation(`eqir`)
	eqri = Operation(`eqri`)
	eqrr = Operation(`eqrr`)
)
