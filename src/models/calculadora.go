package models

//Alineados calcula si 3 planetas estan alineados
func Alineados(a, b, c Planet) bool {
	return Vector(a.X, b.X, c.X) == Vector(a.Y, b.Y, c.Y)
}

//Vector calcula la direccion de un Vector
func Vector(a, b, c float64) float64 {
	return (b - a) / (c - b)
}
