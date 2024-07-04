package main

import (
	handlers "ascii-art-reverse/handlers"
	"flag"
	"fmt"
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

	// Check for incorrect usage of flags
	if len(Args) < 2 {
		fmt.Fprintln(os.Stderr, "Invalid number of arguments")
		flag.Usage()
		os.Exit(1)
	}

	// Handle different command-line options
	switch {
	case *reverseOption != "":
		if err := handlers.HandleReverse(*reverseOption); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case *colorOption != "":
		if err := handlers.HandleColor(Args); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case *outputOption != "":
		if err := handlers.HandleOutput(Args); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case *alignOption != "":
		if err := handlers.HandleAlign(Args); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	default:
		if err := handlers.HandleDefault(Args); err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
}
