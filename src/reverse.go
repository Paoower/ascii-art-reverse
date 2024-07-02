package src

import (
	"fmt"
)

// symbol in ascii-art file(input.txt). Count - count of lines in ascii-art file(input.txt). Start - number of line in font(standard.txt)
func Reverse(font []string, text []string, pos int, count int, start int) {
	if pos != len(text[count]) { // if we are not finished reversing
		l := len(font[start]) // length of candidate for research
		if pos+l <= len(text[count]) {
			if count < 7 {
				if text[count][pos:l+pos] == font[start+count] { // if the font-line and line of ascii-art are equal
					Reverse(font, text, pos, count+1, start) // compare the next line
				} else {
					Reverse(font, text, pos, 0, start+9) // if not equal, try the next symbol in font-file
				}
			} else {
				r := ((start - 1) / 9) + 32 // find and print the letter
				fmt.Printf("%c", r)
				Reverse(font, text, pos+l, 0, 1) // restart counts for next research
			}
		} else {
			Reverse(font, text, pos, 0, start+9)
		}
	}
}
