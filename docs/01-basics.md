# Go Basics

## Packages

Every Go program is made up of packages. Programs start running in package `main`.

```go
package main
```

## Imports

```go
import (
    "fmt"
    "math"
)
```

## Variables

```go
var name string = "Go"
age := 10  // Short declaration
```

## Constants

```go
const Pi = 3.14159
```

## Control Flow

### If-Else

```go
if x > 0 {
    return "positive"
} else if x < 0 {
    return "negative"
} else {
    return "zero"
}
```

### For Loop

```go
// Standard for loop
for i := 0; i < 10; i++ {
    fmt.Println(i)
}

// Range-based
for index, value := range slice {
    fmt.Printf("Index: %d, Value: %v\n", index, value)
}
```

### Switch

```go
switch day {
case "Monday", "Tuesday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown")
}
```

## Next Steps

Continue to [02-types.md](02-types.md) to learn about types.
