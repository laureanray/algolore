package main

import "testing"

func TestFibRecur(t *testing.T) {
	got := FibRecur(12)
	want := 144

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestFibDp(t *testing.T) {
	got := FibDp(12)
	want := 144

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func BenchmarkDp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibDp(30)
	}
}

func BenchmarkFibRecur(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibRecur(30)
	}
}
