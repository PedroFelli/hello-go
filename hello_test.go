package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectNessage := func(t testing.TB, got, want string){
		t.Helper()
		if got != want {
			t.Errorf("got %s, want %s", got, want)
		}
	}
	
	t.Run("should say hello to people", func(t *testing.T) {
		got := Hello("Joao", "")
		want := "Hello, Joao"
		assertCorrectNessage(t, got, want)
	})

	t.Run("should say 'Hello, word' when name is empty", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectNessage(t, got, want)
	})

	t.Run("should say hello in spanish", func (t *testing.T) {
		got := Hello("Joao", "spanish")
		want := "Hola, Joao"
		assertCorrectNessage(t, got, want)
	})

	t.Run("should say hello in french", func (t *testing.T) {
		got := Hello("Joao", "french")
		want := "Salut, Joao"
		assertCorrectNessage(t, got, want)
	})
	
}