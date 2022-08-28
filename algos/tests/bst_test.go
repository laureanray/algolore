package tests

import (
	algo "algolore/impl"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	input := []int{2, 3, 5, 7, 9}
	target := 7
	// this should return the index if found

	got := algo.LinearSearch(input, target)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestBstSearch(t *testing.T) {
	input := []int{2, 3, 5, 7, 9}
	target := 7
	// this should return the index if found

	got := algo.BstSearch(input, target)
	want := 3

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
