// Package cpu contains the core cpu logic for the emulator
package cpu

import (
	"fmt"

	"github.com/pranshuj73/pika.git/internal/memory"
)

type CPU struct {
	registers Registers
	cpsr CPSR
	instructionState InstructionState
	mode Mode
	memory *memory.Memory
}

func cpu() {
	fmt.Println("idk what to even implement here (yet)")
}
