# pi-display

Utility for printing data to a 16x2 LCD screen.

### Details


This project uses a generic 16x2 LCD screen. It uses the following wiring:

```asm
GND -> 20
VCC -> 2
SDA -> 3
SCL -> 5
```

### Usage


You can import `pi-display` as a library, and print static or scrolling strings on your screen.

```go
package main

import (
	"log"

	"github.com/levsthings/go-i2c"
	device "github.com/levsthings/lcd16x2-driver"
	pidisplay "github.com/levsthings/pi-display"
)


func main() {
	i2c, err := i2c.NewI2C(0x27, 1)
	if err != nil {
		log.Fatal(err)
	}
	defer i2c.Close()

	lcd, err := device.NewLcd(i2c, device.LCD_16x2)
		if err != nil {
		log.Fatal(err)
	}

	err = lcd.BacklightOn()
		if err != nil {
		log.Fatal(err)
	}

    // A static string which fits in the 16 character display
	pidisplay.PrintText(lcd, 1, "A static string")
	for {
		// A dynamic string that scrolls and updates it's value at the end of each render
		pidisplay.ScrollText(lcd, 2, updateValue())

	}
}
```

A more complete example lives in the `example` folder which displays data from a temperature/humidity sensor.


