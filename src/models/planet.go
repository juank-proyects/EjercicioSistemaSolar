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

//UpdateDay mueve un planeta
func (p Planet) UpdateDay(dia int) {
	grados := (dia * p.VelocidadGradosDia) % 360
	if !p.SentidoHorario {
		grados = 360 - grados
	}
	res := (float64(grados) * math.Pi) / float64(180)
	p.X = p.Radio * math.Cos(res) //499
	p.Y = p.Radio * math.Sin(res) //8.7

	fmt.Printf("%v ( %g ) => X = %v ; Y = %v \n", p.Name, res, p.X, p.Y)
}

//GetGrados retorna los grados en los que se encuentra el planeta
func (p Planet) GetGrados() float64 {
	res := p.Grados
	if !p.SentidoHorario {
		res = 360 - p.Grados
	}
	return res
}
