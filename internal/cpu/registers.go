package cpu

type (
	Registers [16]uint32
	CPSR      uint32
)

const (
	flagN    CPSR = 1 << 31
	flagZ    CPSR = 1 << 30
	flagC    CPSR = 1 << 29
	flagV    CPSR = 1 << 28
	flagT    CPSR = 1 << 5
	modeMosk        = 0x1F
)

func (c CPSR) Has(f CPSR) bool {
	return c&f != 0
}

func (c *CPSR) Set(f CPSR) {
	*c |= f
}

func (c *CPSR) Clear(f CPSR) {
	*c &^= f
}

func (c *CPSR) Mode() uint8 {
	return uint8(*c & modeMosk)
}

func (c *CPSR) IsThumb() bool {
	return c.Has(flagT)
}

func (c CPSR) ConditionPassed(cond uint8) bool {
	nF, zF, cF, vF := c.Has(flagN), c.Has(flagZ), c.Has(flagC), c.Has(flagV)
	switch cond {
	case 0x0:
		return zF
	case 0x1:
		return !zF
	case 0x2:
		return cF
	case 0x3:
		return !cF
	case 0x4:
		return nF
	case 0x5:
		return !nF
	case 0x6:
		return vF
	case 0x7:
		return !vF
	case 0x8:
		return cF && !zF
	case 0x9:
		return !cF || zF
	case 0xA:
		return nF == vF
	case 0xB:
		return nF != vF
	case 0xC:
		return !zF && nF == vF
	case 0xD:
		return zF || nF != vF
	case 0xE:
		return true
	default:
		return false
	}
}
