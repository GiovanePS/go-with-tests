package iteracao

// Recebe um caractere e repete-o N vezes o valor recebido em reps.
func Repetir(caractere string, reps int) string {
	var output string
	for range reps {
		output += caractere
	}

	return output
}
