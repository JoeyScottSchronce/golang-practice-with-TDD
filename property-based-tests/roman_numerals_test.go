package main

import (
	"fmt"
	"testing"
)

var testcases = []struct {
	Description string
	Arabic      int
	Roman       string
}{
	// {"Description", Arabic, "Want"},
	{"1 converted to I", 1, "I"},
	{"2 converted to II", 2, "II"},
	{"3 converted to III", 3, "III"},
	{"4 converted to IV", 4, "IV"},
	{"5 converted to V", 5, "V"},
	{"6 converted to VI", 6, "VI"},
	{"7 converted to VII", 7, "VII"},
	{"8 converted to VIII", 8, "VIII"},
	{"9 converted to IX", 9, "IX"},
	{"10 gets converted to X", 10, "X"},
	{"14 gets converted to XIV", 14, "XIV"},
	{"18 gets converted to XVIII", 18, "XVIII"},
	{"20 gets converted to XX", 20, "XX"},
	{"39 gets converted to XXXIX", 39, "XXXIX"},
	{"40 gets converted to XL", 40, "XL"},
	{"47 gets converted to XLVII", 47, "XLVII"},
	{"49 gets converted to XLIX", 49, "XLIX"},
	{"50 gets converted to L", 50, "L"},
	{"99 gets converted to XCIX", 99, "XCIX"},
	{"100 gets converted to C", 100, "C"},
	{"499 gets converted to CDXCIX", 499, "CDXCIX"},
	{"500 gets converted to D", 500, "D"},
	{"999 gets converted to CMXCIX", 999, "CMXCIX"},
	{"1000 gets converted to M", 1000, "M"},
	{"1984 gets converted to MCMLXXXIV", 1984, "MCMLXXXIV"},
	{"3999 gets converted to MMMCMXCIX", 3999, "MMMCMXCIX"},
}

func TestConvertToRomanNumerals(t *testing.T) {

	for _, test := range testcases {
		t.Run(test.Description, func(t *testing.T) {
			got := ConvertToRoman(test.Arabic)

			if got != test.Roman {
				t.Errorf("got %q, want %q", got, test.Roman)
			}
		})
	}
}

func TestConvertToArabic(t *testing.T) {
	for _, test := range testcases {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Roman, test.Arabic), func(t *testing.T) {
			got := ConvertToArabic(test.Roman)

			if got != test.Arabic {
				t.Errorf("got %d, want %d", got, test.Arabic)
			}
		})
	}
}
