package main

import "testing"

func TestOla(t *testing.T) {
	result := Hello("Joao")
	expected := "Olá, Joao"

	if result != expected{
		t.Errorf("result '%s', expected '%s'", result, expected)
	}
}