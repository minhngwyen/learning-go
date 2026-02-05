package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func (t *testing.T){
		got := hello("Minh")
		want := "Nihao, Minh"
		
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Nihao, World' when an empty string is supplied", func(t *testing.T){
		got := hello("")
		want := "Nihao, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string){
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}