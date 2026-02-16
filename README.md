
# Learn Go (Golang) 

This repository contains a concise learning guide to follow along the YouTube playlist: https://www.youtube.com/playlist?list=PLYrn63eEqAzYOwqstIZWW4Q8dQ3Ncchbw

Whether you are brand-new to Go or coming from another language, this README gives a runnable learning path, recommended exercises, and small projects to solidify your skills.

## What you'll learn
- Go toolchain basics: `go` command, `go run`, `go build`, `go test`, and `go fmt`
- Package structure and modules with `go mod`
- Core language: types, structs, interfaces, slices, maps, functions, and error handling
- Concurrency: goroutines, channels, and common concurrency patterns
- Standard library usage: `net/http`, `encoding/json`, `io`, and more
- Testing and simple benchmarking

## Recommended learning path
1. Watch the playlist videos in order — the series is structured progressively.
2. Pause frequently and type the examples instead of copying — typing helps retention.
3. After each concept, write one short exercise (see Exercises below).
4. Build the small projects at the end of the playlist to combine concepts.

## Quick setup
1. Install Go: follow the official installer at https://go.dev/dl/
2. Verify installation:

```
go version
```

## Exercises (short, repeatable)
- Implement helper functions for common slice operations (map, filter, reduce).
- Parse JSON into structs and handle missing/optional fields.
- Build a small CLI that accepts flags and prints formatted output.
- Write unit tests for each exercise using `testing` package.

## Small Projects (apply learned concepts)
- Todo API: `net/http` + JSON + in-memory storage + simple routing
- Web scraper: `net/http` + `goquery` (or `encoding/xml`) to fetch and parse pages
- Concurrent worker pool: start N goroutines processing a job queue via channels

## Tips & Best Practices
- Keep functions small and focused.
- Prefer returning errors instead of panics for recoverable problems.
- Use `go fmt` and `gofmt` or `gofumpt` for consistent formatting.
- Write tests for core logic; use table-driven tests for many cases.
- Read Go's effective idioms at https://go.dev/doc/effective_go

## Resources
- Official docs: https://go.dev/doc/
- Effective Go: https://go.dev/doc/effective_go
- Go by Example: https://gobyexample.com/
- The YouTube playlist that structures this guide: https://www.youtube.com/playlist?list=PLYrn63eEqAzYOwqstIZWW4Q8dQ3Ncchbw

- Tour of Go (interactive tutorial): https://tour.golang.org/
- Getting started tutorial: https://go.dev/doc/tutorial/get-started
- Learn Go (official learn pages): https://go.dev/learn
- Package documentation (pkg.go.dev): https://pkg.go.dev/

