package models

import (
	"fmt"
	"math"
)

//Galaxia es la estructura de un Galaxia
type Galaxia struct {
	dias     int
	planetas []Planet
	sol      Planet
}

//Iniciar carga todos los planetas en la galaxia
func (g Galaxia) Iniciar() Galaxia {
	ferengi := Planet{Radio: 500, VelocidadGradosDia: 1, SentidoHorario: true, Name: "ferengi", X: 500, Y: 0, Grados: 0}
	betasoide := Planet{Radio: 2000, VelocidadGradosDia: 3, SentidoHorario: true, Name: "betasoide", X: 200, Y: 0, Grados: 0}
	vulcano := Planet{Radio: 1000, VelocidadGradosDia: 5, SentidoHorario: false, Name: "vulcano", X: 500, Y: 0, Grados: 0}
	g.planetas = append(g.planetas, ferengi)
	g.planetas = append(g.planetas, betasoide)
	g.planetas = append(g.planetas, vulcano)
	g.sol = Planet{X: 0, Y: 0}
	g.dias = 100
	return g
}

//MoverGalaxia es la funcion que calcula el clima de los proximos N Dias
func (g Galaxia) MoverGalaxia(anios int) {
	cantDias := g.dias * anios
	cantPeriodsRain := 0
	lastRes := false
	var periodo Periodo
	for i := 0; i <= cantDias; i++ {
		g = g.ActualizarGalaxia(i)
		fmt.Printf("DIA: %v \n", i)
		if g.IsSunIn() {
			if !lastRes {
				periodo = Periodo{}
				periodo.Inicio = i
				lastRes = true
				periodo.Lluvia = true
			} else {
				periodo.Fin = i
			}
		} else {
			if lastRes {
				cantPeriodsRain++
				periodo.Imprimir()
			}
			lastRes = false
		}
	}
	fmt.Printf("cantidad temporadas de lluvia %v \n", cantPeriodsRain)
}

//ActualizarGalaxia actualzia la ubicacion del los 3 planetas de la BD
func (g Galaxia) ActualizarGalaxia(dias int) Galaxia {
	for i := range g.planetas {
		g.planetas[i] = g.planetas[i].ActualizarUbicacion(dias)
	}
	return g
}

//IsSunIn dato 3 planetas calcula si el sol se encuentra en el medio
func (g Galaxia) IsSunIn() bool {
	var res = false
	a := g.planetas[0]
	b := g.planetas[1]
	c := g.planetas[2]

	var areaPlanetas = g.GetArea(a, b, c)
	var areaA = g.GetArea(a, b, g.sol)
	var areaB = g.GetArea(a, c, g.sol)
	var areaC = g.GetArea(b, c, g.sol)
	fmt.Printf("SIN REDONDEAR-ABC: %v areaPlanetas: %v \n", (areaA + areaB + areaC), areaPlanetas)
	var areaABC = math.Round((areaA+areaB+areaC)*100) / 100
	areaPlanetas = math.Round((areaPlanetas)*100) / 100
	fmt.Printf("areaABC: %v areaPlanetas: %v \n", areaABC, areaPlanetas)
	if areaABC == areaPlanetas {
		res = true
		fmt.Print("ESTA DENTRO\n")
	} else {
		fmt.Print("esta fuera\n")
	}

	return res
}

//GetArea obtiene el area de 3 puntos
func (g Galaxia) GetArea(a Planet, b Planet, c Planet) float64 {
	var area = (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2
	fmt.Print("Planetas\n")
	fmt.Printf("%+v\n", a)
	fmt.Printf("%+v\n", b)
	fmt.Printf("%+v\n", c)
	fmt.Printf("Area de los planetas %+v\n", area)
	fmt.Print("Fin Planetas\n")
	return math.Abs(area)
}
