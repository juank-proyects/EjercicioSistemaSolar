package models

import (
	"testing"

	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func TestAlineados(t *testing.T) {
	var cases = []models.Planet{
		models.Planet{X: -1.18, Y: 3.28},
		models.Planet{X: 1.03, Y: 1.16},
		models.Planet{X: 2.75, Y: -0.49},
	}
	var calc models.Calculadora
	res := calc.Alineados(cases[0], cases[1], cases[2])
	wanted := true
	if res != wanted {
		t.Errorf("Alineados was incorrect, got: %v, want: %v. \n", res, wanted)
	}
}
