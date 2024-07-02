package src

import (
	"os"
	"strings"
)

const LETTER_HEIGHT = 8
const NONE = "\033[0m"

func GetColoredAscii(input, style, substr, color string) ([]string, error) {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		return nil, err
	}

	lines := make([]string, 0)
	inputLines := strings.Split(input, "\\n")

	//Find all indexes where theres substr

	times := strings.Count(input, substr)
	var indexes []int

	t := input
	for i := 0; i < times; i++ {
		idx := strings.Index(t, substr)
		indexes = append(indexes, idx)
		t = strings.Replace(t, substr, strings.Repeat("a", len(substr)), 1)
	}

	for _, line := range inputLines {
		if line == "" {
			lines = append(lines, "")
			continue
		}
		linesOfline := getLine(line, bannerFile, substr, color, indexes)
		lines = append(lines, linesOfline...)
	}

	return lines, nil
}

func GetLetter(content string, ascii int) string {
	if ascii == 32 {
		s := ""
		for i := 0; i < 8; i++ {
			if i != 7 {
				s += "    " + "\n"
				continue
			}
			s += "    "
		}
		return s
	}
	str := ""
	lines := strings.Split(content, "\n")
	place := ascii - 31
	times := (place - 1) * LETTER_HEIGHT
	beginning := (ascii - 30) + times
	for i := beginning; i < beginning+LETTER_HEIGHT; i++ {
		if i != (beginning+LETTER_HEIGHT)-1 {
			str += lines[i-1] + "\n"
		} else {
			str += lines[i-1]
		}
	}
	return str
}

func GetAscii(input, style string) string {
	bannerFile, err := GetBannerFile(style)
	if err != nil {
		return ""
	}
	lines := make([]string, 0)
	words := strings.Split(input, "\\n")
	for _, word := range words {
		if word == "" {
			lines = append(lines, "")
			continue
		}
		lines = append(lines, GetWord(word, bannerFile)...)
	}
	str := ""
	for _, line := range lines {
		str += line + "\n"
	}
	return str
}

func fileOpen(filename string) string {
	f, err := os.ReadFile(filename)
	if err != nil {
		return ""
	}
	// Normalize line endings to Unix-style
	content := strings.ReplaceAll(string(f), "\r\n", "\n")
	return content
}

func GetWord(input string, bannerFile string) []string {
	lines := make([]string, 8)
	content := fileOpen(bannerFile)
	for _, char := range input {
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			lines[i] += c[i]
		}
	}
	return lines
}

func getLine(input, bannerFile, substr, color string, indexes []int) []string {
	lines := make([]string, 8)

	f, err := os.ReadFile(bannerFile)
	if err != nil {
		return []string{}
	}
	// Normalize line endings to Unix-style
	content := strings.ReplaceAll(string(f), "\r\n", "\n")

	for j, char := range input {
		c := strings.Split(GetLetter(content, int(char)), "\n")
		for i := 0; i < len(lines); i++ {
			for _, n := range indexes {
				if j == n {
					c[i] = color + c[i]
				}
				if j == n+len(substr)-1 {
					c[i] += NONE
				}
			}
			lines[i] += c[i]
		}
	}

	return lines
}
