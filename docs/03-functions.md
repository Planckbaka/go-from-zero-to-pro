# Functions and Methods

## Functions

```go
func add(a, b int) int {
    return a + b
}
```

## Multiple Return Values

```go
func divide(a, b int) (int, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}
```

## Named Return Values

```go
func rectangle(width, height int) (area, perimeter int) {
    area = width * height
    perimeter = 2 * (width + height)
    return // naked return
}
```

## Variadic Functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

## Closures

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
```

## Methods

Methods are functions with a receiver:

```go
type Rectangle struct {
    Width, Height float64
}

func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Pointer receiver (can modify the struct)
func (r *Rectangle) Scale(factor float64) {
    r.Width *= factor
    r.Height *= factor
}
```

## Next Steps

Continue to [04-interfaces.md](04-interfaces.md) to learn about interfaces.
