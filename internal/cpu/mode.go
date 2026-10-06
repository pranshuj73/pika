package cpu

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

