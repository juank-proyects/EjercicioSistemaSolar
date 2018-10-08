package models

import (
	"fmt"
	"math"
)

//Planeta es la estructura de un planeta
type Planeta struct {
	X                  float64
	Y                  float64
	VelocidadGradosDia int
	Grados             float64
	Radio              float64
	SentidoHorario     bool
	Name               string
}

//ActualizarUbicacion mueve un planeta
func (p Planeta) ActualizarUbicacion(dia int) Planeta {
	angulo := p.GetAngulo(dia)
	p.X = p.Radio * math.Cos(angulo)
	p.Y = p.Radio * math.Sin(angulo)
	fmt.Printf("%+v\n", p)
	fmt.Printf("angulo %v  dia %v \n", angulo, dia)
	return p
}

//GetAngulo retorna el angulo en el que se encuentra el planeta
func (p Planeta) GetAngulo(dia int) float64 {
	grados := (dia * p.VelocidadGradosDia) % 360
	if !p.SentidoHorario {
		grados = 360 - grados
	}
	return (float64(grados) * math.Pi) / float64(180)
}
