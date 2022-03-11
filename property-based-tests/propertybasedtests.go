package propertybasedtests

import (
	"fmt"
	"strings"
)

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, n := range allRomanNumerals {
		for arabic >= n.Value {
			result.WriteString(n.Roman)
			arabic -= n.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) (arabic uint16) {
	for i := 0; i < len(roman); i++ {
		char := roman[i]
		atEnd := i+1 == len(roman)

		if !atEnd && canBeSubtractive(char) {
			twoChar := fmt.Sprintf("%s%s", string(char), string(roman[i+1]))
			val := allRomanNumerals.ValueOf(twoChar)
			if val > 0 {
				// twoChar value exists, so add it
				arabic += val
				// and increment to skip the second char
				i++
				// and jump to next for-loop iteration
				continue
			}
		}
		arabic += allRomanNumerals.ValueOf(string(char))
	}
	return
}

type RomanNumeral struct {
	Value uint16
	Roman string
}

type RomanNumerals []RomanNumeral

func (r RomanNumerals) ValueOf(char string) uint16 {
	for _, roman := range r {
		if roman.Roman == char {
			return roman.Value
		}
	}
	return 0
}

var allRomanNumerals = RomanNumerals{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func canBeSubtractive(c byte) bool {
	return c == 'I' || c == 'X' || c == 'C'
}
