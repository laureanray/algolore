package tests

import (
	algo "algolore/impl"
	"testing"
)

func TestFibRecur(t *testing.T) {
	got := algo.FibRecur(12)
	want := 144

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestFibDp(t *testing.T) {
	got := algo.FibDp(12)
	want := 144

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkDp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.FibDp(30)
	}
}

func BenchmarkFibRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		algo.FibRecur(30)
	}
}
