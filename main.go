package main

import (
	"github.com/juank-proyects/EjercicioSistemaSolar/src/api"
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func main() {
	LoadApi()
	//MoverGalaxia()
}

func LoadApi() {
	api.Start()
}

func MoverGalaxia() {
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	galaxia.MoverGalaxia()
}
