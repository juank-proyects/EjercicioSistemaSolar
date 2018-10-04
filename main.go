package main

import (
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

var ferengi models.Planet
var betasoide models.Planet
var vulcano models.Planet
var anios = 1

func main() {
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	galaxia.MoverGalaxia(anios)
}
