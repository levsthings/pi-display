package pidisplay

import (
	"errors"
	"time"

	device "github.com/levsthings/lcd16x2-driver"
)

const (
	chars = 16
	speed = time.Second
)

// ScrollText is used to render ASCII strings which are longer than 16 characters.
// It returns after scrolling through all characters of the string. It's supposed
// to be run in an infinite loop by the caller. You can update the data source in each iteration.
// The default display line is 1, you can pass 2 to display on the second line.
func ScrollText(lcd *device.Lcd, line int, text string) error {
	renderLine := device.SHOW_LINE_1
	if line == 2 {
		renderLine = device.SHOW_LINE_2
	}

	s := text + " " + text[:chars]
	// scrolls text, len(text)+2 to account for index & added empty " "
	for i := 0; i < len(text)+2; i++ {
		err := lcd.ShowMessage(s[i:i+chars], device.ShowOptions(renderLine))
		if err != nil {
			return errors.New("error printing scrolling text to display")
		}
		time.Sleep(speed)
	}
	return nil
}

// PrintText is used to render ASCII strings which are 16 or less characters long.
// If your string is longer, the text will overflow so you should use ScrollText for these.
// The default display line is 1, you can pass 2 to display on the second line.
func PrintText(lcd *device.Lcd, line int, text string) error {
	if len(text) >= chars {
		return errors.New("string is longer than 16 characters, use scrollText instead")
	}

	renderLine := device.SHOW_LINE_1
	if line == 2 {
		renderLine = device.SHOW_LINE_2
	}

	err := lcd.ShowMessage(text, device.ShowOptions(renderLine))
	if err != nil {
		return errors.New("error printing text to display")
	}

	return nil
}
