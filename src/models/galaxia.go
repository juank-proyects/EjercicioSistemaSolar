package models

import (
	"fmt"
	"math"
)

//Galaxia es la estructura de un Galaxia
type Galaxia struct {
	dias     int
	planetas []Planet
}

//Iniciar carga todos los planetas en la galaxia
func (g Galaxia) Iniciar() Galaxia {
	ferengi := Planet{Radio: 500, VelocidadGradosDia: 1, SentidoHorario: true, Name: "ferengi", X: 500, Y: 0, Grados: 0}
	betasoide := Planet{Radio: 2000, VelocidadGradosDia: 3, SentidoHorario: true, Name: "betasoide", X: 200, Y: 0, Grados: 0}
	vulcano := Planet{Radio: 1000, VelocidadGradosDia: 5, SentidoHorario: false, Name: "vulcano", X: 500, Y: 0, Grados: 0}
	g.planetas = append(g.planetas, ferengi)
	g.planetas = append(g.planetas, betasoide)
	g.planetas = append(g.planetas, vulcano)
	return g
}

//MoverGalaxia es la funcion que calcula el clima de los proximos N Dias
func (g Galaxia) MoverGalaxia(anios int) {
	cantDias := g.dias * anios
	cantPeriodsRain := 0
	lastRes := false
	for i := 0; i <= cantDias; i++ {
		g.ActualizarGalaxia(i)
		if g.IsSunIn() {
			if !lastRes {
				cantPeriodsRain++
				lastRes = true
			}
		} else {
			lastRes = false
		}
	}
	fmt.Printf("cantidad temporadas de lluvia %v \n", cantPeriodsRain)
}

//ActualizarGalaxia actualzia la ubicacion del los 3 planetas de la BD
func (g Galaxia) ActualizarGalaxia(dias int) {
	for i := range g.planetas {
		g.planetas[i].ActualizarUbicacion(dias)
	}
}

//IsSunIn dato 3 planetas calcula si el sol se encuentra en el medio
func (g Galaxia) IsSunIn() bool {
	a := g.planetas[0]
	b := g.planetas[1]
	c := g.planetas[2]

	var sol Planet
	var res = false
	sol.X = 0
	sol.Y = 0

	var areaPlanetas = g.GetArea(a, b, c)
	var areaA = g.GetArea(a, b, sol)
	var areaB = g.GetArea(a, c, sol)
	var areaC = g.GetArea(b, c, sol)
	var areaABC = areaA + areaB + areaC
	if areaABC == areaPlanetas {
		res = true
		fmt.Print("esta dentro\n")
	} else {
		fmt.Print("esta fuera\n")
	}

	return res
}

//GetArea obtiene el area de 3 puntos
func (g Galaxia) GetArea(a Planet, b Planet, c Planet) float64 {
	var area = (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2
	return math.Abs(area)
}
