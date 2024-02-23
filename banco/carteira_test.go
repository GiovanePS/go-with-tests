package banco

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	verifyBalance := func(t *testing.T, carteira Carteira, expected Bitcoin) {
		t.Helper()

		output := carteira.Saldo()

		if output != expected {
			t.Errorf("output %v, expected %v", output, expected)
		}
	}
	t.Run("Depositar bitcoin", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10.0))

		expected := Bitcoin(10)

		verifyBalance(t, carteira, expected)
	})

	t.Run("Sacar bitcoin", func(t *testing.T) {
		carteira := Carteira{Bitcoin(100)}
		carteira.Sacar(Bitcoin(50))

		expected := Bitcoin(50.0)

		verifyBalance(t, carteira, expected)
	})

	t.Run("Sacar com saldo insuficiente", func(t *testing.T) {
		carteira := Carteira{Bitcoin(10)}
		err := carteira.Sacar(Bitcoin(20))
		expected := Bitcoin(10)

		verifyBalance(t, carteira, expected)

		if err != nil {
			t.Error("Expected error, but none occured")
		}
	})
}
