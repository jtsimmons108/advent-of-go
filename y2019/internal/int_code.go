package internal

import (
	"fmt"
)

type parameter struct {
	value int64
	mode  ParameterMode
}

type IntCodeProgram struct {
	//internals
	debug        bool
	inputQueue   []int64
	pointer      int64
	pauseForIO   bool
	relativeBase int64

	Program []int64
	Outputs []int64
}

func NewIntCodeProgram(instructions []int64, inputQueue ...int64) *IntCodeProgram {
	program := append([]int64{}, instructions...)
	program = append(program, make([]int64, 100000)...)

	ip := IntCodeProgram{
		pointer:      0,
		relativeBase: 0,
		inputQueue:   inputQueue,

		Program: program,
		Outputs: []int64{},
	}

	return &ip
}

func (ip *IntCodeProgram) DebugMode() *IntCodeProgram {
	ip.debug = true
	return ip
}

func (ip *IntCodeProgram) WithIOPauses() *IntCodeProgram {
	ip.pauseForIO = true
	return ip
}

func (ip *IntCodeProgram) AddInput(in int64) {
	ip.inputQueue = append(ip.inputQueue, in)
}

func (ip *IntCodeProgram) GetOutput() int64 {
	if len(ip.Outputs) == 0 {
		panic(`Trying to access output with no outputs`)
	}

	out := ip.Outputs[0]
	ip.Outputs = ip.Outputs[1:]

	return out
}

func (ip *IntCodeProgram) Run() (status Status) {

	status = StatusContinue

	for status == StatusContinue {
		status = ip.step()
		if status == StatusFault {
			return
		}
	}

	return
}

func (ip *IntCodeProgram) step() Status {

	op := OpCode(ip.Program[ip.pointer] % 100)

	switch op {
	case ADD:
		ip.opAdd()
	case MUL:
		ip.opMul()
	case INPUT:
		if ip.pauseForIO && len(ip.inputQueue) == 0 {
			return StatusInput
		}
		ip.opInput()
	case OUTPUT:
		ip.opOutput()

		if ip.pauseForIO {
			return StatusOutput
		}

	case JIT:
		ip.opJumpIfTrue()
	case JIF:
		ip.opJumpIfFalse()
	case LT:
		ip.opLessThan()
	case EQ:
		ip.opEqual()
	case OFFSET:
		ip.opOffset()
	case HALT:
		return StatusComplete
	default:
		return StatusFault
	}

	return StatusContinue

}

func (ip *IntCodeProgram) calcParams(length int) []parameter {
	params := []parameter{}

	opCode := ip.Program[ip.pointer]
	mode := opCode / 100

	if ip.debug {
		fmt.Printf("At pos [%d] with opCode [%d] -> mode [%d]\n", ip.pointer, opCode, mode)
	}
	i := 1

	for i < length {
		params = append(params, parameter{
			value: ip.Program[ip.pointer+int64(i)],
			mode:  ParameterMode(mode % 10),
		})
		mode /= 10
		i++
	}

	if ip.debug {
		fmt.Printf("Calculated params %v\n", params)
	}

	return params
}

func (ip *IntCodeProgram) calcParamValue(p parameter) int64 {
	val := p.value
	switch p.mode {
	case PositionMode:
		val = ip.Program[val]
	case RelativeMode:
		val = ip.Program[val+ip.relativeBase]
	default:
		// no-op
	}

	if ip.debug {
		fmt.Printf("Calculated value [%d] from %v\n", val, p)
	}

	return val
}

func (ip *IntCodeProgram) opAdd() {
	params := ip.calcParams(4)
	v1 := ip.calcParamValue(params[0])
	v2 := ip.calcParamValue(params[1])

	posMode := params[2].mode
	if posMode == ImmediateMode {
		panic(`Got ImmediateMode for position setter`)
	}

	pos := params[2].value
	if posMode == RelativeMode {
		pos += ip.relativeBase
	}

	if ip.debug {
		fmt.Printf("Taking [%d + %d] and assigning to pos [%d]\n", v1, v2, pos)
	}

	ip.Program[pos] = v1 + v2
	ip.pointer += 4
}

