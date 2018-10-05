package models

import (
	"testing"

	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func TestAlineados(t *testing.T) {
	var cases = []models.Planet{
		models.Planet{X: 1, Y: 1},
		models.Planet{X: 2, Y: 2},
		models.Planet{X: 3, Y: 3},
	}
	var calc models.Calculadora
	res := calc.Alineados(cases[0], cases[1], cases[2])
	wanted := true
	if res != wanted {
		t.Errorf("Alineados was incorrect, got: %v, want: %v. \n", res, wanted)
	}
}
