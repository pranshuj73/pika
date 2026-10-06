package cpu

type InstructionState uint8

const (
	ARM InstructionState = iota
	THUMB
)
