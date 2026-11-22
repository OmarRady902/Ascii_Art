# ASCII Art

A command-line tool written in Go that renders text as large ASCII art banners.

---

## Description

`ascii-art` takes a string as a command-line argument and prints it as ASCII art. Every character is an 8-line tall drawing read from a banner file, and the drawings are printed side by side to form words. Three banner styles are included. The project uses only the Go standard library.

---

## Features

- Three banner styles: `standard`, `shadow`, `thinkertoy`
- Multi-line input with the `\n` escape sequence
- Empty input prints nothing; input made only of `\n` prints blank lines
- Input validation with clear error messages and exit codes
- A table-driven unit test and benchmarks
- Standard library only

---

## Authors

- **Omar Rady** — [@orady](https://01.nextera.education/git/orady)
- **Hazem Elmahdy** — [@helmahdy](https://01.nextera.education/git/helmahdy)

---

## Usage

```bash
go run . "<text>" [banner]
```

| Argument | Required | Description |
|----------|----------|-------------|
| `text` | Yes | The text to render. Use `\n` for a line break |
| `banner` | No | `standard`, `shadow` or `thinkertoy`. Default: `standard` |

Run the program from the project root, because the banner files are read from the `banners/` directory.

---

## Examples

```bash
go run . "Hello"
```

```text
 _    _          _   _
| |  | |        | | | |
| |__| |   ___  | | | |   ___
|  __  |  / _ \ | | | |  / _ \
| |  | | |  __/ | | | | | (_) |
|_|  |_|  \___| |_| |_|  \___/
```

```bash
go run . "Hi" shadow
```

```text
_|    _| _|
_|    _|
_|_|_|_| _|
_|    _| _|
_|    _| _|
```

```bash
go run . "Hi" thinkertoy
```

```text
o  o
|  | o
O--O
|  | |
o  o |
```

```bash
go run . "Hello\nWorld"      # two words, one under the other
go run . "a\n\nb"            # an empty part becomes one blank line
```

---

## How It Works

```
go run . "Hello\nWorld" shadow
        │
        ▼
  main.go
        ├─ check the arguments .............. wrong count -> usage message, exit 1
        ├─ banner.LoadBannerFile ............ banners/<name>.txt -> string
        ├─ banner.BuildCharacterMap ......... string -> map[rune][]string
        ├─ replace the \n escape with a real newline, validate the characters
        └─ for each line of the text
              └─ render.RenderText ........... 8 rows, glyph by glyph
        │
        ▼
     stdout
```

### Banner file format

Each banner file stores the 95 printable ASCII characters, from space (32) to `~` (126), in order. Every character takes 9 lines: one separator line followed by its 8 art rows, so a file has 95 x 9 = 855 lines.

```
line 0       separator
lines 1-8    rows of the first character (space)
line 9       separator
lines 10-17  rows of the second character (!)
...
```

`BuildCharacterMap` walks the file in steps of 9 and stores the 8 rows of each character in a map, with the character itself as the key.

---

## Project Structure

```
ascii-art/
├── main.go                  # entry point: arguments, validation, printing
├── go.mod                   # module ASCII_ART
├── banner/
│   ├── banner.go            # LoadBannerFile, BuildCharacterMap
│   ├── banner_test.go       # table-driven test for BuildCharacterMap
│   └── benchmark_test.go    # benchmarks for loading and parsing
├── render/
│   ├── render.go            # RenderText
│   └── benchmark_test.go    # benchmark for RenderText
└── banners/
    ├── standard.txt
    ├── shadow.txt
    └── thinkertoy.txt
```

---

## Implementation Details

| Package | Function | Responsibility |
|---------|----------|----------------|
| `banner` | `LoadBannerFile` | Reads a banner file and returns its content |
| `banner` | `BuildCharacterMap` | Builds a `map[rune][]string` from the file content |
| `render` | `RenderText` | Builds the 8 output rows for one line of text and returns them as a string |
| `main` | `main` | Reads the arguments, validates the input, splits on newlines, prints the result |

Design decisions:

- **`map[rune][]string`** gives direct access to a character's drawing by the character itself, with no index arithmetic at render time.
- **`rune` keys** match what `for range` yields when it walks a string, so no conversion is needed.
- **`strings.Builder`** with `Grow` builds the output without repeatedly copying strings.
- **Separation of concerns**: `RenderText` only returns a string and never prints, so `main` stays the only place that touches the terminal and the function stays easy to test.

---

## Error Handling

| Situation | Behaviour |
|-----------|-----------|
| Missing or extra arguments | Prints the usage message and exits with code 1 |
| Empty text | Prints nothing and exits normally |
| Text made only of `\n` | Prints one blank line per `\n` |
| Banner file missing or unreadable | Prints the error and exits with code 1 |
| Character outside ASCII 32-126 | Prints `Error: unsupported character` and exits with code 1 |

---

## Testing

```bash
go test ./...                    # run the tests
go test -v ./banner              # see each case
go test -bench=. -benchmem ./... # run the benchmarks
```

| Kind | Name | What it covers |
|------|------|----------------|
| Unit test (table-driven) | `TestBuildCharacterMap` | First and last row of two characters, so the 9-line stepping is checked |
| Benchmark | `BenchmarkLoadBannerFile` | Reading a banner file |
| Benchmark | `BenchmarkBuildCharacterMap` | Parsing the standard banner |
| Benchmark | `BenchmarkRenderText` | Rendering `Hello Omar!` |
