package main

import "testing"

func TestOla(t *testing.T) {
	output := Ola("Giovane")
	expected := "Olá, Giovane!"

	if output != expected {
		t.Errorf("output %q, expected %q", output, expected)
	}
}
