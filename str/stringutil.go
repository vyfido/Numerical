//Implement the routines auxiliaries based in Ruby language
package str

import (
	"strings"
	"unicode"
)

//Change the order the string putting the right to left
func Reverse(s string) string {
	var r string = ""
	var max int = len(s) - 1
	for i := max; i >= 0; i-- {
		r += string(s[i])
	}
	return r
}

//Convert the string in minuscules alone put the first character in mayuscules
func Capitalize(s string) string {
	if len(s) > 0 {
		return strings.ToUpper(string(s[0])) + strings.ToLower(s[1:])
	} else {
		return s
	}
}

//Change the letters the lower to upper and the upper to lower
func Invertir(s string) string {
	var r string = ""
	for i := 0; i < len(s); i++ {
		l := rune(s[i])
		if unicode.IsLower(l) {
			r += strings.ToUpper(string(l))
		} else if unicode.IsUpper(l) {
			r += strings.ToLower(string(l))
		} else {
			r += string(l)
		}
	}
	return r
}

//Permit rotate the characters the a string the right to left
func Rotate(s string, pos int) string {
	var ind int = pos % len(s)
	if pos > 0 {
		return s[ind:] + s[:ind]
	} else {
		ind *= -1
		return s[:ind] + s[ind:]
	}
}
