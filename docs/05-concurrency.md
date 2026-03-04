# Concurrency Patterns

## Goroutines

Lightweight threads managed by Go runtime:

```go
go func() {
    fmt.Println("Running in a goroutine")
}()
```

## Channels

Typed conduits for communication:

```go
ch := make(chan int)

// Send
ch <- 42

// Receive
value := <-ch
```

## Buffered Channels

```go
ch := make(chan int, 100)
```

## Select

Wait on multiple channel operations:

```go
select {
case msg := <-ch1:
    fmt.Println("Received from ch1:", msg)
case msg := <-ch2:
    fmt.Println("Received from ch2:", msg)
case <-time.After(time.Second):
    fmt.Println("Timeout")
}
```

## Worker Pool Pattern

```go
func worker(id int, jobs <-chan int, results chan<- int) {
    for j := range jobs {
        results <- j * 2
    }
}
```

## Pipeline Pattern

```go
func generate(nums ...int) <-chan int {
    out := make(chan int)
    go func() {
        for _, n := range nums {
            out <- n
        }
        close(out)
    }()
    return out
}

func square(in <-chan int) <-chan int {
    out := make(chan int)
    go func() {
        for n := range in {
            out <- n * n
        }
        close(out)
    }()
    return out
}
```

## Context

For cancellation and timeouts:

```go
ctx, cancel := context.WithTimeout(context.Background(), time.Second)
defer cancel()
```

## Next Steps

Continue to [90-tooling.md](90-tooling.md) to learn about Go tooling.
