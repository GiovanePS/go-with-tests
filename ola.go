package main

import "fmt"

const prefixoOla = "Olá, "

func Ola(s string) string {
	if s == "" {
		s = "mundo"
	}

	return fmt.Sprintf("%s%s!", prefixoOla, s)
}
