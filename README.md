# ASCII Art Reverse

This Go program converts ASCII art back into its original text representation. It supports multiple styles of ASCII art and handles command-line arguments for both generating and reversing ASCII art.

## Arborescence

- `main.go`
- `asciiart/`
  - `asciiart.go`
- `banners/`
    - `standard.txt`
    - `shadow.txt`
    - `thinkertoy.txt`
## Usage

To clone the repository:
```bash
git clone https://zone01normandie.org/git/mtrebert/ascii-art-reverse.git
```

To run the program to generate ASCII art from a string:
```bash
cd ascii-art-reverse

go run . "hello world"
```

To reverse ASCII art from a file:
```bash
cd ascii-art-reverse

go run . --reverse=file.txt
```

## Files

### [main.go](main.go)

Handles command-line arguments and interacts with the `asciiart` package to either generate ASCII art or reverse it.

### [asciiart/](asciiart/)

- **[asciiart.go](asciiart/asciiart.go)**:
  Contains functions for generating ASCII art from a string (`GetColoredAscii`) and reversing ASCII art from a file (`ReverseAsciiArt`).

### [banners/](banners/)
- **[standard.txt](banners/standard.txt)**:
  ASCII representations of letters in standard style.

- **[shadow.txt](banners/shadow.txt)**:
  ASCII representations of letters in shadow style.

- **[thinkertoy.txt](banners/thinkertoy.txt)**:
  ASCII representations of letters in thinkertoy style.

## Example

Given a sample input file `file.txt` with ASCII art:
```
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
```

Running the command:
```bash
go run . --reverse=file.txt
```

Will output:
```
hello
```

