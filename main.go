package main

import (
	"fmt"
	"math"

	"github.com/juankis/SistemaSolar/src/models"
)

var ferengi models.Planet
var betasoide models.Planet
var vulcano models.Planet
var dias int = 10
var anios int = 1

func main() {
	ferengi.Radio = 500
	ferengi.VelocidadGradosDia = 1
	ferengi.SentidoHorario = true
	ferengi.Name = "ferengi"
	ferengi.X = 0
	ferengi.Y = ferengi.Radio
	ferengi.Grados = 0

	betasoide.Radio = 2000
	betasoide.VelocidadGradosDia = 3
	betasoide.SentidoHorario = true
	betasoide.Name = "betasoide"
	betasoide.X = 0
	betasoide.Y = betasoide.Radio
	betasoide.Grados = 0

	vulcano.Radio = 1000
	vulcano.VelocidadGradosDia = 5
	vulcano.SentidoHorario = false
	vulcano.Name = "vulcano"
	vulcano.X = 0
	vulcano.Y = vulcano.Radio
	vulcano.Grados = 0

	//IsSunIn(ferengi, betasoide, vulcano)
	MoveGalaxy()
}

//isSunIn dato 3 planetas calcula si el sol se encuentra en el medio
func IsSunIn(a models.Planet, b models.Planet, c models.Planet) bool {
	var sol models.Planet
	var res bool = false
	sol.X = 0
	sol.Y = 0

	var areaPlanetas float64 = GetArea(a, b, c)
	var areaA float64 = GetArea(a, b, sol)
	var areaB float64 = GetArea(a, c, sol)
	var areaC float64 = GetArea(b, c, sol)
	var areaABC float64 = areaA + areaB + areaC
	if areaABC == areaPlanetas {
		res = true
		fmt.Print("esta dentro\n")
	} else {
		fmt.Print("esta fuera\n")
	}

	return res
}

//getArea obtiene el area de 3 puntos
func GetArea(a models.Planet, b models.Planet, c models.Planet) float64 {
	var area float64 = (a.X*(b.Y-c.Y) + b.X*(c.Y-a.Y) + c.X*(a.Y-b.Y)) / 2
	return math.Abs(area)
}

func MoveGalaxy() {
	cantDias := dias * anios
	cantPeriodsRain := 0
	lastRes := false
	for i := 1; i <= cantDias; i++ {
		UpdateUbications()
		if IsSunIn(ferengi, betasoide, vulcano) {
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

func UpdateUbications() {
	ferengi = ferengi.UpdateDay()
	betasoide = betasoide.UpdateDay()
	vulcano = vulcano.UpdateDay()
}
