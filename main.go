package main

import (
	handlers "ascii-art-reverse/handlers"
	"flag"
	"os"
)

func main() {
	// Define flags for command-line options
	alignOption := flag.String("align", "", "Alignment option")
	outputOption := flag.String("output", "", "Output option")
	colorOption := flag.String("color", "", "Color option")
	reverseOption := flag.String("reverse", "", "Reverse option")

	// Parse command-line arguments
	flag.Parse()
	Args := os.Args
	// Handle different command-line options
	switch {
	case *reverseOption != "":
		handlers.HandleReverse(*reverseOption)
	case *colorOption != "":
		handlers.HandleColor(Args)
	case *outputOption != "":
		handlers.HandleOutput(Args)
	case *alignOption != "":
		handlers.HandleAlign(Args)
	default:
		handlers.HandleDefault(Args)
	}
}
