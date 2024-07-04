package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"log"
	"os"
	"strings"
)

// HandleOutput handles the --output=<OPTION> option
func HandleOutput(args []string) error {
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-output:")

	if len(args) == 2 {
		output := strings.Split(src.GetAscii(args[1], "standard"), "\n")
		for i := 0; i < len(output)-1; i++ {
			fmt.Println(output[i])
		}
		return nil
	}
	usageMessage := "Usage: go run . [OPTION] [STRING] [BANNER]\nEX: go run . --output=<fileName.txt> something standard"
	// Checking if the correct number of arguments is provided
	if len(args) != 4 {
		return fmt.Errorf(usageMessage)
	}

	if !strings.Contains(args[1], "--output=") {
		return fmt.Errorf(usageMessage)
	}

	outputFilename := args[1][9:]

	// Get the input string and banner style from command line arguments
	input := args[2]
	style := args[3]

	output := src.GetAscii(input, style)
	if output == "" {
		return fmt.Errorf("could not generate ASCII art")
	}

	err := os.WriteFile(outputFilename, []byte(output+"\n"), 0600)
	if err != nil {
		return fmt.Errorf("could not save file: %v", err)
	}
	fmt.Printf("File saved at %v\n", outputFilename)
	return nil
}
