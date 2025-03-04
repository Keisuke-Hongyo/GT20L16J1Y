package OLED

import (
	"GT20L16J1Y/GT20L16J1Y"
	"image/color"
	"tinygo.org/x/drivers/ssd1306"
)

// Display wraps
type Display struct {
	device ssd1306.Device
	XPos   int16
	YPos   int16
	font   *GT20L16J1Y.Device
}

func NewDisplay(device ssd1306.Device, font *GT20L16J1Y.Device) Display {
	return Display{device: device, font: font}
}

func (d *Display) LcdPrint(x int16, y int16, str string) {
	d.XPos = x // set position X
	d.YPos = y // set position Y
	d.printText(str)
}

func (d *Display) printText(str string) {
	var f GT20L16J1Y.Fonts
	tmp := d.XPos
	f = d.font.ReadFonts(str)
	for i := 0; i < len(f); i++ {
		// Font Data Output
		d.printChar(f[i])
		d.XPos += f[i].FontWidth
	}
	d.XPos = tmp
	_ = d.device.Display()
}

func (d *Display) printChar(font GT20L16J1Y.Font) {
	var x, y int16
	for y = 0; y < font.FontHeight; y++ {
		for x = 0; x < font.FontWidth; x++ {
			if font.FontData[x]&(0x01<<y) != 0x00 {
				d.device.SetPixel(x+d.XPos, y+d.YPos, color.RGBA{255, 255, 255, 0})
			} else {
				d.device.SetPixel(x+d.XPos, y+d.YPos, color.RGBA{0, 0, 0, 0})
			}
		}
	}
}
