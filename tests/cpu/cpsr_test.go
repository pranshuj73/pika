package cpu_test

import (
	"testing"

	"github.com/pranshuj73/pika.git/internal/cpu"
)

// Bit masks for the CPSR flags. The cpu package's flagN/flagZ/... constants
// are unexported, so an external test package uses the literals directly.
const (
	flagN = 1 << 31
	flagZ = 1 << 30
	flagC = 1 << 29
	flagV = 1 << 28
)

// Truth table per GBATEK: expected result of ConditionPassed(cond)
// for every combination of the four flags N,Z,C,V.
// Rows: bitmask of flags set (bit0=N, bit1=Z, bit2=C, bit3=V).
func TestConditionPassedTruthTable(t *testing.T) {
	for mask := 0; mask < 16; mask++ {
		var c cpu.CPSR
		if mask&1 != 0 {
			c.Set(cpu.CPSR(flagN))
		}
		if mask&2 != 0 {
			c.Set(cpu.CPSR(flagZ))
		}
		if mask&4 != 0 {
			c.Set(cpu.CPSR(flagC))
		}
		if mask&8 != 0 {
			c.Set(cpu.CPSR(flagV))
		}
		n := mask&1 != 0
		z := mask&2 != 0
		cf := mask&4 != 0
		v := mask&8 != 0

		expected := map[uint8]bool{
			0x0: z,            // EQ
			0x1: !z,           // NE
			0x2: cf,           // CS
			0x3: !cf,          // CC
			0x4: n,            // MI
			0x5: !n,           // PL
			0x6: v,            // VS
			0x7: !v,           // VC
			0x8: cf && !z,     // HI
			0x9: !cf || z,     // LS
			0xA: n == v,       // GE
			0xB: n != v,       // LT
			0xC: !z && n == v, // GT
			0xD: z || n != v,  // LE
			0xE: true,         // AL
			0xF: false,        // reserved, never
		}

		for cond, want := range expected {
			if got := c.ConditionPassed(cond); got != want {
				t.Errorf("flags N=%v Z=%v C=%v V=%v, cond %#x: got %v, want %v",
					n, z, cf, v, cond, got, want)
			}
		}
	}
}

// Extraction move from lesson 2 §7: a word whose top hex digit is D
// carries condition D (13) in bits 31-28.
func TestCondExtraction(t *testing.T) {
	for _, cond := range []uint8{0, 2, 7, 0xD, 0xE} {
		instr := uint32(cond) << 28 // cond in the top four bits, rest zero
		got := uint8(instr >> 28)
		if got != cond {
			t.Errorf("extraction: put cond %#x in bits 31-28, got %#x back", cond, got)
		}
	}
}
