# ASCII Art Reverse

This Go program converts ASCII art back into its original text representation. It supports multiple styles of ASCII art and handles command-line arguments for both generating and reversing ASCII art.

## Arborescence

- `main.go`
- `color.sh`
- `output.sh`
- `fs.sh`
- `reverse.sh`
- `test/`
  - `example00.txt`
  - `example01.txt`
  - `example02.txt`
  - `example03.txt`
  - `example04.txt`
  - `example05.txt`
  - `example06.txt`
  - `example07.txt`
- `src/`
  - `ascii.go`
  - `banners.go`
  - `colored.go`
  - `justify.go`
  - `reversed.go`
- `handlers/`
  - `align.go`
  - `color.go`
  - `default.go`
  - `output.go`
  - `reverse.go`
- `banners/`
  - `shadow.txt`
  - `standard.txt`
  - `thinkertoy.txt`
  - `train.txt`

## Usage

To clone the repository:
```bash
git clone https://zone01normandie.org/git/mtrebert/ascii-art-reverse.git

cd ascii-art-reverse
```

To run the program to generate ASCII art from a string:
```bash
go run . "hello world"
```

To reverse ASCII art from a file:
```bash
go run . --reverse=file.txt
```

To Test the program:
- Test reverse :
```bash
./reverse.sh
```
- Test color :
```bash
./color.sh
```

- Test output.sh :
```bash
./output.sh

rm -r outputtest
```
- Test fs :
```bash
./fs.sh
```


## Files

### [main.go](main.go)

Handles command-line arguments and interacts with the `handlers` package to manage ASCII art generation, coloring, output, alignment, and reversal based on user input.

### [src/](src/)

Contains source files implementing core functionalities.

- **[ascii.go](src/ascii.go)**: Functions for generating ASCII art.
- **[banners.go](src/banners.go)**: Functions for handling different banner styles.
- **[colored.go](src/colored.go)**: Functions for coloring ASCII art.
- **[justify.go](src/justify.go)**: Functions for justifying text in ASCII art.
- **[reversed.go](src/reversed.go)**: Functions for reversing ASCII art.

### [handlers/](handlers/)

Contains handlers for various command-line options.

- **[align.go](handlers/align.go)**: Handles alignment of ASCII art.
- **[color.go](handlers/color.go)**: Handles coloring of ASCII art.
- **[default.go](handlers/default.go)**: Default handler for ASCII art generation.
- **[output.go](handlers/output.go)**: Handles output file generation for ASCII art.
- **[reverse.go](handlers/reverse.go)**: Handles reversing ASCII art from files.

### [banners/](banners/)

Contains ASCII representations of letters in various styles.

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

