package src

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	chunkSize int = 9
	padding   int = 3
)

var windowSize int = GetWindowSize()

// Get the terminal window size.
func GetWindowSize() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdout
	out, _ := cmd.CombinedOutput()
	f := strings.Fields(string(out))
	i, _ := strconv.Atoi(f[1])
	return i
}

// Check theme validity.
func ThemeIsValid(i int, args []string) bool {
	return args[i] == "shadow" ||
		args[i] == "standard" ||
		args[i] == "thinkertoy"
}

// Get theme name from arguments.
func CheckArgs(args []string) (string, []string, string, error) {
	if len(args) == 4 && ThemeIsValid(3, args) {
		return strings.TrimPrefix(args[1], "--align="), GetTextInput(args[2]), args[3], nil
	} else if len(args) == 3 && strings.HasPrefix(args[1], "--align=") {
		return strings.TrimPrefix(args[1], "--align="), GetTextInput(args[2]), "standard", nil
	} else if len(args) == 3 && !strings.HasPrefix(args[1], "--align=") && ThemeIsValid(2, args) {
		return "", GetTextInput(args[1]), args[2], nil
	} else if len(args) == 2 {
		return "", GetTextInput(args[1]), "standard", nil
	} else {
		return "", nil, "", errors.New("usage: go run .  [OPTION] [STRING] [BANNER]\nexample: go run . --align=right  something  standard")
	}
}

// Convert .txt file of the theme into lines.
func ThemeToLines(s string) []string {
	data, _ := os.ReadFile("templates/" + s + ".txt")
	return strings.Split(string(data), "\n")
}

// Get text input from arguments and split it into differente lines.
func GetTextInput(s string) []string {
	inputCorrected := strings.Replace(s, "\\n", "\n", -1)
	return strings.Split(inputCorrected, "\n")
}

// Print the ascii-art.
func PrintAsciiArt(input, lines []string, option string) {
	for _, inputLine := range input {
		if inputLine == "" {
			fmt.Println()
		} else {
			for line := 1; line <= chunkSize; line++ {
				isFirstSpace := true
				s := ""
				switch option {
				case "center":
					AlignCenter(inputLine, s, lines, line)
				case "right":
					AlignRight(inputLine, s, lines, line)
				case "justify":
					AlignJustify(inputLine, s, lines, line, isFirstSpace)
				case "left":
					fallthrough
				default:
					s = Printline(inputLine, s, lines, line)
					fmt.Println(s)
				}
			}
		}
	}
}

// Get the real lenght of the standard output.
func GetOutputLenght(inputLine, s string, lines []string, line int) int {
	o := ""
	o = Printline(inputLine, s, lines, line)
	return len(o)
}

// Verify if the real lenght of the output is not larger than the window size.
func PrintSafe(s string) {
	if len(s) > windowSize {
		fmt.Println("You better buy a larger screen")
		os.Exit(0)
	}
}

// The basic line printing function.
func Printline(inputLine, s string, lines []string, line int) string {
	for char := 0; char < len(inputLine); char++ {
		if !(inputLine[char] >= 32 && inputLine[char] <= 126) {
			fmt.Println("Please only use basic ascii characters")
			os.Exit(0)
		}
		characterStart := (int(inputLine[char]) - 32) * chunkSize
		if characterStart+line < len(lines) {
			s += lines[characterStart+line]
		}
	}
	PrintSafe(s)
	return s
}

// Align center.
func AlignCenter(inputLine, s string, lines []string, line int) {
	s = Printline(inputLine, s, lines, line)
	spaces := (windowSize - len(s) - padding) / 2
	s = fmt.Sprintf("%*s%s", spaces, " ", s)
	PrintSafe(s)
	fmt.Println(s)
}

// Align right.
func AlignRight(inputLine, s string, lines []string, line int) {
	outputLenght := GetOutputLenght(inputLine, s, lines, line) + 1
	spaces := (windowSize - outputLenght)
	s += fmt.Sprintf("%*s", spaces, " ")
	s = Printline(inputLine, s, lines, line)
	PrintSafe(s)
	fmt.Println(s)
}

// Align justify.
func AlignJustify(inputLine, s string, lines []string, line int, isFirstSpace bool) {
	spaceCount := 0
	for _, s := range inputLine {
		if s == ' ' {
			spaceCount++
		}
	}
	if spaceCount == 0 {
		s = Printline(inputLine, s, lines, line)
		fmt.Println(s)
	} else {
		outputLenght := GetOutputLenght(inputLine, s, lines, line) - spaceCount
		spaces := (windowSize - outputLenght - padding) / spaceCount
		for char := 0; char < len(inputLine); char++ {
			characterStart := (int(inputLine[char]) - 32) * chunkSize
			if inputLine[char] != ' ' {
				s += lines[characterStart+line]
			} else {
				if (outputLenght+spaceCount)+(spaces*spaceCount) < windowSize-1 && isFirstSpace {
					s += lines[characterStart+line] + fmt.Sprintf("%*s", spaces+1, " ")
					isFirstSpace = false
				} else {
					s += lines[characterStart+line] + fmt.Sprintf("%*s", spaces, " ")
				}
			}
		}
		PrintSafe(s)
		fmt.Println(s)
	}
}
