package banco

import (
	"testing"
)

func TestCarteira(t *testing.T) {
	carteira := Carteira{}
	carteira.Depositar(Bitcoin(10.0))

	output := carteira.Saldo()
	expected := Bitcoin(10)

	expected++

	if output != expected {
		t.Errorf("output: %v, expected: %v", output, expected)
	}
}
