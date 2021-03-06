package models

import (
	"testing"

	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func TestGetArea(t *testing.T) {
	var cases = []models.Planeta{
		models.Planeta{X: 1, Y: 1},
		models.Planeta{X: 2, Y: 3},
		models.Planeta{X: 100, Y: 5},
	}

	var galaxia models.Galaxia

	res := galaxia.GetArea(cases[0], cases[1], cases[2])
	if res != 97 {
		t.Errorf("GetArea was incorrect, got: %v, want: %v. \n", res, 1)
	}
}
