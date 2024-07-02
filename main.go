package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"syscall"
	"unsafe"

	src "ascii-art-reverse/src"
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
		handleReverse(*reverseOption)
	case *colorOption != "":
		handleColor(Args)
	case *outputOption != "":
		handleOutput(Args)
	case *alignOption != "":
		handleAlign(Args)
	default:
		handleDefault(flag.Args())
	}
}

// Handle --reverse=<fileName> option
func handleReverse(fileName string) {
	// Read font art
	fContent, err := os.ReadFile("banners/standard.txt")
	if err != nil {
		panic(err)
	}
	fontData := string(fContent)
	font := strings.Split(fontData, "\n")

	// Read art file for reverse
	textContent, err := os.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	textData := string(textContent)
	text := strings.Split(textData, "\n")

	// Process the text using the Reverse function from asciiart package
	if len(text) > 9 {
		for i := 0; i < len(text)-1; {
			if len(text[i]) > 0 {
				src.Reverse(font, text[i:i+8], 0, 0, 1)
				fmt.Println()
				i = i + 8
			} else {
				fmt.Println()
				i = i + 1
			}
		}
	} else {
		src.Reverse(font, text, 0, 0, 1)
		fmt.Println()
	}
}

// Handle --color=<color> option
func handleColor(args []string) {
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

// Handle --output=<OPTION> option
func handleOutput(args []string) {
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

// Handle --align=<OPTION> option
func handleAlign(args []string) {
	if len(args) < 2 || len(args) > 4 {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
		return
	}
	var input, style, align string
	style = "standard"
	align = "left"
	argIndex := 1
	for argIndex < len(os.Args) {
		arg := args[argIndex]
		switch {
		case strings.HasPrefix(arg, "--align="):
			align = strings.TrimPrefix(arg, "--align=")
			if align != "left" && align != "center" && align != "right" && align != "justify" {
				fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
				return
			}
		case input == "":
			input = arg
		case style == "standard":
			style = arg
		default:
			fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --align=right something standard")
			return
		}
		argIndex++
	}
	// Use correct Banner files
	bannerFile, err := src.GetBannerFile(style)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")
	if align == "justify" {
		numOfSpaces := strings.Count(input, " ")
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW := src.GetWord(Trim(word), bannerFile)
				lines = append(lines, getW...)
			}
		}
		lenOfOneLine := len(lines[0])
		spacesToRemain := getTerminalWidth() - lenOfOneLine
		lines = []string{}
		lines = append(lines, src.GetJustifiedAscii(input, style, spacesToRemain, numOfSpaces)...)
	}
	if align == "center" {
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW := src.GetWord(word, bannerFile)

				lines = append(lines, getW...)
			}
		}
		lenOfOneLine := len(lines[0])
		spacesToRemain := getTerminalWidth() - lenOfOneLine
		lines = []string{}
		lines = append(lines, src.GetCenteredAscii(input, style, spacesToRemain, 2)...)
	}
	if align == "right" {
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW := src.GetWord(word, bannerFile)
				lines = append(lines, getW...)
			}
		}
		lenOfOneLine := len(lines[0])
		spacesToRemain := getTerminalWidth() - lenOfOneLine
		lines = []string{}
		lines = append(lines, src.GetCenteredAscii(input, style, spacesToRemain, 1)...)
	}
	if align == "left" {
		for _, word := range words {
			if word == "" {
				lines = append(lines, "")
			} else {
				getW := src.GetWord(word, bannerFile)
				lines = append(lines, getW...)
			}
		}
	}
	for _, line := range lines {
		fmt.Println(line)
	}
}

// Handle default case
func handleDefault(args []string) {
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

// Get the width of the terminal
func getTerminalWidth() int {
	var dimensions [4]uint16
	retCode, _, errno := syscall.Syscall6(syscall.SYS_IOCTL,
		uintptr(syscall.Stdout),
		uintptr(syscall.TIOCGWINSZ),
		uintptr(unsafe.Pointer(&dimensions)),
		0, 0, 0)
	if int(retCode) == -1 {
		fmt.Println("Error getting terminal size:", errno)
		return 80 // default width
	}
	return int(dimensions[1])
}
func Trim(s string) string {
	str := ""
	for _, char := range s {
		if char != ' ' {
			str += string(char)
		}
	}
	return str
}
