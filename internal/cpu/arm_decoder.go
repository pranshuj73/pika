package cpu

import "math/bits"

// DPInstr is one data-processing instruction, split into its fields.
type DPInstr struct {
    Cond uint8
    I    bool
    Op   uint8
    S    bool
    Rn   uint8
    Rd   uint8
    Imm  uint32
}

func decodeDP(instr uint32) DPInstr {
	var res DPInstr
	res.Cond = uint8(uint32(instr) >> 28 & 0xF)
	res.I = bool(uint32(instr) >> 25 & 0x1 != 0)
	res.Op = uint8(uint32(instr) >> 21 & 0xF)
	res.S = bool(uint32(instr) >> 20 & 0x1 != 0)
	res.Rn = uint8(uint32(instr) >> 16 & 0xF)
	res.Rd = uint8(uint32(instr) >> 12 & 0xF)

	if res.I {
		value := uint32(uint32(instr) & 0xFF)
		rotate := int(uint32(instr) >> 8 & 0xF) * -2
		res.Imm = uint32(bits.RotateLeft32(value, rotate))
	} else {
		res.Imm = 0
	}

	return res
}
