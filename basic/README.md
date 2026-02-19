# Packages

Every Go source file starts with a `package` declaration. Packages group related code and form the basic unit of compilation and visibility.

## `package` declaration

At the top of each file you write the package name:

```go
package math
```

Files in the same directory should use the same package name (with one notable exception described below).

## `main` package vs libraries

- `package main` declares an executable program. It must provide a `func main()` as the entry point.
- Other package names produce libraries that can be imported by other code.

Example `main`:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello from main")
}
```

## Importing packages

Use `import` to reuse code from the standard library or other modules:

```go
import (
	"fmt"
	"net/http"
)
```

Then call exported identifiers from the package (exported names start with a capital letter):

```go
fmt.Println("text")
http.ListenAndServe(":8080", nil)
```

## The `fmt` package

The standard `fmt` package provides formatted I/O for printing and string formatting. Common functions:

- `Print`, `Println`, `Printf` — print to stdout
- `Sprintf` — format into a string
- `Fprintf` — write formatted output to an `io.Writer`
- `Errorf` — create formatted errors

Common verbs:

- `%v` — default format
- `%+v` — include struct field names
- `%#v` — Go-syntax representation
- `%T` — type of the value
- `%s`, `%d`, `%f`, `%t` — string, integer, float, boolean

Example:

```go
package main

import "fmt"

type User struct{ Name string; Age int }

func main() {
	u := User{"Gopher", 5}
	fmt.Printf("user: %+v\n", u)
	s := fmt.Sprintf("hello, %s", "gopher")
	fmt.Println(s)
}
```

Notes: `fmt` is great for quick output and debugging; favor structured logging libraries for production logging.

## Exported vs unexported

- Identifiers beginning with a capital letter are exported (public) and visible to importing packages.
- Identifiers beginning with a lowercase letter are unexported (package-private).

## Internal vs package main test files

It's common to write package-level tests in the same package (e.g., `package math`) and black-box tests using the package name with `_test` suffix (e.g., `package math_test`).

## Package layout and module

- Keep one package per directory for clarity.
- Use `go mod init` at the repository root to create a module and enable imports across your project.

| Command           | Purpose                   |
| ----------------- | ------------------------- |
| `go mod init`     | Initialize new module     |
| `go mod tidy`     | Clean & sync dependencies |
| `go mod download` | Download dependencies     |
| `go mod vendor`   | Copy deps locally         |
| `go list -m all`  | List modules              |


## Conventions and tips
- Name packages short and descriptive (e.g., `http`, `json`, `store`).
- Keep package APIs small and focused.
- Prefer composition over large exported structs.

## Module initialization and the `init()` function

### `go mod init`

To make your code an importable module and manage dependencies, initialize a module at your project root:

```sh
go mod init example.com/myproject
go mod tidy
```

This creates a `go.mod` file declaring the module path and the Go version, and lets you `import` your own packages by their module-qualified paths.

Run `go mod init` once per repository (or per module root). Use `go mod tidy` to add/remove indirect dependencies after editing imports.

### The `init()` function

`init()` is a special function that runs before `main()` and is useful for package-level setup (registration, default config, etc.). Example:

```go
package config

import "log"

var Default string

func init() {
	Default = "production"
	log.Println("config init: default set")
}
```

Notes about `init()`:
- Multiple `init()` functions may appear across files; they execute in file order within a package.
- Package initialization (including `init()`) follows the import dependency graph — imported packages initialize first.
- Keep `init()` lightweight and avoid long-running or complex logic; prefer explicit initialization functions your callers control when appropriate.

## Further reading
- Official packages overview: https://go.dev/doc/code
- Effective Go: https://go.dev/doc/effective_go

## Common build/run/install commands

These commands are the basic workflow for compiling and running Go programs locally.

- `go run` — compile and run in one step (no persistent binary):

```sh
go run .           # run main in current module directory
go run main.go     # run a specific file
```

- `go build` — compile packages and produce a binary in the current directory:

```sh
go build ./cmd/app          # builds binary for the package
go build -o myapp ./cmd/app # write binary to `myapp`
```

Notes: `go build` does not install the binary to your PATH; it produces a local artifact.

- `go install` — compile and install the binary to `$GOBIN` (or `$GOPATH/bin`):

```sh
go install ./cmd/app
# or with a module-aware version (install remote tool):
go install github.com/user/tool@latest
```

Notes: use `go install` to place executable in a bin directory so it's runnable from your shell.

## Variables

Go has explicit variable declarations and a short declaration form. Variables have zero values when not initialized.

Examples:

```go
var x int           // zero value 0
var s string = "a" // explicit type
var a, b = 1, "x"  // type inferred

// short declaration (inside functions only)
count := 10
```

Constants:

```go
const Pi = 3.14
const (
	A = iota // 0
	B         // 1
)
```

Notes:
- Use the short `:=` form inside functions for brevity.
- Prefer explicit types in package-level declarations for clarity.

## For loops

`for` is Go's only loop construct and supports multiple forms.

Traditional C-style:

```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

Condition-only (like `while`):

```go
n := 0
for n < 5 {
	n++
}
```

Infinite loop:

```go
for {
	// run forever
}
```

Range loop (iterate slices, maps, strings, channels):

```go
nums := []int{1,2,3}
for i, v := range nums {
	fmt.Println(i, v)
}

m := map[string]int{"a":1}
for k, v := range m {
	fmt.Println(k, v)
}
```

Control:
- `break`, `continue`, and labeled loops are supported for complex flow control.

Best practice: prefer `range` for collections and keep loop bodies small and clear.

