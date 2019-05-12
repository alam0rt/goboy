package main

import (
	"fmt"
)

var cpu CPU // declare the cpu

type CPU struct {
	registers Registers
	bus       memoryBus
	sp        uint16 // stack pointer
	pc        uint16 // program counter

}

type memoryBus struct {
	memory [0xFFFF]uint8 // 16bit memory range
}

// Registers is a struct which defines the CPU registers
type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	h uint8
	l uint8
	f flagsRegister // flags register
}

// Instruction set

// add does just that. Overflow is handled by
// checking if either operand is greater than
// the result, if it is, we set the carry flag

func (r *Registers) add(a *uint8, b *uint8) {
	if *a|*b > (*a + *b) {
		r.f.carry = true
	}
	// if result is 0, set zero flag
	if *a+*b == 0 {
		r.f.zero = true
	}
	// if the lower nibble of both operands
	// when added do not overflow into the
	// higher nibble, set flag
	r.f.halfCarry = (*a&0xF)+(*b&0xF) > 0xF

	*a = *a + *b
}

func (c *CPU) printRegisters() {
	fmt.Printf(`
a: %d
b: %d
c: %d
d: %d
e: %d
h: %d
l: %d
sp: %d
pc: %d
zero_flag: %t
subtract: %t
half_carry: %t
carry: %t
`,
		c.registers.a,
		c.registers.b,
		c.registers.c,
		c.registers.d,
		c.registers.e,
		c.registers.h,
		c.registers.l,
		c.sp,
		c.pc,
		c.registers.f.zero,
		c.registers.f.subtract,
		c.registers.f.halfCarry,
		c.registers.f.carry,
	)
}

// the below constants define the bit position
// of the flags in the flag register
const ZERO_FLAG_BYTE_POSITION uint8 = 7
const SUBTRACT_FLAG_BYTE_POSITION uint8 = 6
const HALF_CARRY_BYTE_POSITION uint8 = 5
const CARRY_FLAG_BYTE_POSITION uint8 = 4

// flagsRegister is a struct which makes sense
// of the f (flags) register
type flagsRegister struct {
	zero      bool
	subtract  bool
	halfCarry bool
	carry     bool
}

func (f *flagsRegister) convUInttoFlag(i uint8) {
	f.zero = ((i >> ZERO_FLAG_BYTE_POSITION & 1) != 0)
	f.subtract = ((i >> SUBTRACT_FLAG_BYTE_POSITION & 1) != 0)
	f.halfCarry = ((i >> HALF_CARRY_BYTE_POSITION & 1) != 0)
	f.carry = ((i >> CARRY_FLAG_BYTE_POSITION & 1) != 0)
}

// convFlagToUInt8 takes an f Flag struct and performs
// bitwise SHIFT, OR to set the top 4 bits to the correct
// values which represent flags. Lower nibble always
// zero.
// 1000 0000 = zero flag
// 0100 0000 = subtract flag
// 0010 0000 = half carry
// 0001 0000 = carry flag
func (r *flagsRegister) convFlagToUInt8() uint8 {
	var f uint8
	if r.zero {
		f = f | 1<<ZERO_FLAG_BYTE_POSITION
	}

	if r.subtract {
		f = f | 1<<SUBTRACT_FLAG_BYTE_POSITION
	}

	if r.halfCarry {
		f = f | 1<<HALF_CARRY_BYTE_POSITION
	}

	if r.carry {
		f = f | 1<<CARRY_FLAG_BYTE_POSITION
	}
	return f

}

// getAF reads the a, f registers and
// returns a combined "virtual" register
// of type uint16
func (r *Registers) getAF() uint16 {
	af := uint16(r.a)<<8 | uint16(r.f.convFlagToUInt8())
	return af

}

// setAF takes a the register struct and turns the
// input uint16 value into x 2 uint8 values
func (r *Registers) setAF(value uint16) {
	r.a = uint8((value & 0xFF00) >> 8)
	r.f.convUInttoFlag(uint8(value & 0xFF))
}

// getDE reads the d, e registers and
// returns a combined "virtual" register
// of type uint16
func (r *Registers) getDE() uint16 {
	de := uint16(r.d)<<8 | uint16(r.e)
	return de

}

// setDE takes a the register struct and turns the
// input uint16 value into x 2 uint8 values
func (r *Registers) setDE(value uint16) {
	r.d = uint8((value & 0xFF00) >> 8)
	r.e = uint8(value & 0xFF)
}

// getBC reads the b, c registers and
// returns a combined "virtual" register
// of type uint16
func (r *Registers) getBC() uint16 {
	bc := uint16(r.b)<<8 | uint16(r.c)
	return bc

}

// setBC takes a the register struct and turns the
// input uint16 value into x 2 uint8 values
func (r *Registers) setBC(value uint16) {
	r.b = uint8((value & 0xFF00) >> 8)
	r.c = uint8(value & 0xFF)
}

// getHL reads the h, l registers and
// returns a combined "virtual" register
// of type uint16
func (r *Registers) getHL() uint16 {
	hl := uint16(r.h)<<8 | uint16(r.l)
	return hl

}

// setSP takes a uint16 and sets it in the register
func (c *CPU) setSP(value uint16) {
	c.sp = value
}

// setPC takes a uint16 and sets it in the register
func (c *CPU) setPC(value uint16) {
	c.pc = value
}

// setHL takes a the register struct and turns the
// input uint16 value into x 2 uint8 values
func (r *Registers) setHL(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8(value & 0xFF)
}

func main() {
	cpu.registers.a = 255
	cpu.registers.c = 255
	cpu.registers.l = 255
	a := &cpu.registers.a
	l := &cpu.registers.l
	cpu.registers.f.convUInttoFlag(0)
	cpu.registers.add(l, a)

	cpu.printRegisters()

}
