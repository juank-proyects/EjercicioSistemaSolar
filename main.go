package main

import (
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
	"github.com/juank-proyects/EjercicioSistemaSolar/src/api"
)

func main() {
	LoadApi()
}

func LoadApi() {
	api.Start()
}

func MoverGalaxia(){
	var anios = 1
	var galaxia models.Galaxia
	galaxia = galaxia.Iniciar()
	galaxia.MoverGalaxia(anios)
}