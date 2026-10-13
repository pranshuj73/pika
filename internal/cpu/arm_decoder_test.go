package cpu

import "testing"

// assembleDP builds a data-processing instruction word field by field.
// Each field is placed at its architectural bit position; anything left 0
// is "field absent". This mirrors how an assembler encodes instructions.
func assembleDP(cond, op, rn, rd uint8, i, s bool, operand2 uint32) uint32 {
	w := uint32(cond) << 28
	if i {
		w |= 1 << 25
	}
	w |= uint32(op) << 21
	if s {
		w |= 1 << 20
	}
	w |= uint32(rn) << 16
	w |= uint32(rd) << 12
	w |= operand2 & 0xFFF
	return w
}

func TestDecodeDPFields(t *testing.T) {
	// Every field gets a distinct nonzero value so a field swap can't hide.
	instr := assembleDP(0xE, 0xD, 3, 4, true, false, 42)
	dp := decodeDP(instr)

	if dp.Cond != 0xE {
		t.Errorf("Cond: got %#x, want 0xE", dp.Cond)
	}
	if !dp.I {
		t.Error("I: got false, want true (bit 25 set)")
	}
	if dp.Op != 0xD {
		t.Errorf("Op: got %#x, want 0xD (MOV)", dp.Op)
	}
	if dp.S {
		t.Error("S: got true, want false (bit 20 clear)")
	}
	if dp.Rn != 3 {
		t.Errorf("Rn: got %d, want 3", dp.Rn)
	}
	if dp.Rd != 4 {
		t.Errorf("Rd: got %d, want 4", dp.Rd)
	}
	if dp.Imm != 42 {
		t.Errorf("Imm: got %d, want 42 (rotate=0: operand is the raw 8-bit value)", dp.Imm)
	}
}

func TestDecodeDPImmediateRotation(t *testing.T) {
	// operand2 with rotate=1, value=0xFF: immediate is 0xFF rotated right by 2*1.
	instr := assembleDP(0xE, 0xD, 0, 4, true, false, (1<<8)|0xFF)
	dp := decodeDP(instr)
	want := uint32(0xFF)>>2 | (uint32(0xFF)&3)<<30 // 0xC000003F
	if dp.Imm != want {
		t.Errorf("Imm rotated: got %#x, want %#x (8-bit value 0xFF rotated right by 2)", dp.Imm, want)
	}
}

func TestDecodeDPFieldIsolation(t *testing.T) {
	// Bits of adjacent fields must not bleed into each other: set S, clear I.
	instr := assembleDP(0x0, 0x4, 0xF, 0xF, false, true, 0xFFF)
	dp := decodeDP(instr)
	if dp.I {
		t.Error("I: got true, want false — bit 25 bleeding in from a neighbor?")
	}
	if !dp.S {
		t.Error("S: got false, want true (bit 20 set)")
	}
	if dp.Rn != 0xF || dp.Rd != 0xF {
		t.Errorf("Rn/Rd: got %d/%d, want 15/15", dp.Rn, dp.Rd)
	}
}
