package main

import (
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func main() {
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	galaxia.MoverGalaxia()
}
