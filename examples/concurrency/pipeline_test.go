package concurrency

import "testing"

func TestPipeline(t *testing.T) {
	in := []int{1, 2, 3}
	got := Pipeline(in)

	if len(got) != len(in) {
		t.Fatalf("unexpected length: got %d want %d", len(got), len(in))
	}
	for i := range in {
		if got[i] != in[i] {
			t.Fatalf("unexpected value at %d: got %d want %d", i, got[i], in[i])
		}
	}
}
