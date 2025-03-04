package main

import (
	"GT20L16J1Y/GT20L16J1Y"
	"GT20L16J1Y/OLED"
	"fmt"
	"machine"
	"time"
	"tinygo.org/x/drivers/ssd1306"
)

func main() {
	var err error
	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	led.High()

	//font library init
	spi := machine.SPI0
	err = spi.Configure(machine.SPIConfig{})
	if err != nil {
		panic(err)
	}

	csn := machine.D7 // Digital Input	SPI Chip Select

	gt := GT20L16J1Y.New(&spi, &csn)

	gt.Initialize()

	// OLED初期化
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

	display := OLED.NewDisplay(dev, gt)

	display.LcdPrint(0, 0+0, "本行　駿介")
	display.LcdPrint(0, 0+20, "本行　すず")
	cnt := 0

	for {
		led.Low()
		time.Sleep(100 * time.Millisecond)
		led.High()
		time.Sleep(100 * time.Millisecond)
		str := fmt.Sprintf("CNT:%3d", cnt)
		display.LcdPrint(0, 40, str)
		cnt++
		//gt.PrintTerminal(gt.ReadFonts(s))
	}
}
