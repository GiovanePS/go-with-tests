package soma

func SomaTodoOResto(numeros ...[]int) []int {
	var somas []int

	for _, n := range numeros {
		if len(n) == 0 {
			somas = append(somas, 0)
		} else {
			final := n[1:]
			somas = append(somas, Soma(final))
		}
	}

	return somas
}
