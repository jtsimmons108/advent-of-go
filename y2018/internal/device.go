package internal

import (
	"errors"
)

type Device struct {
	Registers []int
}

func NewDevice() *Device {
	return &Device{
		Registers: make([]int, 4),
	}
}

func (d *Device) SetRegisters(regs []int) error {
	if len(regs) != 4 {
		return errors.New(`registers must have length of 4`)
	}

	d.Registers = regs
	return nil
}

func (d *Device) Addr(a, b, c int) {
	d.Registers[c] = d.Registers[a] + d.Registers[b]
}

func (d *Device) Addi(a, b, c int) {
	d.Registers[c] = d.Registers[a] + b
}

func (d *Device) Mulr(a, b, c int) {
	d.Registers[c] = d.Registers[a] * d.Registers[b]
}

func (d *Device) Muli(a, b, c int) {
	d.Registers[c] = d.Registers[a] * b
}

func (d *Device) Banr(a, b, c int) {
	d.Registers[c] = d.Registers[a] & d.Registers[b]
}

func (d *Device) Bani(a, b, c int) {
	d.Registers[c] = d.Registers[a] & b
}

func (d *Device) Borr(a, b, c int) {
	d.Registers[c] = d.Registers[a] | d.Registers[b]
}

func (d *Device) Bori(a, b, c int) {
	d.Registers[c] = d.Registers[a] | b
}

func (d *Device) Setr(a, b, c int) {
	d.Registers[c] = d.Registers[a]
}

func (d *Device) Seti(a, b, c int) {
	d.Registers[c] = a
}

func (d *Device) Gtir(a, b, c int) {
	val := 0
	if a > d.Registers[b] {
		val = 1
	}
	d.Registers[c] = val
}

func (d *Device) Gtri(a, b, c int) {
	val := 0
	if d.Registers[a] > b {
		val = 1
	}
	d.Registers[c] = val
}

func (d *Device) Gtrr(a, b, c int) {
	val := 0
	if d.Registers[a] > d.Registers[b] {
		val = 1
	}
	d.Registers[c] = val
}

func (d *Device) Eqir(a, b, c int) {
	val := 0
	if a == d.Registers[b] {
		val = 1
	}
	d.Registers[c] = val
}

func (d *Device) Eqri(a, b, c int) {
	val := 0
	if d.Registers[a] == b {
		val = 1
	}
	d.Registers[c] = val
}

func (d *Device) Eqrr(a, b, c int) {
	val := 0
	if d.Registers[a] == d.Registers[b] {
		val = 1
	}
	d.Registers[c] = val
}

func (d *Device) Operate(op Operation, a, b, c int) {
	switch op {
	case addr:
		d.Addr(a, b, c)
	case addi:
		d.Addi(a, b, c)
	case mulr:
		d.Mulr(a, b, c)
	case muli:
		d.Muli(a, b, c)
	case banr:
		d.Banr(a, b, c)
	case bani:
		d.Bani(a, b, c)
	case borr:
		d.Borr(a, b, c)
	case bori:
		d.Bori(a, b, c)
	case setr:
		d.Setr(a, b, c)
	case seti:
		d.Seti(a, b, c)
	case gtir:
		d.Gtir(a, b, c)
	case gtri:
		d.Gtri(a, b, c)
	case gtrr:
		d.Gtrr(a, b, c)
	case eqir:
		d.Eqir(a, b, c)
	case eqri:
		d.Eqri(a, b, c)
	case eqrr:
		d.Eqrr(a, b, c)
	default:
		panic(`Unrecognized operation`)
	}
}

func (d *Device) RegistersEqual(other []int) bool {
	if len(d.Registers) != len(other) {
		return false
	}

	for i := range d.Registers {
		if d.Registers[i] != other[i] {
			return false
		}
	}

	return true
}
