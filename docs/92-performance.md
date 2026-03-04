# Performance Optimization

## Profiling

### CPU Profiling

```go
import "runtime/pprof"

f, _ := os.Create("cpu.prof")
pprof.StartCPUProfile(f)
defer pprof.StopCPUProfile()
```

### Memory Profiling

```go
f, _ := os.Create("mem.prof")
pprof.WriteHeapProfile(f)
f.Close()
```

### Using pprof

```bash
go tool pprof cpu.prof
go tool pprof -http=:8080 cpu.prof
```

## Benchmarking

```bash
go test -bench=. -benchmem
go test -bench=. -cpuprofile=cpu.prof
```

## Common Optimizations

### 1. Reduce Allocations

```go
// Bad: creates new slice each time
func process(items []int) []int {
    result := make([]int, 0)
    for _, item := range items {
        result = append(result, item*2)
    }
    return result
}

// Better: pre-allocate
func process(items []int) []int {
    result := make([]int, len(items))
    for i, item := range items {
        result[i] = item * 2
    }
    return result
}
```

### 2. Use sync.Pool

```go
var pool = sync.Pool{
    New: func() interface{} {
        return new(bytes.Buffer)
    },
}

buf := pool.Get().(*bytes.Buffer)
defer pool.Put(buf)
```

### 3. Avoid Interface Overhead

When performance-critical, use concrete types.

### 4. String Builder

```go
// Bad
s := ""
for i := 0; i < n; i++ {
    s += str // creates new string each time
}

// Better
var sb strings.Builder
for i := 0; i < n; i++ {
    sb.WriteString(str)
}
s := sb.String()
```

## Escape Analysis

Check what escapes to heap:

```bash
go build -gcflags="-m"
```

## Next Steps

Continue to [93-security.md](93-security.md) to learn about security.
