package main

import "testing"

func TestOla(t *testing.T) {
	verificaMensagemCorreta := func(t testing.TB, output string, expected string) {
		t.Helper()
		if output != expected {
			t.Errorf("output %q, expected %q", output, expected)
		}
	}
	t.Run("diz olá para as pessoas", func(t *testing.T) {
		output := Ola("Giovane")
		expected := "Olá, Giovane!"

		verificaMensagemCorreta(t, output, expected)
	})

	t.Run("diz 'Olá, mundo' quando uma string vazia for passada", func(t *testing.T) {
		output := Ola("")
		expected := "Olá, mundo!"

		verificaMensagemCorreta(t, output, expected)
	})
}
