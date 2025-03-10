# ASCII Art Reverse

A utility that converts ASCII art back to plain text.

## Description

This program does the reverse operation of the ascii-art tools. It takes ASCII art as input and attempts to convert it back to the original text string that would generate such output.

## Features

- Reverse conversion from ASCII art to text
- Support for standard, shadow, and thinkertoy banner styles
- Input from files or standard input
- Pattern recognition algorithms

## Usage

```bash
go run . [OPTION] [BANNER]
```

Options:
- `--reverse=<filename>`: Specify the file containing ASCII art to reverse

Examples:
```bash
go run . --reverse=ascii-art.txt standard
cat ascii-art.txt | go run . --reverse standard
```

## Implementation Details

- Pattern matching against known banner templates
- Character recognition for each ASCII art block
- Support for multi-line ASCII art
- Error handling for unknown or malformed patterns

## Requirements

- Go 1.13 or higher
