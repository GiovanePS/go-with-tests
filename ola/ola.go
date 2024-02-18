package main

const prefixoOlaPortugues = "Olá, "
const prefixoOlaEspanhol = "Hola, "
const prefixoOlaFrances = "Bonjour, "
const sufixoExclamacao = "!"

func Ola(nome string, idioma string) string {
	if nome == "" {
		nome = "mundo"
	}

	switch idioma {
	case "espanhol":
		return prefixoOlaEspanhol + nome + sufixoExclamacao
	case "frances":
		return prefixoOlaFrances + nome + sufixoExclamacao
	default:
		return prefixoOlaPortugues + nome + sufixoExclamacao
	}
}
