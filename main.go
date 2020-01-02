package gotest_area

import "math"

const Pi = 3.1416

// veja a funcao
func Circ(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// outra funcao
func Rect(base, altura float64) float64 {
	return base * altura
}

func _TrianguloEq(base, altura float64) float64 {
	return (base * altura)
}
