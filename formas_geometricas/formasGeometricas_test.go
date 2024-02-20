package formasGeometricas

import (
	"testing"
)

func TestPerimetro(t *testing.T) {
	verifyPerimetro := func(t testing.TB, output, expected float64) {
		t.Helper()

		if output != expected {
			t.Errorf("output: %.2f, expected: %.2f", output, expected)
		}
	}

	t.Run("calcula o perimetro de um retangulo dado altura e largura", func(t *testing.T) {
		retangulo := Retangulo{10.0, 10.0}
		output := retangulo.Perimetro()
		expected := 40.0

		verifyPerimetro(t, output, expected)
	})
}

func TestArea(t *testing.T) {
	testesArea := []struct {
		nome    string
		forma   Forma
		temArea float64
	}{
		{"Retangulo", Retangulo{12.0, 6.0}, 72.0},
		{"Círculo", Circulo{10}, 314.1592653589793},
		{"Triângulo", Triangulo{12, 6}, 36.0},
	}

	for _, tt := range testesArea {
		t.Run(tt.nome, func(t *testing.T) {
			output := tt.forma.Area()
			if output != tt.temArea {
				t.Errorf("%#v > output: %.2f, expected: %.2f", tt.forma, output, tt.temArea)
			}
		})
	}
}
