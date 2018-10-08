package models

import (
	"testing"

	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func TestGetAngulo(t *testing.T) {
	planeta := models.Planeta{X:321.39380484326966, 
					  Y:383.022221559489, 
					  VelocidadGradosDia:1, 
					  Grados:0, 
					  Radio:500, 
					  SentidoHorario:true, 
					  Name:"ferengi"}
	caso := 50
	esperado := 0.8726646259971648
	resultado := planeta.GetAngulo(caso)
	if resultado != esperado {
		t.Errorf("GetArea was incorrect, got: %v, want: %v. \n", resultado, esperado)
	}
}

func TestActualizarUbicacion(t *testing.T) {
	caso := models.Planeta{
		VelocidadGradosDia:1, 
		Grados:0, 
		Radio:500, 
		SentidoHorario:true, 
		Name:"ferengi"}
	esperado := models.Planeta{
		X:321.39380484326966, 
		Y:383.022221559489, 
		VelocidadGradosDia:1, 
		Grados:0, 
		Radio:500, 
		SentidoHorario:true, 
		Name:"ferengi"}
	resultado := caso.ActualizarUbicacion(50)
	if resultado != esperado {
		t.Errorf("GetArea was incorrect, got: %v, want: %v. \n", resultado, esperado)
	}
}
