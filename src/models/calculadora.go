package models

import (
	//"fmt"
	"math"
)

//Calculadora es la estructura de un Calculadora
type Calculadora struct {
}

//Alineados calcula si 3 planetas estan alineados
func (calc Calculadora) Alineados(a, b, c Planeta) bool {
	res := false
	if (a.X == b.X && b.X == c.X) || (a.Y == b.Y && b.Y == c.Y) {
		res = true
	} else {
		res = calc.CalcularDentroDeRango(calc.Vector(a.X, b.X, c.X), calc.Vector(a.Y, b.Y, c.Y))
	}
	return res
}

//Vector calcula la direccion de un Vector
func (calc Calculadora) Vector(a, b, c float64) float64 {
	res := (b - a) / (c - b)
	res = calc.Redondear(res)
	//fmt.Printf("vector %g \n", res)
	return res
}

//Redondear funcion que redondea los valores a 2 decimales
func (calc Calculadora) Redondear(num float64) float64 {
	return math.Round((num)*100) / 100
}

//CalcularDentroDeRango funcion que calcula si la diferencia de 2 numeros son mayores a un X
func (calc Calculadora) CalcularDentroDeRango(a, b float64) bool {
	return math.Abs(math.Abs(a)-math.Abs(b)) < 0.05
}
