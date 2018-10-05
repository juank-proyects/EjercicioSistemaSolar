package models

//Calculadora es la estructura de un Calculadora
type Calculadora struct {
}

//Alineados calcula si 3 planetas estan alineados
func (calc Calculadora) Alineados(a, b, c Planet) bool {
	res := false
	if (a.X == b.X && b.X == c.X) || (a.Y == b.Y && b.Y == c.Y) {
		res = true
	} else {
		res = calc.Vector(a.X, b.X, c.X) == calc.Vector(a.Y, b.Y, c.Y)
	}
	return res
}

//Vector calcula la direccion de un Vector
func (calc Calculadora) Vector(a, b, c float64) float64 {
	return (b - a) / (c - b)
}
