package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"os"
	"strings"
)

// Handle default case
func HandleDefault(args []string) {
	// Checking if the correct number of arguments is provided
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]\nEX: go run . something standard")
		return
	}
	// Get the input string and banner style from command line arguments
	input := os.Args[1]
	style := "standard"
	if len(args) == 3 {
		style = args[2]
	}

	// Use correct Banner files
	bannerFile, err := src.GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Store line and get line of input
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
		} else {
			lines = append(lines, src.GetWord(word, bannerFile)...)
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}
