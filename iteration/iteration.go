package iteration

import "strings"

func Repeat(char string, times int) string {
	return strings.Repeat(char, times)
}

const repeatCount = 5

func Repeater(c string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(c)
	}
	return repeated.String()
}
