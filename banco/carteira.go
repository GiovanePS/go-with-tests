package banco

import (
	"errors"
	"fmt"
)

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

var ErroSaldoInsuficiente = errors.New("não é possível retirar: saldo insuficiente")

func (c *Carteira) Sacar(valor Bitcoin) error {
	if valor > c.saldo {
		return ErroSaldoInsuficiente
	}

	c.saldo -= valor
	return nil
}

func (c Carteira) Saldo() Bitcoin {
	return c.saldo
}
