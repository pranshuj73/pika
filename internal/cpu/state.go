package cpu

type InstructionState uint8

const (
	ARM InstructionState = iota
	THUMB
)

type Mode uint8

const (
	User Mode = iota
	FIQ
	IRQ
	Supervisor
	Abort
	Undefined
	System
)
