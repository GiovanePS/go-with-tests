package inteiros

import (
	"fmt"
	"testing"
)

func TestAdicionador(t *testing.T) {
	soma := Adiciona(2, 2)
	expected := 4

	if soma != expected {
		t.Errorf("output: %d, expected: %d", soma, expected)
	}
}

func ExampleAdiciona() {
	soma := Adiciona(1, 5)
	fmt.Println(soma)
	// Output: 6
}
