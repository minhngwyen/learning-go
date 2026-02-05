package main

import "testing"

func TestHello(t *testing.T) {
	got := hello("Minh")
	want := "Nihao, Minh"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}