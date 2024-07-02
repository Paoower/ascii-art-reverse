package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"log"
	"os"
	"strings"
)

// Handle --output=<OPTION> option
func HandleOutput(args []string) {
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-output:")

	if len(args) == 2 {
		output := strings.Split(src.GetAscii(args[1], "standard"), "\n")
		for i := 0; i < len(output)-1; i++ {
			fmt.Println(output[i])
		}
		return
	}

	// Checking if the correct number of arguments is provided
	if len(args) != 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	if !strings.Contains(args[1], "--output=") {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("EX: go run . --output=<fileName.txt> something standard")
		return
	}

	outputFilename := args[1][9:]

	// Get the input string and banner style from command line arguments
	input := args[2]
	style := args[3]

	output := src.GetAscii(input, style)
	if output == "" {
		log.Fatal("Could not save file")
	}

	err := os.WriteFile(outputFilename, []byte(output+"\n"), 0600)
	if err != nil {
		log.Fatal("Could not save file")
	}
	fmt.Printf("File saved at %v.txt\n", outputFilename)
}
