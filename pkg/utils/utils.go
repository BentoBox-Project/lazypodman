package utils

import (
	"fmt"
	"regexp"

	"github.com/fatih/color"
)

// Decolorise strips a string of color
func Decolorise(str string) string {
	re := regexp.MustCompile(`\x1B\[([0-9]{1,2}(;[0-9]{1,2})?)?[m|K]`)
	return re.ReplaceAllString(str, "")
}

// Max returns the maximum of two integers
func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// ColoredString takes a string and a colour attribute and returns a colored
// string with that attribute
func ColoredString(str string, colorAttribute color.Attribute) string {
	// fatih/color does not have a color.Default attribute, so unless we fork that repo the only way for us to express that we don't want to color a string different to the terminal's default is to not call the function in the first place, but that's annoying when you want a streamlined code path. Because I'm too lazy to fork the repo right now, we'll just assume that by FgWhite you really mean Default, for the sake of supporting users with light themed terminals.
	if colorAttribute == color.FgWhite {
		return str
	}
	colour := color.New(colorAttribute)
	return ColoredStringDirect(str, colour)
}

// MultiColoredString takes a string and an array of colour attributes and returns a colored
// string with those attributes
func MultiColoredString(str string, colorAttribute ...color.Attribute) string {
	colour := color.New(colorAttribute...)
	return ColoredStringDirect(str, colour)
}

// ColoredStringDirect used for aggregating a few color attributes rather than
// just sending a single one
func ColoredStringDirect(str string, colour *color.Color) string {
	return colour.SprintFunc()(fmt.Sprint(str))
}
