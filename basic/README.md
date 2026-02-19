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

## If / Else statements

Go's conditional statements are straightforward. The `if` condition does not require parentheses and can include a short statement before the condition.

Examples:

```go
if x > 0 {
	fmt.Println("positive")
} else if x < 0 {
	fmt.Println("negative")
} else {
	fmt.Println("zero")
}

// short statement before condition
if err := doSomething(); err != nil {
	fmt.Println("error:", err)
}
```

Notes:
- Use the short statement form to limit the scope of temporary variables (like `err`).
- Avoid overly complex `if` chains; prefer early returns for clearer control flow.

## Switch / Case

`switch` in Go provides a clear way to select between multiple alternatives. It supports value switches, expressionless switches, `fallthrough`, and type switches.

Value switch:

```go
switch s := status; s {
case "ok":
	fmt.Println("all good")
case "warn", "minor":
	fmt.Println("warning")
default:
	fmt.Println("unknown")
}
```

Expressionless switch (like `if` chains):

```go
switch {
case x < 0:
	fmt.Println("neg")
case x == 0:
	fmt.Println("zero")
default:
	fmt.Println("pos")
}
```

`fallthrough` forces execution to the next case (use sparingly):

```go
switch n {
case 0:
	fmt.Println("zero")
	fallthrough
case 1:
	fmt.Println("also run for 0")
}
```

Type switch (inspect dynamic type of an interface):

```go
var i interface{} = 42
switch v := i.(type) {
case int:
	fmt.Println("int", v)
case string:
	fmt.Println("string", v)
default:
	fmt.Printf("unknown type %T\n", v)
}
```

Notes:
- Prefer `switch` over long `if`/`else` chains for readability.
- Avoid unnecessary `fallthrough`; it can make control flow harder to follow.

## Arrays and slices

In Go, arrays are fixed-size, contiguous sequences of elements. Slices are the more commonly used, flexible view over arrays.

Arrays:

```go
var a [3]int           // array of 3 ints, zero-valued
a[0] = 1
// length is part of the type
var b = [2]string{"x", "y"}
```

Notes about arrays:
- Arrays have a fixed size and their length is part of the type (`[3]int` != `[4]int`).
- Arrays are rarely used directly; they're useful when fixed-size storage is required.

Slices:

```go
// slice literal (backed by array)
nums := []int{1, 2, 3}

// make creates a slice with length and capacity
buf := make([]byte, 10, 20) // len=10, cap=20

// slicing an array or slice
s := nums[1:3] // elements at indices 1 and 2
```

Key slice concepts:
- Slices are descriptors: `(pointer, length, capacity)` that reference an underlying array.
- Appending grows a slice; if capacity is exceeded Go allocates a new underlying array:

```go
s = append(s, 4)
```

- Passing a slice to a function passes the slice header by value but allows modifying the underlying array.

When to use which:
- Use slices for most variable-length collections.
- Use arrays only when you specifically need a fixed-size value type.

Further reading: https://go.dev/doc/faq#slices

### Slice details

A slice in Go is a descriptor for a contiguous segment of an underlying array and consists of three parts: a pointer to the array, a length, and a capacity.

- `len(s)` returns the number of elements currently accessible in the slice.
- `cap(s)` returns the maximum number of elements the slice can grow to without allocating a new underlying array (measured from the slice start).

Example:

```go
s := make([]int, 3, 4) // len=3, cap=4
s[0], s[1], s[2] = 10, 20, 30
fmt.Println(len(s), cap(s)) // prints: 3 4

s = append(s, 40)           // len becomes 4, cap stays 4
s = append(s, 50)           // len becomes 5 -> new underlying array allocated; cap grows
```

Slicing an existing slice affects the new slice's len and cap:

```go
t := s[1:3]                // len(t) == 2 (indices 1 and 2)
// cap(t) == cap(s) - 1     // capacity measured from index 1 to end of underlying array
```

Important behaviours and tips:
- Appending past `cap` causes allocation of a new underlying array and copies the old data; the exact growth strategy is implementation-dependent.
- Multiple slices can share the same underlying array; modifying one slice can affect others. Use `copy` to make an independent copy:

```go
dup := make([]int, len(s))
copy(dup, s)
```

- Preallocate capacity with `make([]T, 0, n)` when you know expected size to reduce reallocations.
- Use `len` for safe iteration and `cap` for optimization decisions; avoid relying on exact capacity growth rules.

## Maps

Maps are Go's built-in hash table type for storing key → value pairs. The type is written `map[K]V`.

Declaration and literals:

```go
var m map[string]int            // nil map, cannot write until initialized
m2 := map[string]int{"A": 1}  // literal, ready to use
m3 := make(map[string]int, 10)  // preallocate space for ~10 entries
```

Access and existence check:

```go
v := m2["A"]           // zero value if key missing
v, ok := m2["Alice"]   // ok==true if key present
```

Delete and length:

```go
delete(m2, "A")        // remove key safely (no panic if missing)
fmt.Println(len(m2))    // number of keys
```

Iteration:

```go
for k, v := range m2 {
	fmt.Println(k, v)
}
```

Notes and best practices:
- A nil map behaves like an empty map on reads but will panic on writes; use `make` before writing.
- Map iteration order is intentionally randomized; do not rely on any ordering.
- Maps are reference-like: assigning a map to another variable copies the header but both refer to the same underlying data.
- Maps are NOT safe for concurrent writes. Use `sync.RWMutex` or `sync.Map` for concurrent access.
- For performance-sensitive code, pre-size with `make(map[K]V, n)` when you know approximate number of keys.

Example combined usage:

```go
users := map[string]int{"alice": 30}
if _, ok := users["bob"]; !ok {
	users["bob"] = 25
}
for name := range users {
	fmt.Println(name)
}
```



