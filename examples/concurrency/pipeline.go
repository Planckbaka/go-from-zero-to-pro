package concurrency

// Generate sends values to a channel.
func Generate(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// Square receives integers and sends their squares.
func Square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// Sum consumes integers from a channel and returns their sum.
func Sum(done <-chan struct{}, in <-chan int) int {
	total := 0
	for {
		select {
		case n, ok := <-in:
			if !ok {
				return total
			}
			total += n
		case <-done:
			return total
		}
	}
}
