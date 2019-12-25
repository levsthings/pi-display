package pidisplay

import (
	"log"
	"time"

	device "github.com/levsthings/lcd16x2-driver"
)

const (
	// OPTIONS
	chars = 16
	speed = time.Second
)

// ScrollText is used to render strings which are longer than 16 characters
// It returns after scrolling through all characters of the string. It's supposed
// to be run in an infinite loop by the caller. You can update the data source in each iteration.
// The default display line is 1.
func ScrollText(lcd *device.Lcd, line device.ShowOptions, text string) {
	renderLine := device.SHOW_LINE_1

	if line == 2 {
		renderLine = device.SHOW_LINE_2
	}

	s := text + " " + text[:chars]

	// scrolls text, len(text)+2 to account for index & added empty " "
	for i := 0; i < len(text)+2; i++ {
		err := lcd.ShowMessage(s[i:i+chars], device.ShowOptions(renderLine))
		CheckError(err)

		time.Sleep(speed)
	}
}

// PrintText is used to render static strings which are 16 characters long. If the string
// is longer, the text will overflow. You should use ScrollText for these.
// The default display line is 1.
func PrintText(lcd *device.Lcd, line int, text string) {
	renderLine := device.SHOW_LINE_1

	if line == 2 {
		renderLine = device.SHOW_LINE_2
	}

	renderLine = device.SHOW_LINE_1

	err := lcd.ShowMessage(text, device.ShowOptions(renderLine))
	CheckError(err)
}

// CheckError is a generic error checking function which prints and exits in case
// of an error.
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
