package main

import (
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

var ferengi models.Planeta
var betasoide models.Planeta
var vulcano models.Planeta
var anios = 1

func main() {
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	galaxia.MoverGalaxia(anios)
}
