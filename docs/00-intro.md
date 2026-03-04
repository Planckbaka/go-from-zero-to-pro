# Introduction to Go

## What is Go?

Go (also known as Golang) is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson.

## Why Go?

- **Simple and Readable**: Minimal syntax, easy to learn
- **Fast Compilation**: Quick build times
- **Built-in Concurrency**: Goroutines and channels
- **Strong Standard Library**: Rich set of packages
- **Cross-platform**: Compile for multiple targets

## Installing Go

Visit [go.dev/dl](https://go.dev/dl/) and download the installer for your platform.

Verify installation:

```bash
go version
```

## Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```

Run it:

```bash
go run main.go
```

## Next Steps

Continue to [01-basics.md](01-basics.md) to learn the fundamentals.
