package handlers

import (
	"fmt"
	"log"
	"strings"

	src "ascii-art-reverse/src"
)

func HandleColor(args []string) error {
	log.SetFlags(log.Ltime)
	log.SetPrefix("ascii-art-color:")

	var lines []string
	var err error

	usageMessage := "Usage: go run . --color=<color> <substring to be colored> \"something\""

	if len(args) > 5 || len(args) == 1 {
		return fmt.Errorf(usageMessage)
	}

	switch len(args) {
	case 2:
		lines, err = src.GetColoredAscii(args[1], "standard", "", "NONE")
		if err != nil {
			return fmt.Errorf("error getting colored ASCII: %v", err)
		}
	case 3:
		if strings.HasPrefix(args[1], "--color=") {
			color := src.GetColor(args[1])
			lines, err = src.GetColoredAscii(args[2], "standard", args[2], color)
			if err != nil {
				return fmt.Errorf("error getting colored ASCII: %v", err)
			}
		} else {
			lines, err = src.GetColoredAscii(args[1], args[2], "", "NONE")
			if err != nil {
				return fmt.Errorf("error getting colored ASCII: %v", err)
			}
		}
	case 4:
		color := src.GetColor(args[1])
		lines, err = src.GetColoredAscii(args[3], "standard", args[2], color)
		if err != nil {
			return fmt.Errorf("error getting colored ASCII: %v", err)
		}
	case 5:
		color := src.GetColor(args[1])
		lines, err = src.GetColoredAscii(args[3], args[4], args[2], color)
		if err != nil {
			return fmt.Errorf("error getting colored ASCII: %v", err)
		}
	default:
		return fmt.Errorf(usageMessage)
	}

	for _, line := range lines {
		fmt.Println(line)
	}

	return nil
}
