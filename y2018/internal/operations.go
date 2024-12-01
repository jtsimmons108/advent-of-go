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

	MappedOperations = map[int]Operation{
		0:  bori,
		1:  borr,
		2:  seti,
		3:  mulr,
		4:  setr,
		5:  addr,
		6:  gtir,
		7:  eqir,
		8:  gtri,
		9:  bani,
		10: muli,
		11: gtrr,
		12: banr,
		13: eqri,
		14: addi,
		15: eqrr,
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
