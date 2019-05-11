package main

import (
	"fmt"
)

// Registers is a struct which defines the CPU registers
type Registers struct {
	a uint8
	b uint8
	c uint8
	d uint8
	e uint8
	f uint8 // flags register
	h uint8
	l uint8
}

// Instructions
// Add:
func (r *Registers) add(a *uint8, b *uint8) {
	x := *a
	y := *b
	v := x + y
	fmt.Println(v)
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
func convFlagToUInt8(r flagsRegister) uint8 {
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
	af := uint16(r.a)<<8 | uint16(r.f)
	return af

}

// setAF takes a the register struct and turns the
// input uint16 value into x 2 uint8 values
func (r *Registers) setAF(value uint16) {
	r.a = uint8((value & 0xFF00) >> 8)
	r.f = uint8(value & 0xFF)
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

// setHL takes a the register struct and turns the
// input uint16 value into x 2 uint8 values
func (r *Registers) setHL(value uint16) {
	r.h = uint8((value & 0xFF00) >> 8)
	r.l = uint8(value & 0xFF)
}

func main() {
	wow := Registers{1, 9, 9, 4, 5, 6, 7, 8}
	wow.setBC(2314)

	flag := flagsRegister{false, true, true, true}
	f := convFlagToUInt8(flag)
	flag.convUInttoFlag(112)
	fmt.Println(flag)
	fmt.Println(f) // should be 240 - aka 0b11110000
	var g uint8 = 9
	wow.add(&g, &wow.b)

}
