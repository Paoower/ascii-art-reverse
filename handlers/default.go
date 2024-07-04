package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"strings"
)

// HandleDefault handles the default case for the program
func HandleDefault(args []string) error {
	// Checking if the correct number of arguments is provided
	if len(args) < 2 || len(args) > 3 {
		return fmt.Errorf("usage: go run . [string] [banner]\nex: go run . something standard")
	}

	// Get the input string and banner style from command line arguments
	input := args[1]
	style := "standard"
	if len(args) == 3 {
		style = args[2]
	}

	// Use correct Banner files
	bannerFile, err := src.GetBannerFile(style)
	if err != nil {
		return fmt.Errorf("failed to get banner file: %v", err)
	}

	// Store lines and get lines of input
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")

	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
		} else {
			wordLines := src.GetWord(word, bannerFile)
			lines = append(lines, wordLines...)
		}
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
