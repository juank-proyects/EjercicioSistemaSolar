package models

import "fmt"

//Periodo es la estructura de un periodo de tiempo
type Periodo struct {
	Inicio int
	Fin    int
	Pico   int
	Clima  string
}

//Imprimir imprime el periodo
func (p Periodo) Imprimir() {
	fmt.Printf("Guardado\n %+v\n", p)
}

//Guardar guarda el periodo en l BD
func (p Periodo) Guardar() {
	p.Imprimir()
}