package main

import "testing"

func TestOla(t *testing.T) {	
	t.Run("should say hello to people", func(t *testing.T) {
		result := Hello("Joao")
		expected := "Hello, Joao"
	
		if result != expected{
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})

	t.Run("should say 'Hello, word' when name is empty", func(t *testing.T) {
		result := Hello("")
		expected := "Hello, word"
	
		if result != expected{
			t.Errorf("result '%s', expected '%s'", result, expected)
		}
	})
	
}