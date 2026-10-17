# Quad

A set of Go programs that draw rectangles ("quads") in the terminal. Each program prints the same rectangle shape using a different set of characters, based on a width `x` and a height `y`.

Built with Go and the [`z01`](https://github.com/01-edu/z01) package for printing runes.

## Project structure

```
back-Quad/
├── go.mod          # module "quad" (Go 1.22.2), depends on github.com/01-edu/z01
├── go.sum
├── quadA/main.go   # Quad A
├── quadB/main.go   # Quad B
├── quadC/main.go   # Quad C
├── quadD/main.go   # Quad D
├── quadE/main.go   # Quad E
├── test/main.go    # Test runner that calls all five quads
├── onlya.go        # Scratch helper (only1)
└── only1           # Scratch notes (not part of the build)
```

## Requirements

- Go 1.22.2 or newer
- Internet access on first run, to download `github.com/01-edu/z01`

## Usage

Run any quad from the project root:

```bash
go run ./quadA
go run ./quadB
go run ./quadC
go run ./quadD
go run ./quadE
```

Each program's `main` calls its function with four sample sizes: `(5,3)`, `(5,1)`, `(1,1)` and `(1,5)`.

To use a quad with your own size, edit the call in `main()`:

```go
QuadA(8, 4)
```

## The quads

Every quad has the signature `QuadX(x, y int)` where `x` is the width and `y` the height. If `x` or `y` is negative, nothing is printed.

| Quad | Corners (top-left, top-right, bottom-left, bottom-right) | Top/bottom edge | Side edge |
|------|----------------------------------------------------------|-----------------|-----------|
| A    | `o` `o` `o` `o`                                          | `-`             | `\|`      |
| B    | `/` `\` `\` `/`                                          | `*`             | `*`       |
| C    | `A` `A` `C` `C`                                          | `B`             | `B`       |
| D    | `A` `C` `A` `C`                                          | `B`             | `B`       |
| E    | `A` `C` `C` `A`                                          | `B`             | `B`       |

### Example output for `(5, 3)`

**Quad A**
```
o---o
|   |
o---o
```

**Quad B**
```
/***\
*   *
\***/
```

**Quad C**
```
ABBBA
B   B
CBBBC
```

**Quad D**
```
ABBBC
B   B
ABBBC
```

**Quad E**
```
ABBBC
B   B
CBBBA
```

### Edge cases

| Size  | Result                                              |
|-------|-----------------------------------------------------|
| `1×1` | A single corner character (the top-left one)        |
| `5×1` | One row: left corner, edge characters, right corner |
| `1×5` | One column: top corner, side characters, bottom corner |

## Known issues

- `test/main.go` imports `quad` as a package, but the quad functions live in `package main` inside each `quadX/` folder, so the test runner will not compile as it is. To fix it, move the functions into a shared non-`main` package (for example a `quad` package) and have each `quadX/main.go` call it.
- `onlya.go` is `package main` in the project root with no `main` function, which breaks `go build ./...`. Move it into its own folder or delete it.
- `only1` has no extension and contains invalid Go code. It looks like a scratch file.