func (ip *IntCodeProgram) opMul() {
	params := ip.calcParams(4)
	v1 := ip.calcParamValue(params[0])
	v2 := ip.calcParamValue(params[1])

	posMode := params[2].mode
	if posMode == ImmediateMode {
		panic(`Got ImmediateMode for position setter`)
	}

	pos := params[2].value
	if posMode == RelativeMode {
		pos += ip.relativeBase
	}

	if ip.debug {
		fmt.Printf("Taking [%d * %d] and assigning to pos [%d]\n", v1, v2, pos)
	}

	ip.Program[pos] = v1 * v2
	ip.pointer += 4
}

func (ip *IntCodeProgram) opInput() {
	if len(ip.inputQueue) == 0 {
		panic(`Trying to input with no inputs in queue`)
	}
	params := ip.calcParams(2)
	mode := params[0].mode

	if mode == ImmediateMode {
		panic(fmt.Sprintf(`Got immediate mode for position setter [%d]`, mode))
	}

	pos := params[0].value
	if mode == RelativeMode {
		pos += ip.relativeBase
	}

	val := ip.inputQueue[0]
	ip.inputQueue = ip.inputQueue[1:]

	if ip.debug {
		fmt.Printf("Read [%d] from inputQueue and assigning to pos [%d]\n", val, pos)
	}

	ip.Program[pos] = val
	ip.pointer += 2
}

func (ip *IntCodeProgram) opOutput() {
	params := ip.calcParams(2)

	val := ip.calcParamValue(params[0])

	if ip.debug {
		fmt.Printf("Adding [%d] to output queue\n", val)
	}
	ip.Outputs = append(ip.Outputs, val)
	ip.pointer += 2
}

func (ip *IntCodeProgram) opJumpIfTrue() {
	params := ip.calcParams(3)

	val1 := ip.calcParamValue(params[0])
	val2 := ip.calcParamValue(params[1])

	if ip.debug {
		fmt.Printf("Jumping to pos [%d] if [%d] is non-zero\n", val2, val1)
	}

	if val1 > 0 {
		ip.pointer = val2
	} else {
		ip.pointer += 3
	}
}

func (ip *IntCodeProgram) opJumpIfFalse() {
	params := ip.calcParams(3)

	val1 := ip.calcParamValue(params[0])
	val2 := ip.calcParamValue(params[1])

	if ip.debug {
		fmt.Printf("Jumping to pos [%d] if [%d] is zero\n", val2, val1)
	}

	if val1 == 0 {
		ip.pointer = val2
	} else {
		ip.pointer += 3
	}
}

func (ip *IntCodeProgram) opLessThan() {
	params := ip.calcParams(4)

	val1 := ip.calcParamValue(params[0])
	val2 := ip.calcParamValue(params[1])

	posMode := params[2].mode
	if posMode == ImmediateMode {
		panic(`Got ImmediateMode for position setter`)
	}

	pos := params[2].value
	if posMode == RelativeMode {
		pos += ip.relativeBase
	}

	if ip.debug {
		fmt.Printf("Assigning result of [%d < %d] to pos [%d]\n", val1, val2, pos)
	}

	result := 0
	if val1 < val2 {
		result = 1
	}

	ip.Program[pos] = int64(result)
	ip.pointer += 4
}

func (ip *IntCodeProgram) opEqual() {
	params := ip.calcParams(4)

	val1 := ip.calcParamValue(params[0])
	val2 := ip.calcParamValue(params[1])

	posMode := params[2].mode
	if posMode == ImmediateMode {
		panic(`Got ImmediateMode for position setter`)
	}

	pos := params[2].value
	if posMode == RelativeMode {
		pos += ip.relativeBase
	}

	if ip.debug {
		fmt.Printf("Assigning result of [%d == %d] to pos [%d]\n", val1, val2, pos)
	}

	result := 0
	if val1 == val2 {
		result = 1
	}

	ip.Program[pos] = int64(result)
	ip.pointer += 4
}

func (ip *IntCodeProgram) opOffset() {
	params := ip.calcParams(2)

	val := ip.calcParamValue(params[0])

	if ip.debug {
		fmt.Printf("Offsetting relative base by [%d]\n", val)
	}

	ip.relativeBase += val
	ip.pointer += 2
}
