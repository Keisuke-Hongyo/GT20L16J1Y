package main

import (
	"GT20L16J1Y/GT20L16J1Y"
	"GT20L16J1Y/OLED"
	"machine"
	"tinygo.org/x/drivers/ssd1306"
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

	_ = machine.I2C0.Configure(machine.I2CConfig{
		Frequency: 400 * machine.KHz,
	})
	dev := ssd1306.NewI2C(machine.I2C0)

	dev.Configure(ssd1306.Config{
		Address: 0x3C,
		Width:   128,
		Height:  64,
	})

	dev.ClearBuffer()
	dev.ClearDisplay()

	//font library init
	display := OLED.NewDisplay(dev, gt)

	//s := "テストプログラム Ver1"
	for {
		// led.Low()
		// time.Sleep(100 * time.Millisecond)
		// led.High()
		// time.Sleep(100 * time.Millisecond)

		display.LcdPrint(0, 0, "本行 圭介")
		display.LcdPrint(0, 0+18, "本行 駿介")

		//gt.PrintTerminal(gt.ReadFonts(s))
	}
}
