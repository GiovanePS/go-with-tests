package dicionario

import (
	"testing"
)

func TestBusca(t *testing.T) {
	dicionario := Dicionario{"test": "this is only a test"}

	t.Run("busca em um dicionário", func(t *testing.T) {
		searchArg := "test"
		output, err := dicionario.Busca(searchArg)
		expected := "this is only a test"

		verifyErr(t, err)
		verifyOutputExpectedGiven(t, output, expected, searchArg)
	})

	t.Run("busca de uma palavra que não está em um dicionário", func(t *testing.T) {
		searchArg := "not exist"
		_, err := dicionario.Busca(searchArg)

		verifyTextErr(t, err, ErrWordNotFound)
		verifyAssertErr(t, err, "searching inexisting key")
	})
}

func TestAdiciona(t *testing.T) {
	key, value := "value", "teste value"
	t.Run("adiciona um valor em um dicionário", func(t *testing.T) {
		dicionario := Dicionario{}

		dicionario.Adicionar(key, value)

		output, err := dicionario.Busca(key)
		expected := value

		verifyErr(t, err)
		verifyOutputExpected(t, output, expected)
	})

	t.Run("adicionar um valor existente deve gerar um erro", func(t *testing.T) {
		dicionario := Dicionario{key: value}

		err := dicionario.Adicionar(key, value)

		verifyAssertErr(t, err, "adding inexisting key")
		verifyTextErr(t, err, ErrWordExistsOnDictionary)
	})
}

func TestAtualiza(t *testing.T) {
	key, otherKey, value, otherValue := "value", "other value", "teste value", "other test value"

	t.Run("atualiza um valor em um dicionário", func(t *testing.T) {
		dicionario := Dicionario{key: value}
		dicionario.Atualizar(key, otherValue)

		output, err := dicionario.Busca(key)
		verifyErr(t, err)

		expected := otherValue

		verifyOutputExpected(t, output, expected)
	})

	t.Run("atualizar um valor inexistente deve gerar um erro", func(t *testing.T) {
		dicionario := Dicionario{key: value}

		err := dicionario.Atualizar(otherKey, otherValue)

		verifyAssertErr(t, err, "updating inexisting key")
	})
}

func TestRemove(t *testing.T) {
	key, value := "value", "teste value"

	t.Run("remove um valor em um dicionário", func(t *testing.T) {
		dicionario := Dicionario{key: value}
		dicionario.Remover(key)
		_, err := dicionario.Busca(key)

		verifyAssertErr(t, err, "on deleting")
	})
}

func verifyOutputExpectedGiven(t *testing.T, output string, expected string, given string) {
	t.Helper()

	if output != expected {
		t.Errorf("output: %q, expected: %q, given %q", output, expected, given)
	}
}

func verifyOutputExpected(t *testing.T, output string, expected string) {
	t.Helper()

	if output != expected {
		t.Errorf("output: %q, expected: %q", output, expected)
	}
}

func verifyErr(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatalf("error not expected ocurred: %v", err.Error())
	}
}

func verifyAssertErr(t *testing.T, err error, s string) {
	t.Helper()

	if err == nil {
		t.Fatalf("expected error %s, but none occured", s)
	}
}

func verifyTextErr(t *testing.T, output error, expected error) {
	t.Helper()

	if output != expected {
		t.Errorf("output error: %s, expected: %s", output, expected)
	}
}
