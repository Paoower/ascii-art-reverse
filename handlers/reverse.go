package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"os"
	"strings"
)

// HandleReverse handles the --reverse=<fileName> option
func HandleReverse(fileName string) error {
	// Read font art
	fContent, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		return fmt.Errorf("could not read banner file: %v", err)
	}
	fontData := string(fContent)
	font := strings.Split(fontData, "\n")

	// Read art file for reverse
	content, err := os.ReadFile(fileName)
	if err != nil {
		return fmt.Errorf("could not read input file: %v", err)
	}
	textContent := strings.ReplaceAll(string(content), "$", "")
	text := strings.Split(textContent, "\n")

	// Process the text using the Reverse function from asciiart package
	if len(text) > 9 {
		for i := 0; i < len(text)-1; {
			if len(text[i]) > 0 {
				src.Reverse(font, text[i:i+8], 0, 0, 1)
				fmt.Println()
				i += 8
			} else {
				fmt.Println()
				i++
			}
		}
	} else {
		src.Reverse(font, text, 0, 0, 1)
		fmt.Println()
	}
	return nil
}
