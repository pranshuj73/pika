// Package cpu contains the core cpu logic for the emulator
package cpu

import (
	"fmt"
)

type CPU struct {
	registers *Registers
	cpsr *CPSR
}

func cpu() {
	fmt.Println("idk what to even implement here (yet)")
}
