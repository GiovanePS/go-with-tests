package inteiros

import "testing"

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	expected := 4

	if soma != expected {
		t.Errorf("output: %d, expected: %d", soma, expected)
	}
}
