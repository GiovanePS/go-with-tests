package iteracao

import (
	"fmt"
	"testing"
)

func TestRepetir(t *testing.T) {
	verifyCorrectOutput := func(t testing.TB, output, expected string) {
		t.Helper()

		if output != expected {
			t.Errorf("Output: %q, expected %q", output, expected)
		}
	}

	t.Run("repete um caractere N vezes passado como argumento", func(t *testing.T) {
		repeats := Repetir("a", 11)
		expected := "aaaaaaaaaaa"

		verifyCorrectOutput(t, repeats, expected)
	})
}

func BenchmarkRepetir(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repetir("a", 1)
	}
}

func ExampleRepetir() {
	output := Repetir("a", 3)
	fmt.Println(output)
	// Output: aaa
}
