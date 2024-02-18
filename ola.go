package main

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const sufixoExclamacao = "!"

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	if idioma == "espanhol" {
		return prefixoOlaEspanhol + nome + sufixoExclamacao
	}

	return prefixoOlaPortugues + nome + sufixoExclamacao
}
