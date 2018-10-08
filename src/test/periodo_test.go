package models

import (
	"testing"

	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func TestGuardar(t *testing.T) {
	periodo := models.Periodo{
		Inicio:1,
		Fin:3, 
		Pico:0, 
		Clima:"Lluvia", 
		}
	
	esperado := "1"
	resultado := periodo.Guardar()
	if resultado != esperado {
		t.Errorf("TestGuardar was incorrect, got: %v, want: %v. \n", resultado, esperado)
	}
}