package models

import "fmt"

//Periodo es la estructura de un periodo de tiempo
type Periodo struct {
	Inicio int
	Fin    int
	Pico   int
	Lluvia bool
}

//Imprimir imprime el periodo
func (p Periodo) Imprimir() {
	fmt.Printf("%+v\n", p)
}
