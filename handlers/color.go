package handlers

import (
	src "ascii-art-reverse/src"
	"fmt"
	"log"
	"strings"
)

// Handle --color=<color> option
func HandleColor(args []string) {
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-color:")
	var lines []string
	var err error

	if len(args) > 5 || len(args) == 1 {
		fmt.Println("Usage: go run . [OPTION] [STRING]")
		fmt.Println()
		fmt.Println(`EX: go run . --color=<color> <substring to be colored> "something"`)
	}

	if len(args) == 2 {
		lines, err = src.GetColoredAscii(args[1], "standard", "", "NONE")
		if err != nil {
			return
		}
	}

	if len(args) == 3 {
		if strings.HasPrefix(args[1], "--color=") {
			lines, err = src.GetColoredAscii(args[2], "standard", args[2], src.GetColor(args[1]))
			if err != nil {
				return
			}
		} else {
			lines, err = src.GetColoredAscii(args[1], args[2], "", "NONE")
			if err != nil {
				return
			}
		}
	}

	color := src.GetColor(args[1])

	if len(args) == 4 {
		lines, err = src.GetColoredAscii(args[3], "standard", args[2], color)
		if err != nil {
			return
		}
	}

	if len(args) == 5 {
		lines, err = src.GetColoredAscii(args[3], args[4], args[2], color)
		if err != nil {
			return
		}
	}

	for _, l := range lines {
		fmt.Println(l)
	}
}
