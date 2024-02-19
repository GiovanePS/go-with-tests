package soma

func SomaTudo(numeros ...[]int) []int {
	var somas []int

	for _, n := range numeros {
		somas = append(somas, Soma(n))
	}

	return somas
}
