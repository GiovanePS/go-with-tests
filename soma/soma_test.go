package soma

import "testing"

func TestSoma(t *testing.T) {
	verifySoma := func(t testing.TB, output int, expected int, numeros []int) {
		if output != expected {
			t.Errorf("output %d, expected %d given %v", output, expected, numeros)
		}
	}

	t.Run("soma os inteiros de um array de qualquer tamanho", func(t *testing.T) {
		numeros := []int{1, 2, 3}

		output := Soma(numeros)
		expected := 6

		verifySoma(t, output, expected, numeros)
	})
}
