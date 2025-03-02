package main

// This is the most minimal blinky example and should run almost everywhere.

import (
	"GT20L16J1Y/GT20L16J1Y"
	"machine"
	"time"
)

func main() {
	var err error

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.High()

	spi := machine.SPI0
	err = spi.Configure(machine.SPIConfig{})
	if err != nil {
		panic(err)
	}
	csn := machine.D7 // Digital Input	SPI Chip Select
	gt := GT20L16J1Y.New(&spi, &csn)
	gt.Initialize()

	s := "テストプログラム Ver1"
	for {
		led.Low()
		time.Sleep(100 * time.Millisecond)
		led.High()
		time.Sleep(100 * time.Millisecond)
		//gt.ReadFonts(s)
		gt.PrintTerminal(gt.ReadFonts(s))
	}
}
