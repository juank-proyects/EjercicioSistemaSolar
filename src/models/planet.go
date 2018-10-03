package models

import (
	"math"
)

//Planet es la estructura de un planeta
type Planet struct {
	X                  float64
	Y                  float64
	VelocidadGradosDia float64
	Grados             float64
	Radio              float64
	SentidoHorario     bool
	Name               string
}

//UpdateDay mueve un planeta
func (p Planet) UpdateDay() Planet {
	p.Grados = p.Grados + p.VelocidadGradosDia
	p.Grados = (p.Grados * math.Pi) / 180
	p.X = p.Radio * math.Cos(p.Grados) //499
	p.Y = p.Radio * math.Sin(p.Grados) //8.7
	//fmt.Printf("%v ( %g, %v ) => X = %v ; Y = %v \n", p.Name, p.Grados, math.Cos(p.Grados), p.X, p.Y)
	return p
}
