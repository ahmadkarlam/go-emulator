package main

type Bus struct {
	ram map[uint16]uint8
}

func NewBus() *Bus {
	return &Bus{
		ram: make(map[uint16]uint8),
	}
}

func (d *Bus) Write(address uint16, data uint8) {
	d.ram[address] = data
}

func (d *Bus) Read(address uint16) uint8 {
	if val, ok := d.ram[address]; ok {
		return val
	}

	return 0x00
}
