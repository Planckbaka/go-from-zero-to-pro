# Types and Data Structures

## Basic Types

```go
bool
string
int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr
byte // alias for uint8
rune // alias for int32 (Unicode code point)
float32 float64
complex64 complex128
```

## Arrays

Fixed-size sequences:

```go
var arr [5]int
arr := [5]int{1, 2, 3, 4, 5}
```

## Slices

Dynamic-size views into arrays:

```go
slice := []int{1, 2, 3}
slice = append(slice, 4, 5)

// make creates slices with length and capacity
s := make([]int, 5, 10)
```

## Maps

Key-value pairs:

```go
m := make(map[string]int)
m["key"] = 42

// Literal syntax
m := map[string]int{
    "one":   1,
    "two":   2,
}
```

## Structs

Custom data types:

```go
type Person struct {
    Name string
    Age  int
}

p := Person{Name: "Alice", Age: 30}
```

## Pointers

```go
i := 42
p := &i
fmt.Println(*p) // dereference
```

## Next Steps

Continue to [03-functions.md](03-functions.md) to learn about functions.
