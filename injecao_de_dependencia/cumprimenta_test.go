package injecao

import (
	"bytes"
	"testing"
)

func TestCumprimenta(t *testing.T) {
	buffer := bytes.Buffer{}
	Cumprimenta(&buffer, "Chris")

	output := buffer.String()
	expected := "Olá, Chris"

	if output != expected {
		t.Errorf("output %q, expected %q", output, expected)
	}
}
