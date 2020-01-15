package main

import (
	"fmt"

	"github.com/levsthings/go-i2c"
	device "github.com/levsthings/lcd16x2-driver"
	pidisplay "github.com/levsthings/pi-display"
	pitemp "github.com/levsthings/pi-temp"
)

const (
	// I2C Options
	adr = 0x27
	bus = 1
)

func main() {
	i2c, err := i2c.NewI2C(adr, bus)
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't establish a new i2c connection",
		})
	}

	defer i2c.Close()

	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't connect to display",
		})
	}

	err = lcd.BacklightOn()
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't turn display back light on",
		})
	}

	// example usage for PrintText - A static string
	err = pidisplay.PrintText(lcd, 1, "It's kinda hot!!")
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't print text",
		})
	}

	for {
		// example usage for ScrollText - the dynamic data source
		// is called to update values at each iteration.
		err = pidisplay.ScrollText(lcd, 2, updateValue())
		if err != nil {
			logError(errorOutput{
				err,
				"couldn't print scroll text",
			})
		}
	}
}

func updateValue() string {
	d, err := pitemp.GetData()
	if err != nil {
		logError(errorOutput{
			err,
			"couldn't get temp data from pi-temp",
		})
	}

	text := fmt.Sprintf("Temp: %.1fC, Humidity: %.1f%%", d.Temperature, d.Humidity)

	return text
}
