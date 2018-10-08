package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/juank-proyects/EjercicioSistemaSolar/src/models"
)

func Start() {

	router := gin.Default()
	router.GET("/clima", handleStats)
	router.Run(":9990")

}

func handleStats(c *gin.Context) {
	dia := c.Query("dia")
	i, err := strconv.Atoi(dia)
	if err != nil {
		c.String(400, "400-Bad-Request: Asegurese de ingresar un numero entero")
		return
	}
	var periodo models.Periodo
	periodo = periodo.ObtenerPeriodo(i)
	if periodo == (models.Periodo{}) {
		c.JSON(200, gin.H{
			"dia":   dia,
			"clima": "No se cuenta con este dato, no esta dentro los proximos 10 anios",
		})
		return
	}

	c.JSON(200, gin.H{
		"dia":   dia,
		"clima": periodo.Clima,
	})
	return
}
