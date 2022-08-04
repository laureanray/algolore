package main

import "testing"

func TestImpl(t *testing.T) {
	got := Impl(1, 2)
	want := 3

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
