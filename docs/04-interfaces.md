# Interfaces and Polymorphism

## Basic Interface

```go
type Writer interface {
    Write([]byte) (int, error)
}
```

## Implementing Interfaces

Interfaces are implemented implicitly:

```go
type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
    *c += ByteCounter(len(p))
    return len(p), nil
}

// ByteCounter now implements Writer
```

## Interface Composition

```go
type ReadWriter interface {
    Reader
    Writer
}
```

## Empty Interface

The empty interface can hold any value:

```go
func PrintAnything(v interface{}) {
    fmt.Println(v)
}
```

## Type Assertions

```go
var i interface{} = "hello"

s, ok := i.(string)
if ok {
    fmt.Println(s)
}
```

## Type Switches

```go
switch v := i.(type) {
case int:
    fmt.Printf("Integer: %d\n", v)
case string:
    fmt.Printf("String: %s\n", v)
default:
    fmt.Printf("Unknown: %v\n", v)
}
```

## Common Interfaces

- `fmt.Stringer` - like String() in other languages
- `io.Reader` / `io.Writer` - I/O operations
- `sort.Interface` - sorting collections
- `error` - error handling

## Next Steps

Continue to [05-concurrency.md](05-concurrency.md) to learn about concurrency.
