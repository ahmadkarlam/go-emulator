package main

type Flag uint8

const (
	C Flag = 1 << 0 // Carry Bit
	Z Flag = 1 << 1 // Zero
	I Flag = 1 << 2 // Disable Interrupts
	D Flag = 1 << 3 // Decimal Mode (unused in this implementation)
	B Flag = 1 << 4 // Break
	U Flag = 1 << 5 // Unused
	V Flag = 1 << 6 // Overflow
	N Flag = 1 << 7 // Negative
)

// CPU 6502
type CPU struct {
	bus *Bus

	Accumulator    uint8
	XRegister      uint8
	YRegister      uint8
	StackPointer   uint8
	ProgramCounter uint16
	Status         Flag

	Fetched uint8
}

func NewCPU(bus *Bus) *CPU {
	return &CPU{
		bus:            bus,
		Accumulator:    0x00,
		XRegister:      0x00,
		YRegister:      0x00,
		StackPointer:   0x00,
		ProgramCounter: 0x0000,
		Status:         0x00,
		Fetched:         0x00,
	}
}

func (c *CPU) Write(address uint16, data uint8) {
	c.bus.Write(address, data)
}

func (c *CPU) Read(address uint16) uint8 {
	return c.bus.Read(address)
}

func (c *CPU) GetFlag(flag Flag) uint8 {
	result := c.Status & flag
	if result > 0 {
		return 1
	}

	return 0
}

func (c *CPU) SetFlag(flag Flag, value uint8) {
	if value > 0 {
		c.Status |= flag
	} else {
		c.Status &= ^flag
	}
}

func (c *CPU) GetFetched() uint8 {
	return c.Fetched
}
