package soma

import (
	"reflect"
	"testing"
)

func TestSoma(t *testing.T) {
	verifySoma := func(t testing.TB, output int, expected int, numeros []int) {
		t.Helper()

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

func TestSomaTudo(t *testing.T) {
	verifySomaTudo := func(t testing.TB, output, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(output, expected) {
			t.Errorf("output %v, expected %v", output, expected)
		}
	}

	t.Run("soma os slices passados como argumento e devolve um slice contendo a soma de cada slice passado, em ordem.", func(t *testing.T) {
		output := SomaTudo([]int{1, 2}, []int{0, 9})
		expected := []int{3, 9}

		verifySomaTudo(t, output, expected)
	})
}

func TestSomaTodoOResto(t *testing.T) {
	verifySomaTodoOResto := func(t testing.TB, output, expected []int) {
		t.Helper()

		if !reflect.DeepEqual(output, expected) {
			t.Errorf("output %v, expected %v", output, expected)
		}
	}

	t.Run("soma todos os valores após o valor de índice 0 de um slice.", func(t *testing.T) {
		output := SomaTodoOResto([]int{1, 2}, []int{2, 9})
		expected := []int{2, 9}

		verifySomaTodoOResto(t, output, expected)
	})

	t.Run("soma slices vazios de forma segura", func(t *testing.T) {
		output := SomaTodoOResto([]int{}, []int{2, 9})
		expected := []int{0, 9}

		verifySomaTodoOResto(t, output, expected)
	})
}
