package main

import (
	"fmt"
	"log"

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
	checkError(err)
	defer i2c.Close()

	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
	checkError(err)

	err = lcd.BacklightOn()
	checkError(err)

	// example usage for PrintText - A static string
	pidisplay.PrintText(lcd, 1, "It's kinda hot!!")
	for {
		// example usage for ScrollText - the dynamic data source
		// is called to update values at each iteration.
		pidisplay.ScrollText(lcd, 2, updateTemp())

	}
}

func updateTemp() string {
	d := pitemp.GetData()
	text := fmt.Sprintf("Temp: %.1fC, Humidity: %.1f%%", d.Temperature, d.Humidity)

	return text
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
