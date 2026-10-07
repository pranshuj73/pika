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
