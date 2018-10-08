package models

import (
	"fmt"
	
	"github.com/juank-proyects/EjercicioSistemaSolar/src/db"
)	

//Periodo es la estructura de un periodo de tiempo
type Periodo struct {
	tableName struct{} `sql:"periodo"`
	Id 	   int `sql:",pk"`
	Inicio int `sql:"inicio"`
	Fin    int
	Pico   int
	Clima  string
}

//Imprimir imprime el periodo
func (p Periodo) Imprimir() {
	fmt.Printf("Guardado\n %+v\n", p)
}

//Guardar guardar un periodo
func (p Periodo) Guardar() string {
	db := db.Connect()
	defer db.Close()
	err := db.Insert(&p)
	if err != nil {
		return ("Error inserting: " + err.Error())
	} else {
		return "1"
	}
}

//Se guardar un periodo
func (p Periodo) ObtenerClima(dia int) Periodo {
	db := db.Connect()
	defer db.Close()
	err := db.Model(&p).
		Column("clima").
		Where("inicio >= ?", dia).
		Where("inicio <= ?", dia).
		Limit(1).
		Returning("clima").
		Select()
	if err != nil {
		panic(err)
	} else {
		return p
	}
}