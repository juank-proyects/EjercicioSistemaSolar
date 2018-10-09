package models

import (
	"fmt"

	"github.com/juank-proyects/EjercicioSistemaSolar/src/db"
)

//Periodo es la estructura de un periodo de tiempo
type Periodo struct {
	tableName struct{} `sql:"periodos"`
	ID        int      `sql:",pk"`
	Inicio    int      `sql:"inicio"`
	Fin       int
	Pico      int
	Perimetro float32
	Clima     string
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

//ObtenerPeriodo obtiene el perido para el dia X
func (p Periodo) ObtenerPeriodo(dia int) Periodo {
	db := db.Connect()
	defer db.Close()
	db.Model(&p).
		Column("clima").
		Where("inicio >= ?", dia).
		Where("inicio <= ?", dia).
		Limit(1).
		Returning("clima").
		Select()
	return p
}

//CantidadPeriodos retorna la cantidad de periodos
func (p Periodo) CantidadPeriodos() (int, error) {
	db := db.Connect()
	defer db.Close()
	count, err := db.Model(&Periodo{}).Count()
	if err != nil {
		return 0, err
	}
	return count, nil
}

//EliminarPeriodos Elimina todos los periodos en la BD
func (p Periodo) EliminarPeriodos() int {
	db := db.Connect()
	defer db.Close()
	query := "DELETE FROM periodos"
	res, err := db.Exec(query)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
	fmt.Printf("filas afetadas %v\n", res.RowsAffected())
	return res.RowsAffected()
}

//CantidadPeriodosClima retorna la cantidad de periodos para determinado clima
func (p Periodo) CantidadPeriodosClima(clima string) (int, error) {
	db := db.Connect()
	defer db.Close()
	cantidad, err := db.Model(&p).
		Where("clima = ?", clima).
		Count()
	if err != nil {
		return 0, err
	}
	return cantidad, nil
}

//PicoMaximoLluvia retorna el dia del pico maximo de lluvia
func (p Periodo) PicoMaximoLluvia() (int, error) {
	db := db.Connect()
	defer db.Close()
	var periodo Periodo
	err := db.Model(periodo).Where("Clima = ?", "Lluvia").Order("perimetro ASC").Limit(1).Select()
	if err != nil {
		return 0, err
	}
	return periodo.Pico, nil
}
