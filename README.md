# go-reloaded

A text completion, editing, and auto-correction tool written in Go.

## Overview

`go-reloaded` reads a text file, applies a set of transformations — number base conversion, casing, article correction, punctuation formatting, and quote cleanup — and writes the result to an output file.

## Usage

```bash
go run . <input_file> <output_file>
```

**Example:**

```bash
go run . sample.txt result.txt
```

## Transformations

### Number Conversion

| Tag | Description | Example |
|-----|-------------|---------|
| `(hex)` | Converts the preceding hexadecimal word to decimal | `1E (hex)` → `30` |
| `(bin)` | Converts the preceding binary word to decimal | `10 (bin)` → `2` |

Tags can be chained. The output of one conversion becomes the input for the next:

```
a (hex) (bin)  →  2
```

Invalid hex/bin words are left unchanged.

### Casing

| Tag | Description | Example |
|-----|-------------|---------|
| `(up)` | Uppercases the preceding word | `go (up)` → `GO` |
| `(low)` | Lowercases the preceding word | `STOP (low)` → `stop` |
| `(cap)` | Capitalizes the preceding word | `bridge (cap)` → `Bridge` |
| `(up, N)` | Uppercases the preceding N words | `so exciting (up, 2)` → `SO EXCITING` |
| `(low, N)` | Lowercases the preceding N words | `IN BED (low, 2)` → `in bed` |
| `(cap, N)` | Capitalizes the preceding N words | `the great gatsby (cap, 3)` → `The Great Gatsby` |

Casing tags can also be chained and their effects stack left to right:

```
hello (low, 1)(cap, 1)(up, 1)  →  HELLO
```

### Article Correction

`a` and `an` are automatically corrected based on the first letter of the following word.

- `a` → `an` before words starting with a vowel (`a e i o u`) or `h`
- `an` → `a` before words starting with a consonant

```
A amazing rock  →  An amazing rock
It was a hour   →  It was an hour
an cat          →  a cat
```

Case is preserved (`A` → `An`, `An` → `A`).

### Punctuation Formatting

Punctuation marks (`. , ! ? : ;`) are attached to the preceding word with a space before the next word:

```
hello , world !  →  hello, world!
Hello:world      →  Hello: world
```

Groups of the same punctuation are collapsed together:

```
I was thinking .  .    .  →  I was thinking...
BAMM !  !  !              →  BAMM!!!
```

Mixed punctuation groups are also supported:

```
what do you think !?  →  what do you think!?
```

### Quote Cleanup

Single (`'`) and double (`"`) quote pairs have their inner spacing removed:

```
' awesome '                    →  'awesome'
' I am here '                  →  'I am here'
"  I can't believe  "          →  "I can't believe"
```

In-word apostrophes (e.g. `ma'am`, `can't`) are left untouched.

## Project Structure

```
go-reloaded/
├── main.go
├── go.mod
└── transforms/
    ├── apply.go        # Orchestrates all transformations
    ├── numbers.go      # (hex) and (bin) conversion
    ├── casing.go       # (up), (low), (cap) tags
    ├── articles.go     # a/an correction
    ├── punctuation.go  # Punctuation spacing and grouping
    └── quotes.go       # Quote margin cleanup
```

## Running Tests

```bash
go test ./...
```

Test cases are in `main_test.go` and cover number conversion, casing, article correction, punctuation, quotes, edge cases, and chained/nested tags.

## Requirements

- Go 1.21+
- Standard library only