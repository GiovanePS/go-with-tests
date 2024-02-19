package soma

func Soma(n []int) int {
	soma := 0
	for _, v := range n {
		soma += v
	}

	return soma
}
