package banco

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	t.Run("Depositar bitcoin", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(10.0))

		expected := Bitcoin(10)

		verifyBalance(t, carteira, expected)
	})

	t.Run("Sacar bitcoin", func(t *testing.T) {
		carteira := Carteira{Bitcoin(100)}
		err := carteira.Sacar(Bitcoin(50))

		expected := Bitcoin(50.0)

		verifyBalance(t, carteira, expected)
		verifyNoError(t, err)
	})

	t.Run("Sacar com saldo insuficiente", func(t *testing.T) {
		carteira := Carteira{Bitcoin(10)}
		err := carteira.Sacar(Bitcoin(20))
		expected := Bitcoin(10)

		verifyBalance(t, carteira, expected)
		verifyError(t, err, ErroSaldoInsuficiente.Error())
	})
}

func verifyBalance(t *testing.T, carteira Carteira, expected Bitcoin) {
	t.Helper()

	output := carteira.Saldo()

	if output != expected {
		t.Errorf("output %v, expected %v", output, expected)
	}
}

func verifyError(t *testing.T, err error, expected string) {
	t.Helper()

	if err == nil {
		t.Fatal("Expected error, but none occured")
	}

	output := err.Error()

	if output != expected {
		t.Errorf("Output %s, expected %s", output, expected)
	}
}

func verifyNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("Unexpected error occured.")
	}
}
