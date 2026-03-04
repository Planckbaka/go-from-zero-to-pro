package concurrency

import (
	"testing"
)

func TestPipeline(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	// Pipeline: generate -> square -> sum
	nums := Generate(done, 1, 2, 3, 4, 5)
	squared := Square(done, nums)
	result := Sum(done, squared)

	// 1 + 4 + 9 + 16 + 25 = 55
	expected := 55
	if result != expected {
		t.Errorf("Sum = %d; want %d", result, expected)
	}
}

func TestGenerate(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	nums := Generate(done, 1, 2, 3)
	var result []int
	for n := range nums {
		result = append(result, n)
	}

	if len(result) != 3 {
		t.Errorf("Generate produced %d numbers; want 3", len(result))
	}
}

func TestSquare(t *testing.T) {
	done := make(chan struct{})
	defer close(done)

	in := Generate(done, 2, 3, 4)
	out := Square(done, in)

	expected := []int{4, 9, 16}
	var i int
	for n := range out {
		if n != expected[i] {
			t.Errorf("Square = %d; want %d", n, expected[i])
		}
		i++
	}
}
