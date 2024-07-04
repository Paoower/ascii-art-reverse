package handlers

import (
	src "ascii-art-reverse/src"
)

// Handle --align=<OPTION> option
func HandleAlign(args []string) error {
	option, input, themeName, err := src.CheckArgs(args)
	if err != nil {
		return err
	}
	lines := src.ThemeToLines(themeName)
	src.PrintAsciiArt(input, lines, option)
	return nil
}
