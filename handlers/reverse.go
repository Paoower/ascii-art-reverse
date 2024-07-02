package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"os"
	"strings"
)

// Handle --reverse=<fileName> option
func HandleReverse(fileName string) {
	// Read font art
	fContent, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		panic(err)
	}
	fontData := string(fContent)
	font := strings.Split(fontData, "\n")

	// Read art file for reverse

	Content, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	textContent := strings.ReplaceAll(string(Content), "$", "")
	textData := string(textContent)
	text := strings.Split(textData, "\n")

	// Process the text using the Reverse function from asciiart package
	if len(text) > 9 {
		for i := 0; i < len(text)-1; {
			if len(text[i]) > 0 {
				src.Reverse(font, text[i:i+8], 0, 0, 1)
				fmt.Println()
				i = i + 8
			} else {
				fmt.Println()
				i = i + 1
			}
		}
	} else {
		src.Reverse(font, text, 0, 0, 1)
		fmt.Println()
	}
}
