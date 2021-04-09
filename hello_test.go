package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Zach")
	want := "Hello, Zach"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
