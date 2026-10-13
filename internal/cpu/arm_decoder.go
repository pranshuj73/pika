package cpu

// DPInstr is one data-processing instruction, split into its fields.
type DPInstr struct {
    Cond uint8   // TODO: bits 31-28 — a number 0-15
    I    bool    // TODO: bit 25 — true when operand2 is an immediate
    Op   uint8   // TODO: bits 24-21 — a number 0-15
    S    bool    // TODO: bit 20
    Rn   uint8   // TODO: bits 19-16 — register INDEX
    Rd   uint8   // TODO: bits 15-12 — register INDEX
    Imm  uint32  // TODO: when I=1, the rotated immediate value; 0 otherwise
}

// decodeDP splits a 32-bit ARM instruction word into DPInstr fields.
// TODO: extract each field with the shift-then-mask move (lesson 2, §4):
//       shift right by the field's LOW bit number, mask with as many 1s as the field is wide.
// TODO: for the I bit and S bit, "field is 1 bit wide" — after shifting,
//       compare against 1 to decide true/false.
// TODO: when I is set, split the low 12 bits: bottom 8 = value, top 4 = rotate;
//       the immediate is the value rotated RIGHT by (2 * rotate). Start with
//       rotate=0 meaning "the value as-is" and add rotation once MOV works.
func decodeDP(instr uint32) DPInstr {
	
}
