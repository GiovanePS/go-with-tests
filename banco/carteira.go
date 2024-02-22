package banco

import "fmt"

type Bitcoin float64

type Carteira struct {
	saldo Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%f BTC", b)
}

func (c *Carteira) Depositar(valor Bitcoin) {
	c.saldo += valor
}

func (c Carteira) Saldo() Bitcoin {
	return c.saldo
}
