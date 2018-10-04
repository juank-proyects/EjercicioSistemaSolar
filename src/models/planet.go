package models

import (
	"fmt"
	"math"
)

//Planet es la estructura de un planeta
type Planet struct {
	X                  float64
	Y                  float64
	VelocidadGradosDia int
	Grados             float64
	Radio              float64
	SentidoHorario     bool
	Name               string
}

//ActualizarUbicacion mueve un planeta
func (p Planet) ActualizarUbicacion(dia int) Planet {
	angulo := p.GetAngulo(dia)
	p.X = p.Radio * math.Cos(angulo)
	p.Y = p.Radio * math.Sin(angulo)
	fmt.Printf("%+v\n", p)
	return p
}

//GetAngulo retorna el angulo en el que se encuentra el planeta
func (p Planet) GetAngulo(dia int) float64 {
	grados := (dia * p.VelocidadGradosDia) % 360
	if !p.SentidoHorario {
		grados = 360 - grados
	}
	return (float64(grados) * math.Pi) / float64(180)
}
