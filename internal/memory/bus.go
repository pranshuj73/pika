package memory

type Memory struct {
	BIOS *BIOS
	IWRAM *IWRAM
	EWRAM *EWRAM
	// cartridge
}

func (m *Memory) Read(width int, addr uint32) uint32 {
	var val uint32

	return val
}
