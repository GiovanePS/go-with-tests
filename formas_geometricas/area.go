package formasGeometricas

import "math"

func (r Retangulo) Area() float64 {
	return r.Altura * r.Largura
}

func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}
