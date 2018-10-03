package models

import (
	"fmt"
	"math"
)

type Planet struct {
	X                  float64
	Y                  float64
	VelocidadGradosDia float64
	Grados             float64
	Radio              float64
	SentidoHorario     bool
	Name               string
}

func (p Planet) UpdateDay() Planet {
	p.Grados = p.Grados + p.VelocidadGradosDia
	p.X = p.Radio * math.Cos(p.Grados) //499
	p.Y = p.Radio * math.Sin(p.Grados) //8.7
	fmt.Printf("%v ( %g, %g ) => X = %v ; Y = %v \n", p.Name, p.Grados, math.Cos(p.Grados), p.X, p.Y)
	return p
}
