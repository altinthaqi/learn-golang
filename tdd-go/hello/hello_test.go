package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {

		got := Hello("Altin", "English")
		want := "Hello, Altin"

		assertCorrectMessage(t, got, want)
	})
	t.Run("say 'Hello, World' when an empty string i supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	t.Run("defaults to English when no language provided", func(t *testing.T) {
		got := Hello("Altin", "")
		want := "Hello, Altin"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Altin", "Spanish")
		want := "Hola, Altin"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Altin", "French")
		want := "Bonjour, Altin"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
