package db

import (
	"github.com/go-pg/pg"
)

//Connect funcion para conectarse a la BD
func Connect() *pg.DB {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5432",
		User:     "postgres",
		Password: "postgres",
		Database: "clima",
	})
	return db
}
